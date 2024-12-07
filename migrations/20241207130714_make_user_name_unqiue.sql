-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users"
    ADD CONSTRAINT user_unique_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users"
DROP CONSTRAINT IF EXISTS user_unique_name;
-- +goose StatementEnd
