CREATE TABLE gumsite_products (
	id bigserial not null primary key,
	name varchar not null unique,
	description varchar not null,
	cost int
);