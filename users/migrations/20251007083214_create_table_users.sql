-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS user_profiles (
    id text NOT NULL,
    email text NOT NULL,
    nickname text NOT NULL,
    bio text,
    avatar_url text,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE (email),
    UNIQUE (nickname)
);

comment on table user_profiles is 'Таблица профайлов пользователей';
comment on column user_profiles.id is 'Уникальный идентификатор пользователя';
comment on column user_profiles.email is 'email пользователя';
comment on column user_profiles.nickname is 'Псевдоним пользователя';
comment on column user_profiles.bio is 'Биография пользователя';
comment on column user_profiles.avatar_url is 'Ссылка на аватарку пользователя';
comment on column user_profiles.created_at is 'Когда создан';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists user_profiles;
-- +goose StatementEnd
