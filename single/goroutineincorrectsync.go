package main

/*var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	println(b)
	println(a)
}

func main() {
	go f()
	g()
	//fmt.Scanln()
}
*/

/*var a string
var done bool
var once sync.Once

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	fmt.Scanln()
}
*/

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
