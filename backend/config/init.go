package config

import (
	"Kogalym/backend/app/controller"
	"Kogalym/backend/app/repository"
	"Kogalym/backend/app/service"
)

type Initialization struct {
	authSvc  service.AuthService
	AuthCtrl controller.AuthController

	groupRepo repository.GroupRepository
	groupSvc  service.GroupService
	GroupCtrl controller.GroupController
}

func NewInitialization(
	authCtrl controller.AuthController,
	authSvc service.AuthService,
	groupRepo repository.GroupRepository,
	groupService service.GroupService,
	groupCtrl controller.GroupController,
) *Initialization {
	return &Initialization{
		authSvc:   authSvc,
		AuthCtrl:  authCtrl,
		groupRepo: groupRepo,
		groupSvc:  groupService,
		GroupCtrl: groupCtrl,
	}
}
