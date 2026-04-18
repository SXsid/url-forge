--+goose up
--+goose StatementBegin
ALTER TABLE url ALTER COLUMN code DROP NOT NULL;
--+goose StatementEnd
