package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)


var wg sync.WaitGroup

func countUptoTen(s string){
	for i := 0; i<=10;i++ {
		fmt.Printf("%s - %d\n", s,i)
		time.Sleep(time.Millisecond * 200)
	}	

	wg.Done()
}


func countDowntoZero(s string){
	for i := 10; i>=0;i-- {
		fmt.Printf("%s - %d\n", s,i)
		time.Sleep(time.Millisecond * 500)
	}	

	wg.Done()
}


func main() {
		fmt.Println("CPU's - ", runtime.NumCPU())
		fmt.Println("GoRoutines - ", runtime.NumGoroutine())

		wg.Add(2)

		go countUptoTen("hello")
		go countDowntoZero("Goodbye")

		wg.Wait()
		fmt.Println("GoRoutines - ", runtime.NumGoroutine())

}
