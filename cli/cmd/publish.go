package cmd

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
)

var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a message to an MQTT topic",
	Run: func(cmd *cobra.Command, args []string) {

		if mqttClient == nil {
			opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
			mqttClient = mqtt.NewClient(opts)
			if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
				fmt.Println(token.Error())
				return
			}
			defer mqttClient.Disconnect(250)
		}

		token := mqttClient.Publish(topic, 0, false, message)
		token.Wait()
		fmt.Printf("Published message '%s' to topic '%s'\n", message, topic)
	},
}

func init() {
	PublishCmd.Flags().StringVarP(&broker, "broker", "b", "tcp://localhost:1883", "MQTT broker address")
	PublishCmd.Flags().StringVarP(&clientID, "clientID", "i", "mqtt-cli", "Client ID for MQTT connection")
	PublishCmd.Flags().StringVarP(&topic, "topic", "t", "", "MQTT topic to publish to")
	PublishCmd.Flags().StringVarP(&message, "message", "m", "", "Message to publish")
	PublishCmd.MarkFlagRequired("topic")
	PublishCmd.MarkFlagRequired("message")
}
