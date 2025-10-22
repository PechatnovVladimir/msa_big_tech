-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists chat_members (
    chat_id text not null,
    user_id text not null,
    unique (chat_id,user_id)

);
comment on table chat_members is 'Таблица владельцев чатов';
comment on column chat_members.chat_id is 'Уникальный идентификатор чата';
comment on column chat_members.user_id is 'Уникальный идентификатор пользователя';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists chat_members;
-- +goose StatementEnd
