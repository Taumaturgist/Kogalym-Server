package repository

import (
	"Kogalym/backend/app/domain/dao"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GroupRepository interface {
	FindAllGroup() ([]dao.Group, error)
	FindGroupById(id int) (dao.Group, error)
	Save(group *dao.Group) (dao.Group, error)
	DeleteGroupById(id int) error
}

type GroupRepositoryImpl struct {
	db *gorm.DB
}

func (g GroupRepositoryImpl) FindAllGroup() ([]dao.Group, error) {
	var groups []dao.Group

	var err = g.db.Preload("Role").Find(&groups).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return groups, nil
}

func (g GroupRepositoryImpl) FindGroupById(id int) (dao.Group, error) {
	group := dao.Group{
		ID: id,
	}
	err := g.db.Preload("Role").First(&group).Error
	if err != nil {
		log.Error("Got and error when find group by id. Error: ", err)
		return dao.Group{}, err
	}
	return group, nil
}

func (g GroupRepositoryImpl) Save(group *dao.Group) (dao.Group, error) {
	var err = g.db.Save(group).Error
	if err != nil {
		log.Error("Got an error when save group. Error: ", err)
		return dao.Group{}, err
	}
	return *group, nil
}

func (g GroupRepositoryImpl) DeleteGroupById(id int) error {
	err := g.db.Delete(&dao.Group{}, id).Error
	if err != nil {
		log.Error("Got an error when delete group. Error: ", err)
		return err
	}
	return nil
}

func GroupRepositoryInit(db *gorm.DB) *GroupRepositoryImpl {
	db.AutoMigrate(&dao.Group{})
	return &GroupRepositoryImpl{
		db: db,
	}
}
