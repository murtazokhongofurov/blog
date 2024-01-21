package db

import (
	"testing"

	"github.com/blog/internal/controller/http/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateComment(t *testing.T) {
	createComment(t)
}

func createComment(t *testing.T) *models.CommentRes {
	faker := gofakeit.New(0)
	user := createUser(t)
	post := createPost(t)
	comment := models.CommentReq{
		PostId: post.Id,
		UserId: user.Id,
		Text:   faker.Product().Description,
	}
	commentInfo, err := testMain.CommentPost(&comment)
	require.NoError(t, err)
	require.NotEmpty(t, commentInfo)
	require.Equal(t, comment.PostId, commentInfo.PostId)
	require.Equal(t, comment.Text, commentInfo.Text)
	require.Equal(t, comment.UserId, commentInfo.UserId)

	require.NotZero(t, commentInfo.Id)

	return commentInfo
}

func TestCommentGet(t *testing.T) {
	comment1 := createComment(t)
	comment2, err := testMain.CommentGet(comment1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)
	require.Equal(t, comment1.Id, comment2.Id)
	require.Equal(t, comment1.PostId, comment2.PostId)
	require.Equal(t, comment1.Text, comment2.Text)
	require.Equal(t, comment1.UserId, comment2.UserId)
}

func TestCommentUpdate(t *testing.T) {
	comment := createComment(t)
	faker := gofakeit.New(0)
	data := models.CommentRes{
		Id:   comment.Id,
		Text: faker.Product().Description,
	}
	err := testMain.CommentPut(&data)
	require.NoError(t, err)
}

func TestCommentDelete(t *testing.T) {
	comment := createComment(t)
	err := testMain.CommentDel(comment.Id)
	require.NoError(t, err)
}
