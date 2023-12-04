package groups

var Groups = []Group{
	{Id: 1, Name: `Группа 1`},
	{Id: 2, Name: `Группа 2`},
}

func getAll() []Group {
	return Groups
}

func GetById(id int) Group {
	var result Group

	for _, group := range Groups {
		if group.Id == id {
			result = group
			break
		}
	}

	return result
}

func create(name string) Group {
	return Group{
		Id:   10,
		Name: name,
	}
}

func update(id int, name string) Group {
	group := GetById(id)

	group.Name = name

	return group
}
