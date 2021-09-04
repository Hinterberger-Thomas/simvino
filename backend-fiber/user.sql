CREATE TABLE IF NOT EXISTS "user"
(
 userID   int NOT NULL,
 email    varchar(255) NOT NULL,
 password bytea NOT NULL,
 CONSTRAINT PK_user PRIMARY KEY ( userID )
);



