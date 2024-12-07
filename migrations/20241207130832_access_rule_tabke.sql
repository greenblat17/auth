-- +goose Up
-- +goose StatementBegin
CREATE TABLE "access_rules"
(
    role     int  NOT NULL,
    endpoint text NOT NULL,
    PRIMARY KEY (role, endpoint)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "access_rules";
-- +goose StatementEnd