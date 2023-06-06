package models

type Projects struct {
	Id int `json:"id"`
	ProjectName string `json:"projectName"`
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	Description string `json:"description"`
	Technologies interface{} `json:"technologies"`
	Image string `json:"image"`

}
