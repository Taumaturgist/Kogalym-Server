package controller

import (
	"Kogalym/backend/app/service"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
)

type GroupController interface {
	GetAllGroupData(c *gin.Context)
	AddGroupData(c *gin.Context)
	GetGroupById(c *gin.Context)
	UpdateGroupData(c *gin.Context)
	IndexGroups(c *gin.Context)
}

type GroupControllerImpl struct {
	svc service.GroupService
}

func (u GroupControllerImpl) GetAllGroupData(c *gin.Context) {
	u.svc.GetAllGroup(c)
}

func (u GroupControllerImpl) AddGroupData(c *gin.Context) {
	u.svc.AddGroupData(c)
}

func (u GroupControllerImpl) GetGroupById(c *gin.Context) {
	u.svc.GetGroupById(c)
}

func (u GroupControllerImpl) UpdateGroupData(c *gin.Context) {
	u.svc.UpdateGroupData(c)
}

func (u GroupControllerImpl) IndexGroups(c *gin.Context) {
	c.HTML(http.StatusOK, "groups.index.html", gin.H{
		"csrf": csrf.GetToken(c),
	})
}

func GroupControllerInit(groupService service.GroupService) *GroupControllerImpl {
	return &GroupControllerImpl{
		svc: groupService,
	}
}
