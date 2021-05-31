package main
import ( "fmt" )
func test() int {
	arr := [] int {}
	defer func(){
		v := recover()
		fmt.Println("recovered:", v)
	}()
	x := arr[5]
	return x }

func main() {
	result := test()
	fmt.Printf("smooth exit %d", result)
}
