use gonotify;
drop table if exists gn_session;
create table gn_session (
	id SERIAL PRIMARY KEY NOT NULL AUTO_INCREMENT,
	created TIMESTAMP NOT NULL DEFAULT now(),
	uid BIGINT NOT NULL,
	scookie BLOB NOT NULL
);