create table "user"
(
    id uuid not null,
    full_name varchar(100) not null,
    username varchar(50) not null,
    email varchar(50) not null,
    password varchar(255) not null,
    phone_number numeric,
    birth_of_date date,
    balance float default 0.0,
    profile_picture text,
    profile_desc varchar(140),
    token text,
    refresh_token text,
    type int2 not null
);

create unique index user_email_uindex
    on "user" (email);

create unique index user_id_uindex
    on "user" (id);

create unique index user_phone_number_uindex
    on "user" (phone_number);

create unique index user_username_uindex
    on "user" (username);

create unique index user_refresh_token_uindex
    on "user" (refresh_token);

create unique index user_token_uindex
    on "user" (token);

alter table "user"
    add constraint user_pk
        primary key (id);
