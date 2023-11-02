create table clients
(
    id         integer                            not null
        constraint clients_pk primary key autoincrement,
    login      TEXT                               not null,
    password   TEXT                               not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null
);

create unique index clients_uindex
    on clients (login);

