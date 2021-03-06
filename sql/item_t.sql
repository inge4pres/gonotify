use gonotify;
drop table if exists gn_item;
create table gn_item (
	id serial PRIMARY KEY NOT NULL AUTO_INCREMENT,
	date TIMESTAMP DEFAULT now(),
	level VARCHAR(10) NOT NULL,
	recipient VARCHAR(100),
	sender VARCHAR(100),
	subject TINYTEXT,
	message MEDIUMTEXT,
	archived BOOL DEFAULT FALSE,
	UNIQUE INDEX rcpt_idx USING BTREE (id, recipient)
);