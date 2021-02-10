// Christopher Lewis Assignment #2
// go version go1.15.6 darwin/amd64

package main

import (
	"fmt"
	"os"
)


func createAndInitMatrix(matrixSize int) (twoDim [][]int) {
	outerArray := make([][]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		outerArray[i] = make([]int, matrixSize)
	}
	return outerArray;
}

func printMatrix(matrix [][]int, otherdiag int, input int) {
	fmt.Printf("The number you selected was %d", input);
	fmt.Printf(", and the matrix is:\n\n");
	for row := 1; row <= input; row++ {
		fmt.Printf("     ")
		for col := 1; col <= input; col++ {
			fmt.Printf("%5d", matrix[row][col])
		}
		fmt.Printf(" = %5d\n", matrix[row][0])
	}

	for col := 0; col <= input; col++ {
		fmt.Printf("-----");
	}
	fmt.Printf("\n%5d", otherdiag);
	for col := 1; col <= input; col++ {
		fmt.Printf("%5d", matrix[0][col])
	}
	fmt.Printf("   %5d\n", matrix[0][0])
}

func calculateMatrix(matrix [][]int, input int) ([][]int, int){
	row := 1;
	col := input / 2 + 1;
	otherdiag := 0;

	for value := 1; value <= input*input; value++ {
		if matrix[row][col] > 0 {
			row += 2;
			if row > input {
				row -= input
			}
			col--;
			if (col < 1) {
				col = input
			}
		}
		matrix[row][col] = value;

		matrix[0][col] += value;
		matrix[row][0] += value;
		if row == col {
			matrix[0][0] += value
		}

		if row+col == input + 1 {
			otherdiag += value
		}

		row--;
		if row < 1 {
			row = input;
		}
		col++;
		if col > input {
			col = 1
		}
	}

	return matrix, otherdiag
}

func main() {
			/*                                                                        */
		/* Print introduction of what this program is all about.                  */
		/*                                                                        */
		fmt.Println("\nMagic Squares: This program produces an NxN matrix where");
		fmt.Println("N is some positive odd integer.  The matrix contains the ");
		fmt.Println("values 1 to N*N.  The sum of each row, each column, and ");
		fmt.Println("the two main diagonals are all equal.  Not only does this ");
		fmt.Println("program produces the matrix, it also computes the totals for");
		fmt.Println("each row, column, and two main diagonals.");

		fmt.Println("\nBecause of display constraints, this program can work with");
		fmt.Println("values up to 13 only.");
		fmt.Println("")
	    for {
        var input int;
		fmt.Print("Enter a positive, odd integer (-1 to exit program): ");
        fmt.Scanln(&input);

        switch {
		case (input == -1):
            fmt.Println("Bye Bye");
			os.Exit(1);
        case (input <= 0):
			fmt.Println("Sorry, but the integer has to be positive.");
			fmt.Println("\nEnter a positive, odd integer (-1 to exit program):");
			break;
        case (input > 13):
			fmt.Println("Sorry, but the integer has to be less than 15.");
			fmt.Println("\nEnter a positive, odd integer (-1 to exit program):");
			break;
		case (input % 2 == 0):
			fmt.Println("Sorry, but the integer has to be odd.");
			fmt.Println("\nEnter a positive, odd integer (-1 to exit program):");
			break;
		default:
            matrix := createAndInitMatrix(input + 1);
			var otherdiag int;	
			matrix, otherdiag = calculateMatrix(matrix, input);

			printMatrix(matrix, otherdiag, input)

			fmt.Println("\nEnter a positive, odd integer (-1 to exit program):");
        }
		fmt.Println("Bye bye!")
    }
}