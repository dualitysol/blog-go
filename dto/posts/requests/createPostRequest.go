package requests

type CreatePost struct {
	Title       string `validate:"required,min=1,max=120", json:"title"`
	Description string `json:"description"`
	Content     string `validate:"required", json:"content"`
	AuthorId    string `validate:"required", json:"authorId"`
	Media       string `json:"media"`
}
