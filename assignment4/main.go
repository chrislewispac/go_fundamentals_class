// Christopher Lewis Assignment #4
// go version go1.15.6 darwin/amd64

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack interface {
	pop() *LLNode;
	push();
	print();
	contains() bool;
}

// Employee employee
type Employee struct {
	name string;
	age int;
	salary int;
}

// LLNode Linked List node
type LLNode struct {
	value *Employee;
	prev *LLNode;
	next *LLNode;
}

// LList Linked List node
type LList struct {
	length int;
	dummyHead *LLNode;
	dummyTail *LLNode;
}

func (ll *LList) push(employee *Employee) {
	newNode := LLNode{
		value: employee,
	}

	if ll.length == 0 {
		ll.dummyHead.next = &newNode;
		ll.dummyTail.prev = &newNode;
		newNode.prev = ll.dummyHead;
		newNode.next = ll.dummyTail;
	} else {
		curr := ll.dummyHead.next;
		for curr.next != nil && newNode.value.name > curr.value.name {
			curr = curr.next;
		}
		curr.prev.next = &newNode;
		newNode.prev = curr.prev;
		curr.prev = &newNode;
		newNode.next = curr;
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

func (ll *LList) remove() bool {
	node := ll.find()
	if node == nil {
		return false
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	return true
}

func (ll *LList) find() *LLNode {
	var firstName string;          
	var lastName string;                   
	fmt.Print("Please enter employee first name:  ");
	fmt.Scanln(&firstName);
	fmt.Print("Please enter employee last name:  ");
	fmt.Scanln(&lastName);

	if ll.length == 0 {
		return nil;
	}

	curr := ll.dummyHead.next;

	for curr.next != nil {
		if curr.value.name == firstName + " " + lastName {
			fmt.Println(curr.value.name, curr.value.age, curr.value.salary)
			return curr;
		}
		curr = curr.next;
	}

	return nil;
}

func writeFile(ll *LList, fileName string) {
	
    f, err := os.Create(fileName)

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	if ll.length == 0 {
		_, _ = f.WriteString("")
		return;
	}

	curr := ll.dummyHead.next;

	for curr != nil {
		if curr.next != nil {
			_, _ = f.WriteString(curr.value.name + ";" + strconv.Itoa(curr.value.age) + ";" + strconv.Itoa(curr.value.salary) + "\n")
		}
		
		curr = curr.next;
	}
}

func displayEmployees(ll *LList) {
	if ll.length == 0 {
		fmt.Println("Employee list is empty")
	}

	curr := ll.dummyHead.next;

	for curr != nil {

		if curr.next != nil {
			fmt.Println(curr.value.name, curr.value.age, curr.value.salary);
		}
		
		curr = curr.next;
	}

	fmt.Println("");
}


func addEmployee(ll *LList) {
	var firstName string;          
	var lastName string;          
	var a string;          
	var s string;          
	fmt.Print("Please enter employee first name:  ");
	fmt.Scanln(&firstName);
	fmt.Print("Please enter employee last name:  ");
	fmt.Scanln(&lastName);
	fmt.Print("Please enter employee age:  ");
	fmt.Scanln(&a);
	fmt.Print("Please enter employee salary:  ");
	fmt.Scanln(&s);
	age, _ := strconv.Atoi(a) 
	salary, _ := strconv.Atoi(strings.ReplaceAll(s, " ", ""))
	e := &Employee{
		name: firstName + " " + lastName,
		age: age,
		salary: salary,
	}
	ll.push(e)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ERROR: No file name provided as a command line argument")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ll := &LList{
		length: 0,
		dummyHead: &LLNode{},
		dummyTail: &LLNode{},
	};   
	for scanner.Scan() {
		res := strings.SplitN(scanner.Text(), ";", -1)
		age, _ := strconv.Atoi(res[1]) 
		salary, _ := strconv.Atoi(strings.ReplaceAll(res[2], " ", ""))
		e := &Employee{
			name: res[0],
			age: age,
			salary: salary,
		}
		ll.push(e)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
        var input string;      
		fmt.Println("Menu Options:")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Delete Employee ")
		fmt.Println("3. Search Employee ")
		fmt.Println("4. List All Employees")
		fmt.Println("5. Save Employee Database")
		fmt.Println("6. Exit Employee Database")
		fmt.Print("Enter Your Choice: ")
        fmt.Scanln(&input);

		if input == "1" {
			addEmployee(ll);
		} else if input == "2" {
			if ll.remove() {
				fmt.Println("Employee Deleted\n")
			} else {
				fmt.Println("Employee Not Found\n")
			}
		} else if input == "3" {
			if ll.find() == nil {
				fmt.Println("Employee not found \n")
			} 
		} else if input == "4" {
			displayEmployees(ll);
		} else if input == "5" {
			writeFile(ll, fileName)
		} else if input == "6" {
			fmt.Println("Exiting Employee Database")
			os.Exit(1);
		} else {
			fmt.Println("Invalid input, please choose from the menu.\n")
		}
    }
}