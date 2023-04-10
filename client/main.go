package main

import (
	"FinalDesign/Model"
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/xuri/excelize/v2"
)

func main() {
	//TestLike()
	Model.InitElasticSearch(true)
	Model.GlobalES.GetRelevanceDocByProName("北川")
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
