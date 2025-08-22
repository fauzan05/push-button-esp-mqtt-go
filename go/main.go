package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// ButtonMessage represents the structure of the message from ESP8266
type ButtonMessage struct {
	Device    string `json:"device"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

// MQTT configuration
const (
	broker   = "localhost"
	port     = 1883
	topic    = "esp8266/button"
	clientID = "golang-mqtt-client"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("\n=== NEW MESSAGE RECEIVED ===\n")
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Raw Message: %s\n", string(msg.Payload()))

	// Parse JSON message
	var buttonMsg ButtonMessage
	if err := json.Unmarshal(msg.Payload(), &buttonMsg); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Display parsed data
	fmt.Printf("Device: %s\n", buttonMsg.Device)
	fmt.Printf("Action: %s\n", buttonMsg.Action)
	fmt.Printf("ESP8266 Timestamp: %d ms\n", buttonMsg.Timestamp)
	fmt.Printf("Status: %s\n", buttonMsg.Status)
	fmt.Printf("Received at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("============================")

	// You can add your business logic here
	handleButtonPress(buttonMsg)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("✅ Connected to MQTT broker!")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("❌ Connection lost: %v\n", err)
}

func handleButtonPress(msg ButtonMessage) {
	// Add your custom logic here
	fmt.Printf("🔘 Processing button press from %s\n", msg.Device)

	// Example: Count button presses, save to database, trigger other actions, etc.
	// For now, just a simple response
	switch msg.Action {
	case "button_pressed":
		fmt.Println("🎉 Button was pressed! Executing custom logic...")
		// Add your custom code here
		// e.g., save to database, send notification, control other devices
	default:
		fmt.Printf("⚠️ Unknown action: %s\n", msg.Action)
	}
}

func main() {
	fmt.Println("🚀 Starting Golang MQTT Button Receiver...")

	// MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// Create and start the client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", token.Error())
	}

	// Subscribe to the topic
	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to topic: %v", token.Error())
	}
	fmt.Printf("📡 Subscribed to topic: %s\n", topic)
	fmt.Println("⏳ Waiting for button presses... (Press Ctrl+C to exit)")

	// Wait for interrupt signal to gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("\n🛑 Shutting down...")
	client.Disconnect(250)
	fmt.Println("✅ Goodbye!")
}
