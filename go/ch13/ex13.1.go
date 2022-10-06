package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "경기도 성남시 ..."
	house.Size = 28
	house.Price = 15.5
	house.Type = "아파트"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원 \n", house.Price)
	fmt.Println("타입:", house.Type)

	var house2 House = House{"서울시 강동구", 28, 9.80, "아파트"}

	fmt.Println(house2)

	var house3 House = House{
		Size: 28,
		Type: "아파트",
	}

	fmt.Println(house3)
}
