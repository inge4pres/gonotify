use gonotify;
drop table if exists gn_pwd;
create table gn_pwd (
	id SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
	modified TIMESTAMP NOT NULL DEFAULT now(),
	uid BIGINT NOT NULL,
	pwd BLOB NOT NULL,
	UNIQUE INDEX uid_idx USING BTREE (uid),
	FOREIGN KEY (uid) REFERENCES gn_user(id)
)