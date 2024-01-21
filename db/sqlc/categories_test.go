package db

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/gioSmith25/expense-tracker/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Name:         utils.RandomString(5),
		IsForIncomes: utils.RadomBool(),
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

func TestListCategories(t *testing.T) {

}

func TestUpdateCategory(t *testing.T) {
	category := CreateRandomCategory(t)

	arg := UpdateCategoryParams{
		ID:           category.ID,
		Name:         utils.RandomString(10),
		IsForIncomes: utils.RadomBool(),
	}

	err := testQueries.UpdateCategory(context.Background(), arg)

	require.NoError(t, err)

	updatedCategory, err := testQueries.GetCategory(context.Background(), category.ID)

	require.Equal(t, category.ID, updatedCategory.ID)
	require.Equal(t, arg.Name, updatedCategory.Name)
	require.Equal(t, arg.IsForIncomes, updatedCategory.IsForIncomes)
	require.Equal(t, category.CreatedAt, updatedCategory.CreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category := CreateRandomCategory(t)

	err := testQueries.DeleteCategory(context.Background(), category.ID)

	arg, err := testQueries.GetCategory(context.Background(), category.ID)

	require.Error(t, err)
	require.EqualError(t, err, "sql: no rows in result set")
	require.Empty(t, arg)
}

type CategoriesFile struct {
	Categories []CreateCategoryParams
}

func TestCreateCategoriesBulk(t *testing.T) {
	file, err := os.Open("../../categories.json")

	require.NoError(t, err)

	defer file.Close()

	decoder := json.NewDecoder(file)
	var data CategoriesFile
	err = decoder.Decode(&data)

	for _, category := range data.Categories {
		arg := CreateCategoryParams{
			Name:         category.Name,
			IsForIncomes: category.IsForIncomes,
		}
		newCategory, err := testQueries.CreateCategory(context.Background(), arg)
		require.NoError(t, err)

		require.NotEmpty(t, newCategory)
		require.Equal(t, arg.Name, newCategory.Name)
		require.Equal(t, arg.IsForIncomes, newCategory.IsForIncomes)
	}
}
