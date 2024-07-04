//--------------------------------
//#1
//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("Hello, World!")
//}
//--------------------------------



//--------------------------------
//№2
//package main
//import (
//	"fmt"
//)
//
//func sum(a, b int) int {
//	return a + b
//}
//
//func main() {
//	var a int
//	var b int
//	fmt.Scan(&a, &b)
//	fmt.Println(sum(a, b))
//}
//--------------------------------



//--------------------------------
// №3
//package main
//
//import (
//	"fmt"
//)
//
//func IsEven(a int) bool {
//	if a%2 == 0 {
//		return true
//	}
//	return false
//}
//
//func main() {
//	var a int
//	fmt.Scan(&a)
//	if IsEven(a) {
//		fmt.Println("Even")
//		return
//	}
//	fmt.Println("Not even")
//}
//--------------------------------



//--------------------------------
//№4
//package main
//
//import (
//	"fmt"
//)
//
//func Max(a, b, c int) int {
//	return max(a, b, c)
//}
//
//func main() {
//	var a, b, c int
//	fmt.Scan(&a, &b, &c)
//	fmt.Println(Max(a, b, c))
//}
//--------------------------------



//--------------------------------
//#5
//package main
//
//import "fmt"
//
//func fact(a int64) int64 {
//	if a == 1 {
//		return a
//	} else {
//		return a * fact(a-1)
//	}
//}
//
//func main() {
//	var a int64
//	fmt.Scan(&a)
//	fmt.Println(fact(a))
//}
//--------------------------------




//--------------------------------
//#6
//package main
//
//import "fmt"
//
//const vow string = "aeiouyAEIOUY"
//
//func vowel(a rune) bool {
//	for _, b := range vow {
//		if a == b {
//			return true
//		}
//	}
//	return false
//}
//
//func main() {
//	var a string
//	fmt.Scan(&a)
//	if vowel(rune(a[0])) {
//		fmt.Println("Гласная")
//		return
//	}
//	fmt.Println("Не гласная")
//}
//--------------------------------




//--------------------------------
// #7
//package main
//
//import (
//	"fmt"
//)
//
//func isPrime(n int) bool {
//	if n <= 1 {
//		return false
//	}
//	for i := 2; i*i <= n; i++ {
//		if n%i == 0 {
//			return false
//		}
//	}
//	return true
//}
//
//func primes(a int) {
//	for i := 2; i <= a; i++ {
//		if isPrime(i) {
//			fmt.Println(i)
//		}
//	}
//}
//
//func main() {
//	var a int
//	fmt.Scan(&a)
//	primes(a)
//}
//--------------------------------




//--------------------------------
//№8
//package main
//
//import "fmt"
//
//func reverse(a string) {
//	for i := len(a) - 1; i >= 0; i-- {
//		fmt.Print(string(a[i]))
//	}
//}
//
//func main() {
//	var a string
//	fmt.Scan(&a)
//	reverse(a)
//}
//--------------------------------




//--------------------------------
// #9
//package main
//import "fmt"
//
//func summ(mass []int) int {
//	sum := 0
//	for _, value := range mass {
//		sum += value
//	}
//	return sum
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите кол-во чисел массива:")
//	fmt.Scan(&num)
//	var mass []int
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//	fmt.Println(summ(mass))
//}
//--------------------------------




//--------------------------------
// #10
//package main
//
//import "fmt"
//
//type triangle struct {
//	height float64
//	width  float64
//}
//
//func area(height, width float64) float64 {
//	return height * width
//}
//
//func main() {
//	var t triangle
//	fmt.Print("Enter height: ")
//	fmt.Scan(&t.height)
//	fmt.Print("Enter width: ")
//	fmt.Scan(&t.width)
//	fmt.Println("Area of triangle is:", area(t.height, t.width))
//}
//--------------------------------




//--------------------------------
// #11
//package main
//
//import "fmt"
//
//func far(a float64) float64 {
//	return ((9 / 5.0) * a) + 32
//}
//
//func main() {
//	var a, b float64
//	fmt.Scan(&a)
//	b = far(a)
//	fmt.Println(b)
//}
//--------------------------------




//--------------------------------
// #12
//package main
//
//import "fmt"
//
//func main() {
//	var a int
//	fmt.Scan(&a)
//	for i := a; i >= 1; i-- {
//		fmt.Print(i, " ")
//	}
//}
//--------------------------------




