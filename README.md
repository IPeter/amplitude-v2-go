# amplitude-v2-go
Amplitude client for Go

## Installation

	$ go get github.com/IPeter/amplitude-v2-go

## Example

```go
	client := amplitude.New("key")

	client.SendEvents([]amplitude.Event{
		amplitude.Event{
			UserID: "id",
			EventType: "type",
		},
	})
```