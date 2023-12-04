package groups

import (
	"github.com/gin-gonic/gin"
	"kogalym-backend/helpers"
	"net/http"
	"strconv"
)

type Group struct {
	Id   int
	Name string
}

func IndexGroups(c *gin.Context) {
	c.HTML(http.StatusOK, "groups.index.html", gin.H{})
}

func GetGroups(c *gin.Context) {
	groups := getAll()

	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

type UpdateGroupRequest struct {
	Name string `json:"Name" binding:"required"`
}

func UpdateGroup(c *gin.Context) {
	var request UpdateGroupRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.WebValidateError(c, err)
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	helpers.CheckErr(err)

	result := update(idInt, request.Name)

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

	group := create(request.Name)

	c.JSON(http.StatusOK, gin.H{
		"data": group,
	})
}
