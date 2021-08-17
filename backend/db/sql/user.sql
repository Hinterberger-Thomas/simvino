CREATE TABLE IF NOT EXISTS user 
(
 email    varchar(255) NOT NULL AUTO_INCREMENT,
 userID   int NOT NULL ,
 password varchar(255) NOT NULL ,

PRIMARY KEY (email, userID)
);
