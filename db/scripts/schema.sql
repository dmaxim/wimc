CREATE TABLE `wimc`.`CloudResource` (
    CloudResourceId INT NOT NULL AUTO_INCREMENT,
    CloudId VARCHAR(255) NOT NULL,
    Location VARCHAR(50) NOT NULL,
    Name VARCHAR(100) NOT NULL,
    Notes VARCHAR(255) NULL,
    PRIMARY KEY (CloudResourceId)

)


