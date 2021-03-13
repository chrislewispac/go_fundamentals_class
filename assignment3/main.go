// Christopher Lewis Assignment #3
// go version go1.15.6 darwin/amd64

package main

import (
	"fmt"
	"os"
)

type stack interface {
	pop() *LLNode;
	push();
	print();
	contains() bool;
}

// LLNode Linked List node
type LLNode struct {
	value int;
	prev *LLNode;
	next *LLNode;
}

// LList Linked List node
type LList struct {
	length int;
	dummyHead *LLNode;
	dummyTail *LLNode;
}

func (ll *LList) pop() *LLNode {
	if ll.length == 0 {
		return nil;
	}

	outNode := ll.dummyHead.next;
	ll.dummyHead.next = ll.dummyHead.next.next

	ll.length--;
	return outNode;
}

func (ll *LList) push(value int) {
	newNode := LLNode{
		value: value,
	}

	if ll.length == 0 {
		ll.dummyHead.next = &newNode;
		ll.dummyTail.prev = &newNode;
	} else {
		newNode.next = ll.dummyHead.next;
		ll.dummyHead.next.prev = &newNode;
		ll.dummyHead.next = &newNode;
	}

	ll.length++;
}

func (ll *LList) print() {
	if ll.length == 0 {
		return;
	}

	curr := ll.dummyHead.next;

	for curr != nil {

		if curr.next != nil {
			fmt.Print(curr.value, " -> ");
		} else {
			fmt.Print(curr.value);
		}
		
		curr = curr.next;
	}

	fmt.Println("");
}

func (ll *LList) contains(target int) bool {
	if ll.length == 0 {
		return false;
	}

	curr := ll.dummyHead.next;

	for curr != nil {
		if curr.value == target {
			return true;
		}
		curr = curr.next;
	}

	return false;
}




func main() {

	ticketNum := 0;

	alleyA := LList{
		length: 0,
		dummyHead: &LLNode{},
		dummyTail: &LLNode{},
	};

	alleyB := LList{
		length: 0,
		dummyHead: &LLNode{},
		dummyTail: &LLNode{},
	};

	for {
        var input string;
		fmt.Print("D) isplay P) ark R) etrieve Q) uit: ");
        fmt.Scanln(&input);

		if input == "D" || input == "d" {
			fmt.Print("Alley A: ");
			alleyA.print();
			fmt.Println();

		} else if input == "P" || input == "p" {
			ticketNum++;
			if alleyA.length == 5 {
				fmt.Println("My Lot is Full");
				continue;
			} 

			alleyA.push(ticketNum);
			fmt.Println("Issued ticket number", ticketNum, "to customer and parked car.")
			
		} else if input == "R" || input == "r" {
			if alleyA.length == 0 {
				fmt.Println("Lot is Empty, No cars to Retrieve");
				continue;
			}

			var stubNum int;
			fmt.Print("Please provide your stub number: ");
			fmt.Scanln(&stubNum);

			if !alleyA.contains(stubNum) {
				fmt.Println("CAR NOT PARKED IN MY LOT.");
				continue;
			} 

			node := alleyA.pop()

			for node.value != stubNum {
				alleyB.push(node.value);
				node = alleyA.pop();
			}

			fmt.Println("Successfully retrieved car #", node.value);

			for alleyB.length > 0 {
				replaceNode := alleyB.pop()
				alleyA.push(replaceNode.value);
			}
		} else if input == "Q" || input == "q" {
			fmt.Println("Good Bye!");
			os.Exit(1);
		} else {
			fmt.Println("Invalid input, please choose from the menu.")
		}
    }
}