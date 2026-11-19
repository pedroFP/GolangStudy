package main

import "fmt"

type Product struct {
	Id    int
	Price float64
	Title string
}

func main() {
	// prices := []float64{10.99, 8.99}
	//
	// // fmt.Println(prices)
	//
	// // fmt.Println(prices[0:1])
	//
	// // prices[1] = 9.99
	// // prices[2] = 1.12
	//
	// prices = append(prices, 10.12)
	//
	// fmt.Println(prices)
	//
	// prices = append(prices[:1], prices[1+1:]...)
	//
	// fmt.Println(prices)

	hobbies := [3]string{"Climbing", "Guitar", "Reading"}

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println("hobbies[1:3]", hobbies[1:3])

	newList := hobbies[1:]
	fmt.Println(newList)

	newSlice := hobbies[0:1]
	// newSlice := hobbies[:1]
	fmt.Println(newSlice)

	// newSlice[1] = hobbies[1]
	// newSlice[2] = hobbies[2]

	newSlice = append(newSlice, hobbies[1], hobbies[2])

	fmt.Println(newSlice)

	newDynamicArray := []string{"my", "goals"}
	newDynamicArray[1] = "another"
	newDynamicArray = append(newDynamicArray, "another goal")

	var products []Product

	p1 := Product{
		Id:    1,
		Title: "product1",
		Price: 1.90,
	}

	p2 := Product{
		Id:    2,
		Title: "product2",
		Price: 18.90,
	}

	products = append(products, p1, p2)

	fmt.Println(products)
}
