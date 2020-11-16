package main
import "fmt"

// func main() {
// 	const (
// 		a = iota
// 		b
// 		c
// 		d = "ha"
// 		e
// 		f = 100
// 		g 
// 		h = iota 
// 		i 
// 	)
// 	fmt.Println(a,b,c,d,e,f,g,h,i)
// }

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}