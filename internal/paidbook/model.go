package paidbook

type PaidBook struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Summary     *string `json:"summary"`
	Author      string  `json:"author"`
	Cover_image string  `json:"cover_image"`
	Pdf_file    string  `json:"pdf_file"`
	Pages       int     `json:"pages"`
	Price       int     `json:"price"`
}
