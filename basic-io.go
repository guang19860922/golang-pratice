package main
import "fmt"
func main(){
	// fmt.Println(3, "Hello", true)
	/* var x int
	fmt.Println("請輸入一個數字")
	fmt.Scanln(&x)
	fmt.Println(x) */
	/* var x, y int
	fmt.Print("請輸入第一個數字")
	fmt.Scanln(&x)
	fmt.Print("請輸入第二個數字")
	fmt.Scanln(&y)
	var result int = x + y
	fmt.Print(result) */
	var x, y int
	fmt.Println("請輸入兩個數字用空格隔開")
	fmt.Scanln(&x, &y)
	var result int = x + y
	fmt.Print(result)
	if true {
		fmt.Println("Yes")
	}
	else{
		fmt.Println("NO")
	}
}