-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists messages (
  id text not null default gen_random_uuid(),
  chat_id text not null,
  sender_id text not null,
  text text,
  created_at timestamp with time zone not null default now() ,
  primary key (id)
);

comment on table messages is 'Таблица сообщений';
comment on column messages.id is 'Уникальный идентификатор сообщения';
comment on column messages.chat_id is 'Идентификатор чата';
comment on column messages.sender_id is 'Идентификатор отправителя';
comment on column messages.text is 'Текст сообщения';
comment on column messages.created_at is 'Время создания сообщения';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists messages;
-- +goose StatementEnd
