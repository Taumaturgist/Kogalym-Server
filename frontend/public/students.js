import axios from "axios";

window.students = () => {
    return {
        students: [],
        groups: [],
        editStudent: {Id: null, Name: ''},
        createStudent: {Name: '', GroupId: null},
        selectedGroup: null,
        showUpdateModal: false,
        showCreateModal: false,
        error: '',

        loadData: function () {
            this.loadGroups();

            fetch('/api/students')
                .then(response => response.json())
                .then(data => {
                    if (data && Array.isArray(data.data)) {
                        this.students = data.data;
                    }
                });
        },

        loadGroups: function () {
            fetch('/api/groups')
                .then(response => response.json())
                .then(data => {
                    if (data && Array.isArray(data.data)) {
                        this.groups = data.data;
                    }

                    if (!this.createStudent.GroupId && this.groups.length > 0) {
                        this.createStudent.GroupId = this.groups[0].Id;
                    }
                });
        },

        getGroupName: function (groupId) {
            const group = this.groups.find(group => group.Id === groupId);
            return group.Name;
        },

        openEditModal: function (student) {
            this.editStudent = {...student};
            this.showUpdateModal = true;
        },

        openCreateModal: function () {
            this.showCreateModal = true;
        },

        create: function () {
            axios.post('/api/students', {
                Name: this.createStudent.Name,
                GroupId: parseInt(this.createStudent.GroupId),
            })
                .then(response => {
                    this.students.push(response.data.data);
                    this.showCreateModal = false;
                })
                .catch(error => {
                    if (error.response?.data?.errors) {
                        this.error = Object.values(error.response.data.errors);
                    }
                });
        },

        update: function () {
            axios.put('/api/students/' + this.editStudent.Id, {
                Name: this.editStudent.Name,
                GroupId: this.editStudent.GroupId,
            }, {
                headers: {
                    'Content-Type': 'application/json',
                }
            })
                .then(response => {
                    const updatedStudent = response.data.data;
                    const index = this.students.findIndex(student => student.Id === updatedStudent.Id);
                    if (index !== -1) {
                        this.students[index] = updatedStudent;
                    }
                    this.showUpdateModal = false;
                })
                .catch(error => {
                    if (error.response && error.response.data && error.response.data.errors) {
                        this.error = Object.values(error.response.data.errors)
                            .flatMap(messages => Array.isArray(messages) ? messages : [messages]);
                    }
                });
        },
    };
};