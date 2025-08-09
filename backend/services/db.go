package services

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "host=postgres user=" + os.Getenv("POSTGRES_USER") +
       " password=" + os.Getenv("POSTGRES_PASSWORD") +
       " dbname=" + os.Getenv("POSTGRES_DB") +
       " port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: ", err)
	}

	DB = database
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
}