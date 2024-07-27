package requests

type UpdatePost struct {
	ID          uint   `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	AuthorId    string `json:"authorId,omitempty"`
	Media       string `json:"media,omitempty"`
}
