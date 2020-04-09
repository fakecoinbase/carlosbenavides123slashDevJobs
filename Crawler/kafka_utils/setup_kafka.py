from kafka import KafkaProducer
from kafka.errors import KafkaError
import json

class KafkaMsg():
    def __init__(self):
        self.producer = None
        self.servers = ['192.168.1.66:19092',
                        '192.168.1.66:32181']

    def setup_json_producer(self):
        self.producer = KafkaProducer(
                        value_serializer=lambda m: json.dumps(m).encode('ascii'),
                        bootstrap_servers='192.168.1.66:19092'
                        )

    def send_json_message(self, topic, data):
        # self.producer.send('foo', b'some_message_bytes')
        self.producer.send(topic, data)
    