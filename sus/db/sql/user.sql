CREATE TABLE IF NOT EXISTS `user`
(
 `userID`   int NOT NULL AUTO_INCREMENT,
 `password` varchar(255) NOT NULL,
 `email`    varchar(255) NOT NULL UNIQUE,

PRIMARY KEY (`userID`)
);