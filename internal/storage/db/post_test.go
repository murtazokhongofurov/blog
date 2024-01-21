package db

import (
	"testing"

	"github.com/blog/internal/controller/http/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	createPost(t)
}

func createPost(t *testing.T) *models.PostRes {
	faker := gofakeit.New(0)

	post := models.PostReq{
		Title:     faker.Book().Title,
		Text:      faker.Product().Description,
		PostImage: faker.ImageURL(500, 500),
	}
	postInfo, err := testMain.PostCreate(&post)
	require.NoError(t, err)
	require.NotEmpty(t, postInfo)
	require.Equal(t, post.Title, postInfo.Title)
	require.Equal(t, post.Text, postInfo.Text)
	require.Equal(t, post.PostImage, postInfo.PostImage)

	require.NotZero(t, postInfo.Id)

	return postInfo
}

func TestPostGet(t *testing.T) {
	post1 := createPost(t)
	post2, err := testMain.PostGet(post1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, post2)
	require.Equal(t, post1.Id, post2.Id)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Text, post2.Text)
	require.Equal(t, post1.PostImage, post2.PostImage)
}

func TestPostUpdate(t *testing.T) {
	post := createPost(t)
	faker := gofakeit.New(0)
	data := models.PostRes{
		Id:        post.Id,
		Title:     faker.Book().Title,
		Text:      faker.Product().Description,
		PostImage: faker.ImageURL(500, 500),
	}
	err := testMain.PostPut(&data)
	require.NoError(t, err)
}

func TestPostDelete(t *testing.T) {
	post := createPost(t)
	err := testMain.PostDel(post.Id)
	require.NoError(t, err)
}
