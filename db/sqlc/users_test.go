package db

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/gioSmith25/expense-tracker/utils"
	"github.com/stretchr/testify/require"
)

type user struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type RandomUsersFile struct {
	Users []user
}

func TestCreateUsersBulk(t *testing.T) {
	file, err := os.Open("../../random_users.json")

	require.NoError(t, err)

	defer file.Close()

	decoder := json.NewDecoder(file)
	var data RandomUsersFile
	err = decoder.Decode(&data)

	require.NoError(t, err)

	for _, user := range data.Users {
		hashedPassword, err := utils.HashPassword(user.Password)

		require.NoError(t, err)
		arg := CreateUserParams{HashedPassword: hashedPassword, Email: user.Email, FullName: user.FullName}
		_, err = testQueries.CreateUser(context.Background(), arg)
		require.NoError(t, err)
	}
}
