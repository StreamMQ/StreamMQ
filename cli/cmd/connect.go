package cmd

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var broker   string
var clientID string
var topic    string
var	message  string
var mqttClient mqtt.Client

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to the MQTT broker",
	Run: func(cmd *cobra.Command, args []string) {
		opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
		mqttClient = mqtt.NewClient(opts)
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println("Error connecting to broker:", token.Error())
			return
		}
		defer mqttClient.Disconnect(250)

		fmt.Println("Connected to MQTT broker successfully!")
	},
}

func init() {
	ConnectCmd.Flags().StringVarP(&broker, "broker", "b", "tcp://localhost:1883", "MQTT broker address")
	ConnectCmd.Flags().StringVarP(&clientID, "clientID", "i", "mqtt-cli", "Client ID for MQTT connection")
}