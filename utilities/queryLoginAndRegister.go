package utilities

import (
	"context"
	"fmt"
	"os"
	"stage1/configs"
	"stage1/models"
)

func FindUsers() ([]models.User, error) {
	query := "SELECT * FROM tb_projects;"
	data, err := configs.Conn.Query(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	var result []models.User
	var eachError error
	for data.Next() {

		var each = models.User{}
		eachError = data.Scan(&each.Id, &each.Username, &each.Email, &each.Password)
		result = append(result, each)
	}

	return result, eachError

}

func FindOneUser(email string) (models.User, error) {

	data := configs.Conn.QueryRow(context.Background(), "SELECT * FROM tb_users WHERE email = $1", email)

	var result models.User

	err := data.Scan(&result.Id, &result.Username, &result.Email, &result.Password)

	return result, err
}

func InsertUser(username string, email string, password string) error {
	_, err := configs.Conn.Exec(context.Background(), "INSERT INTO tb_users (username, email, password) VALUES($1, $2, $3);", username, email, password)
	return err
}
