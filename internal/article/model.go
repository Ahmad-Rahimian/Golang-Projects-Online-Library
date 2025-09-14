package article

type Article struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Short_summary string `json:"short_summary"`
	Full_text     string `json:"full_text"`
	Author        string `json:"author"`
	Cover_image   string `json:"cover_image"`
	Reading_time  *int   `json:"reading_time"`
}
