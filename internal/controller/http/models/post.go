package models

type PostReq struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	PostImage string `json:"post_image"`
}

type PostRes struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	PostImage string `json:"post_image"`
	CreatedAt string `json:"created_at"`
}
