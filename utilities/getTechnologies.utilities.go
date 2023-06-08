package utilities

func GetTechnologies(reactjs string, nextjs string, nodejs string, typescript string) []string {

	var checkboxData []string

	if reactjs != "" {
		checkboxData = append(checkboxData, "reactjs")
	}
	if nextjs != "" {
		checkboxData = append(checkboxData, "nextjs")
	}
	if nodejs != "" {
		checkboxData = append(checkboxData, "nodejs")		
	}
	if typescript != "" {
		checkboxData = append(checkboxData, "typescript")
	}

	return checkboxData
}


func GetTechnologiesValue(technologiesData []string) []string {

	var checkboxValue []string

	for _, data := range technologiesData {
		
		if data == "reactjs" {
			checkboxValue = append(checkboxValue, "reactjs")
		}
		if data == "nextjs"{
			checkboxValue = append(checkboxValue, "nextjs")
		}
		if data == "nodejs" {
			checkboxValue = append(checkboxValue, "nodejs")		
		}
		if data == "typescript" {
			checkboxValue = append(checkboxValue, "typescript")
		}
	}
	

	return checkboxValue
}

func GetTechnologiesChecked(t []string, s string) bool {
	isChecked := false

	for _, data := range t {
		if data == s {
			isChecked = true
		}
	}
	return isChecked
}