sudo docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic new_job --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

sudo docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic del_job --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181


CREATE DATABASE devjobs;

CREATE TABLE jobs (
    UUID varchar(255),
    CompanyName varchar(255),
    JobLink varchar(255),
    JobPosted varchar(255),
    JobFound varchar(255),
    Active TINYINT(1),
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (UUID)
);