package service

import (
	"Kogalym/backend/app/constant"
	"Kogalym/backend/app/domain/dao"
	"Kogalym/backend/app/pkg"
	"Kogalym/backend/app/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GroupService interface {
	GetAllGroup(c *gin.Context)
	GetGroupById(c *gin.Context)
	AddGroupData(c *gin.Context)
	UpdateGroupData(c *gin.Context)
}

type GroupServiceImpl struct {
	groupRepository repository.GroupRepository
}

func (u GroupServiceImpl) UpdateGroupData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update group data by id")
	groupID, _ := strconv.Atoi(c.Param("groupID"))

	var request dao.Group
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := u.groupRepository.FindGroupById(groupID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	//data.RoleID = request.RoleID
	//data.Email = request.Email
	//data.Name = request.Password
	//data.Status = request.Status
	u.groupRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u GroupServiceImpl) GetGroupById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get group by id")
	groupID, _ := strconv.Atoi(c.Param("groupID"))

	data, err := u.groupRepository.FindGroupById(groupID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u GroupServiceImpl) AddGroupData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data group")
	var request dao.Group
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	//hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	//request.Password = string(hash)

	data, err := u.groupRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (u GroupServiceImpl) GetAllGroup(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data group")

	data, err := u.groupRepository.FindAllGroup()
	if err != nil {
		log.Error("Happened Error when find all group data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func GroupServiceInit(groupRepository repository.GroupRepository) *GroupServiceImpl {
	return &GroupServiceImpl{
		groupRepository: groupRepository,
	}
}
