package service

import (
	"time"
)

func CurrentUTime() string {
	return time.Now().String()[:19]
}

func IsStartAfterTimeEnd(start string, end string) bool {
	time1 := StringToTime(start)
	time2 := StringToTime(end)

	return time2.After(time1)
}

func StringToTime(times string) time.Time {
	time1, _ := time.Parse("2006-01-02 15:04:05", times)
	return time1
}

func IsStartBeforeTimeEnd(start string, end string) bool {
	time1 := StringToTime(start)
	time2 := StringToTime(end)

	return time2.Before(time1)
}

func IsStartBeforeTimeEndWithDelayed(start string, end string, duration time.Duration) bool {
	time1 := StringToTime(start)
	// fmt.Println(time1)
	time2 := StringToTime(end).Add(duration)
	// fmt.Println(time2)
	// fmt.Println(time2.Before(time1))

	return time2.Before(time1)
}
