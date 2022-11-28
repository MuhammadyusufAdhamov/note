create table if not exists users (
    id serial primary key not null ,
    first_name varchar(30),
    last_name varchar(30),
    phone_number varchar(15),
    email varchar,
    image_url varchar,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

create table if not exists notes (
    id serial primary key not null,
    user_id int,
    title varchar,
    description varchar,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);