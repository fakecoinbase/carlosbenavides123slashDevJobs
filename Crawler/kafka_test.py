from kafka import KafkaProducer
producer = KafkaProducer(bootstrap_servers='192.168.0.120:19092')
for _ in range(100):
    producer.send('foo', b'some_message_bytes')