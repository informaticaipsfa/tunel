package sssifanb

//RedSocial
type RedSocial struct {
	Twitter   string `json:"twitter,omitempty" bson:"twitter"`
	Facebook  string `json:"facebook,omitempty" bson:"facebook"`
	Instagram string `json:"instagram,omitempty" bson:"instagram"`
	Linkedin  string `json:"linkedin,omitempty" bson:"linkedin"`
}
