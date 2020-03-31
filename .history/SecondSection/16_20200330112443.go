package main 


func main() [
	i := flag.Int("int", 0, "int flag")
	a := 24 / *i

	flag.Parse()
]