create table graph_data
(
    id          integer                            not null
        constraint graph_data_pk
            primary key autoincrement,
    filename    TEXT                               not null,
    graph_name  TEXT                               not null,
    description TEXT,
    unit        TEXT,
    data        TEXT                               not null,
    created_at  DATETIME default CURRENT_TIMESTAMP not null,
    constraint graph_data_filename_graph_name_unique
        unique (filename, graph_name)
);

