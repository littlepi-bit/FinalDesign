package main

import (
	"FinalDesign/Model"
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/hankcs/gohanlp/hanlp"
	"github.com/olivere/elastic/v7"
	"github.com/shakinm/xlsReader/xls"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc"

	hp "FinalDesign/protos/hanlp"
)

func main() {
	//TestLike()
	//Model.InitElasticSearch(false)
	//Model.GlobalES.GetRelevantDocByProName("北川羌族自治县传染病专科医院集中隔离留观中心项目")
	//Model.GlobalES.GetRelevanceByProName("北川", "成都德川友邦印务有限公司新建厂区一期项目")
	//TestDoc()
	//TestTermQuery()
	//TestQueryMatch()
	//TestNewExcel()
	// TestNlp()
	//TestDeleteInfo()
	// TestInfo()
	// TestGraph()
	TestSmartSearch()
	// TestPyRpc()
	//fmt.Println(len(Model.UnitProjectRows["单价措施项目清单与计价表"]))
	//TestSplitExcel()
}

func TestDeleteInfo() {
	Model.InitElasticSearch(false)
	termQuery := elastic.NewTermQuery("ProName.keyword", "成都德川友邦印务有限公司新建厂区一期项目")
	boolQuery := elastic.NewBoolQuery()
	boolQuery.MustNot(termQuery)
	res, err := Model.GlobalES.Client.DeleteByQuery("info").Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(res.Deleted)
}

func TestPyRpc() {
	//连接服务器
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Println("连接服务器失败", err)
	}
	defer conn.Close()

	cli := hp.NewHanlpServerClient(conn)

	//远程调用方法
	reply, err := cli.Similarity(context.Background(), &hp.HanlpRequest{Search: "甲方为阳光建筑公司的项目有哪些"})
	if err != nil {
		log.Println("服务器错误：", err)
		return
	}
	fmt.Println("reply=", reply)
}

func TestSmartSearch() {
	Model.InitElasticSearch(false)
	ans, _ := Model.SmartSearch("成都德川友邦印务有限公司新建厂区一期项目的建筑类型是什么")
	fmt.Println(ans)
}

func TestInfo() {
	Model.InitElasticSearch(false)
	info := Model.GlobalES.GetAllInfo()
	fmt.Println(info)
}

func TestGraph() {
	Model.InitElasticSearch(false)
	graph := Model.GlobalES.GetGraph()
	fmt.Println(graph)
}

func PrintCon(cons []hanlp.ConTuple) {
	if cons == nil {
		return
	}
	for _, con := range cons {
		fmt.Println("key = " + con.Key)
		PrintCon(con.Value)
	}
}

func PrintRes(s *hanlp.HanResp) {
	fmt.Println("Con:")
	//PrintCon(s.Con)
	fmt.Println(s)
	// fmt.Println("Dep:")
	// fmt.Println(s.Dep)
	fmt.Println("NerPku")
	fmt.Println(s.NerPku)
	fmt.Println("NerMsra")
	fmt.Println(s.NerMsra)
	fmt.Println("Pos863")
	fmt.Println(s.Pos863)
	fmt.Println("s.PosPku")
	fmt.Println(s.PosPku)
	fmt.Println("TokFine")
	fmt.Println(s.TokFine)
	fmt.Println("TokCoarse")
	fmt.Println(s.TokCoarse)
	for i, _ := range s.TokFine {
		for j, _ := range s.TokFine[i] {
			fmt.Printf("%s: %s\n", s.TokFine[i][j], s.PosPku[i][j])
		}
	}
}

func TestNlp() {
	client := hanlp.HanLPClient(hanlp.WithAuth("MjQ0NEBiYnMuaGFubHAuY29tOktkc3hOc2RaZklPdnFzd3I=")) // auth不填则匿名
	s, _ := client.ParseObj("绵阳职业技术学院孵化楼", hanlp.WithLanguage("zh"), hanlp.WithTasks("ner/ontonotes"))
	fmt.Println(s.NerOntonotes)
	client.Parse("", hanlp.WithTokens())
}

