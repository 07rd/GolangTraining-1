package main
import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(ch chan string, link string) {
	_, err := http.Get(link)
	if err!=nil {
		ch <- link+"  Might Be DOWN"		
	} else {
	   ch <- link+"   Is Up"
	}
}

func checkGoRoutineLink(link string) {
	_,err := http.Get(link)
	if err!=nil {
		fmt.Println(link,":-  Might Be DOWN")
		
	} else {
		fmt.Println(link,":- Is Up")
	}
}

func main() {
	links := []string {
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.golang.org",
	}

	//buffered channel
	fmt.Println("=========buffered channel===========")
	ch := make(chan string, 2)
	for _, link :=  range links{
	   go checkLink(ch, link)
	   fmt.Println(<-ch)
	   time.Sleep(1*time.Second)
	}

	//unbuffered channel
	fmt.Println("=========Non-buffered channel===========")
	ch_unbuffered := make(chan string)
	for _,link :=  range links{
	   go checkLink(ch_unbuffered, link)
	   time.Sleep(1*time.Second)
	}
	fmt.Println(<-ch_unbuffered)
	fmt.Println(<-ch_unbuffered)
	fmt.Println(<-ch_unbuffered)

	//buffered channel
	fmt.Println("========= Go routine ===========")
	//ch := make(chan string, 2)
	for _, link :=  range links{
	   go checkGoRoutineLink(link)
	   time.Sleep(1*time.Second)
	}
	
}