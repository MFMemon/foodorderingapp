create table users(
    id int generated always as identity primary key,
    name text,
    email text unique,
    password text
);

create index on users(email);