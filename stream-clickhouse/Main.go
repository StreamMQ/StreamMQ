package main

import (
	"fmt"
	"log"

	"github.com/StreamMQ/StreamMQ/stream-clickhouse/services"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	mqttBroker = "tcp://localhost:1883"
	mqttTopic  = "measurements"
)

func main() {
	log.Printf("Starting Mqtt To Clickhouse Streaming...")

	// Create a new MQTT client options object
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker).SetClientID("mqtt-to-clickhouse")

	// Set the callback function for receiving messages
	opts.SetDefaultPublishHandler(services.OnMessageReceived)

	// Create a new MQTT client instance
	client := mqtt.NewClient(opts)

	// Connect to the MQTT broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v\n", token.Error())
	}
	fmt.Println("Connected to MQTT broker")

	// Subscribe to the MQTT topic
	if token := client.Subscribe(mqttTopic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to MQTT topic: %v\n", token.Error())
	}
	fmt.Printf("Subscribed to MQTT topic: %s\n", mqttTopic)

	// Keep the program running
	select {}
}
