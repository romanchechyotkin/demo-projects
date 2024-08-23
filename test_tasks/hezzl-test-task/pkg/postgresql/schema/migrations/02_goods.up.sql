create table if not exists goods
(
    id serial primary key,
    project_id serial references projects (id),
    name text not null,
    description text not null,
    priority int not null check (priority >= 1),
    removed boolean default false,
    created_at timestamp with time zone default now()
);