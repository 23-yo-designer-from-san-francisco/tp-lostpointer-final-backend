create table if not exists "mentor" (
    id serial not null unique,
    name text not null,
    surname text not null,
    email text not null unique,
    password text not null,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

create table if not exists "template" (
    id serial not null unique,
    mentor_id int references "mentor"(id) on delete cascade not null,
    name text not null,
    imgUUID text not null unique
);

create table if not exists "personal_image" (
    id serial not null unique,
    mentor_id int references "mentor"(id) on delete cascade not null,
    imgUUID text not null unique,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

create table if not exists "child" (
    id serial not null unique,
    mentor_id int references "mentor"(id) on delete cascade not null,
    name text not null,
    date_of_birth date,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

create table if not exists "schedule_day" (
    id serial not null unique,
    child_id int references "child"(id) on delete cascade not null,
    name text not null,
    day date not null,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

CREATE TABLE IF NOT EXISTS "card_day" (
    id serial not null unique,
    schedule_id int references "schedule_day"(id) on delete cascade not null,
    name text,
    done boolean default false,
    imgUUID text,
    startTime time,
    endTime time,
    orderPlace int not null,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

create table if not exists "schedule_lesson" (
    id serial not null unique,
    child_id int references "child"(id) on delete cascade not null,
    name text not null,
    duration int,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

CREATE TABLE IF NOT EXISTS "card_lesson" (
    id serial not null unique,
    schedule_id int references "schedule_lesson"(id) on delete cascade not null,
    name text,
    done boolean default false,
    imgUUID text,
    duration int,
    orderPlace int not null,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);