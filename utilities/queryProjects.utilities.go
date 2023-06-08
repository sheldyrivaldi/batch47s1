package utilities

import (
	"context"
	"fmt"
	"os"
	"stage1/configs"
	"stage1/models"
	"strings"
)

func FindProjects() interface{} {
	query := "SELECT * FROM tb_projects"
	data, err := configs.Conn.Query(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		}
		var result []models.Projects
		for data.Next(){
			
			var each = models.Projects{}
			err = data.Scan(&each.Id, &each.ProjectName, &each.StartDateTime, &each.EndDateTime, &each.Description, &each.Technologies, &each.Image)
			if err != nil {
				fmt.Println(err.Error())
			}
			
			each.StartDate = each.StartDateTime.Format("2006-01-02")
			each.EndDate = each.EndDateTime.Format("2006-01-02")
			result = append(result, each)
		}
	return result
	
}

func FindOneProject (id string) interface{} {
	query := "SELECT * FROM tb_projects WHERE id = " + id
	data, err := configs.Conn.Query(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		}
		var result []models.Projects
		for data.Next(){
			
			var each = models.Projects{}
			err = data.Scan(&each.Id, &each.ProjectName, &each.StartDateTime, &each.EndDateTime, &each.Description, &each.Technologies, &each.Image)
			if err != nil {
				fmt.Println(err.Error())
			}
			
			each.StartDate = each.StartDateTime.Format("2006-01-02")
			each.EndDate = each.EndDateTime.Format("2006-01-02")
			result = append(result, each)
		}
	return result
}

func InsertProject (ProjectName string, StartDate string, EndDate string, Description string, Technologies []string, Image string) error {
	for i,data:= range Technologies {
		Technologies[i] = fmt.Sprintf("'%s'", data)
	}

	technologyJoined := strings.Join(Technologies, ", ")
	
	query := fmt.Sprintf("INSERT INTO tb_projects (name, start_date, end_date, description, technologies, image) VALUES ('%s', '%s', '%s', '%s', ARRAY[%v], 'project-list1.png')", ProjectName, StartDate, EndDate, Description, technologyJoined)
	_, err := configs.Conn.Query(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		}
		fmt.Println("Berhasil menambahkan project!")

	return err
}

func DeleteProject (id string) error {
	query := fmt.Sprintf("DELETE FROM tb_projects WHERE id = %s", id)
	_, err := configs.Conn.Query(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		}
		fmt.Println("Berhasil menghapus project!")

	return err
}

func UpdateProject (id string, ProjectName string, StartDate string, EndDate string, Description string, Technologies []string, Image string) error {
	for i,data:= range Technologies {
		Technologies[i] = fmt.Sprintf("'%s'", data)
	}

	technologyJoined := strings.Join(Technologies, ", ")
	
	query := fmt.Sprintf("UPDATE tb_projects SET name = '%s', start_date = '%s', end_date = '%s', description = '%s', technologies = ARRAY[%v], image = 'project-list1.png' WHERE id = '%s'", ProjectName, StartDate, EndDate, Description, technologyJoined, id)
	_, err := configs.Conn.Query(context.Background(), query)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		}
		fmt.Println("Berhasil mengupdate project!")

	return err
}