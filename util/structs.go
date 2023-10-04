package util

type Member struct {
	Email     string `json:"Email"`
	Pass      string `json:"Pass"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Member_Since string
}

type LoginInfo struct {
	Email string `json:"Email"`
	Pass  string `json:"Pass"`
}