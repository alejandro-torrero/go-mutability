package main

import "fmt"

type group struct {
	id   int
	name string
}

func main() {
	fmt.Println("Start")

	// By default, slices, maps and structs are mutables

	fmt.Println("--------------- Slices --------------")
	// slices
	slice := []string{
		"1", "dos", "drei",
	}

	fmt.Println("Slice state", slice)

	editSlice(slice, 0, "uno")

	fmt.Println("Slice state", slice)

	// Immutable
	fmt.Println("--------------- Int --------------")

	x := 10

	fmt.Println("x state", x)

	editInt(x)

	fmt.Println("x state", x)

	// But, if we use pointers
	editIntPointer(&x)

	fmt.Println("x state", x)

	fmt.Println("--------------- Struct --------------")
	// Structs

	gr := group{
		id:   15,
		name: "Math 101",
	}

	fmt.Println("gr state", gr)

	editGroupName(gr, "Science")

	fmt.Println("gr state", gr)

	editGroupNamePointer(&gr, "science")

	fmt.Println("gr state", gr)

	fmt.Println("--------------- Pointer --------------")

	a := 10
	b := &a

	fmt.Printf("a: %v b: %v\n", a, *b)

	// Now we change a
	a = 15

	fmt.Printf("a: %v b: %v\n", a, *b)

}

func editSlice(slice []string, index int, value string) {
	slice[index] = value
}

func editInt(x int) {
	x = 2
}

func editIntPointer(x *int) {
	// Derrefencing the value
	*x = 2
}

func editGroupName(gr group, name string) {
	gr.name = name
}

func editGroupNamePointer(gr *group, name string) {
	gr.name = name
}
