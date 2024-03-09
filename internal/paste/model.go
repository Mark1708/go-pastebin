package paste

type ResponseDTO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Visibility string `json:"visibility"`
	CreatedAt  string `json:"created_at"`
	ExpiresAt  string `json:"expires_at"`
	User       string `json:"user"`
	Content    string `json:"content"`
}

type RequestDTO struct {
	Title      string `json:"title" minLength:"1" maxLength:"256" example:"Some Title" validate:"required"`
	Visibility string `json:"visibility" enum:"PUBLIC,PRIVATE,UNLISTED" default:"PUBLIC" validate:"optional"`
	Content    string `json:"content" maxLength:"2048" example:"Some Content" validate:"required"`
	Expires    string `json:"expires" validate:"optional" example:"2024-01-02T15:04:05.999Z03:00"`
}
