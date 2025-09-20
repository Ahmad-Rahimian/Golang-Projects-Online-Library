package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// root path of project (where main.go is)
	root := "./"

	// routes to check
	routesToCheck := map[string]int{
		`/auth/send-otp`:   0,
		`/auth/verify-otp`: 0,
	}

	// function to scan files
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fs := token.NewFileSet()
		node, err := parser.ParseFile(fs, path, nil, parser.AllErrors)
		if err != nil {
			return nil
		}

		ast.Inspect(node, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if sel.Sel.Name == "POST" || sel.Sel.Name == "GET" || sel.Sel.Name == "PUT" || sel.Sel.Name == "DELETE" {
				if len(call.Args) > 0 {
					if lit, ok := call.Args[0].(*ast.BasicLit); ok {
						pathVal := strings.Trim(lit.Value, `"`)
						if _, exists := routesToCheck[pathVal]; exists {
							routesToCheck[pathVal]++
						}
					}
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print results
	for route, count := range routesToCheck {
		if count > 1 {
			fmt.Printf("Warning: route %s is declared %d times\n", route, count)
		} else {
			fmt.Printf("OK: route %s is declared %d time(s)\n", route, count)
		}
	}
}
