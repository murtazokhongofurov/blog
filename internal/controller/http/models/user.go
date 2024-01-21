package models

type UserReq struct {
	FullName string `json:"full_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type UserRes struct {
	Id        int    `json:"id"`
	FullName  string `json:"full_name"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
}
