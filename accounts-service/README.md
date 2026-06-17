Responsibilities : 

- Store user free & locked balance & spots
- Expose gRPC calls to increase, decrease, lock, unlock balance & spots
- Listen to event's comming from queue and based on fills created lock/unlock users balance/spots
- Emit all transaction events to queue for db-writter to write in database
- Store the exchagne fund & insurance fund


`Generate gRPC Clients : protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     protobufs/user_balance.proto`
