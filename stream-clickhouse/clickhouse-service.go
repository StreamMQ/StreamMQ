package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/eclipse/paho.mqtt.golang"
	_ "github.com/ClickHouse/clickhouse-go"
)

var (
    clickHouseEndpoint = "tcp://localhost:9000"
	clickHouseTable    = "measurements"
)

type Measurement struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}


func onMessageReceived(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	fmt.Printf("Received message: %s on topic: %s\n", payload, msg.Topic())

	// Parse the JSON array payload
	var measurements []Measurement
	err := json.Unmarshal([]byte(payload), &measurements)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return
	}

	fmt.Println("Parsed message: ", measurements)

	// Open a connection to ClickHouse
	db, err := sql.Open("clickhouse", clickHouseEndpoint)
	if err != nil {
		log.Printf("Error connecting to ClickHouse: %v\n", err)
		return
	}
	defer db.Close()
	
	fmt.Println("Connection opened to Database")

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error beginning transaction: %v\n", err)
		return
	}
	defer tx.Rollback() // Rollback transaction if not committed

	fmt.Println("Begin Transation!!!")

	// Prepare the INSERT statement
	stmt, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s (name, value) VALUES (?, ?)", clickHouseTable))
	if err != nil {
		log.Printf("Error preparing INSERT statement: %v\n", err)
		return
	}
	defer stmt.Close()

	// Insert data
	for _, m := range measurements {
		_, err = stmt.Exec(m.Name, m.Value)
		if err != nil {
			log.Printf("Error inserting data into ClickHouse: %v\n", err)
			return
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v\n", err)
		return
	}

	fmt.Println("Data inserted into ClickHouse successfully")


}