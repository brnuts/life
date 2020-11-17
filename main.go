package main

import (
    tm "github.com/buger/goterm"
    "time"
)
func main() {
    tm.Clear() // Clear current screen
    for {
        // By moving cursor to top-left position we ensure that console output
        // will be overwritten each time, instead of adding new.
        tm.MoveCursor(1, 1)
        for x:=6; x<=100; x++ {
                tm.MoveCursor(x, 1)
                tm.Printf("*")
                time.Sleep(time.Second/100)
                tm.MoveCursor(x-10, 1)
                tm.Printf(" ")
                tm.Flush()
        }
        tm.Println("Current Time:", time.Now().Format(time.RFC1123))
        tm.Flush() // Call it every time at the end of rendering
        time.Sleep(time.Second)
    }
}
