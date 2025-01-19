package database

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type User struct {
	ID             bson.ObjectID `bson:"_id"`
	SocialProvider string        `bson:"social_provider"`
	SocialUID      string        `bson:"social_uid"`
	Nickname       *string       `bson:"nickname"`
	CreatedAt      time.Time     `bson:"created_at"`
	UpdatedAt      time.Time     `bson:"updated_at"`
}

func (u *User) GetID() string {
	return u.ID.Hex()
}
