Responsibilities : 

- Store user free & locked balance & spots
- Expose gRPC calls to increase, decrease, lock, unlock balance & spots
- Listen to event's comming from queue and based on fills created lock/unlock users balance/spots
- Emit all transaction events to queue for db-writter to write in database
- Store the exchagne fund & insurance fund


Steps to Run : 

1. Create dir :  `<HOME_DIR>/.accounts-service`
2. `Generate gRPC Clients : protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     protobufs/user_balance.proto`


TODO : 

1. (Implement Kafka First) Sending events to kafka(reading them from the logs)(tracking last log read)
2. (Implement snapshot service & uploading to S3)Implement recreation of database
3. (Implement matching engine)Consuming matching engine events
