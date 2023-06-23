package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"store-product-restapi/models"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type response struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Message string    `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return db
}

func Addstore(store models.Store) string {
	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO store (sname) VALUES ($1) RETURNING sname`

	var name string
	err := db.QueryRow(sqlStatement, store.Name).Scan(&name)

	if err != nil {
		log.Fatalf("Unable to execute the store insert query. %v", err)
	}

	fmt.Printf("Inserted a new store %s\n", name)
	return name
}

func Addproduct(product models.Product, metadata models.Metadata) string {

	db := CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO product (pname, created_at, expiry_date, metadata) VALUES ($1, $2, $3, $4) RETURNING pname`

	var name string
	jsonb, err := json.Marshal(metadata)
	
	if err != nil{
		log.Fatalf("Unable to encode json. %v", err)
	}

	var metadataJSONb interface{}
	err = json.Unmarshal(jsonb, &metadataJSONb)
	if err != nil {
		log.Fatalf("Unable to unmarshal json. %v", err)
	}

	err = db.QueryRow(sqlStatement, product.Name, product.Created_at, product.Expiry_date, pq.Array([]byte(jsonb))).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No rows were returned from the database.")
		}
		log.Fatalf("Unable to execute the product insert query. %v", err)
	}

	fmt.Printf("Inserted a new product %s\n", name)
	return name
}
