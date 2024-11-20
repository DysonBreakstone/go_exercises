package main
import ("fmt")

// int, float32, string, bool

/*
variable names:
	Must start iwth a letter or underscore
	cannot start with a digit
	can only contain alpha-numeric charactersa nd underscores
	case-sensitive
	no limit on length
	cannot contain spaces
	cannot be any go keywords		
*/



func main() {
	// int:
	var myfirstint int = 10
	fmt.Println(myfirstint)
	mysecondint := 10 // no 'var' keyword and can only be called inside a function
	fmt.Println(mysecondint)
	var mythirdint int
	fmt.Println(mythirdint)
	var myfourthint = 10
	fmt.Println(myfourthint)
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a + b + c + d)
	var e, f = 10, "pidgeons"
	fmt.Println(e, f)
	var (
		g int
		h int = 1
		i string = "yaaa"
	)
	fmt.Println(g, h, i)


	// Constants

	const PI = 3.14
	fmt.Println(PI)
	const A int = 1
	fmt.Println(A)
}