package domain

type PaidBook struct {
	ID       int
	Title    string
	Summery  *string
	Author   string
	CoverURL string
	FileURL  string
	Pages    int
	price    int
}
