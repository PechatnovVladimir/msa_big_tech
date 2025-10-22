-- +goose Up
-- +goose NO TRANSACTION

-- CREATE INDEX CONCURRENTLY cannot run inside a transaction block (SQLSTATE 25001)
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_messages_chat_id_created_at
ON messages (chat_id,created_at);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop index if exists idx_messages_chat_id_created_at;
-- +goose StatementEnd
