package business

import (
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"kogalym-backend/helpers"
	"net/http"
	"strconv"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"csrf":  csrf.GetToken(c),
		"title": "Добро пожаловать на главную страницу",
	})
}

type Group struct {
	Id   int
	Name string
}

func Groups(c *gin.Context) {
	Groups := [...]Group{
		{Id: 1, Name: `Группа 1`},
		{Id: 2, Name: `Группа 2`},
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Groups,
	})
}

type UpdateGroupRequest struct {
	Name string `json:"Name" binding:"required"`
}

func UpdateGroup(c *gin.Context) {
	var request UpdateGroupRequest
	var result Group

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ValidateError(c, err)
		return
	}
	name := request.Name
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	helpers.CheckErr(err)

	Groups := [...]Group{
		{Id: 1, Name: `Группа 1`},
		{Id: 2, Name: `Группа 2`},
	}

	for _, group := range Groups {
		if group.Id == idInt {
			result = group
			break
		}
	}

	result.Name = name

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

type CreateGroupRequest struct {
	Name string `json:"Name" binding:"required"`
}

func CreateGroup(c *gin.Context) {
	var request CreateGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.WebValidateError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": Group{
			Id:   10,
			Name: request.Name,
		},
	})
}

func WebGroups(c *gin.Context) {
	Groups := [...]Group{
		{Id: 1, Name: `Группа 1`},
		{Id: 2, Name: `Группа 2`},
	}
	c.HTML(http.StatusOK, "groups.index.html", gin.H{
		"csrf":   csrf.GetToken(c),
		"Groups": Groups,
	})
}
