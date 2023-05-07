package Model

type Msg struct {
	FileInfo []GeneralFile `json:"filesInfo"`
}

type SearchContent struct {
	Content string `json:"content"`
}
