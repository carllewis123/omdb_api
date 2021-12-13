# omdb_api

1. Run Query to Create Table
CREATE TABLE `log` (
  `protocol` varchar(50) DEFAULT NULL,
  `request` text DEFAULT NULL,
  `response` text DEFAULT NULL,
  `insert_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

2. Setting Connection DB in directory config/config.yaml


3. How to start Http Server and GRPC Server
- Opent directory omdb_api
- run command go run main.go in pacakge omdb_api

2. How to run grpc client
- Open directory omdb_api/client
- run command : go run client.go