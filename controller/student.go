package controller

// GetStudents
//func GetStudents(c echo.Context) error {
//	students, _ := GetRepoStudents()
//	return c.JSON(http.StatusOK, students)
//}
//
//func GetRepoStudents() ([]model.Students, error) {
//	db := storage.GetDBInstance()
//	var students []model.Students
//
//	if err := db.Find(&students).Error; err != nil {
//		return nil, err
//	}
//
//	return students, nil
//}
