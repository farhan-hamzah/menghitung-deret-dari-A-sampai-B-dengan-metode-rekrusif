// Menghitung Sebuah Deret Dari A ke B
package main
import "fmt"
func rekrusi(a, b int)int{
	if a > b {
		return 0
	}
	fmt.Print(a, " ")
	return a + rekrusi(a+1, b)
}

func main(){
	var x, y int
	fmt.Scan(&x, &y)
	rekrusi(x, y)
}
