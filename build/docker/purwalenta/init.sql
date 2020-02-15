create table public.user
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
    status smallint default 2 not null,
    type smallint default 3 not null
);

alter table public.user owner to purwalenta;

create unique index user_email_uindex
    on public.user (email);

create unique index user_id_uindex
    on public.user (id);

create unique index user_username_uindex
    on public.user (username);

create unique index user_phone_number_uindex
    on public.user (phone_number);
