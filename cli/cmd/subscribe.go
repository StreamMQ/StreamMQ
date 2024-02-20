package cmd

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var SubscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to an MQTT topic",
	Run: func(cmd *cobra.Command, args []string) {
		opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return
		}
		defer client.Disconnect(250)

		fmt.Printf("Subscribing to topic '%s'...\n", topic)
		if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("Received message on topic '%s': %s\n", msg.Topic(), msg.Payload())
		}); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			return
		}

		// Block the program by waiting for a signal (e.g., Ctrl+C) to exit
		select {}
	},
}

func init() {
	SubscribeCmd.Flags().StringVarP(&broker, "broker", "b", "tcp://localhost:1883", "MQTT broker address")
	SubscribeCmd.Flags().StringVarP(&clientID, "clientID", "i", "mqtt-cli", "Client ID for MQTT connection")
	SubscribeCmd.Flags().StringVarP(&topic, "topic", "t", "", "MQTT topic to subscribe to")
	SubscribeCmd.MarkFlagRequired("topic")
}