//package groups
//
//import (
//	"Kogalym/backend/helpers"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//)
//
//func IndexGroups(c *gin.Context) {
//	c.HTML(http.StatusOK, "groups.index.html", gin.H{})
//}
//
//func GetGroups(c *gin.Context) {
//	groupRepo := NewGroupRepository()
//	defer groupRepo.DB.CloseDB()
//
//	groups, _ := groupRepo.GetAllGroups()
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": groups,
//	})
//}
//
//type UpdateGroupRequest struct {
//	Name string `json:"Name" binding:"required"`
//}
//
//func UpdateGroup(c *gin.Context) {
//	var request UpdateGroupRequest
//
//	if err := c.ShouldBindJSON(&request); err != nil {
//		helpers.WebValidateError(c, err)
//		return
//	}
//
//	id := c.Param("id")
//	idInt, err := strconv.Atoi(id)
//	helpers.CheckErr(err)
//
//	groupRepo := NewGroupRepository()
//	defer groupRepo.DB.CloseDB()
//
//	result := update(groupRepo, idInt, request.Name)
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": result,
//	})
//}
//
//type CreateGroupRequest struct {
//	Name string `json:"Name" binding:"required"`
//}
//
//func CreateGroup(c *gin.Context) {
//	var request CreateGroupRequest
//
//	if err := c.ShouldBindJSON(&request); err != nil {
//		helpers.WebValidateError(c, err)
//		return
//	}
//
//	group := create(request.Name)
//
//	c.JSON(http.StatusOK, gin.H{
//		"data": group,
//	})
//}