func TestExcelize() {
	f, err := excelize.OpenFile("./client/test2.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取工作表中指定单元格的值
	sheet := f.GetSheetList()
	for _, list := range sheet {
		fmt.Println(list)
	}
	cell, err := f.GetCellValue("L.2 承包人提供主要材料和工程设备一览表(表-21)【市政~", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
}

// func TestExcelize2() {
// 	f := excelize.NewFile()
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			fmt.Println(err)
// 		}
// 	}()

// 	excel := Model.Excel{}
// 	excel.AnalyseExcel("./client/test4.xls")
// 	tmp := Model.Sheet{}
// 	for i, _ := range excel.Sheets {
// 		tmp = excel.Sheets[i]
// 		index, err := f.NewSheet(tmp.SheetName)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		for _,row := range tmp.Row{

// 		}
// 	}
// }

func TestNewExcel() {
	file, err := xls.OpenFile("./test5.xls")
	if err != nil {
		log.Fatalf("open excel file err: %v", err)
	}
	//s0, _ := file.GetSheet(0)
	s1, _ := file.GetSheet(3)
	fmt.Println(Model.GetStrByRL(s1, s1.GetNumberRows()-1, 3))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(strconv.Itoa(i) + "=" + Model.GetStrByRL(s1, 4, i))
	// }
}

func TestAgg() {
	Model.InitElasticSearch(false)
	aggs := elastic.NewTermsAggregation().Field("ProId")
	termQuery := elastic.NewTermQuery("Col1.keyword", "合计")
	result, err := Model.GlobalES.Client.Search().Index("sheet2").Aggregation("pro", aggs).Query(termQuery).Size(0).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	var tmp Model.Sheet2
	for _, item := range result.Each(reflect.TypeOf(tmp)) {
		t := item.(Model.Sheet2)
		fmt.Printf("%#v\n", t)
	}
	agg, found := result.Aggregations.Terms("pro")
	if found {
		for _, bucket := range agg.Buckets {
			var m int
			m = int(bucket.Key.(float64))
			fmt.Printf("%d\n", m)
		}
	}
}

func TestLike() {
	Model.InitElasticSearch(false)
	mlt := elastic.NewMoreLikeThisQuery()
	tmp := Model.GlobalES.QueryByProjectName("成都")[0]
	doc := elastic.NewMoreLikeThisQueryItem().Doc(tmp)
	mlt.MinimumShouldMatch("90%").MinDocFreq(2).MaxQueryTerms(100).LikeItems(doc)
	var pro Model.Project
	res, err := Model.GlobalES.Client.Search().Profile(true).Human(true).Index("management").Type("project").Query(mlt).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v\n", res.Profile.Shards[0].Searches[0].Query[0].Breakdown)
	for i, v := range res.Each(reflect.TypeOf(pro)) {
		t := v.(Model.Project)
		fmt.Printf("%s: %f\n", t.ProjectName, *res.Hits.Hits[i].Score)
	}
	// for _, v := range res.Each(reflect.TypeOf(pro)) {
	// 	t := v.(Model.Project)
	// 	fmt.Printf("项目名称：%s, 分数：%d", t.ProjectName, res.Profile.Shards[0].Searches[0].Query[0].Breakdown["score"])
	// }
}

func TestGroup() {
	Model.OpenDatabase(false)
	defer Model.CloseDatabase()
	var result []myResult1
	res := Model.GlobalConn.Table("sheet2").
		Select("sheet2.pro_name,project.total_cost_lower as price, sum(col6) as measure_price").
		Joins("left join project on project.project_name=sheet2.pro_name").
		Where("sheet2.col1=?", "合计").Group("sheet2.pro_name").Find(&result)
	//fmt.Printf("%v\n", result)
	fmt.Println(res.Error)
	var result2 []myResult2
	res = Model.GlobalConn.Table("unit").
		Select("unit.pro_name, sum(fees) as rule_price").
		Joins("left join project on project.project_name=unit.pro_name").
		Group("unit.pro_name").Find(&result2)
	fmt.Println(res.Error)
	tmp := make(map[string]myResult2)
	for _, v := range result2 {
		tmp[v.ProName] = v
	}
	for _, v := range result {
		v.Total, _ = strconv.ParseFloat(v.Price[:len(v.Price)-3], 64)
		v.MeasurePresent = v.MeasurePrice / v.Total
		v.RulePrice = tmp[v.ProName].RulePrice
		v.RulePresent = v.RulePrice / v.Total
		fmt.Println(v)
	}
}

type myResult1 struct {
	ProName        string
	Total          float64
	Price          string
	MeasurePrice   float64
	RulePrice      float64
	MeasurePresent float64
	RulePresent    float64
}

type myResult2 struct {
	ProName   string
	RulePrice float64
}

func TestDoc() {
	Model.InitElasticSearch(false)
	res, err := Model.GlobalES.Client.Search("prodoc").Type("reldoc").Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var p Model.ProjectDoc
	for _, v := range res.Each(reflect.TypeOf(p)) {
		t := v.(Model.ProjectDoc)
		fmt.Printf("%s\n", t.ProName)
		for _, k := range t.ReleventDoc {
			fmt.Printf("%s(%f): %v\n", k.ProName, k.Score, k.TermScore)
		}
	}
}

func TestTermQuery() {
	Model.InitElasticSearch(false)
	res := Model.GlobalES.TermQueryByProjectName("变电站、电力通道完善工程")
	fmt.Printf("%v\n", res)
}

func TestQueryMatch() {
	Model.InitElasticSearch(false)
	pro := Model.GlobalES.TermQueryByProjectName("成都德川友邦印务有限公司新建厂区一期项目")
	Model.GlobalES.GetRelevanceByPro(pro[0])
}

func TestSplitExcel() {
	file, err := xls.OpenFile("./client/test5.xls")
	if err != nil {
		log.Fatalf("open excel file err: %v", err)
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//s0, _ := file.GetSheet(0)
	s1, _ := file.GetSheet(3)
	tmp, err := s1.GetRow(1)
	row := strconv.Itoa(s1.GetNumberRows())
	err = f.SetSheetName("Sheet1", s1.GetName())
	if err != nil {
		log.Fatal(err)
	}
	rows := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for j := 0; j < s1.GetNumberRows(); j++ {
		for i := 0; i < len(tmp.GetCols()); i++ {
			//fmt.Println(strconv.Itoa(i) + "=" + Model.GetStrByRL(s1, j, i))
			f.SetCellValue(s1.GetName(), rows[i]+strconv.Itoa(j+1), Model.GetStrByRL(s1, j, i))
		}
	}
	f.MergeCell(s1.GetName(), "A1", "F1")
	B2, err := f.GetCellValue(s1.GetName(), "B2")
	if B2 == "" {
		C2, _ := f.GetCellValue(s1.GetName(), "C2")
		f.SetCellValue(s1.GetName(), "B2", C2)
	}
	f.MergeCell(s1.GetName(), "B2", "D2")
	f.MergeCell(s1.GetName(), "D3", "f3")
	f.MergeCell(s1.GetName(), "A"+row, "B"+row)
	f.MergeCell(s1.GetName(), "A3", "A4")
	f.MergeCell(s1.GetName(), "B3", "B4")
	f.MergeCell(s1.GetName(), "C3", "C4")
	if err := f.SaveAs("./store/" + s1.GetName() + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	buff, err := f.WriteToBuffer()
	buff.Bytes()
}
