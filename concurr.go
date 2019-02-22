package main

import (  
    "fmt"
    "time"
    "sync"
)

func hello() {  
    fmt.Println("Hello world goroutine")
}

func print(done chan bool) {
	fmt.Println("Hello world function print")
	done <- true// write to channel done
}


/*
data := <- a // read from channel a  
a <- data // write to channel a  
*/

func sum( num1 int, num2 int, sum chan int){
	sum <- num1 + num2
}

func mul( num1 int, num2 int, mul chan int){
	mul <- num1 * num2
}


func server1(ch chan string) {  
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {  
    time.Sleep(3 * time.Second)
    ch <- "from server2"

}


var x  = 0  
func increment(wg *sync.WaitGroup, m2 *sync.Mutex) {  
    m2.Lock()
    x = x + 1
    m2.Unlock()
    wg.Done()   
}


func main() {  
    go hello()//Goroutines
    fmt.Println("main function")
    time.Sleep(1 * time.Second)


    done := make(chan bool)//channels
    go print(done)
    <-done 
    fmt.Println("goroutine print done/not")


    csum := make(chan int)
    cmul := make(chan int)

    go sum(1,2,csum)
    s :=<- csum
    go mul(2,3,cmul)
    m :=<- cmul

    fmt.Println(s + m)


    bufch := make(chan string , 2)
    bufch <- "a"
    bufch <- "b"


    fmt.Println(<-bufch)    
    fmt.Println(<-bufch)


    // select
    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }

    //mutex
    var w sync.WaitGroup
    var mu sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w, &mu)
    }
    w.Wait()
    fmt.Println("final value of x", x)



}

/*

Mutex vs Channels
We have solved the race condition problem using both mutexes and channels. So how do we decide what to use when. The answer lies in the problem you are trying to solve. If the problem you are trying to solve is a better fit for mutexes then go ahead and use mutex. Do not hesitate to use mutex if needed. If the problem seems to be a better fit for channels, then use it :).

Most Go newbies try to solve every concurrency problem using a channel as it is a cool feature of the language. This is wrong. The language gives us the option to either use Mutex or Channel and there is no wrong in choosing either.

In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine should access the critical section of code.

In the case of the problem which we solved above, I would prefer to use mutex since this problem does not require any communication between the goroutines. Hence mutex would be a natural fit.

My advice would be to choose the tool for the problem and do not try to fit the problem for the tool :)

This brings us to an end of this tutorial. Have a great day.

*/






