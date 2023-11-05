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

    // Publish a message to the "golang" topic
    client.Publish("golang", 0, false, "Hello, MQTT")

    time.Sleep(3 * time.Second)

    // Disconnect from the broker
    client.Disconnect(250)
}
