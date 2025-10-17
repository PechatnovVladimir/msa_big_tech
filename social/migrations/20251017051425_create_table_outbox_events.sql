-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS outbox_events (
                                                    id              UUID        NOT NULL default gen_random_uuid(),
                                                    aggregate_type  TEXT        NOT NULL,
                                                    aggregate_id    TEXT        NOT NULL,
                                                    event_type      TEXT        NOT NULL,
                                                    payload         JSONB       NOT NULL,
                                                    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
                                                    published_at    TIMESTAMPTZ,
                                                    retry_count     INT         NOT NULL DEFAULT 0,
                                                    next_attempt_at TIMESTAMPTZ,
                                                    primary key (created_at, id)
) PARTITION BY RANGE (created_at);

COMMENT ON TABLE  outbox_events IS 'Таблица событий для паттерна Transactional Outbox (partitioned)';
COMMENT ON COLUMN outbox_events.id               IS 'Уникальный идентификатор события (UUID)';
COMMENT ON COLUMN outbox_events.aggregate_id     IS 'Идентификатор агрегата (например, request_id)';
COMMENT ON COLUMN outbox_events.aggregate_type   IS 'Тип агрегата (например, friend_request)';
COMMENT ON COLUMN outbox_events.event_type       IS 'Тип события (например, social.friend.updated)';
COMMENT ON COLUMN outbox_events.payload          IS 'Тело события JSONB';
COMMENT ON COLUMN outbox_events.created_at       IS 'Время записи события в outbox';
COMMENT ON COLUMN outbox_events.published_at     IS 'Время успешной публикации события';
COMMENT ON COLUMN outbox_events.retry_count      IS 'Количество попыток публикации';
COMMENT ON COLUMN outbox_events.next_attempt_at  IS 'Время следующей попытки публикации (NULL = сразу)';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists outbox_events
-- +goose StatementEnd
