package Model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
)

type ElasticSearch struct {
	Client *elastic.Client
	Host   string
}

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

type DocScore struct {
	ProName   string             `json:"proName"`
	Score     float64            `json:"score"`
	TermScore map[string]float64 `json:"termScore"`
}

type ProjectDoc struct {
	ProName     string
	ReleventDoc []DocScore
}

var GlobalES *ElasticSearch

func InitElasticSearch(remote bool) {
	GlobalES = NewElasticSearch(remote)
	GlobalES.Init(remote)
}

func NewElasticSearch(remote bool) *ElasticSearch {
	if remote {
		return &ElasticSearch{
			Host: "http://120.77.12.35:9200/",
		}
	}
	return &ElasticSearch{
		Host: "http://127.0.0.1:9200/",
	}
}

//清空索引
func EmptyES() {
	GlobalES.Delete()
}

//初始化
func (e *ElasticSearch) Init(remote bool) {
	var err error
	if remote {
		e.Client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(e.Host), elastic.SetBasicAuth("elastic", "elastic"))
	} else {
		e.Client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(e.Host))
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	_, _, err = e.Client.Ping(e.Host).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = e.Client.ElasticsearchVersion(e.Host)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (e *ElasticSearch) Gets() {
	//通过id查找
	get1, err := e.Client.Get().Index("megacorp").Type("employee").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		var bb Employee
		err := json.Unmarshal(get1.Source, &bb)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(bb.FirstName)
		fmt.Println(string(get1.Source))
	}
}

func (e *ElasticSearch) Delete() {
	query := elastic.NewBoolQuery()
	res, err := e.Client.DeleteByQuery("management").Query(query).Do(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("delete result %s\n", res)
	res, err = e.Client.DeleteByQuery("gfile").Query(query).Do(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("delete result %s\n", res)
	for i := 0; i < 7; i++ {
		index := fmt.Sprintf("sheet%d", i)
		res, err = e.Client.DeleteByQuery(index).Query(query).Do(context.Background())
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Printf("delete result %s\n", res)
	}
	res, err = e.Client.DeleteByQuery("prodoc").Query(query).Do(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("delete result %s\n", res)
}

func (e *ElasticSearch) Query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = e.Client.Search("management").Type("project").Do(context.Background())
	printProjects(res, err)
	res, err = e.Client.Search("gfile").Type("gfile").Do(context.Background())
	printGFiles(res, err)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

//打印查询到的Project
func printProjects(res *elastic.SearchResult, err error) (pros []Project) {
	if err != nil {
		log.Fatal(err.Error())
	}
	var typ Project
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Project)
		pros = append(pros, t)
		//fmt.Printf("%#v\n", t)
		fmt.Println(t.ProjectName)
	}
	return pros
}

//打印查询到的通用文件
func printGFiles(res *elastic.SearchResult, err error) (gFiles []GeneralFile) {
	if err != nil {
		log.Fatal(err.Error())
	}
	var typ GeneralFile
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(GeneralFile)
		gFiles = append(gFiles, t)
		fmt.Printf("%#v\n", t)
	}
	return
}

//打印查询到的sheet5
func printSheet5(res *elastic.SearchResult, err error) (s []Sheet5) {
	if err != nil {
		log.Fatal(err.Error())
	}
	var typ Sheet5
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Sheet5)
		s = append(s, t)
		//fmt.Printf("%#v\n", t)
	}
	return
}

//打印查询到的sheet6
func printSheet6(res *elastic.SearchResult, err error) (s []Sheet6) {
	if err != nil {
		log.Fatal(err.Error())
	}
	var typ Sheet6
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(Sheet6)
		s = append(s, t)
		//fmt.Printf("%#v\n", t)
	}
	return
}

//添加项目信息
func (e *ElasticSearch) InsertProject(pro Project) {
	put, err := e.Client.Index().Index("management").Type("project").BodyJson(pro).Refresh("true").Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Indexed tweet %s to index s%s, type %s\n", put.Id, put.Index, put.Type)
}

//添加文件信息
func (e *ElasticSearch) InsertFile(file File) {
	gFile := FileToGFile(file)
	e.InsertGFile(gFile)
}

//添加文件夹信息
func (e *ElasticSearch) InsertFolder(folder Folder) {
	gFile := FolderToGFile(folder)
	e.InsertGFile(gFile)
}

