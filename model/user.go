package model

type User struct {
	// ID      bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Name    string `json:"Name"`
	Age     int64  `json:"Age"`
	Address string `json:"Address"`
}
