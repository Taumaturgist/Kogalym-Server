import axios from "axios";

window.groups = () => {
    return {
        groups: [],
        editGroup: {Id: null, Name: ''},
        createGroupName: null,
        showUpdateModal: false,
        showCreateModal: false,
        error: [],

        loadGroups: function () {
            fetch('/api/groups')
                .then(response => response.json())
                .then(data => {
                    if (data && Array.isArray(data.data)) {
                        this.groups = data.data;
                    } else {
                        console.error('Неверный формат ответа');
                    }
                })
                .catch(error => {
                    console.error('Ошибка при загрузке данных:', error);
                });
        },

        openEditModal: function (group) {
            this.editGroup = {...group};
            this.showUpdateModal = true;
        },

        openCreateModal: function () {
            this.showCreateModal = true;
        },

        createGroup: function () {
            axios.post('/api/groups', {
                Name: this.createGroupName,
            })
                .then(response => {
                    this.groups.push(response.data.data);
                    this.showCreateModal = false;
                })
                .catch(error => {
                    if (error.response?.data?.errors) {
                        this.error = Object.values(error.response.data.errors);
                    }
                });
        },

        saveChanges: function () {
            axios.put('/api/groups/' + this.editGroup.Id, {
                Name: this.editGroup.Name,
            }, {
                headers: {
                    'Content-Type': 'application/json',
                }
            })
                .then(response => {
                    const updatedGroup = response.data.data;
                    const index = this.groups.findIndex(group => group.Id === updatedGroup.Id);
                    if (index !== -1) {
                        this.groups[index] = updatedGroup;
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