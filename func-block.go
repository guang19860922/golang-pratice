package main	
import "fmt"	
	
func main(){	
	var s string
	var n int
	fmt.Scanln(&s)
	show(s)
	fmt.Println("請輸入一個數字:")
	fmt.Scanln(&n)
	fmt.Println(greeting(n))
}

func show(s string){
	fmt.Print("你輸入的文字是:")
	fmt.Println(s)
}

func greeting(n int){	
	return n*n
}