package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50}

	fmt.Println("====[mySlice]====")
	fmt.Println("Data: ", mySlice)
	fmt.Println("Len: ", len(mySlice))
	fmt.Println("Cap: ", cap(mySlice))
	fmt.Println()

	slice1 := mySlice[0:3]
	slice2 := mySlice[3:]

	fmt.Println("====[Slice1]====")
	fmt.Println("Data: ", slice1)
	fmt.Println("Len: ", len(slice1))
	fmt.Println("Cap: ", cap(slice1))
	fmt.Println()

	fmt.Println("====[Slice2]====")
	fmt.Println("Data: ", slice2)
	fmt.Println("Len: ", len(slice2))
	fmt.Println("Cap: ", cap(slice2))
	fmt.Println()

	slice1 = append(slice1, 60)

	fmt.Println("====[Setelah Append ke Slice1]====")
	fmt.Println("Data Slice 1 :", slice1)
	fmt.Println("Len:", len(slice1))
	fmt.Println("Cap:", cap(slice1))
	fmt.Println()
	fmt.Println("Data Slice 2: ", slice2)
	fmt.Println("Len:", len(slice2))
	fmt.Println("Cap:", cap(slice2))
	fmt.Println()

	slice2 = append(slice2, 70)

	fmt.Println("====[Setelah Append ke Slice2]====")
	fmt.Println("Data Slice 1 :", slice1)
	fmt.Println("Len:", len(slice1))
	fmt.Println("Cap:", cap(slice1))
	fmt.Println()
	fmt.Println("Data Slice 2: ", slice2)
	fmt.Println("Len:", len(slice2))
	fmt.Println("Cap:", cap(slice2))
	fmt.Println()

}
