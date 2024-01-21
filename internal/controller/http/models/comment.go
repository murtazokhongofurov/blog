package models

type CommentReq struct {
	PostId int    `json:"post_id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
}

type CommentRes struct {
	Id        int    `json:"id"`
	PostId    int    `json:"post_id"`
	UserId    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
