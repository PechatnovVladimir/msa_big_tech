-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists friends (
    user_id text not null,
    friend_user_id text not null,
    created_at timestamp with time zone not null default now(),
    UNIQUE (user_id,friend_user_id)
);
comment on table friends is 'Таблица друзей';
comment on column friends.user_id is 'ID пользователя';
comment on column friends.friend_user_id is 'ID пользователя друга';
comment on column friends.created_at is 'Время возникновения дружбы';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists friends;
-- +goose StatementEnd
