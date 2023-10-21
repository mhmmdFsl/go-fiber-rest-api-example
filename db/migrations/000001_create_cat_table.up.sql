CREATE TABLE if not exists cats(
    id SERIAL primary key,
    name varchar(255) not null ,
    breed varchar(255) not null ,
    color varchar(255) not null ,
    birth_at timestamptz default current_timestamp,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp
);