CREATE TABLE IF NOT EXISTS `test_case`(
   `problem_id` BIGINT,
   `case_id` INT,
   `case` MEDIUMTEXT,
   `expected` MEDIUMTEXT,
   PRIMARY KEY (`problem_id`, `case_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;