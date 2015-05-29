use gonotify;
drop table if exists gn_user;
create table gn_user (
	id SERIAL NOT NULL AUTO_INCREMENT PRIMARY KEY,
	modified TIMESTAMP NOT NULL DEFAULT now(),
	uname VARCHAR(24) NOT NULL,
	rname VARCHAR(100) NOT NULL,
	mail VARCHAR(100) NOT NULL UNIQUE KEY,
	pwd VARCHAR(256) NOT NULL,
	islogged BOOL DEFUALT FALSE,
	UNIQUE INDEX mail_idx USING HASH (id,mail),
	UNIQUE INDEX uname_idx USING BTREE (id,uname),
	CONSTRAINT uni_uname UNIQUE (uname),
	CONSTRAINT uni_mail UNIQUE KEY(mail)
);