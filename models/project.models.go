package models

type Projects struct {
	Id           int         `form:"id"`
	ProjectName  string      `form:"projectName"`
	StartDate    string      `form:"startDate"`
	EndDate      string      `form:"endDate"`
	Duration     string      `form:"duration"`
	Description  string      `form:"description"`
	Technologies []string	 `form:"technologies"`
	Image        string      `form:"image"`
}

