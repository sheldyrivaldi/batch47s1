package utilities

import (
	"context"
	"fmt"
	"stage1/configs"
	"stage1/models"
	"strings"
)

func FindProjects() ([]models.Projects, error) {
	data, err := configs.Conn.Query(context.Background(), "SELECT * FROM tb_projects ORDER BY id;")
	var result []models.Projects
	for data.Next() {

		var each = models.Projects{}
		eachError := data.Scan(&each.Id, &each.ProjectName, &each.StartDateTime, &each.EndDateTime, &each.Description, &each.Technologies, &each.Image, &each.UserId)
		if eachError != nil {
			fmt.Println(err.Error())
		}
		each.StartDate = each.StartDateTime.Format("2006-01-02")
		each.EndDate = each.EndDateTime.Format("2006-01-02")
		each.Duration = GetDuration(each.StartDate, each.EndDate)
		result = append(result, each)
	}

	return result, err
}
func FindProjectsWithUserId(userId int) ([]models.Projects, error) {
	data, err := configs.Conn.Query(context.Background(), "SELECT * FROM tb_projects WHERE user_id = $1 ORDER BY id;", userId)
	var result []models.Projects
	for data.Next() {

		var each = models.Projects{}
		eachError := data.Scan(&each.Id, &each.ProjectName, &each.StartDateTime, &each.EndDateTime, &each.Description, &each.Technologies, &each.Image, &each.UserId)
		if eachError != nil {
			fmt.Println(err.Error())
		}
		each.StartDate = each.StartDateTime.Format("2006-01-02")
		each.EndDate = each.EndDateTime.Format("2006-01-02")
		each.Duration = GetDuration(each.StartDate, each.EndDate)
		result = append(result, each)
	}

	return result, err
}

func FindOneProject(id int) (models.Projects, error) {

	data := configs.Conn.QueryRow(context.Background(), "SELECT * FROM tb_projects WHERE id = $1;", id)

	var result models.Projects

	err := data.Scan(&result.Id, &result.ProjectName, &result.StartDateTime, &result.EndDateTime, &result.Description, &result.Technologies, &result.Image, &result.UserId)

	result.StartDate = result.StartDateTime.Format("2006-01-02")
	result.EndDate = result.EndDateTime.Format("2006-01-02")
	result.Duration = GetDuration(result.StartDate, result.EndDate)

	return result, err
}

func InsertProject(ProjectName string, StartDate string, EndDate string, Description string, Technologies []string, Image string, UserId int) error {
	for i, data := range Technologies {
		Technologies[i] = fmt.Sprintf("'%s'", data)
	}

	technologyJoined := strings.Join(Technologies, ", ")

	query := fmt.Sprintf("INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image, user_id) VALUES ('%s', '%s', '%s', '%s', ARRAY[%v], '%s', %d)", ProjectName, StartDate, EndDate, Description, technologyJoined, Image, UserId)
	_, err := configs.Conn.Exec(context.Background(), query)

	return err
}

func DeleteProject(id int) error {
	_, err := configs.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id = $1", id)

	return err
}

func UpdateProject(id int, ProjectName string, StartDate string, EndDate string, Description string, Technologies []string, Image string) error {

	_, err := configs.Conn.Exec(context.Background(), "UPDATE tb_projects SET name = $1, start_date = $2, end_date = $3, description = $4, technologies = $5, image = $6 WHERE id=$7", ProjectName, StartDate, EndDate, Description, Technologies, Image, id)

	return err
}
