package model

type User struct {
	Name   string `json:"name"`
	Target string `json:"target"`
	Career string `json:"career"`
	State  string `json:"state"`
}

type VarSet struct {
	AllNum   int `json:"allNum"`
	MafiaNum int `json:"mafiaNum"`
	DocNum   int `json:"docNum"`
	PolNum   int `json:"polNum"`
}
