use gonotify;
drop table if exists gn_item;
create table gn_item (
	id serial PRIMARY KEY NOT NULL AUTO_INCREMENT,
	date TIMESTAMP DEFAULT now(),
	level VARCHAR(5) NOT NULL,
	recipient VARCHAR(100),
	sender VARCHAR(100),
	subject TINYTEXT,
	message MEDIUMTEXT,
	archived BOOL DEFAULT FALSE
)