package models

type Teste struct {
	ID    uint `gorm: "primaryKey; autoIncrement:true"`
	Texto string
}
