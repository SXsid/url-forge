--+goose up 
--+goose statement begin
CREATE TABLE url(id SEARIAL PRIMARY KEY NOT NULL,
code VARCHAR(
    6)NOT NULL,
    url TEXT NOT NULL,
    number_of_req INTEGER DEFAULT 0,
    avg_rep_time DECIMAL(
        3,
        2)NOT NULL DEFAULT 0;
))--+goose statement end
