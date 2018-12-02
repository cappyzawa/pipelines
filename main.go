package main

import (
	"fmt"
	"os"
	"time"

)

func main() {
	nowHour := time.Now().Hour()
	fmt.Fprintf(os.Stdout, "Good %s!!\n", Greet(nowHour))
}

func Greet(nowHour int) string {
	if 0 <= nowHour && nowHour <= 9 {
		return "Morning"
	} else if 10 <= nowHour && nowHour <= 17 {
		return "Afternoon"
	} else {
		return "Evening"
	}
}
