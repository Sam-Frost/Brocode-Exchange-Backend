Directoy in container : /opt/kafka/bin

Creating Topic : 
`./kafka-topics.sh --bootstrap-server localhost:9092 --create --topic btc-spot-orders`

Producing to topic : 
`./kafka-console-producer.sh --bootstrap-server localhost:9092 --topic  btc-spot-orders`
