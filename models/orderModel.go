package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date time.Time          `json:"order_date"validate:"required"`
	Created_at time.Time          `json:"create_at"`
	Updated_at time.Time          `json:"update_at"`
	Order_id   string             `json:"order_id"`
	Table_id   *string            `json:"table_id"validate:"required"`
}
