package main

import "fmt"

func calculate(i int, n int) int {

    var number int

    if(i <= 0) {
	return 0
    }	

   // fmt.Println("Enter # ", n-i+1);

    fmt.Scan(&number)

    if number < 0 {
	 number = 0
    } else {
	number = number*number
    }

    return number + calculate(i-1, n)
}

func make_set(i int, N int){
    if i <= 0{
	 return
    }

    var n int

//    fmt.Println("Enter number of elements in set ",N-i+1," : ")
    fmt.Scan(&n)
//    fmt.Println()

    ans := calculate(n, n)

//    fmt.Println("Answer for set ", N-i+1," is : ",ans)
      fmt.Println(ans)
//    fmt.Println()

    make_set(i-1, N)
}


func main(){

    var N int

//    fmt.Println("Enter number of sets: ")
    fmt.Scan(&N)
//    fmt.Println()

    if N <= 100 && N >= 1 {
	 make_set(N, N)
    } else {
	fmt.Println("Only 1 to 100 sets allowed\n")
    }

}