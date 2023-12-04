package students

var students = []Student{
	{
		Id:      1,
		Name:    `Студент 1`,
		GroupId: 1,
	},
	{
		Id:      2,
		Name:    `Студент 2`,
		GroupId: 2,
	},
}

type Student struct {
	Id      int
	Name    string
	GroupId int
}

func getAll() []Student {
	return students
}

func getById(id int) Student {
	var result Student

	for _, student := range students {
		if student.Id == id {
			result = student
			break
		}
	}

	return result
}

func create(name string, groupId int) Student {
	return Student{
		Id:      10,
		Name:    name,
		GroupId: groupId,
	}
}

func update(id int, name string, groupId int) Student {
	student := getById(id)

	student.Name = name
	student.GroupId = groupId

	return student
}
