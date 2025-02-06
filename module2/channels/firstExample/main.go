package main

import (
	"fmt"
	"time"
)

func main(){
	message1 := make(chan string)
	message2 := make(chan string)

	go func()  {
		for{
			message1 <- "Ilyas Azamat channel 1"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for{
			message2 <- "Zhumabay Dias channel 2"
			time.Sleep(time.Second)
		}
	}()

	
	for{
		select{
		case msg := <- message1:
			fmt.Println(msg)
		case msg := <- message2:
			fmt.Println(msg)
		default:
		}
	}

	//Не правильно работает
	// for{
	// 	fmt.Println(<-message1)
	// 	fmt.Println(<-message2)
	// }

}