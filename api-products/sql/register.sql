
drop table permissions;

/*
I WILL WORK ON HOW TO ASSIGN GROUP AND PERMISSIONS

create table if not exists groups (
    id bigserial primary key,
    name varchar(150) not null,
    permission_id bigint not null,
    removed_at varchar(150) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint groups_permission_id foreign key(permission_id) references permissions(id)
        on delete cascade on update cascade
);
*/

create table if not exists permissions (
    id bigserial primary key,
    role_name varchar(150) not null,
    removed_status boolean default false,
    removed_at varchar(150) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp

);

create table if not exists users (
    id bigserial primary key,
    first_name varchar(150) not null,
    last_name varchar(150) not null,
    email varchar(150) not null unique,
    password varchar(250) not null,
    telephone varchar(150) not null unique,
    token varchar(250) not null,
    active boolean default false,
    verify boolean default false,
    permission_id bigint not null,
    last_login varchar(150) not null,
    removed_at varchar(150) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint users_permission_id foreign key(permission_id) references permissions(id)
        on delete cascade on update cascade
);

create table if not exists adminusers (
    id bigserial primary key,
    email varchar(150) not null unique,
    password varchar(250) not null
    phone varchar(150) not null unique,
    active boolean default false,
    group_id bigint not null,
    permission_id bigint not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint adminusers_group_id foreign key(group_id) references groups(id)
        on delete cascade on update cascade,
    constraint adminusers_permission_id foreign key(permission_id) references permissions(id)
        on delete cascade on update cascade
);


create table if not exists products(
    id bigint primary key auto_increment,
    name varchar(512) not null unique,
    price decimal(10,2) default 0.0,
    quantity bigint default 0,
    status char(1) default 0,
    category_id bigint not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp,
    constraint products_category_id foreign key(category_id) references categories(id)
    on delete cascade on update cascade
)

