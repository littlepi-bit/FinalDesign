package Model

type Info struct {
	ProName string `json:"ProName"`
	Size    string `json:"Size"`
	//建筑类型
	BType string `json:"BType"`
	//造价业务类型
	PType string `json:"PType"`
	//甲方
	PartyA   string `json:"PartyA"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Price    string `json:"Price"`
	//承包单位
	PartyB string `json:"PartyB"`
	//投资类型
	IType string `json:"IType"`
	//建设性质
	CType string `json:"CType"`
	//场地类型
	SType string `json:"SType"`
	//容积率
	Plot string `json:"Plot"`
	//绿地率
	Greening string `json:"Greening"`
	//建筑密度
	Density string `json:"Density"`
}
