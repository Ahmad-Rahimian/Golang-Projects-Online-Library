package domain

type Article struct {
	ID        int
	Title     string
	ShortDesc string
	Author    string
	CoverURL  string
	Content   string
	ReadTime  int
}
