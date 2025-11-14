create table if not exists urls (
	id int auto_increment primary key,
	urn varchar(255) not null unique,
	redirect_url text(65536) not null,
	created_at datetime not null default current_timestamp()
);
