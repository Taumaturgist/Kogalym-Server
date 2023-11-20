create table graph_data
(
    id          integer                            not null
        constraint graph_data_pk
            primary key autoincrement,
    mine_name   TEXT                               not null,
    graph_name  TEXT                               not null,
    description TEXT,
    start_depth REAL                               not null,
    stop_depth  REAL                               not null,
    step_depth  REAL                               not null,
    units       TEXT,
    data        TEXT                               not null,
    created_at  DATETIME default CURRENT_TIMESTAMP not null,
    constraint graph_data_filename_graph_name_unique
        unique (mine_name, graph_name)
);

