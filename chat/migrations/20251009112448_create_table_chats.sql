-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists chats (
    id text not null default gen_random_uuid(),
    created_at timestamp with time zone default now(),
    primary key(id)
);
comment on table chats is 'Таблица чатов';
comment on column chats.id is 'Уникальный идентификатор чата';
comment on column chats.created_at is 'Время создания чата'
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists chats;
-- +goose StatementEnd
