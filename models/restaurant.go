package models

type Restaurant struct {
	Isdelivering bool   `json:"is_delivering"`
	Mino         int    `json:"mino"`
	Del          int    `json:"del"`
	Cu           string `json:"cu"`
	City         string `json:"city"`
	Na           string `json:"na"`
	Cs           string `json:"cs_phone"`
	Address      string `json:"ad"`
	Info         string `json:"info"`
}