//--------------------------------
// #13
//package main
//
//import "fmt"
//
//func length(a string, b int) int {
//	for range a {
//		b++
//	}
//	return b
//}
//
//func main() {
//	var a string
//	var b int = 0
//	fmt.Scan(&a)
//	fmt.Print(length(a, b))
//}
//--------------------------------




//--------------------------------
// #14
//package main
//
//import "fmt"
//
//func main() {
//	var num int
//	fmt.Println("Введите кол-во элементов массива:")
//	fmt.Scan(&num)
//	var mass []string
//	for i := 0; i < num; i++ {
//		var mas string
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//	var a string
//	fmt.Scan(&a)
//	for i := range mass {
//		if mass[i] == a {
//			fmt.Println("YES")
//			return
//		}
//	}
//}
//--------------------------------





//--------------------------------
//№15
//package main
//
//import (
//	"fmt"
//)
//
//func average(arr []int) float64 {
//	if len(arr) == 0 {
//		return 0
//	}
//	sum := 0
//	for _, value := range arr {
//		sum += value
//	}
//	return float64(sum) / float64(len(arr))
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите кол-во элементов массива:")
//	fmt.Scan(&num)
//	var mass []int
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//	avg := average(mass)
//	fmt.Print("Среднее: ", avg)
//}
//--------------------------------




//--------------------------------
//#16
//package main
//
//import "fmt"
//
//func table(num int) {
//	for i := 1; i <= 10; i++ {
//		fmt.Print(num, " * ", i, " = ")
//		fmt.Println(i * num)
//	}
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите число")
//	fmt.Scan(&num)
//	table(num)
//}
//--------------------------------





//--------------------------------
// #17
//package main
//
//import "fmt"
//
//func reverse(s string) string {
//	reversed := ""
//	for i := len(s) - 1; i >= 0; i-- {
//		reversed += string(s[i])
//	}
//	return reversed
//}
//
//func main() {
//	var input string
//	fmt.Scan(&input)
//
//	if reverse(input) == input {
//		fmt.Println("Palindrome")
//	} else {
//		fmt.Println("Not palindrome")
//	}
//}
//--------------------------------




//--------------------------------
// #18
//package main
//
//import (
//	"fmt"
//)
//
//func findMinMax(arr []int) (int, int) {
//	if len(arr) == 0 {
//		return 0, 0
//	}
//
//	min := arr[0]
//	max := arr[0]
//
//	for _, value := range arr {
//		if value < min {
//			min = value
//		}
//		if value > max {
//			max = value
//		}
//	}
//
//	return min, max
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num)
//
//	var mass []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//
//	min, max := findMinMax(mass)
//	fmt.Println("Минимум:", min)
//	fmt.Println("Максимум:", max)
//}
//--------------------------------




//--------------------------------
// #19
//package main
//
//import (
//	"fmt"
//)
//
//func removeElement(arr []int, index int) []int {
//	return append(arr[:index], arr[index+1:]...)
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num)
//
//	var mass []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//
//	var index int
//	fmt.Println("Введите индекс для удаления начиная с 0:")
//	fmt.Scan(&index)
//
//	mass = removeElement(mass, index)
//	fmt.Println("Массив после удаления элемента:", mass)
//}
//--------------------------------





//--------------------------------
//#20
//package main
//
//import (
//	"fmt"
//)
//
//func linearSearch(arr []int, target int) int {
//	for index, value := range arr {
//		if value == target {
//			return index
//		}
//	}
//	return -1
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num)
//
//	var mass []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//
//	var target int
//	fmt.Println("Введите элемент для поиска:")
//	fmt.Scan(&target)
//
//	index := linearSearch(mass, target)
//	if index != -1 {
//		fmt.Print("Элемент найден на индексе: ", index)
//	} else {
//		fmt.Print(-1)
//	}
//}
//--------------------------------





//--------------------------------
// №21
//package main
//
//import (
//	"fmt"
//)
//
//func removeDuplicates(arr []int) []int {
//	var result []int
//	for _, value := range arr {
//		if !contains(result, value) {
//			result = append(result, value)
//		}
//	}
//	return result
//}
//
//func contains(slice []int, item int) bool {
//	for _, v := range slice {
//		if v == item {
//			return true
//		}
//	}
//	return false
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num)
//
//	var mass []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//
//	uniqueMass := removeDuplicates(mass)
//	fmt.Println("Массив без дубликатов:", uniqueMass)
//}
//--------------------------------





