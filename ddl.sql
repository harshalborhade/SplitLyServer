create table user_credentials
(
    id         serial
        primary key,
    username   text not null
        unique,
    email      text not null
        unique,
    password   text not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);

create table user_dimension
(
    id                  integer      not null
        constraint user_dimension_pk
            primary key
        references user_credentials,
    first_name          varchar(200) not null,
    last_name           varchar(200),
    profile_picture_url varchar(200)
);