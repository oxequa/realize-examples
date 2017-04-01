package main

import (
	"fmt"
	"github.com/tockins/realize-examples/scheduler/schedule"
	"time"
)

func main() {
	f := func() { fmt.Println("ciao") }
	schedule.Command(f).Timezone("America/New_York")
	//t := time.Now()
	//fmt.Println("Location:", t.Location(), ":Time:", t)
	//utc, err := time.LoadLocation("America/New_York")
	//if err != nil {
	//	fmt.Println("err: ", err.Error())
	//}
	//fmt.Println("Location:", utc, ":Time:", t.In(utc))
}