//--------------------------------
// #22
//package main
//
//import (
//	"fmt"
//)
//
//func bubbleSort(arr []int) {
//	n := len(arr)
//	for i := 0; i < n; i++ {
//		for j := 0; j < n-i-1; j++ {
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num)
//
//	var mass []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//
//	bubbleSort(mass)
//	fmt.Println("Отсортированный массив:", mass)
//}
//--------------------------------





//--------------------------------
// #23
//package main
//
//import (
//	"fmt"
//)
//
//func generateFibonacci(n int) []int {
//	if n <= 0 {
//		return []int{}
//	}
//	fib := make([]int, n)
//	fib[0] = 0
//	if n > 1 {
//		fib[1] = 1
//	}
//
//	for i := 2; i < n; i++ {
//		fib[i] = fib[i-1] + fib[i-2]
//	}
//
//	return fib
//}
//
//func main() {
//	var num int
//	fmt.Println("Введите количество элементов последовательности Фибоначчи:")
//	fmt.Scan(&num)
//
//	fibSequence := generateFibonacci(num)
//	fmt.Println("Последовательность Фибоначчи:", fibSequence)
//}
//--------------------------------





//--------------------------------
// №24
//package main
//
//import "fmt"
//
//func contains(slice []int, item int) int {
//	cnt := 0
//	for _, v := range slice {
//		if v == item {
//			cnt++
//		}
//	}
//	return cnt
//}
//
//func main() {
//	var num, target int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num)
//
//	var mass []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		mass = append(mass, mas)
//	}
//	fmt.Println("Введите элемент для проверки:")
//	fmt.Scan(&target)
//	fmt.Println("Этот элемент сожержится в массиве", contains(mass, target), "раз")
//}
//--------------------------------





//--------------------------------
// #25
//package main
//
//import (
//	"fmt"
//)
//
//func intersection(arr1, arr2 []int) []int {
//	var result []int
//
//	for _, num1 := range arr1 {
//		for i, num2 := range arr2 {
//			if num1 == num2 {
//				result = append(result, num1)
//				arr2 = append(arr2[:i], arr2[i+1:]...)
//				break
//			}
//		}
//	}
//
//	return result
//}
//
//func main() {
//	var num1, num2 int
//	fmt.Println("Введите количество элементов первого массива:")
//	fmt.Scan(&num1)
//
//	var arr1 []int
//	fmt.Println("Введите элементы первого массива:")
//	for i := 0; i < num1; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		arr1 = append(arr1, mas)
//	}
//
//	fmt.Println("Введите количество элементов второго массива:")
//	fmt.Scan(&num2)
//
//	var arr2 []int
//	fmt.Println("Введите элементы второго массива:")
//	for i := 0; i < num2; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		arr2 = append(arr2, mas)
//	}
//
//	result := intersection(arr1, arr2)
//	fmt.Println("Пересечение массивов:", result)
//}
//--------------------------------





//--------------------------------
// #26
//package main
//
//import (
//	"fmt"
//	"sort"
//	"strings"
//)
//
//func sortStr(s string) string {
//	chars := strings.Split(s, "")
//	sort.Strings(chars)
//	return strings.Join(chars, "")
//}
//
//func areAnags(s1, s2 string) bool {
//	s1 = strings.ToLower(s1)
//	s2 = strings.ToLower(s2)
//
//	if len(s1) != len(s2) {
//		return false
//	}
//
//	sortedS1 := sortStr(s1)
//	sortedS2 := sortStr(s2)
//
//	return sortedS1 == sortedS2
//}
//
//func main() {
//	var s1, s2 string
//	fmt.Println("Введите первую строку:")
//	fmt.Scan(&s1)
//	fmt.Println("Введите вторую строку:")
//	fmt.Scan(&s2)
//
//	if areAnags(s1, s2) {
//		fmt.Println("Анаграммы")
//	} else {
//		fmt.Println("Не анаграммы")
//	}
//}
//--------------------------------





