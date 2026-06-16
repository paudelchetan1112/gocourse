package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	///specific time
	specificTime:=time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	//Parse Time
	parsedTime,_:=time.Parse("2006-01-02","2020-05-01" ) //Mon Jan 2 2006 15:04:05 MST 2006
	parsedTime1,_:=time.Parse("06-01-02","20-05-01") ////Mon Jan 2 2006 15:04:05 MST 2006
	fmt.Println(parsedTime1)

	fmt.Println(parsedTime);

t:=time.Now()
// fmt.Println("Formatted Time", t.Format("06-01-02 04-15-05"))
// oneDayLater:=t.Add(time.Hour * 24)
// fmt.Println(oneDayLater)
// fmt.Println(oneDayLater.Weekday())
// fmt.Println(t.Weekday())

// fmt.Println("RoundTime:", t.Round(time.Hour))
// loc, _:=time.LoadLocation("Asia/kathmandu")
// fmt.Println(loc)
// tLocal:=t.In(loc)


// //perform rounding
// roundedTime:=t.Round(time.Hour)

// roundedTimeLocal:=roundedTime.In(loc)
// fmt.Println("Orginal Time (UTC):", t)
// fmt.Println("Orginal Time (UTC):", tLocal)
// fmt.Println("Orginal Time (UTC):", roundedTime)
// fmt.Println("Orginal Time (UTC):", roundedTimeLocal)


fmt.Println("Truncated Time:", t.Truncate(time.Hour))
loc, _:=time.LoadLocation("America/New_York")
tInNy:=time.Now().In(loc)
fmt.Println(loc, tInNy)

t1:=time.Date(2024, time.July, 4, 12 , 0, 0, 0, time.UTC)
t2:=time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
fmt.Println(t2==t1)

}