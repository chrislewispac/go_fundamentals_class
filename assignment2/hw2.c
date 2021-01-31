/* MAGIC SQUARE - An NxN matrix containing values from 1 to N*N that are  */
/* ordered so that the sum of the rows, columns, and the major diagonals  */
/* are all equal.  There is a particular algorithm for odd integers, and  */
/* this program constructs that matrix, up to 13 rows and columns.  This   */
/* program also adds the sums of the row, columns, and major diagonals.   */

#include <stdio.h>

main()
{
   int input;                               /* User defined integer       */
   int loc[14][14];                         /* Array holding all          */
                                            /*  input*input values.       */
   int row;                                 /* Determines row of matrix   */
   int col;                                 /* Determines col of matrix   */
   int value;                               /* Integer between 1 and      */
                                            /*  input*input               */
   int otherdiag;                            /* Total of one matrix diagonal*/

/*                                                                        */
/* Print introduction of what this program is all about.                  */
/*                                                                        */
   printf("\nMagic Squares: This program produces an NxN matrix where\n");
   printf("N is some positive odd integer.  The matrix contains the \n");
   printf("values 1 to N*N.  The sum of each row, each column, and \n");
   printf("the two main diagonals are all equal.  Not only does this \n");
   printf("program produces the matrix, it also computes the totals for\n");
   printf("each row, column, and two main diagonals.\n");

   printf("\nBecause of display constraints, this program can work with\n");
   printf("values up to 13 only.\n\n");

    printf("Enter a positive, odd integer (-1 to exit program):\n");
   while (scanf("%d",&input) == 1)
   {

/*                                                                        */
/*    If input = -1, then exit program.                                   */
/*                                                                        */
      if (input == -1)
	 break;
/*                                                                        */
/*    Validity check for input: Must be a positive odd integer < 13.      */
/*                                                                        */
      if (input <= 0)
      {
         printf("Sorry, but the integer has to be positive.\n");
	 printf("\nEnter a positive, odd integer (-1 to exit program):\n");
  	 continue;
      }
      if (input > 13)
      {
	 printf("Sorry, but the integer has to be less than 15.\n");
	 printf("\nEnter a positive, odd integer (-1 to exit program):\n");
	 continue;
      }
      if (input%2 == 0)
      {
	 printf("Sorry, but the integer has to be odd.\n");
	 printf("\nEnter a positive, odd integer (-1 to exit program):\n");
	 continue;
      }
/*                                                                        */
/*    Initialize matrix, row, col, and otherdiag                          */
/*                                                                        */
      for (row = 0; row <= input; row++)    /* Initialize matrix with     */
         for (col = 0; col <= input; col++) /*  all zeroes.               */
            loc[row][col] = 0;              /* Values will reside within  */
                                            /*  rows 1 to input*input and */
                                            /*  columns 1 to input*input. */
                                            /* Row totals will reside in  */
                                            /*  loc[row][0], where row is */
                                            /*  the row number, while the */
                                            /*  column totals will reside */
                                            /*  in loc[0][col], where col */
                                            /*  is the column number.     */  
      row = 1;                              /* First value gets to sit on */
      col = input/2 + 1;                    /*  1st row, middle of matrix.*/
      otherdiag = 0;

 /*                                                                        */
/*    Loop for every value up to input*input, and position value in matrix*/
/*                                                                        */
      for (value = 1; value <= input*input; value++)
      {                                     /* Loop for all values.       */
         if (loc[row][col] > 0)             /* If some value already      */
         {                                  /*  present, then             */
            row += 2;                       /*  move down 1 row of prev.  */
            if (row > input)                /*  If exceeds side, then     */    
               row -= input;                /*   go to other side.        */

            col--;                          /*  move left 1 column.       */
            if (col < 1)                    /*  If exceeds side, then     */
               col = input;                 /*   go to other side.        */
         }

         loc[row][col] = value;             /* Assign value to location.  */

/*                                                                        */
/*       Add to totals                                                    */ 
/*                                                                        */
         loc[0][col] += value;              /* Add to its column total.   */
         loc[row][0] += value;              /* Add to its row total.      */
         if (row == col)                    /* Add to diagonal total if   */
            loc[0][0] += value;             /*  it falls on the diagonal. */

         if (row+col == input+1)            /* Add to other diagonal if   */
            otherdiag += value;             /*  it falls on the line.     */

/*                                                                        */
/*       Determine where new row and col are                              */
/*                                                                        */
         row--;
         if (row < 1)                       /* If row exceeds side then   */
            row = input;                    /*  goto other side.          */
         col++;
         if (col > input)                   /* If col exceeds side then   */
            col = 1;                        /*  goto other side.          */
      }                                     /* End of getting all values. */

/*                                                                        */
/*    Print out the matrix with its totals                                */
/*                                                                        */
      printf("\nThe number you selected was %d",input);
      printf(", and the matrix is:\n\n");
      for (row = 1; row <=input; row++)     /* Loop: print a row at a time*/
      {
         printf("     ");                   /* Create column for diag.total*/
         for (col = 1; col <=input; col++)
            printf("%5d",loc[row][col]);    /* Print values found in a row*/
         printf(" = %5d\n",loc[row][0]);    /* Print total of row.        */
      }

 /*                                                                        */
/*    Print out the totals for each column, starting with diagonal total. */
/*                                                                        */
      for (col = 0; col <=input; col++)     /* Print line separating the  */
         printf("-----");                   /*  value matrix and col totals*/
      printf("\n%5d",otherdiag);            /* Print out the diagonal total*/
      for (col = 1; col <=input; col++)
         printf("%5d",loc[0][col]);         /* Print out the column totals*/
      printf("   %5d\n",loc[0][0]);         /* Print out the other diagonal*/
                                            /*  total                     */
      printf("\nEnter a positive, odd integer (-1 to exit program):\n");

   }                                        /* End of while input>-1 loop */
   printf("\nBye bye!\n");
}

