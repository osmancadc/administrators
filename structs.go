package main

type Administrator struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Owner      string `json:"owner" db:"owner"`
	Citicality string `json:"criticality" db:"criticality"`
}
