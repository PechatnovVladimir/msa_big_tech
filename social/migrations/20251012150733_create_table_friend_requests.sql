-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists friend_requests (
    id text not null default gen_random_uuid(),
    from_user_id text not null,
    to_user_id text not null,
    status int not null,
    created_at timestamp with time zone not null default now(),
    primary key (id)

);
comment on table friend_requests is 'Таблица заявок в друзья';
comment on column friend_requests.id is 'Уникальный идентификатор заявки';
comment on column friend_requests.from_user_id is 'От кого заявка';
comment on column friend_requests.to_user_id is 'К кому заявка';
comment on column friend_requests.status is 'Статус заявки';
comment on column friend_requests.created_at is 'Время создания заявки';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists friend_requests;
-- +goose StatementEnd
