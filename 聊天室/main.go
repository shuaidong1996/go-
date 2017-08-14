package main
import (
	"fmt"
	"time"
)

func main1() {
	now := time.Now()
	hour, minute, second := now.Hour(), now.Minute(), now.Second()
	fmt.Println(now.Unix(),"  ",hour, minute, second, now.Unix() - int64(hour * 3600 + minute * 60 + second))
	fmt.Println(3600 * 22)
}
func main() {
	arr := []int{1,2,3,4,5,6}
	for i,v := range arr {
		fmt.Println(i,v)
	}
}