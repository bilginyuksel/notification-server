package main

import (
	"log"

	"github.com/bilginyuksel/notification-server/server/conn"
	"github.com/bilginyuksel/notification-server/server/entity"
)

func SendNotification(c conn.Connection) error {
	notification := entity.Notification{
		Title:       "Golang Push Server Test",
		Description: "First push notification try"}

	if err := c.SendJSON(notification); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
