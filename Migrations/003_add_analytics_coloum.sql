--+goose up
--+goose StatementBegin
ALTER TABLE url 
ADD COLUMN number_of_req BIGINT DEFAULT 0,
ADD COLUMN avg_res_time DECIMAL(5,
2) DEFAULT 0;
--+goose StatementEnd
