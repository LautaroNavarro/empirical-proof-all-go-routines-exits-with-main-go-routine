
package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
)


func writeIntoFile(){
	f, _ := os.Create("data.txt")
	defer f.Close()
	counter := 0
	for {
		time.Sleep(100 * time.Millisecond)
    	bs := []byte(strconv.Itoa(counter))
		f.Write(bs)
		f.WriteString("\n")
		counter = counter + 1
	}
}

func main() {
	fmt.Println("Starting go routine")
	go writeIntoFile()
	fmt.Println("Goroutine triggered")
	time.Sleep(20 * time.Second)
}