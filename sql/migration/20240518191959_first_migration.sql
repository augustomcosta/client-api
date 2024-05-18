-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS clients (
                                           id uuid PRIMARY KEY,
                                           name VARCHAR(255) NOT NULL,
                                           age INT NOT NULL,
                                           social_info_instagram VARCHAR(255) NOT NULL,
                                           social_info_linkedin VARCHAR(255) NOT NULL,
                                           social_info_facebook VARCHAR(255) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd