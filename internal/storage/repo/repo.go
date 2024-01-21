package repo

import "github.com/blog/internal/controller/http/models"

type BlogService interface {
	UserPost(*models.UserReq) (*models.UserRes, error)
	UserGet(id int) (*models.UserRes, error)
	UserPut(*models.UserRes) error
	UserDel(id int) error

	PostCreate(*models.PostReq) (*models.PostRes, error)
	PostGet(id int) (*models.PostRes, error)
	PostPut(*models.PostRes) error
	PostDel(id int) error

	CommentPost(*models.CommentReq) (*models.CommentRes, error)
	CommentGet(id int) (*models.CommentRes, error)
	CommentPut(*models.CommentRes) error
	CommentDel(id int) error
}
