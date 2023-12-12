package groups

import (
	"Kogalym/backend/helpers"
)

type GroupRepository struct {
	DB *helpers.Database
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{DB: helpers.GetDB()}
}

func (groupRepo *GroupRepository) GetAllGroups() ([]Group, error) {
	var groups []Group
	if err := groupRepo.DB.DB.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func GetById(groupRepo *GroupRepository, id int) (*Group, error) {
	var result Group
	if err := groupRepo.DB.DB.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func create(name string) Group {
	return Group{
		ID:   10,
		Name: name,
	}
}

func update(groupRepo *GroupRepository, id int, name string) *Group {
	group, _ := GetById(groupRepo, id)

	group.Name = name
	groupRepo.DB.DB.Save(&group)

	return group
}
