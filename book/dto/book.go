package dto

type Request struct {
	From   string `bson:"first_name"`
	To     string `bson:"last_name"`
	Helper string `bson:"helper"`
}
