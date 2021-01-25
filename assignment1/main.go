package main

import (
	"fmt"
	"os"
)

const max int = 50;
var rear int = -1;
var front int = -1;

func main() {
    var choice int;
    for {
        fmt.Println("1.Insert element to queue");
        fmt.Println("2.Delete element from queue");
        fmt.Println("3.Display all elements of queue");
        fmt.Println("4.Quit");
        fmt.Print("Enter your choice : ");
        fmt.Scanln(&choice);
        switch choice {
        case 1:
            fmt.Println("INSERT")
        case 2:
            fmt.Println("DELETE")
        case 3:
            fmt.Println("DISPLAY")
        case 4:
            os.Exit(1)
        default:
            fmt.Println("Wrong choice.");
        }
    }
}

func insert() {

}

func delete() {

}