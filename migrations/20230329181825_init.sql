-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE images
(
    id         SERIAL PRIMARY KEY,
    image_url  TEXT        NOT NULL,
    prompt     TEXT        NOT NULL,
    status     TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON images
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS images;
-- +goose StatementEnd
