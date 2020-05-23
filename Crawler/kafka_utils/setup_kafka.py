from kafka import KafkaProducer
from kafka.errors import KafkaError
import json
import time
import requests

class KafkaMsg():
	def __init__(self, query):
		self.producer = None
		self.servers = ['192.168.0.120:19092',
						'192.168.0.120:29092']
		self.query = query

	def setup_json_producer(self):
		self.producer = KafkaProducer(
						value_serializer=lambda m: json.dumps(m).encode('ascii'),
						bootstrap_servers='192.168.0.120:19092'
						)
	def setup_protobuf_producer(self):
		self.producer = KafkaProducer(
						bootstrap_servers=self.servers,
						retries=5,
						acks='all'
						)
	def send_protobuf_message(self, topic, data):
		print("topic", data)
		self.producer.send(topic, data.SerializeToString())
		if topic == "new_job":
			self.send_notif(data)
		time.sleep(0.2)

	def send_json_message(self, topic, data):
		self.producer.send(topic, data)

	def send_notif(self, data):
		url = "https://fcm.googleapis.com/fcm/send"
		headers= {"Authorization": "key=AAAAwRI-MiY:APA91bHExWXJyxdLKzDo-w7iS79nKs2e5QiYr4xIYBMzc4uHvExfXfkuHfrktpAkFEqTRp-uw5h5PBVQLwHmhlMkpF7Ub6SsFG_zXUODCp3OXe5LTB4w-aBsFiGzlWOqhKz_jf6izL-f", "Content-Type": "application/json"}
		res = self.query.get_rows_by_exp(data.ExperienceLevel, data.CompanyUUID)
		title_experience = "intern level"
		if data.ExperienceLevel == 2:
			title_experience = "entry level"
		elif data.ExperienceLevel == 3:
			title_experience = "mid level"
		elif data.ExperienceLevel == 4:
			title_experience = "senior level"
		elif data.ExperienceLevel == 5:
			title_experience = "manager level"
		print(res, "DEVICEIDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")

		for devId in res:
			print(devId[0])
			notif_details = { 
				"to": devId[0],
			 	"data": { 
					"title": "New " + title_experience + " job!", 
					"body": data.CompanyName + " has just posted a new " + title_experience + " job",
					"url": data.JobLink
					}
				}
			notif_details = json.dumps(notif_details)
			r = requests.post(url=url, data=notif_details, headers=headers)
			print(r.content)


