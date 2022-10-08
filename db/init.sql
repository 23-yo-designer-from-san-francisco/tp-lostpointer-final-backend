CREATE TABLE IF NOT EXISTS "card" (
    id serial not null unique,
    name text not null,
    done boolean default false,
    imgUrl text,
    startTime time,
    endTime time,
    createdAt timestamp default now() not null,
    updatedAt timestamp, 
    deletedAt timestamp
)