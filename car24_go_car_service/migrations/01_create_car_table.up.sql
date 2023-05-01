create table if not exists brand  (
    "id" uuid primary key,
    "name" varchar not null,
    "created_at" timestamp default current_timestamp not null,
    "updated_at" timestamp default current_timestamp not null
);

create table if not exists mark  (
    "id" uuid primary key,
    "brand_id" uuid references brand(id),
    "name" varchar not null,
    "created_at" timestamp default current_timestamp not null,
    "updated_at" timestamp default current_timestamp not null
);


create table if not exists car (
    "id" uuid PRIMARY KEY,
    "mark_id" uuid references mark(id),
    "category_id" uuid,
    "investor_id" uuid,
    "state_number" varchar,
    "updated_at" timestamp default current_timestamp not null,
    "created_at" timestamp default current_timestamp not null
);