package response

type PostResult struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	AuthorId    string `json:"authorId"`
	Media       string `json:"media"`
}
