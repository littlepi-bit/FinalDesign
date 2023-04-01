package Model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

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

var GlobalES *ElasticSearch

func InitElasticSearch() {
	GlobalES = NewElasticSearch()
	GlobalES.Init()
}

func NewElasticSearch() *ElasticSearch {
	return &ElasticSearch{
		Host: "http://127.0.0.1:9200/",
	}
}

//清空索引
func EmptyES() {
	GlobalES.Delete()
}

//初始化
func (e *ElasticSearch) Init() {
	var err error
	e.Client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(e.Host))
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

//添加项目信息
func (e *ElasticSearch) InsertProject(pro Project) {
	put, err := e.Client.Index().Index("management").Type("project").BodyJson(pro).Do(context.Background())
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
	put, err := e.Client.Index().Index("gfile").Type("gfile").BodyJson(gFile).Do(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Indexed tweet %s to index s%s, type %s\n", put.Id, put.Index, put.Type)
}

//根据项目名称搜索项目
func (e *ElasticSearch) QueryByProjectName(proName string) []Project {
	matchPhraseQuery := elastic.NewMatchQuery("ProjectName", proName)
	res, err := e.Client.Search("management").Type("project").Query(matchPhraseQuery).Do(context.Background())
	return printProjects(res, err)
}

func (e *ElasticSearch) QueryByIndividualProjectName(indName string) []Project {
	matchPhraseQuery := elastic.NewMatchQuery("IndividualProjects.IndividualProjectName", indName)
	res, err := e.Client.Search("management").Type("project").Query(matchPhraseQuery).Do(context.Background())
	return printProjects(res, err)
}

func (e *ElasticSearch) QueryByUnitProjectName(unitName string) []Project {
	matchPhraseQuery := elastic.NewMatchQuery("IndividualProjects.UnitProjects.UnitName", unitName)
	res, err := e.Client.Search("management").Type("project").Query(matchPhraseQuery).Do(context.Background())
	return printProjects(res, err)
}

//根据项目名称精确查询项目
func (e *ElasticSearch) TermQueryByProjectName(proName string) []Project {
	termQuery := elastic.NewTermQuery("ProjectName", proName)
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
