// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"Kogalym/backend/app/controller"
	"Kogalym/backend/app/repository"
	"Kogalym/backend/app/service"
	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var authServiceSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
)

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var groupServiceSet = wire.NewSet(service.GroupServiceInit,
	wire.Bind(new(service.GroupService), new(*service.GroupServiceImpl)),
)

var groupRepoSet = wire.NewSet(repository.GroupRepositoryInit,
	wire.Bind(new(repository.GroupRepository), new(*repository.GroupRepositoryImpl)),
)

var groupCtrlSet = wire.NewSet(controller.GroupControllerInit,
	wire.Bind(new(controller.GroupController), new(*controller.GroupControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, authServiceSet, authCtrlSet, groupCtrlSet, groupServiceSet, groupRepoSet)
	return nil
}
