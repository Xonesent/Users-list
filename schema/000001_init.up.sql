CREATE TABLE users 
(
	id serial not null unique,
	name varchar(255) not null,
	surname varchar(255) not null,
	patronymic varchar(255) not null,
	age int not null,
	gender varchar(6) not null,
	nationality varchar(2) not null
);