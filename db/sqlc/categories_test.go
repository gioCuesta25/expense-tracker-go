package db

import (
	"context"
	"testing"

	"github.com/gioSmith25/expense-tracker/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Name:         utils.RandomString(5),
		IsForIncomes: false,
	}

	newCategory, err := testQueries.CreateCategory(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, newCategory)
	require.Equal(t, arg.Name, newCategory.Name)
	require.Equal(t, arg.IsForIncomes, newCategory.IsForIncomes)

	return newCategory

}

func TestCreateCategory(t *testing.T) {
	CreateRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	arg := CreateRandomCategory(t)

	category, err := testQueries.GetCategory(context.Background(), arg.ID)

	require.NoError(t, err)

	require.NotEmpty(t, category)
	require.Equal(t, arg.ID, category.ID)
	require.Equal(t, arg.CreatedAt, category.CreatedAt)
	require.Equal(t, arg.Name, category.Name)
	require.Equal(t, arg.IsForIncomes, category.IsForIncomes)
}
