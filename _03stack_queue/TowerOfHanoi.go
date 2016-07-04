package main

import "fmt"

func main() {
	size := 4
	var source *[]int = createSource(size)
	var target *[]int = &([]int{})
	var buffer *[]int = &([]int{})

	fmt.Print("Source -> ")
	fmt.Println(source)
	fmt.Print("Target -> ")
	fmt.Println(target)

	MoveDisk(size, source, buffer, target)

	fmt.Print("Source -> ")
	fmt.Println(source)
	fmt.Print("Target -> ")
	fmt.Println(target)

}

func createSource(num int) (*[]int) {
	source := []int{num}
	for i := 1; i < num; i++ {
		source = append(source, (num - i))
	}
	return &source
}

func MoveDisk(n int, source *[]int, buffer *[]int, target *[]int) {
	if n > 0 {
		MoveDisk(n - 1, source, target, buffer)
		moveTop(source, target)
		MoveDisk(n - 1, buffer, source, target)
	}
}

func moveTop(source *[]int, target *[]int) {
	if len(*source) > 0 {
		disk := (*source)[len(*source) - 1]
		*source = (*source)[:(len(*source) - 1)]
		*target = append((*target), disk)
	}
}





