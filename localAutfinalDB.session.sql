CREATE TABLE IF NOT EXISTS "schedule_before_after" (
    id serial not null unique,
    child_id int references "child"(id) on delete cascade not null,
    name text,
    favourite boolean default false,
    cards_count int default 0,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

CREATE TABLE IF NOT EXISTS "card_before_after" (
    id serial not null unique,
    schedule_id int references "schedule_before_after"(id) on delete cascade not null,
    name text,
    done boolean default false,
    imgUUID text,
    orderPlace int not null,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
);

create or replace function update_schedule_before_after_cards_count_up() returns trigger as $update_schedule_before_after_cards_count_up$
begin
    update schedule_before_after set cards_count = (cards_count + 1) where id = new.schedule_id;
    return new;
end;
$update_schedule_before_after_cards_count_up$ language plpgsql;

drop trigger if exists create_card_before_after ON card_lesson;
create trigger create_card_before_after after insert on card_before_after for each row execute procedure update_schedule_before_after_cards_count_up();

create or replace function update_schedule_before_after_cards_count_down() returns trigger as $update_schedule_before_after_cards_count_down$
begin
    update schedule_before_after set cards_count = (cards_count - 1) where id = new.schedule_id and new.deletedAt is not null;
    return new;
end;
$update_schedule_before_after_cards_count_down$ language plpgsql;

drop trigger if exists delete_card_before_after ON card_lesson;
create trigger delete_card_before_after after update on card_before_after for each row execute procedure update_schedule_before_after_cards_count_down();