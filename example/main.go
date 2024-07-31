package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yj12138/eventbus"
)

const (
	EventId_1 int32 = 1
	EventId_2 int32 = 2
	EventId_3 int32 = 3
)

type Data1 struct {
	Val string
}
type Data2 struct {
	Val int
}
type Data3 struct {
	Val bool
}

func main() {
	eventbus.ListenOne(EventId_1, func(data *Data1) {
		fmt.Println("1", data.Val)
	})
	eventbus.ListenOne(EventId_1, func(data *Data1) {
		fmt.Println("2", data.Val)
	})
	eventbus.ListenTwo(EventId_2, func(data1 *Data1, data2 *Data2) {
		fmt.Println("3", data1.Val, data2.Val)
	})
	eventbus.ListenTwo(EventId_2, func(data1 *Data1, data2 *Data2) {
		fmt.Println("4", data1.Val, data2.Val)
	})
	eventbus.ListenThree(EventId_3, func(data1 *Data1, data2 *Data2, data3 *Data3) {
		fmt.Println("5", data1.Val, data2.Val, data3.Val)
	})
	eventbus.ListenThree(EventId_3, func(data1 *Data1, data2 *Data2, data3 *Data3) {
		fmt.Println("6", data1.Val, data2.Val, data3.Val)
	})
	go func() {
		eventbus.EmitOne(EventId_1, &Data1{Val: "Val world"})
		eventbus.EmitTwo(EventId_2, &Data1{Val: "Val world"}, &Data2{Val: 33})
		eventbus.EmitThree(EventId_3, &Data1{Val: "Val world"}, &Data2{Val: 33}, &Data3{Val: true})
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
