package main

import "fmt"

func main() {
	ids := []int{33, 25, 46, 89, 550}
	// loop over ids
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}
	// not using index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}
	// add ids to together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// range with map
	emails := map[string]string{"bob": "bob@gmil.com"}
	for key, value := range emails {
		fmt.Printf("%s : %s\n", key, value)
	}
	for k := range emails {
		fmt.Println("name is " + k)
	}
}
