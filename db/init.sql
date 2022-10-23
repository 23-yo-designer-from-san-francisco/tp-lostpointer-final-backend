\c postgres
DROP DATABASE IF EXISTS autfinal;
CREATE DATABASE autfinal;
\c autfinal

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
    favourite boolean default false,
    cards_count int default 0,
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
    favourite boolean default false,
    cards_count int default 0,
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

CREATE TABLE IF NOT EXISTS "stock_image" (
    id serial not null unique,
    category text not null,
    imgUUID text not null,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

CREATE TABLE IF NOT EXISTS "stock_image_name" (
    id serial not null unique,
    image_id int references "stock_image"(id) on delete cascade not null,
    name text,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

--alter table schedule_day add column favourite boolean default false;
--alter table schedule_lesson add column favourite boolean default false;
--alter table schedule_lesson add cards_count int default 0;
--alter table schedule_day add cards_count int default 0;

create or replace function update_schedule_lesson_cards_count_up() returns trigger as $update_schedule_lesson_cards_count_up$
begin
    update schedule_lesson set cards_count = (cards_count + 1) where id = new.schedule_id;
    return new;
end;
$update_schedule_lesson_cards_count_up$ language plpgsql;

drop trigger if exists create_card_lesson ON card_lesson;
create trigger create_card_lesson after insert on card_lesson for each row execute procedure update_schedule_lesson_cards_count_up();

create or replace function update_schedule_lesson_cards_count_down() returns trigger as $update_schedule_lesson_cards_count_down$
begin
    update schedule_lesson set cards_count = (cards_count - 1) where id = new.schedule_id and new.deletedAt is not null;
    return new;
end;
$update_schedule_lesson_cards_count_down$ language plpgsql;

drop trigger if exists delete_card_lesson ON card_lesson;
create trigger delete_card_lesson after update on card_lesson for each row execute procedure update_schedule_lesson_cards_count_down();

create or replace function update_schedule_day_cards_count_up() returns trigger as $update_schedule_day_cards_count_up$
begin
    update schedule_day set cards_count = (cards_count + 1) where id = new.schedule_id;
    return new;
end;
$update_schedule_day_cards_count_up$ language plpgsql;

drop trigger if exists create_card_day ON card_lesson;
create trigger create_card_day after insert on card_day for each row execute procedure update_schedule_day_cards_count_up();

create or replace function update_schedule_day_cards_count_down() returns trigger as $update_schedule_day_cards_count_down$
begin
    update schedule_day set cards_count = (cards_count - 1) where id = new.schedule_id and new.deletedAt is not null;
    return new;
end;
$update_schedule_day_cards_count_down$ language plpgsql;

drop trigger if exists delete_card_day ON card_lesson;
create trigger delete_card_day after update on card_day for each row execute procedure update_schedule_day_cards_count_down();

create extension pg_trgm;
create index stock_image_name_idx on stock_image_name using gin (name gin_trgm_ops);