//添加gFile信息
func (e *ElasticSearch) InsertGFile(gFile GeneralFile) {
	id := strconv.Itoa(gFile.GId)
	put, err := e.Client.Index().Index("gfile").Type("gfile").Id(id).BodyJson(gFile).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet0信息
func (e *ElasticSearch) InsertSheet0(s Sheet0) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet0").Type("sheet0").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet1信息
func (e *ElasticSearch) InsertSheet1(s Sheet1) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet1").Type("sheet1").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet2信息
func (e *ElasticSearch) InsertSheet2(s Sheet2) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet2").Type("sheet2").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet3信息
func (e *ElasticSearch) InsertSheet3(s Sheet3) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet3").Type("sheet3").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet4信息
func (e *ElasticSearch) InsertSheet4(s Sheet4) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet4").Type("sheet4").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet5信息
func (e *ElasticSearch) InsertSheet5(s Sheet5) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet5").Type("sheet5").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加sheet6信息
func (e *ElasticSearch) InsertSheet6(s Sheet6) {
	id := strconv.Itoa(s.SheetId)
	_, err := e.Client.Index().Index("sheet6").Type("sheet6").Id(id).BodyJson(s).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

//添加项目关联信息
func (e *ElasticSearch) InsertRelevance(pro Project) {
	rel := e.GetRelevanceByPro(pro)
	proDoc := ProjectDoc{
		ProName:     pro.ProjectName,
		ReleventDoc: rel,
	}
	put, err := e.Client.Index().Index("prodoc").Type("reldoc").BodyJson(proDoc).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)
	docMap := make(map[string]DocScore)
	for _, r := range proDoc.ReleventDoc {
		docMap[r.ProName] = r
	}
	for _, s := range rel {
		if s.ProName == proDoc.ProName {
			continue
		}
		id, tmp := e.QueryByDocName(s.ProName)
		if id == "" {
			continue
		}
		tmp.ReleventDoc = append(tmp.ReleventDoc, DocScore{
			ProName:   proDoc.ProName,
			Score:     docMap[tmp.ProName].Score,
			TermScore: docMap[tmp.ProName].TermScore,
		})
		_, err := e.Client.Update().Index("prodoc").Type("reldoc").Id(id).Doc(tmp).Refresh("true").Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}

}

//根据项目名称搜索项目
func (e *ElasticSearch) QueryByProjectName(proName string) []Project {
	if proName == "" {
		res, err := e.Client.Search("management").Type("project").Do(context.Background())
		return printProjects(res, err)
	} else {
		matchPhraseQuery := elastic.NewMatchQuery("ProjectName", proName)
		res, err := e.Client.Search("management").Type("project").Query(matchPhraseQuery).Do(context.Background())
		return printProjects(res, err)
	}
}

//根据单项工程名搜索项目
func (e *ElasticSearch) QueryByIndividualProjectName(indName string) []Project {
	matchPhraseQuery := elastic.NewMatchQuery("IndividualProjects.IndividualProjectName", indName)
	res, err := e.Client.Search("management").Type("project").Query(matchPhraseQuery).Do(context.Background())
	return printProjects(res, err)
}

//根据单位工程名搜索项目
func (e *ElasticSearch) QueryByUnitProjectName(unitName string) []Project {
	matchPhraseQuery := elastic.NewMatchQuery("IndividualProjects.UnitProjects.UnitName", unitName)
	res, err := e.Client.Search("management").Type("project").Query(matchPhraseQuery).Do(context.Background())
	return printProjects(res, err)
}

//根据项目名称精确查询项目
func (e *ElasticSearch) TermQueryByProjectName(proName string) []Project {
	termQuery := elastic.NewTermQuery("ProjectName.keyword", proName)
	res, err := e.Client.Search("management").Type("project").Query(termQuery).Do(context.Background())
	return printProjects(res, err)
}

//根据文件名搜索通用文件
func (e *ElasticSearch) QueryByGFileName(proId int, gFileName string) []GeneralFile {
	boolQuery := elastic.NewBoolQuery()
	termQuery := elastic.NewTermQuery("proId", proId)
	matchQuery := elastic.NewMatchQuery("name", gFileName)
	boolQuery.Must(termQuery, matchQuery)
	res, err := e.Client.Search("gfile").Type("gfile").Query(boolQuery).Do(context.Background())
	return printGFiles(res, err)
}

//通过文件Id搜索通用文件
func (e *ElasticSearch) QueryByGFileId(proId int, GId int) []GeneralFile {
	boolQuery := elastic.NewBoolQuery()
	termQuery := elastic.NewTermQuery("proId", proId)
	matchQuery := elastic.NewMatchQuery("gId", GId)
	boolQuery.Must(termQuery, matchQuery)
	res, err := e.Client.Search("gfile").Type("gfile").Query(boolQuery).Do(context.Background())
	return printGFiles(res, err)
}

//根据文档名搜索相似度
func (e *ElasticSearch) QueryByDocName(docName string) (string, ProjectDoc) {
	termQuery := elastic.NewTermQuery("ProName.keyword", docName)
	res, err := e.Client.Search("prodoc").Type("reldoc").Query(termQuery).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var p ProjectDoc
	for i, v := range res.Each(reflect.TypeOf(p)) {
		t := v.(ProjectDoc)
		return res.Hits.Hits[i].Id, t
	}
	return "", ProjectDoc{}
}

//删除文件
func (e *ElasticSearch) DeleteFile(proId int, fileName string) {
	boolQuery := elastic.NewBoolQuery()
	termQuery1 := elastic.NewTermQuery("ProId", proId)
	termQuery2 := elastic.NewTermQuery("FileName", fileName)
	boolQuery.Must(termQuery1, termQuery2)
	res, err := e.Client.DeleteByQuery("file").Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%v\n", res)
}

//删除文件
func (e *ElasticSearch) DeleteFolder(proId int, folderName string) {
	boolQuery := elastic.NewBoolQuery()
	termQuery1 := elastic.NewTermQuery("ProId", proId)
	termQuery2 := elastic.NewTermQuery("FileName", folderName)
	boolQuery.Must(termQuery1, termQuery2)
	res, err := e.Client.DeleteByQuery("folder").Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%v\n", res)
}

//通过文件名删除通用文件
func (e *ElasticSearch) DeleteGFileByFileName(proId int, fileName string) {
	boolQuery := elastic.NewBoolQuery()
	termQuery1 := elastic.NewTermQuery("proId", proId)
	termQuery2 := elastic.NewTermQuery("name", fileName)
	boolQuery.Must(termQuery1, termQuery2)
	res, err := e.Client.DeleteByQuery("gfile").Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%v\n", res)
}

//通过文件Id删除通用文件
func (e *ElasticSearch) DeleteGFileByFileId(proId int, fileId int) {
	boolQuery := elastic.NewBoolQuery()
	termQuery1 := elastic.NewTermQuery("proId", proId)
	termQuery2 := elastic.NewTermQuery("gId", fileId)
	boolQuery.Must(termQuery1, termQuery2)
	res, err := e.Client.DeleteByQuery("gfile").Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("%v\n", res)
}

//根据项目名和materialName搜索sheet5
func (e *ElasticSearch) SearchMaterialByProName(proName string, materialName string) []Sheet5 {
	boolQuery := elastic.NewBoolQuery()
	matchQuery1 := elastic.NewMatchQuery("ProName", proName)
	matchQuery2 := elastic.NewMatchQuery("Col2", materialName)
	boolQuery.Must(matchQuery1, matchQuery2)
	service := e.Client.Search("sheet5").Type("sheet5").Size(10000)
	var res *elastic.SearchResult
	var err error
	if proName == "" && materialName == "" {
		res, err = service.Do(context.Background())
	} else if proName == "" {
		res, err = service.Query(matchQuery2).Do(context.Background())
	} else if materialName == "" {
		res, err = service.Query(matchQuery1).Do(context.Background())
	} else {
		res, err = service.Query(boolQuery).Do(context.Background())
	}
	return printSheet5(res, err)
}

//根据项目名和materialName搜索sheet6
func (e *ElasticSearch) SearchGlobalByProName(proName string, globalName string) []Sheet6 {
	boolQuery := elastic.NewBoolQuery()
	matchQuery1 := elastic.NewMatchQuery("ProName", proName)
	matchQuery2 := elastic.NewMatchQuery("Col4", globalName)
	termQuery1 := elastic.NewWildcardQuery("Col4", "*")
	termQuery2 := elastic.NewMatchQuery("Col4", "合计")
	boolQuery2 := elastic.NewBoolQuery()
	boolQuery2.MustNot(termQuery2).Filter(termQuery1)
	boolQuery.Must(matchQuery1, matchQuery2, boolQuery2)
	service := e.Client.Search("sheet6").Type("sheet6").Query(boolQuery2).Size(10000)
	var res *elastic.SearchResult
	var err error
	if proName == "" && globalName == "" {
		res, err = service.Do(context.Background())
	} else if proName == "" {
		res, err = service.Query(matchQuery2).Do(context.Background())
	} else if globalName == "" {
		res, err = service.Query(matchQuery1).Do(context.Background())
	} else {
		res, err = service.Query(boolQuery).Do(context.Background())
	}
	return printSheet6(res, err)
}

//获取关联文档
func (e *ElasticSearch) GetRelevantDocByProName(proName string) (docScore []DocScore) {
	mlt := elastic.NewMoreLikeThisQuery()
	tmp := e.TermQueryByProjectName(proName)
	if len(tmp) == 0 {
		return []DocScore{}
	}
	doc := elastic.NewMoreLikeThisQueryItem().Doc(tmp[0])
	mlt.MinimumShouldMatch("30%").MinDocFreq(2).MaxQueryTerms(100).LikeItems(doc)
	var pro Project
	res, err := e.Client.Search().Index("management").Explain(true).Type("project").Query(mlt).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v\n", res.Profile.Shards[0].Searches[0].Query[0].Breakdown)
	for i, v := range res.Each(reflect.TypeOf(pro)) {
		t := v.(Project)
		fmt.Printf("%s: %f\n", t.ProjectName, *res.Hits.Hits[i].Score)
		termScore := make(map[string]float64)
		for _, v := range res.Hits.Hits[i].Explanation.Details {
			if strings.Contains(v.Description, "IndividualProjects.UnitProjects.UnitName.keyword:") {
				tmp := strings.Split(v.Description, ":")
				strs := strings.Split(tmp[1], " ")
				fmt.Printf("%v %f\n", strs[0], v.Value)
				termScore[strs[0]] = v.Value
			}
		}
		d := DocScore{
			ProName:   t.ProjectName,
			Score:     *res.Hits.Hits[i].Score,
			TermScore: termScore,
		}
		docScore = append(docScore, d)
	}
	return
}

//获取两个文档的关联度
func (e *ElasticSearch) GetRelevanceByProName(proName1 string, proName2 string) DocScore {
	mlt := elastic.NewMoreLikeThisQuery()
	tmp := e.QueryByProjectName(proName1)[0]
	doc := elastic.NewMoreLikeThisQueryItem().Doc(tmp)
	mlt.MinimumShouldMatch("30%").MinDocFreq(1).MaxQueryTerms(100).LikeItems(doc)
	termQuery := elastic.NewTermQuery("ProjectName", proName2)
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(termQuery, mlt)
	var pro Project
	res, err := e.Client.Search().Index("management").Type("project").Query(boolQuery).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := range res.Each(reflect.TypeOf(pro)) {
		fmt.Println(*res.Hits.Hits[i].Score)
	}
	return DocScore{}
}

//获取项目文档的关联度
func (e *ElasticSearch) GetRelevanceByPro(pro Project) (docScore []DocScore) {
	mlt := elastic.NewMoreLikeThisQuery()
	doc := elastic.NewMoreLikeThisQueryItem().Doc(pro)
	mlt.MinimumShouldMatch("30%").MinDocFreq(1).MaxQueryTerms(25).LikeItems(doc)
	res, err := e.Client.Search().Index("management").Explain(true).Type("project").Query(mlt).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v\n", res.Profile.Shards[0].Searches[0].Query[0].Breakdown)
	for i, v := range res.Each(reflect.TypeOf(pro)) {
		t := v.(Project)
		fmt.Printf("%s: %f\n", t.ProjectName, *res.Hits.Hits[i].Score)
		//fmt.Printf("%v\n", res.Profile.Shards[0].Searches[0].Query[0])
		//fmt.Println(res.Hits.Hits[i].Explanation.Description)
		//fmt.Printf("%v\n", res.Hits.Hits[0].Explanation.Details[i*50])
		termScore := make(map[string]float64)
		for _, v := range res.Hits.Hits[i].Explanation.Details {
			if strings.Contains(v.Description, "IndividualProjects.UnitProjects.UnitName.keyword:") {
				tmp := strings.Split(v.Description, ":")
				strs := strings.Split(tmp[1], " ")
				fmt.Printf("%v %f\n", strs[0], v.Value)
				termScore[strs[0]] = v.Value
			}
		}
		d := DocScore{
			ProName:   t.ProjectName,
			Score:     *res.Hits.Hits[i].Score,
			TermScore: termScore,
		}
		docScore = append(docScore, d)
	}
	return
}
