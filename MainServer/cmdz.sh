sudo docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic new_job --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

sudo docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic del_job --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

