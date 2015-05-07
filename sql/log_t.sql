use gonotify;
drop table if exists gn_log;
create table gn_log (
	id serial PRIMARY KEY NOT NULL AUTO_INCREMENT,
	date timestamp DEFAULT now(),
	level VARCHAR(5) NOT NULL,
	message VARCHAR(1024),	
	INDEX level_idx USING BTREE (id,level) 
);