//--------------------------------
// #27
//package main
//
//import (
//	"fmt"
//)
//
//func merge(arr1, arr2 []int) []int {
//	i, j := 0, 0
//	var result []int
//
//	for i < len(arr1) && j < len(arr2) {
//		if arr1[i] < arr2[j] {
//			result = append(result, arr1[i])
//			i++
//		} else {
//			result = append(result, arr2[j])
//			j++
//		}
//	}
//
//	for i < len(arr1) {
//		result = append(result, arr1[i])
//		i++
//	}
//
//	for j < len(arr2) {
//		result = append(result, arr2[j])
//		j++
//	}
//
//	return result
//}
//
//func main() {
//	var num1, num2 int
//	fmt.Println("Введите количество элементов первого массива:")
//	fmt.Scan(&num1)
//	var arr1 []int
//	fmt.Println("Введите элементы первого массива:")
//	for i := 0; i < num1; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		arr1 = append(arr1, mas)
//	}
//
//	fmt.Println("Введите количество элементов второго массива:")
//	fmt.Scan(&num2)
//	var arr2 []int
//	fmt.Println("Введите элементы второго массива:")
//	for i := 0; i < num2; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		arr2 = append(arr2, mas)
//	}
//
//	fmt.Println("Слияние отсортированных массивов:", merge(arr1, arr2))
//}
//--------------------------------





//--------------------------------
// №28
//package main
//
//import "fmt"
//
//func main() {
//	m := map[int]string{}
//	for i := 0; i < 15; i++ {
//		if i%2 == 0 {
//			m[i] = "true"
//		} else {
//			m[i] = "false"
//		}
//	}
//	fmt.Println("Делимость на 2: ")
//	for k, v := range m {
//		fmt.Printf("key: %d, value: %s\n", k, v)
//	}
//}
//--------------------------------





//--------------------------------
// #29
//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func binarySearch(arr []int, target int) int {
//	low := 0
//	high := len(arr) - 1
//
//	for low <= high {
//		mid := (low + high) / 2
//		if arr[mid] == target {
//			return mid
//		} else if arr[mid] < target {
//			low = mid + 1
//		} else {
//			high = mid - 1
//		}
//	}
//	return -1
//}
//
//func main() {
//	var num1, target int
//	fmt.Println("Введите количество элементов массива:")
//	fmt.Scan(&num1)
//	var arr1 []int
//	fmt.Println("Введите элементы массива:")
//	for i := 0; i < num1; i++ {
//		var mas int
//		fmt.Scan(&mas)
//		arr1 = append(arr1, mas)
//	}
//	sort.Ints(arr1)
//
//	fmt.Println("Введите элемент для поиска:")
//	fmt.Scan(&target)
//
//	index := binarySearch(arr1, target)
//	if index != -1 {
//		fmt.Print("Элемент найден на позиции \n", index)
//	} else {
//		fmt.Println("Элемент не найден")
//	}
//}
//--------------------------------






//--------------------------------
// №30
//package main
//
//import (
//	"fmt"
//)
//
//func pop(stack []int) ([]int, int) {
//	if len(stack) == 0 {
//		return stack, 0
//	}
//	value := stack[len(stack)-1]
//	stack = stack[:len(stack)-1]
//	return stack, value
//}
//
//func enqueue(stack1 []int, value int) []int {
//	return append(stack1, value)
//}
//
//func dequeue(stack1, stack2 []int) ([]int, []int, int) {
//	var value int
//	if len(stack2) == 0 {
//		for len(stack1) > 0 {
//			stack1, value = pop(stack1)
//			stack2 = append(stack2, value)
//		}
//	}
//
//	stack2, value = pop(stack2)
//	return stack1, stack2, value
//}
//
//func main() {
//	var stack1, stack2 []int
//	var choice, value int
//
//	for {
//		fmt.Println("1. Добавить в очередь")
//		fmt.Println("2. Убрать из очереди")
//		fmt.Println("3. Закончить набор")
//		fmt.Print("Введите номер команды: ")
//		fmt.Scan(&choice)
//
//		if choice == 1 {
//			fmt.Print("Введите число, которое хотите добавить: ")
//			fmt.Scan(&value)
//			stack1 = enqueue(stack1, value)
//		} else if choice == 2 {
//			if len(stack1) == 0 && len(stack2) == 0 {
//				fmt.Println("Очередь пуста")
//			} else {
//				stack1, stack2, value = dequeue(stack1, stack2)
//				fmt.Println("Убранное число:", value)
//			}
//		} else if choice == 3 {
//			return
//		}
//	}
//}
//--------------------------------