use gonotify;
drop table if exists gn_user;
create table gn_user (
	id SERIAL NOT NULL AUTO_INCREMENT,
	modified TIMESTAMP NOT NULL DEFAULT now(),
	uname VARCHAR(24) NOT NULL,
	rname VARCHAR(100) NOT NULL,
	mail VARCHAR(100) NOT NULL UNIQUE KEY,
	pwd BLOB NOT NULL,
	islogged BOOL DEFAULT FALSE,
	UNIQUE INDEX mail_idx USING HASH (id,mail),
	CONSTRAINT pk_uid PRIMARY KEY (id,uname)
);