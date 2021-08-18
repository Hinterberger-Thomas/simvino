CREATE TABLE IF NOT EXISTS `balance`
(
 `balanceID` int NOT NULL AUTO_INCREMENT ,
 `userID`    int NOT NULL ,
 `currency`  varchar(255) NOT NULL ,
 `value`     int NOT NULL ,

PRIMARY KEY (`balanceID`),
KEY `fkIdx_11` (`userID`),
CONSTRAINT `FK_10` FOREIGN KEY `fkIdx_11` (`userID`) REFERENCES `user` (`userID`)
);
