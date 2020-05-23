package notifications

import (
	"fmt"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/dbconf"
	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	"github.com/carlosbenavides123/DevJobs/MainServer/pb/notifications/notificationreqpb"
	"github.com/golang/protobuf/proto"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func GetNotificationPreference(deviceUUID string) (*Notifications, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()
	stmt, dbPrepareErr := db.Prepare(`select n.company_uuid, n.intern, n.entry, n.mid, n.senior, n.manager 
							from notifications_by_company n
							WHERE n.device_uuid=?`)
	if dbPrepareErr != nil {
		panic(dbPrepareErr.Error())
	}

	res, dbErr := stmt.Query(deviceUUID)

	if dbErr != nil {
		panic(dbErr.Error())
	}

	notifSettings := &Notifications{}
	var companiesUUID []string

	for res.Next() {
		var CompanyUUID string
		var Intern, Entry, Mid, Senior, Manager bool

		scanErr := res.Scan(&CompanyUUID, &Intern, &Entry, &Mid, &Senior, &Manager)
		if scanErr != nil {
			panic(scanErr.Error())
		}
		companiesUUID = append(companiesUUID, CompanyUUID)
		notifSettings.Intern = Intern
		notifSettings.Entry = Entry
		notifSettings.Mid = Mid
		notifSettings.Senior = Senior
		notifSettings.Manager = Manager
	}
	notifSettings.CompaniesUUID = companiesUUID
	return notifSettings, nil
}

func CreateNotificationPreference(notification_list *Notifications, p *kafka.Producer) {
	db := dbconf.DbConn()
	defer db.Close()
	for _, s := range notification_list.CompaniesUUID {
		stmt, dbPrepareErr := db.Prepare(`INSERT INTO notifications_by_company 
		(company_uuid, device_uuid, intern, entry, mid, senior, manager) 
		VALUES(?, ?, ?, ?, ?, ?, ?)`)

		if dbPrepareErr != nil {
			panic(error(dbPrepareErr))
		}
		stmt.Exec(s, notification_list.DeviceUUID, notification_list.Intern, notification_list.Entry, notification_list.Mid, notification_list.Senior, notification_list.Manager)
		notifCreatePbBinary := createNotifPb(s, notification_list.DeviceUUID, "CREATE",
			notification_list.Intern, notification_list.Entry,
			notification_list.Mid, notification_list.Senior,
			notification_list.Manager)
		produceKafkaMessage(p, "NotificationReq", notifCreatePbBinary)
		fmt.Println(s)
	}
}

func UpdateNotificationPreference(notifUpdateReq *NotificationsUpdate, p *kafka.Producer) (*utils.ApplicationSuccess, *utils.ApplicationError) {
	db := dbconf.DbConn()
	defer db.Close()

	stmtCheck, dbPrepCheckErr := db.Prepare(`SELECT 1 FROM notifications_by_company WHERE device_uuid=? AND company_uuid=?`)

	if dbPrepCheckErr != nil {
		panic(dbPrepCheckErr.Error())
	}

	stmtInsert, dbInsertErr := db.Prepare(`INSERT INTO notifications_by_company 
											(company_uuid, device_uuid, intern, entry, mid, senior, manager) 
											VALUES(?, ?, ?, ?, ?, ?, ?)`)

	if dbInsertErr != nil {
		panic(dbInsertErr.Error())
	}

	stmtUpdate, dbUpdateErr := db.Prepare(`UPDATE notifications_by_company 
									SET intern=?, entry=?, mid=?, senior=?, manager=?
									WHERE company_uuid=? AND device_uuid=?`)

	if dbUpdateErr != nil {
		panic(dbUpdateErr.Error())
	}

	stmtDel, dbDelErr := db.Prepare(`DELETE FROM notifications_by_company
										WHERE company_uuid=? AND device_uuid=?`)

	if dbDelErr != nil {
		panic(dbUpdateErr.Error())
	}

	for _, s := range notifUpdateReq.CompaniesUUID {
		res, dbErr := stmtCheck.Query(notifUpdateReq.DeviceUUID, s)

		if dbErr != nil {
			panic(dbErr.Error())
		}

		if res.Next() == false {
			stmtInsert.Exec(s, notifUpdateReq.DeviceUUID, notifUpdateReq.Intern, notifUpdateReq.Entry, notifUpdateReq.Mid, notifUpdateReq.Senior, notifUpdateReq.Manager)
			notifCreatePbBinary := createNotifPb(s, notifUpdateReq.DeviceUUID, "CREATE",
				notifUpdateReq.Intern, notifUpdateReq.Entry,
				notifUpdateReq.Mid, notifUpdateReq.Senior,
				notifUpdateReq.Manager)
			produceKafkaMessage(p, "NotificationReq", notifCreatePbBinary)
		} else {
			stmtUpdate.Exec(notifUpdateReq.Intern, notifUpdateReq.Entry, notifUpdateReq.Mid, notifUpdateReq.Senior, notifUpdateReq.Manager, s, notifUpdateReq.DeviceUUID)
			notifCreatePbBinary := createNotifPb(s, notifUpdateReq.DeviceUUID, "UPDATE",
				notifUpdateReq.Intern, notifUpdateReq.Entry,
				notifUpdateReq.Mid, notifUpdateReq.Senior,
				notifUpdateReq.Manager)
			produceKafkaMessage(p, "NotificationReq", notifCreatePbBinary)
		}
	}

	for _, s := range notifUpdateReq.RemoveCompUUID {
		stmtDel.Exec(s, notifUpdateReq.DeviceUUID)
		notifCreatePbBinary := createNotifPb(s, notifUpdateReq.DeviceUUID, "DELETE",
			notifUpdateReq.Intern, notifUpdateReq.Entry,
			notifUpdateReq.Mid, notifUpdateReq.Senior,
			notifUpdateReq.Manager)
		produceKafkaMessage(p, "NotificationReq", notifCreatePbBinary)
	}

	return &utils.ApplicationSuccess{
		Message:    fmt.Sprintf("Update has been received!"),
		StatusCode: http.StatusAccepted,
		Code:       "Updated",
	}, nil
}

func createNotifPb(companyUUID string, deviceUUID string, action string, intern bool, entry bool, mid bool, senior bool, manager bool) []byte {
	notifPb := &notificationreqpb.NotifReq{
		CompanyUUID: companyUUID,
		DeviceUUID:  deviceUUID,
		Action:      action,
		Intern:      intern,
		Entry:       entry,
		Mid:         mid,
		Senior:      senior,
		Manager:     manager,
	}
	notifPbBinary, err := proto.Marshal(notifPb)
	if err != nil {
		panic(err.Error())
	}
	return notifPbBinary
}

func produceKafkaMessage(p *kafka.Producer, topic string, data []byte) {
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)
}
