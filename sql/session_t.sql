use gonotify;
drop table if exists gn_session;
create table gn_session (
	id SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
	created TIMESTAMP NOT NULL DEFAULT NOW(),
	expires TIMESTAMP NOT NULL,
	uid BIGINT NOT NULL,
	scookie VARCHAR(128) NOT NULL,
	UNIQUE INDEX uid_idx USING BTREE (uid,scookie)
);