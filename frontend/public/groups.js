import axios from "axios";

window.groups = () => {
    return {
        groups: [],
        editGroup: {Id: null, Name: ''},
        createGroupName: null,
        showUpdateModal: false,
        showCreateModal: false,
        error: '',

        loadGroups: function () {
            fetch('/api/groups')
                .then(response => response.json())
                .then(data => {
                    if (data && Array.isArray(data.data)) {
                        this.groups = data.data;
                        console.log(groups);
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
            console.log(this.showUpdateModal)
            this.showUpdateModal = true;
            console.log(this.showUpdateModal)
        },

        openCreateModal: function () {
            console.log(this.showCreateModal)
            this.showCreateModal = true;
            console.log(this.showCreateModal)
        },

        createGroup: function () {
            axios.post('/api/groups', {
                Name: this.createGroupName,
            })
                .then(response => {
                    // Обработка успешного ответа
                    this.groups.push(response.data.data);

                    // Закрываем модальное окно
                    this.showCreateModal = false;
                })
                .catch(error => {
                    // Обработка ошибок
                    if (error.response?.data?.errors) {
                        const errorMessages = Object.values(error.response.data.errors);
                        this.error = errorMessages;
                        console.log(errorMessages);
                    } else {
                        console.log(error);
                    }
                });
        },

        saveChanges: function () {
            // Отправляем запрос на обновление данных (пример с fetch)
            fetch('/api/groups/' + this.editGroup.Id, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    Name: this.editGroup.Name,
                    // Другие поля для обновления, если необходимо
                }),
            })
                .then(response => response.json())
                .then(data => data.data)
                .then(updatedGroup => {
                    // Обновляем данные в массиве groups
                    const index = this.groups.findIndex(group => group.Id === updatedGroup.Id);
                    if (index !== -1) {
                        this.groups[index] = updatedGroup;
                    }

                    // Закрываем модальное окно
                    this.showUpdateModal = false;
                })
                .then(data => {
                    if (data && data.errors) {
                        console.log(data.errors)
                        this.error = data.errors;
                    }
                })
                .catch(error => {
                    console.error('Ошибка при обновлении данных:', error);
                });
        },
    };
};