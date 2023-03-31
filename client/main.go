package main

import (
	"FinalDesign/Model"
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/extrame/xls"
)

func main() {
	fmt.Println("hello")
	fmt.Println(gcd(10, 15))
	// ExcelTest()
	// var e = Model.Excel{
	// 	ExcelName: "test3.xls",
	// 	Sheets:    []Model.Sheet{},
	// }
	// e.AnalyseExcel(e.ExcelName)
	// e.ShowExcel(7)
	// Model.OpenDatabase(false)
	// Model.TestTable()
	// Model.CloseDatabase()
	e := Model.NewElasticSearch()
	e.Init()
	e.Query()
}

func splitNum(num int) int {
	tmp := map[int]int{}
	n := num
	len := 0
	for n > 0 {
		tmp[n%10]++
		n /= 10
		len++
	}
	ans := 0
	num1 := 0
	num2 := 0
	for i := len; i > 0; i -= 2 {
		for j := 0; j < 10; {
			if tmp[j] != 0 {
				num1 = num1*10 + j
				tmp[j]--
			}
			for tmp[j] == 0 {
				j++
			}
			if tmp[j] != 0 {
				num2 = num2*10 + j
				tmp[j]--
			}
		}
	}
	ans = num1 + num2
	return ans
}

func coloredCells(n int) int64 {
	return int64(2*n*n - 2*n + 1)
}

func countWays(ranges [][]int) int {
	tmp := make([][]int, 0)
	for i, _ := range ranges {
		flag := false
		for j, _ := range tmp {
			if cross(ranges[i], tmp[j]) {
				tmp[j][0] = min(ranges[i][0], tmp[j][0])
				tmp[j][1] = max(ranges[i][1], tmp[j][1])
				flag = true
			}
		}
		if flag == false {
			tmp = append(tmp, ranges[i])
		}
	}
	for i, _ := range tmp {
		for j, _ := range tmp {
			if i != j && cross(tmp[i], tmp[j]) {
				return countWays(tmp)
			}
		}
	}
	ans := 1
	for i := 0; i < len(tmp); i++ {
		ans = (ans * 2) % (1e9 + 7)
	}
	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func cross(num1 []int, num2 []int) bool {
	if num1[0] > num2[1] || num1[1] < num2[0] {
		return false
	}
	return true
}

func findValidSplit(nums []int) int {
	tmp := map[int]bool{}
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if gcd(nums[i], nums[j]) != 1 {
				tmp[i] = true
			}
		}
	}
	for k, v := range tmp {
		if v == true {
			return k
		}
	}
	return -1
}

func gcd(a int, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func ExcelTest() {
	file, err := xls.Open("test4.xls", "utf-8")
	if err != nil {
		log.Fatalf("open excel file err: %v", err)
	}
	fmt.Println(file.GetSheet(0).Row(0).Col(1))
	fmt.Println(file.GetSheet(0).Row(2).Col(1))
	fmt.Println(file.GetSheet(0).Row(1).Col(0))
	fmt.Println(file.GetSheet(1).Row(2).Col(4))
	fmt.Println(file.GetSheet(1).Row(3).Col(4))
	for i := 2; i < 3; i++ {
		sheet := file.GetSheet(i)
		fmt.Println("sheetName:" + sheet.Name)
		for j := 0; j < int(sheet.MaxRow); j++ {
			if sheet.Row(j) == nil {
				continue
			}
			row := sheet.Row(j)
			for index := 0; index < row.LastCol(); index++ {
				// fmt.Printf("LastCol=%d\n", row.LastCol()-row.FirstCol())
				if row.Col((index)) != "" {
					fmt.Printf("[%d][%d]%v ", j, index, row.Col(index))
				}
			}
			fmt.Println()
		}
	}
}

func isA(word byte) bool {
	if word == 'a' || word == 'e' || word == 'i' || word == 'o' || word == 'u' {
		return true
	} else {
		return false
	}
}

func vowelStrings(words []string, left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		word := words[i]
		if isA(word[i]) && isA(word[len(word)-1]) {
			res++
		}
	}
	return res
}

func maxScore(nums []int) int {
	sort.Ints(nums)
	sum := 0
	res := 0
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		if sum > 0 {
			res++
		}
	}
	return res
}

func beautifulSubarrays(nums []int) int64 {
	var res int64
	n := len(nums)
	stk := map[int]int{}
	k := 0
	for i := 1; i < n; i++ {
		k ^= nums[i]
		stk[k]++
	}
	for i := 0; i < n; i++ {
		res += int64(stk[nums[i]])
	}
	return res
}

func evenOddBit(n int) []int {
	res := make([]int, 2)
	index := 0
	for n != 0 {
		if n&1 != 0 {
			if index%2 == 0 {
				res[0]++
			} else {
				res[1]++
			}
		}
	}
	return res
}

func checkValidGrid(grid [][]int) bool {
	dx := []int{1, 2, 2, 1, -1, -2, -2, -1}
	dy := []int{2, 1, -1, -2, -2, -1, 1, 2}
	x, y := 0, 0
	index := 0
	n := len(grid)
	for index < n*n-1 {
		index++
		flag := false
		for i := 0; i < len(dx); i++ {
			mx := x + dx[i]
			my := y + dy[i]
			if mx < 0 || mx >= n || my < 0 || my >= n {
				continue
			}
			if grid[mx][my] == index {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

func beautifulSubsets(nums []int, k int) int {
	ans := 0
	n := len(nums)
	if n == 1 {
		return 1
	}
	for i, x := range nums {
		tmp1 := 0
		tmp2 := 0
		dist := make([]map[int]int, 2)
		index := 0
		dist[index][x]++
		for j := i + 1; j < len(nums); j++ {
			if dist[index][nums[i]-k] == 0 && dist[index][nums[i]+k] == 0 {
				if dist[index][nums[j]-k] == 0 && dist[index][nums[j]+k] == 0 {
					dist[index][nums[j]]++
					tmp1++
				} else {
					dist[index+1][nums[j]]++
					tmp2++
				}
			}
		}
		ans += int(math.Pow(2, float64(tmp1))) + int(math.Pow(2, float64(tmp2)))
	}
	return ans
}
