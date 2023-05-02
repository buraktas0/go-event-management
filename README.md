# go-event-management

An event management module written in Golang

## Installation
```
go get "github.com/buraktas0/go-event-management"
```

## Quickstart
* See the [example project](https://github.com/buraktas0/go-event-management/tree/main/example)
## Steps
* Define custom events using `EventType`
* Define handler function that gets an `EventPayload` argument for that event type 
* Create an `EventManager`
* `AddListener` to the `EventManager` for related events using an `EventType` and the event's handler
* `Trigger` the `EventManager` whenever needed passing `EventType` and `EventPayload`
