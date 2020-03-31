package main 

func main () {
	var words []string {"asuka","asumatu","asumatu","asuko","asuofasuko","butachan", "asubu-","asuasu",}
	for i, v := range words {
		if i == i * (len(words)/2) {
			fmt.Println(v + ",")
		}
		fmt.Println(v)
	}
}