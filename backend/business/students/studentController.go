package students

import (
	"Kogalym/backend/helpers"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"strconv"
)

func IndexStudents(c *gin.Context) {
	c.HTML(http.StatusOK, "students.index.html", gin.H{
		"csrf": csrf.GetToken(c),
	})
}

func GetStudents(c *gin.Context) {
	students := getAll()

	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

type UpdateStudentRequest struct {
	Name    string `json:"Name" binding:"required"`
	GroupId int    `json:"groupId" binding:"required"`
}

func UpdateStudent(c *gin.Context) {
	var request UpdateStudentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		//helpers.ValidateError(c, err)
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	helpers.CheckErr(err)

	result := update(idInt, request.Name, request.GroupId)

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

type CreateStudentRequest struct {
	Name    string `json:"Name" binding:"required"`
	GroupId int    `json:"GroupId" binding:"required"`
}

func CreateStudent(c *gin.Context) {
	var request CreateStudentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		//helpers.WebValidateError(c, err)
		return
	}

	student := create(request.Name, request.GroupId)

	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}
