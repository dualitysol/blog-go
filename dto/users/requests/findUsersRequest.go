package requests

type FindUsers struct {
	ID        uint   `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Age       int    `json:"age,omitempty"`
	Gender    string `json:"gender,omitempty"`
	Address   string `json:"address,omitempty"`
	Website   string `json:"wensite,omitempty"`
}
