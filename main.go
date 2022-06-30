package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var stop = false

func main() {
	fmt.Println("sdfdsfdf")
	// evChan := hook.Start()
	// defer hook.End()

	// for ev := range evChan {

	// 	fmt.Println("hook: ", ev.Kind)
	// }

	// mleft := hook.AddEvent("mleft")
	// if mleft != 0 {
	// 	fmt.Println("you press... ", "mouse left button")
	// }
	// robotgo.MouseSleep = 100

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		fmt.Println("---->w")
		stop = true
		hook.End()
	})

	go moveMouse()

	// for {
	// 	robotgo.MoveSmooth(100, 500)
	// 	time.Sleep(3 * time.Second)
	// 	robotgo.MoveSmooth(500, 300)
	// }

	s := hook.Start()

	<-hook.Process(s)

	// robotgo.MouseSleep = 100

	// robotgo.ScrollMouse(10, "up")
	// robotgo.ScrollMouse(20, "right")

	// robotgo.Scroll(0, -10)
	// robotgo.Scroll(100, 0)

	// robotgo.MilliSleep(100)
	// robotgo.ScrollSmooth(-10, 6)
	// // robotgo.ScrollRelative(10, -100)

	// robotgo.Move(10, 20)
	// robotgo.MoveRelative(0, -10)
	// robotgo.DragSmooth(10, 10)

	// robotgo.Click("wheelRight")
	// robotgo.Click("left", true)
	// robotgo.MoveSmooth(100, 200, 1.0, 10.0)

	// robotgo.Toggle("left")
	// robotgo.Toggle("left", "up")
}

func moveMouse() {
	for {
		robotgo.MoveSmooth(100, 500)
		time.Sleep(3 * time.Second)
		robotgo.MoveSmooth(500, 300)
		if stop {
			return
		}
	}
}
