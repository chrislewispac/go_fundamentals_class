package main

import (
	"fmt"
	"os"
)

//MAX size of queue
const MAX int = 50;
var rear int = -1;
var front int = -1;
var queue []int = make([]int, MAX)

func main() {
    for {
        var choice int;
        fmt.Println("1.Insert element to queue");
        fmt.Println("2.Delete element from queue");
        fmt.Println("3.Display all elements of queue");
        fmt.Println("4.Quit");
        fmt.Print("Enter your choice : ");
        fmt.Scanln(&choice);
        switch choice {
            case 1:
                insert();
            case 2:
                delete();
            case 3:
                display();
            case 4:
                os.Exit(1)
            default:
                fmt.Println("Wrong choice.");
        }
    }
}

func insert() {
    var addItem int;
    if rear == MAX - 1 {
        fmt.Println("Queue Overflow");
    } else {
        if front == -1 {
            front = 0;
        }
        fmt.Printf("Insert the element in queue: ")
        fmt.Scanln(&addItem);
        rear++
        queue[rear] = addItem;
    }
}

func delete() {
    if front == -1 || front > rear {
        fmt.Println("Queue Underflow");
        return;
    } 
    
    fmt.Printf("Element deleted from queue is : %d\n", queue[front]);
    front++;
}

func display() {
    var i int;
    if front == -1 {
        fmt.Println("Queue is empty")
    } else {
        fmt.Println("Queue is :")
        for i = front; i <= rear; i++ {
            fmt.Printf("%d", queue[i])
        }
        fmt.Println();
    }


}