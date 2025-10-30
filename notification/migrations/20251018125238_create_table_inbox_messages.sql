-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists inbox_messages (
    id text not null,
    topic text not null,
    partition int not null,
    offsett bigint not null, --ofsett - две tt так как с одной зарезервировано в postgres
    payload jsonb not null ,
    status text not null ,
    attempts int not null,
    last_error text not null,
    received_at timestamptz not null default now(),
    processed_at TIMESTAMPTZ,
    primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists inbox_messages;
-- +goose StatementEnd
