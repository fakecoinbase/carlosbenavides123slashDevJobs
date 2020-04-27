from kafka import KafkaProducer
from kafka.errors import KafkaError
import json
import time

class KafkaMsg():
	def __init__(self):
		self.producer = None
		self.servers = ['192.168.1.66:19092',
						'192.168.1.66:29092']

	def setup_json_producer(self):
		self.producer = KafkaProducer(
						value_serializer=lambda m: json.dumps(m).encode('ascii'),
						bootstrap_servers='192.168.1.66:19092'
						)
	def setup_protobuf_producer(self):
		self.producer = KafkaProducer(
						bootstrap_servers=self.servers,
						retries=5,
						acks='all'
						)
	def send_protobuf_message(self, topic, data):
		print(topic, data)
		self.producer.send(topic, data.SerializeToString())
		# time.sleep(0.1)

	def send_json_message(self, topic, data):
		# self.producer.send('foo', b'some_message_bytes')
		self.producer.send(topic, data)
	