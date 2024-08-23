create table if not exists projects
(
    id serial primary key,
    name text not null,
    created_at timestamp with time zone default now()
);