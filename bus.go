package eventbus

type EventHandlerOne[A any] func(A)
type EventHandlerTwo[A any, B any] func(A, B)
type EventHandlerThree[A any, B any, C any] func(A, B, C)

var (
	eventMap = make(map[int32]*Event, 0)
)

func GetEvent(eventId int32) *Event {
	event, ok := eventMap[eventId]
	if !ok {
		event = NewEvent(eventId)
		eventMap[eventId] = event
	}
	return event
}

func ListenOne[A any](eventId int32, cb EventHandlerOne[A]) error {
	event := GetEvent(eventId)
	return event.AddHandler(cb)
}

func EmitOne[A any](eventId int32, a A) error {
	event := GetEvent(eventId)
	return event.TriggerOne(a)
}

func ListenTwo[A any, B any](eventId int32, cb EventHandlerTwo[A, B]) error {
	event := GetEvent(eventId)
	return event.AddHandler(cb)
}

func EmitTwo[A any, B any](eventId int32, a A, b B) error {
	event := GetEvent(eventId)
	return event.TriggerTwo(a, b)
}

func ListenThree[A any, B any, C any](eventId int32, cb EventHandlerThree[A, B, C]) error {
	event := GetEvent(eventId)
	return event.AddHandler(cb)
}

func EmitThree[A any, B any, C any](eventId int32, a A, b B, c C) error {
	event := GetEvent(eventId)
	return event.TriggeThree(a, b, c)
}
