package paste

type RequestDto struct {
	Title      string `json:"title"`
	Visibility string `json:"visibility"`
	Content    string `json:"content"`
	Expires    string `json:"expires"`
}
