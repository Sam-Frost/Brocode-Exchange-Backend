## Web Server
- Create, login user
- Add balance of user
- Create/Cancel spot order
- Create/Cancel perp
- Get candle data
- Get balance / positions / open orders / fill history / order history / position history
- Get all markets and their data

## Accounts Service
- Maintain user balance
- gRPC with backend
- Log events to file
- Listen to spot matching engine queue to update balance and spots(consuming the fills)
- Send file events to kafka
- Rebuild incase of crash

## Matching Engine
- Receive Order from kafka
- In Memory DB
- Perform Matching
- Lock Free Ring Buffer
- Core Pinninig
- Kafka producer to read from Lock Free Buffer and produce to kafka stream
- DB state rebuild incase of crash
- Thread|DB|Ring Buffer|Kafka Producer Per Market
- Send heartbeat to HB Service

## Snapshot Service
1. Snapshot Accounts Service
- Read data from kafka stream
- Recreate the state
- Snapshot data at set interval
- Upload Data with Kafka Offset to s3
2. Snapshots Matching Engine
- Read data from kafka stream
- Recreate the state
- Snapshot data at set interval
- Upload Data with Kafka Offset to s3


## DB Writer
1. Accounts Service Stream
- Batch and write updates to user
  - money balance
  - spot balance
  - positions
- Incase of server crash and replay, avoid RE-WRITES

2. Spot Matching Engine Stream
- Batch and write fills
- Incase of server crash and replay, avoid RE-WRITES

3. Perp Matching Engine Stream
- Batch and write fills
- Incase of server crash and replay, avoid RE-WRITES

4. Write to timescale DB 
- Store data of OPEN, CLOSE, HIGH, LOW

-------------------------------------------------------------------------------------------------------------------------------

## Web Socket Server 
- Read data from matching engine stream
- Send last traded price through WS
- Send orderbook updates

## Admin Panel
1. Crash matching engine
2. Crash accounts service
3. Create market(name, symbol, quantityStep, priceTick)
4. Crash db writer


Pending : 

## Perps Engine

## Liquidation Engine

## ADL

## Funding Rate

## Market Fees

## Market Fund

## Insurance Fund

## Notifications

## Logs and Metrics

## Deployment on K8s

## Heart Beat Service
- Receives heart beat from the core components of Exchange
  - Accounts Service
  - Spot Matching Engine
  - Perp Matching Engine
  - Liquduation Engine, ADL, Funding Rate(Wherever these are happening)
- Backend constantly check with heartbeat service to ensure system ok before proceding with ordeer


## Infra Components:
1. gRPC
2. Kafka
3. Timescale DB
4. Logs
5. Metrics
