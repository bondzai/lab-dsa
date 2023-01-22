package main

import (
    "fmt"
    "time"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
    // Create an MQTT client
    client := MQTT.NewClient(MQTT.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883"))

    // Connect to the broker
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        return
    }

    // Subscribe to the "golang" topic
    if token := client.Subscribe("golang", 0, func(client MQTT.Client, msg MQTT.Message) {
        fmt.Printf("Received message on topic: %s\n", msg.Topic())
        fmt.Printf("Message: %s\n", msg.Payload())
    }); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        return
    }

	client.Publish("golang", 0, false, "Hello, MQTT")

    time.Sleep(3 * time.Second)

    // Unsubscribe from the "golang" topic
    if token := client.Unsubscribe("golang"); token.Wait() && token.Error() != nil {
        fmt.Println(token.Error())
        return
    }

    // Disconnect from the broker
    client.Disconnect(250)
}
