package main


func twoSum(a, b int) int {
	return a + b
}

func threeSum(a, b, c int) int {
	return a + b + c
}

func main(){

	twoSumResult := twoSum(1, 2)
	threeSumResult := threeSum(1, 2, 3)

	println(twoSumResult)
	println(threeSumResult)

}