package two_number_sum

/*
! Two Number Sum
	Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum.
	If any two numbers in the input array sum up to the target sum, the function should return them in an array, in any order.
	If no two numbers sum up to the target sum, the function should return an empty array.
	Note that the target sum has to be obtained by summing two different integers in the array;
	you can't add a single integer to itself in order to obtain the target sum.
	You can assume that there will be at most one pair of numbers summing up to the target sum.

! Example:
	input: array = [3, 5, -4, 8, 11, 1, -1, 6] targetSum = 10
	output: [-1, 11] // the numbers could be in reverse order
*/

func TwoNumberSum(array []int, target int) (result []int) {
	m := map[int]struct{}{}

	for _, el := range array {
		x := target - el

		if _, ok := m[x]; ok {
			return append(result, el, x)
		}

		m[el] = struct{}{}
	}

	return
}
