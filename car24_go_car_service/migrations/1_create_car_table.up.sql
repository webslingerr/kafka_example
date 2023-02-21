create table if not exists car (
    "id" uuid PRIMARY KEY,
    "category_id" uuid,
    "investor_id" uuid,
    "state_number" varchar,
    "updated_at" timestamp default current_timestamp not null,
    "created_at" timestamp default current_timestamp not null
);