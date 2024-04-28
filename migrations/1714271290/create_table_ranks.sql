CREATE TABLE `ranks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `staff_id` int(11) DEFAULT NULL,
  `deparment_id` int(11) DEFAULT NULL,
  `level` int(11) NOT NULL DEFAULT 1,  
  PRIMARY KEY (`id`)
);