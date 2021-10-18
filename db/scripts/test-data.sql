
INSERT INTO `wimc`.`CloudResource`
(CloudId, Location, Name, Notes)
VALUES
('/resource-one', 'eastus', 'vnet', 'Test notes');

INSERT INTO `wimc`.`CloudResource`
(CloudId, Location, Name)
VALUES
('/resource-two', 'eastus', 'snet-1');

INSERT INTO `wimc`.`CloudResource`
(CloudId, Location, Name, Notes)
VALUES
('/resource-three', 'eastus', 'snet-2', 'Test notes');

INSERT INTO `wimc`.`CloudResource`
(CloudId, Location, Name)
VALUES
('/resource-one', 'eastus', 'vm');