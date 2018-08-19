package base

type Request struct {
	PageSize int `json:"pageSize"`
	PageNo int `json:"pageNo"`
	//DESC or ASC
	Sort string `json:"sort"`
}


type Response struct {
	CurPage int `json:"curPage"`
	PageSize int `json:"pageSize"`
	TotalSize int `json:"totalSize"`
	TotalPage int `json:"totalPage"`
	FirstPage bool `json:"firstPage"`
	LastPage bool `json:"lastPage"`
}



