package api

type User struct {
	ID        ID     `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}
