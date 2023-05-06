package Model

type Graph struct {
	Feature string
	Enties  map[string]Entiy
}

type Entiy struct {
	Name         string
	Attribute    map[string]string
	Relationship map[string]Entiy
}

var GraphPros = []string{"地区", "甲方", "乙方", "建筑类型", "造价类型", "投资类型", "建设性质", "场地类型"}

func NewGraph() *Graph {
	graph := &Graph{
		Feature: "项目信息",
		Enties:  make(map[string]Entiy),
	}
	for _, pro := range GraphPros {
		graph.Enties[pro] = Entiy{
			Name:         pro,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
	}
	return graph
}
func InitGraph() {
	graph := NewGraph()
	GlobalES.InsertGraph(*graph)
}

func (graph *Graph) AddInfo(info Info) {
	proEntiy := Entiy{
		Name:      info.ProName,
		Attribute: make(map[string]string),
	}
	proEntiy.Attribute["Price"] = info.Price
	proEntiy.Attribute["Size"] = info.Size
	entiy, ok := graph.Enties["地区"].Relationship[info.Province]
	if !ok {
		entiy = Entiy{
			Name:         info.Province,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		graph.Enties["地区"].Relationship[info.Province] = entiy
	}
	entiy, ok = entiy.Relationship[info.City]
	if !ok {
		entiy = Entiy{
			Name:         info.City,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		graph.Enties["地区"].Relationship[info.Province].Relationship[info.City] = entiy
	}
	entiy, ok = entiy.Relationship[info.Area]
	if !ok {
		entiy = Entiy{
			Name:         info.Area,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		graph.Enties["地区"].Relationship[info.Province].Relationship[info.City].Relationship[info.Area] = entiy
	}
	entiy, ok = entiy.Relationship[info.ProName]
	if !ok {
		entiy = proEntiy
		graph.Enties["地区"].Relationship[info.Province].Relationship[info.City].
			Relationship[info.Area].Relationship[info.ProName] = entiy
	}
	entiy, ok = graph.Enties["甲方"].Relationship[info.PartyA]
	if !ok {
		entiy = Entiy{
			Name:         info.PartyA,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["甲方"].Relationship[info.PartyA] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
	entiy, ok = graph.Enties["乙方"].Relationship[info.PartyB]
	if !ok {
		entiy = Entiy{
			Name:         info.PartyB,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["乙方"].Relationship[info.PartyB] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
	entiy, ok = graph.Enties["建筑类型"].Relationship[info.BType]
	if !ok {
		entiy = Entiy{
			Name:         info.BType,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["建筑类型"].Relationship[info.BType] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
	entiy, ok = graph.Enties["建设性质"].Relationship[info.CType]
	if !ok {
		entiy = Entiy{
			Name:         info.CType,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["建设性质"].Relationship[info.CType] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
	entiy, ok = graph.Enties["造价类型"].Relationship[info.PType]
	if !ok {
		entiy = Entiy{
			Name:         info.PType,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["造价类型"].Relationship[info.PType] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
	entiy, ok = graph.Enties["投资类型"].Relationship[info.IType]
	if !ok {
		entiy = Entiy{
			Name:         info.IType,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["投资类型"].Relationship[info.IType] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
	entiy, ok = graph.Enties["场地类型"].Relationship[info.SType]
	if !ok {
		entiy = Entiy{
			Name:         info.SType,
			Attribute:    make(map[string]string),
			Relationship: make(map[string]Entiy),
		}
		entiy.Relationship[info.ProName] = proEntiy
		graph.Enties["场地类型"].Relationship[info.SType] = entiy
	} else {
		entiy.Relationship[info.ProName] = proEntiy
	}
}
