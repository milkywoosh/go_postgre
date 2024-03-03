-- CREATE TABLE teacher( 
-- 	id_teacher int GENERATED ALWAYS AS IDENTITY,
-- 	name_teacher varchar(100),
-- 	id_subject int,
-- 	email varchar(150),
-- 	PRIMARY KEY(id_teacher),
-- 	CONSTRAINT fk_subject
-- 	FOREIGN KEY (id_subject)
-- 	REFERENCEs subject(id_subject)	
-- )

-- ALTER TABLE teacher
-- ADD COLUMN id_people int,
-- ADD CONSTRAINT fk_people 
-- FOREIGN KEY("id_people")
-- REFERENCES "people" ("id") 
-- ON DELETE CASCADE




-- INSERT INTO teacher( name_teacher, id_subject, email, id_people)
-- VALUES ('dian', 4, 'dian@x.com', 2),
-- 	('mona', 1, 'mona@x.com', 2),
-- 	('win', 3, 'win@x.com', 2)
	