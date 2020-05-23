package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carlosbenavides123/DevJobs/MainServer/mvc/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	notifications "github.com/carlosbenavides123/DevJobs/MainServer/mvc/domain/Notifications"
)

func GetNotificationPreference(r *http.Request) (*notifications.Notifications, *utils.ApplicationError) {
	deviceUUID := r.URL.Query().Get("deviceUUID")
	if deviceUUID == "" {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("Device UUID is required!"),
			StatusCode: http.StatusBadRequest,
			Code:       "Device UUID Missing.",
		}
	}
	return notifications.GetNotificationPreference(deviceUUID)
}

func CreateNotificationPreference(r *http.Request, p *kafka.Producer) {
	decoder := json.NewDecoder(r.Body)
	var notifRequest *notifications.Notifications
	decodeErr := decoder.Decode(&notifRequest)
	if decodeErr != nil {
		panic(decodeErr)
	}
	fmt.Println(notifRequest)
	notifications.CreateNotificationPreference(notifRequest, p)
}

func UpdateNotificationPreference(r *http.Request, p *kafka.Producer) (*utils.ApplicationSuccess, *utils.ApplicationError) {
	decoder := json.NewDecoder(r.Body)
	var notifUpdateRequest *notifications.NotificationsUpdate
	decodeErr := decoder.Decode(&notifUpdateRequest)
	if decodeErr != nil {
		panic(decodeErr)
	}
	fmt.Println(notifUpdateRequest)
	return notifications.UpdateNotificationPreference(notifUpdateRequest, p)
}
