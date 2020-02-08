create table "user"
(
    id uuid not null
        constraint user_pk
            primary key,
    full_name varchar(100) default ''::character varying,
    username varchar(50) not null,
    email varchar(50) not null,
    password varchar(255) not null,
    phone_number varchar(12) default ''::character varying,
    date_of_birth date,
    balance double precision default 0.0,
    profile_picture text default ''::text,
    profile_desc varchar(140) default ''::character varying,
    refresh_token text default ''::text,
    type smallint default 3 not null
);

alter table "user" owner to postgres;

create unique index user_email_uindex
    on "user" (email);

create unique index user_id_uindex
    on "user" (id);

create unique index user_username_uindex
    on "user" (username);

create unique index user_refresh_token_uindex
    on "user" (refresh_token);

create unique index user_phone_number_uindex
    on "user" (phone_number);
