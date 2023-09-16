package main
import "fmt"

func main(){
	var a = 10+15

	var b = 10/15

	var c = 10*15

	var d = 10%15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)


	var q = 15>10

	fmt.Println(q)

	var rr = 10

	var ll = 11

	if rr>ll {
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}

	var time1 = 30
	var time2 = 60

	if time2>time1{
		fmt.Println("time2 is greater than time1")
	} else if time1<time2 {
		fmt.Println("time1 is greater than time2")
	} else{
		fmt.Println("equal")
	}

}
