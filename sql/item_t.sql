use gonotify;
drop table if exists gn_item;
create table gn_item (
	id serial PRIMARY KEY NOT NULL AUTO_INCREMENT,
	date timestamp DEFAULT now(),
	level VARCHAR(5) NOT NULL,
	recipient_mail VARCHAR(100),
	recipient_user VARCHAR(20),
	message VARCHAR(1024)
)