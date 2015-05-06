use gonotify;
drop table if exists gn_user;
create table gn_user (
	id SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
	modified TIMESTAMP NOT NULL DEFAULT now(),
	uname VARCHAR(24) NOT NULL,
	rname VARCHAR(100) NOT NULL,
	mail VARCHAR(100) NOT NULL UNIQUE KEY,
	UNIQUE INDEX mail_idx USING HASH (mail)
)