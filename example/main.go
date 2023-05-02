package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/buraktas0/go-event-management/event"
)

var eventManager *event.EventManager

func main() {
	eventManager = event.NewEventManager()

	eventManager.AddListener(MessageSent, NotifyMessageSentEvent)
	eventManager.AddListener(MessageProcessed, NotifyMessageProcessedEvent)
	eventManager.AddListener(MessageSent, NotifyMessageSentEvent)

	fmt.Println("Starting...")

	go eventManager.Trigger(MessageSent, "Message to be sent")
	go eventManager.Trigger(MessageSent, "Message to be sent")

	time.Sleep(1 * time.Second)
}

func NotifyMessageSentEvent(p event.EventPayload) {
	fmt.Println("NotifyMessageSentEvent payload is received:", p)
	ProcessMessage(p.(string))
}

func NotifyMessageProcessedEvent(p event.EventPayload) {
	fmt.Println("NotifyMessageProcessedEvent payload is received:", p)
}

func ProcessMessage(m string) {
	fmt.Println("Message is being processed...")
	parts := strings.Split(m, " ")
	eventManager.Trigger(MessageProcessed, parts[0])
}
