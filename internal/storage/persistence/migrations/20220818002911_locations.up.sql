-- create "location" table
CREATE TABLE `location` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(2048) NOT NULL, `description` varchar(2048) NOT NULL, `source` enum('unspecified','minio') NOT NULL, `purpose` enum('unspecified','aip_store') NOT NULL, `uuid` char(36) NOT NULL, PRIMARY KEY (`id`), INDEX `location_name` (`name` (50)), INDEX `location_uuid` (`uuid`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
