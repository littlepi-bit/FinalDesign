package Controller

// func splitNum(num int) int {
// 	tmp := map[int]int{}
// 	n := num
// 	len := 0
// 	for n > 0 {
// 		tmp[n%10]++
// 		n /= 10
// 		len++
// 	}
// 	ans := 0
// 	num1 := 0
// 	num2 := 0
// 	for i := len; i > 0; i -= 2 {
// 		for j := 0; j < 10; {
// 			if tmp[j] != 0 {
// 				num1 = num1*10 + j
// 				tmp[j]--
// 			}
// 			for tmp[j] == 0 {
// 				j++
// 			}
// 			if tmp[j] != 0 {
// 				num2 = num2*10 + j
// 				tmp[j]--
// 			}
// 		}
// 	}
// 	ans = num1 + num2
// 	return ans
// }

// func coloredCells(n int) int64 {
// 	return int64(2*n*n - 2*n + 1)
// }

// func countWays(ranges [][]int) int {
// 	tmp := make([][]int, 0)
// 	for i, _ := range ranges {
// 		flag := false
// 		for j, _ := range tmp {
// 			if cross(ranges[i], tmp[j]) {
// 				tmp[j][0] = min(ranges[i][0], tmp[j][0])
// 				tmp[j][1] = max(ranges[i][1], tmp[j][1])
// 				flag = true
// 			}
// 		}
// 		if flag == false {
// 			tmp = append(tmp, ranges[i])
// 		}
// 	}
// 	for i, _ := range tmp {
// 		for j, _ := range tmp {
// 			if i != j && cross(tmp[i], tmp[j]) {
// 				return countWays(tmp)
// 			}
// 		}
// 	}
// 	ans := 1
// 	for i := 0; i < len(tmp); i++ {
// 		ans = (ans * 2) % (1e9 + 7)
// 	}
// 	return ans
// }

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func cross(num1 []int, num2 []int) bool {
// 	if num1[0] > num2[1] || num1[1] < num2[0] {
// 		return false
// 	}
// 	return true
// }

// func findValidSplit(nums []int) int {
// 	tmp := map[int]bool{}
// 	for i, _ := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if gcd(nums[i], nums[j]) != 1 {
// 				tmp[i] = true
// 			}
// 		}
// 	}
// 	for k, v := range tmp {
// 		if v == true {
// 			return k
// 		}
// 	}
// 	return -1
// }

// func gcd(a int, b int) int {
// 	if b != 0 {
// 		return gcd(b, a%b)
// 	} else {
// 		return a
// 	}
// }

// func ExcelTest() {
// 	file, err := xls.Open("test4.xls", "utf-8")
// 	if err != nil {
// 		log.Fatalf("open excel file err: %v", err)
// 	}
// 	fmt.Println(file.GetSheet(0).Row(0).Col(1))
// 	fmt.Println(file.GetSheet(0).Row(2).Col(1))
// 	fmt.Println(file.GetSheet(0).Row(1).Col(0))
// 	fmt.Println(file.GetSheet(1).Row(2).Col(4))
// 	fmt.Println(file.GetSheet(1).Row(3).Col(4))
// 	for i := 2; i < 3; i++ {
// 		sheet := file.GetSheet(i)
// 		fmt.Println("sheetName:" + sheet.Name)
// 		for j := 0; j < int(sheet.MaxRow); j++ {
// 			if sheet.Row(j) == nil {
// 				continue
// 			}
// 			row := sheet.Row(j)
// 			for index := 0; index < row.LastCol(); index++ {
// 				// fmt.Printf("LastCol=%d\n", row.LastCol()-row.FirstCol())
// 				if row.Col((index)) != "" {
// 					fmt.Printf("[%d][%d]%v ", j, index, row.Col(index))
// 				}
// 			}
// 			fmt.Println()
// 		}
// 	}
// }

// func isA(word byte) bool {
// 	if word == 'a' || word == 'e' || word == 'i' || word == 'o' || word == 'u' {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func vowelStrings(words []string, left int, right int) int {
// 	res := 0
// 	for i := left; i <= right; i++ {
// 		word := words[i]
// 		if isA(word[i]) && isA(word[len(word)-1]) {
// 			res++
// 		}
// 	}
// 	return res
// }

// func maxScore(nums []int) int {
// 	sort.Ints(nums)
// 	sum := 0
// 	res := 0
// 	n := len(nums)
// 	for i := n - 1; i >= 0; i-- {
// 		sum += nums[i]
// 		if sum > 0 {
// 			res++
// 		}
// 	}
// 	return res
// }

// func beautifulSubarrays(nums []int) int64 {
// 	var res int64
// 	n := len(nums)
// 	stk := map[int]int{}
// 	k := 0
// 	for i := 1; i < n; i++ {
// 		k ^= nums[i]
// 		stk[k]++
// 	}
// 	for i := 0; i < n; i++ {
// 		res += int64(stk[nums[i]])
// 	}
// 	return res
// }

// func evenOddBit(n int) []int {
// 	res := make([]int, 2)
// 	index := 0
// 	for n != 0 {
// 		if n&1 != 0 {
// 			if index%2 == 0 {
// 				res[0]++
// 			} else {
// 				res[1]++
// 			}
// 		}
// 	}
// 	return res
// }

// func checkValidGrid(grid [][]int) bool {
// 	dx := []int{1, 2, 2, 1, -1, -2, -2, -1}
// 	dy := []int{2, 1, -1, -2, -2, -1, 1, 2}
// 	x, y := 0, 0
// 	index := 0
// 	n := len(grid)
// 	for index < n*n-1 {
// 		index++
// 		flag := false
// 		for i := 0; i < len(dx); i++ {
// 			mx := x + dx[i]
// 			my := y + dy[i]
// 			if mx < 0 || mx >= n || my < 0 || my >= n {
// 				continue
// 			}
// 			if grid[mx][my] == index {
// 				flag = true
// 				break
// 			}
// 		}
// 		if !flag {
// 			return false
// 		}
// 	}
// 	return true
// }

// func beautifulSubsets(nums []int, k int) int {
// 	ans := 0
// 	n := len(nums)
// 	if n == 1 {
// 		return 1
// 	}
// 	for i, x := range nums {
// 		tmp1 := 0
// 		tmp2 := 0
// 		dist := make([]map[int]int, 2)
// 		index := 0
// 		dist[index][x]++
// 		for j := i + 1; j < len(nums); j++ {
// 			if dist[index][nums[i]-k] == 0 && dist[index][nums[i]+k] == 0 {
// 				if dist[index][nums[j]-k] == 0 && dist[index][nums[j]+k] == 0 {
// 					dist[index][nums[j]]++
// 					tmp1++
// 				} else {
// 					dist[index+1][nums[j]]++
// 					tmp2++
// 				}
// 			}
// 		}
// 		ans += int(math.Pow(2, float64(tmp1))) + int(math.Pow(2, float64(tmp2)))
// 	}
// 	return ans
// }

// func isPrim(a int) bool {
// 	for i := 2; i*i <= a; i++ {
// 		if a%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func diagonalPrime(nums [][]int) int {
// 	ans := 0
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		if isPrim(nums[i][i]) {
// 			ans = max(ans, nums[i][i])
// 		}
// 		if isPrim(nums[i][n-i]) {
// 			ans = max(ans, nums[i][n-i])
// 		}
// 	}
// 	return ans
// }

// func abs(a, b int64) int64 {
// 	if a > b {
// 		return a - b
// 	}
// 	return b - a
// }

// func distance(nums []int) []int64 {
// 	n := len(nums)
// 	left := make([]int64, n)
// 	right := make([]int64, n)
// 	sumleft := make(map[int]int64)
// 	sumright := make(map[int]int64)
// 	lenleft := make(map[int]int)
// 	lenright := make(map[int]int)
// 	tmpleft := make(map[int]int)
// 	leftIndex := make(map[int]int)
// 	rightIndex := make(map[int]int)
// 	for i, num := range nums {
// 		if lenleft[num] == 0 {
// 			tmpleft[num] = i
// 			leftIndex[num] = i
// 			sumleft[num] += int64(i)
// 			lenleft[num]++
// 			continue
// 		}
// 		left[i] = left[tmpleft[num]] + int64(lenleft[num]*i) - sumleft[num]
// 		lenleft[num]++
// 		sumleft[num] += int64(i)
// 		tmpleft[num] = i
// 	}
// 	tmpright := make(map[int]int)
// 	for i := n - 1; i >= 0; i-- {
// 		a := nums[i]
// 		if lenright[a] == 0 {
// 			tmpright[a] = i
// 			rightIndex[a] = i
// 			sumright[a] += int64(i)
// 			lenright[a]++
// 			continue
// 		}
// 		right[i] = right[tmpright[a]] + sumright[a] - int64(lenright[a]*i)
// 		tmpright[a] = i
// 		sumright[a] += int64(i)
// 		lenright[a]++
// 	}
// 	ans := make([]int64, n)
// 	for i, _ := range nums {
// 		ans[i] = left[i] + right[i]
// 	}
// 	return ans
// }

// func TestExcelize() {
// 	f, err := excelize.OpenFile("./client/test2.xlsx")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()
// 	// 获取工作表中指定单元格的值
// 	sheet := f.GetSheetList()
// 	for _, list := range sheet {
// 		fmt.Println(list)
// 	}
// 	cell, err := f.GetCellValue("L.2 承包人提供主要材料和工程设备一览表(表-21)【市政~", "B2")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(cell)
// }

// // func TestExcelize2() {
// // 	f := excelize.NewFile()
// // 	defer func() {
// // 		if err := f.Close(); err != nil {
// // 			fmt.Println(err)
// // 		}
// // 	}()

// // 	excel := Model.Excel{}
// // 	excel.AnalyseExcel("./client/test4.xls")
// // 	tmp := Model.Sheet{}
// // 	for i, _ := range excel.Sheets {
// // 		tmp = excel.Sheets[i]
// // 		index, err := f.NewSheet(tmp.SheetName)
// // 		if err != nil {
// // 			fmt.Println(err)
// // 			return
// // 		}
// // 		for _,row := range tmp.Row{

// // 		}
// // 	}
// // }

// func TestAgg() {
// 	Model.InitElasticSearch()
// 	aggs := elastic.NewTermsAggregation().Field("ProId")
// 	termQuery := elastic.NewTermQuery("Col1.keyword", "合计")
// 	result, err := Model.GlobalES.Client.Search().Index("sheet2").Aggregation("pro", aggs).Query(termQuery).Size(0).Do(context.Background())
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var tmp Model.Sheet2
// 	for _, item := range result.Each(reflect.TypeOf(tmp)) {
// 		t := item.(Model.Sheet2)
// 		fmt.Printf("%#v\n", t)
// 	}
// 	agg, found := result.Aggregations.Terms("pro")
// 	if found {
// 		for _, bucket := range agg.Buckets {
// 			var m int
// 			m = int(bucket.Key.(float64))
// 			fmt.Printf("%d\n", m)
// 		}
// 	}
// }

// func TestLike() {
// 	Model.InitElasticSearch()
// 	mlt := elastic.NewMoreLikeThisQuery()
// 	tmp := Model.GlobalES.QueryByProjectName("成都")[0]
// 	doc := elastic.NewMoreLikeThisQueryItem().Doc(tmp)
// 	mlt.MinimumShouldMatch("60%").MinDocFreq(2).MaxQueryTerms(100).LikeItems(doc)
// 	//var pro Model.Project
// 	res, err := Model.GlobalES.Client.Search().Profile(true).Human(true).Index("management").Type("project").Query(mlt).Do(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%v\n", res.Profile.Shards[0].Searches[0].Query[0])
// 	// for _, v := range res.Each(reflect.TypeOf(pro)) {
// 	// 	t := v.(Model.Project)
// 	// 	fmt.Printf("项目名称：%s, 分数：%d", t.ProjectName, res.Profile.Shards[0].Searches[0].Query[0].Breakdown["score"])
// 	// }
// }

// func TestGroup() {
// 	Model.OpenDatabase(false)
// 	defer Model.CloseDatabase()
// 	var result []myResult1
// 	res := Model.GlobalConn.Table("sheet2").
// 		Select("sheet2.pro_name,project.total_cost_lower as price, sum(col6) as measure_price").
// 		Joins("left join project on project.project_name=sheet2.pro_name").
// 		Where("sheet2.col1=?", "合计").Group("sheet2.pro_name").Find(&result)
// 	//fmt.Printf("%v\n", result)
// 	fmt.Println(res.Error)
// 	var result2 []myResult2
// 	res = Model.GlobalConn.Table("unit").
// 		Select("unit.pro_name, sum(fees) as rule_price").
// 		Joins("left join project on project.project_name=unit.pro_name").
// 		Group("unit.pro_name").Find(&result2)
// 	fmt.Println(res.Error)
// 	tmp := make(map[string]myResult2)
// 	for _, v := range result2 {
// 		tmp[v.ProName] = v
// 	}
// 	for _, v := range result {
// 		v.Total, _ = strconv.ParseFloat(v.Price[:len(v.Price)-3], 64)
// 		v.MeasurePresent = v.MeasurePrice / v.Total
// 		v.RulePrice = tmp[v.ProName].RulePrice
// 		v.RulePresent = v.RulePrice / v.Total
// 		fmt.Println(v)
// 	}
// }

// type myResult1 struct {
// 	ProName        string
// 	Total          float64
// 	Price          string
// 	MeasurePrice   float64
// 	RulePrice      float64
// 	MeasurePresent float64
// 	RulePresent    float64
// }

// type myResult2 struct {
// 	ProName   string
// 	RulePrice float64
// }
