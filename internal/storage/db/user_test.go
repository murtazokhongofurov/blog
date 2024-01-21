package db

import (
	"testing"

	"github.com/blog/internal/controller/http/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createUser(t)
}

func createUser(t *testing.T) *models.UserRes {
	faker := gofakeit.New(0)

	user := models.UserReq{
		FullName: faker.Person().FirstName + faker.Person().LastName,
		Login:    faker.Name(),
		Password: faker.Password(false, true, true, false, false, 10),
	}
	userInfo, err := testMain.UserPost(&user)
	require.NoError(t, err)
	require.NotEmpty(t, userInfo)
	require.Equal(t, user.FullName, userInfo.FullName)
	require.Equal(t, user.Login, userInfo.Login)
	require.Equal(t, user.Password, userInfo.Password)

	require.NotZero(t, userInfo.Id)

	return userInfo
}

func TestUserGet(t *testing.T) {
	user1 := createUser(t)
	user2, err := testMain.UserGet(user1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Id, user2.Id)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Login, user2.Login)
	require.Equal(t, user1.Password, user2.Password)
}

func TestUserUpdate(t *testing.T) {
	user := createUser(t)
	faker := gofakeit.New(0)
	data := models.UserRes{
		Id:       user.Id,
		FullName: faker.Person().FirstName + faker.Person().LastName,
		Login:    faker.Username(),
		Password: faker.Password(false, true, true, false, false, 10),
		Photo:    faker.ImageURL(500, 600),
	}
	err := testMain.UserPut(&data)
	require.NoError(t, err)
}

func TestUserDelete(t *testing.T) {
	user := createUser(t)
	err := testMain.UserDel(user.Id)
	require.NoError(t, err)
}
