package Integer

import "fmt"

func Add(first int , second int) int {
	return first+second
}

func main(){
	fmt.Print("What we have here are numbers")
	fmt.Println(Add(1,2))

}
