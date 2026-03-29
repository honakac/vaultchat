package database

import (
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v3"
)

type SendMessageRequest struct {
	Cuid         string `json:"cuid"`
	ReceiverAddr string `json:"receiver_addr"`
	Payload      string `json:"payload"`
}

func (db *Database) AddInboxMessage(c fiber.Ctx, req SendMessageRequest) error {
	log.Infof("Adding inbox message for receiver: %s, payload size: %d", req.ReceiverAddr, len(req.Payload))

	if err := db.db.Create(&InboxMessage{
		Cuid:         req.Cuid,
		ReceiverAddr: req.ReceiverAddr,
		Payload:      req.Payload,
	}).Error; err != nil {
		log.Errorf("Failed to add inbox message: %v", err)
		return err
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
