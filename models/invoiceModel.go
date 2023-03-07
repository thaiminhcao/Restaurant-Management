package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"ed=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"payment_status" validate:"required,ep=PENDING|EQ=PAID"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"create_at"`
	Updated_at       time.Time          `json:"update_at"`
}