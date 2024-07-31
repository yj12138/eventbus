package eventbus

import (
	"fmt"
	"reflect"
)

type IEvent interface {
	GetEventId() int32
	AddHandler(cb any) error
}

type Event struct {
	eventId        int32
	eventCallBacks []any
}

func (e *Event) GetEventId() int32 {
	return e.eventId
}

func (e *Event) chechHandler(f any) bool {
	fnType := reflect.TypeOf(f)
	if fnType.Kind() != reflect.Func {
		return false
	}
	if len(e.eventCallBacks) > 0 {
		fnType1 := fnType
		fnType2 := reflect.TypeOf(e.eventCallBacks[0])
		if fnType1.Kind() != reflect.Func || fnType2.Kind() != reflect.Func {
			return false
		}
		if fnType1.NumIn() != fnType2.NumIn() || fnType1.NumOut() != fnType2.NumOut() {
			return false
		}
		for i := 0; i < fnType1.NumIn(); i++ {
			if fnType1.In(i) != fnType2.In(i) {
				return false
			}
		}
		for i := 0; i < fnType1.NumOut(); i++ {
			if fnType1.Out(i) != fnType2.Out(i) {
				return false
			}
		}
	}
	return true
}

func (e *Event) AddHandler(cb any) error {
	if e.chechHandler(cb) {
		e.eventCallBacks = append(e.eventCallBacks, cb)
	} else {
		return fmt.Errorf("EventId: %d , func(%v) is illegal", e.GetEventId(), cb)
	}
	return nil
}

func (e *Event) safeCall(f any, args ...interface{}) {
	fnValue := reflect.ValueOf(f)
	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		return
	}
	if len(args) != fnType.NumIn() {
		return
	}
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValue := reflect.ValueOf(arg)
		if argValue.Type() != fnType.In(i) {
			return
		}
		in[i] = argValue
	}
	fnValue.Call(in)
}
func (e *Event) Trigger() error {
	for _, f := range e.eventCallBacks {
		e.safeCall(f)
	}
	return nil
}
func (e *Event) TriggerOne(a any) error {
	for _, f := range e.eventCallBacks {
		e.safeCall(f, a)
	}
	return nil
}
func (e *Event) TriggerTwo(a any, b any) error {
	for _, f := range e.eventCallBacks {
		e.safeCall(f, a, b)
	}
	return nil
}
func (e *Event) TriggeThree(a any, b any, c any) error {
	for _, f := range e.eventCallBacks {
		e.safeCall(f, a, b, c)
	}
	return nil
}

func NewEvent(id int32) *Event {
	return &Event{
		eventId:        id,
		eventCallBacks: make([]any, 0),
	}
}
