package models

type User struct {
	Id        string `bson:"_id,omitempty" json:"id,omitempty"` //bson:omitempty will auto create id in mongo
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
