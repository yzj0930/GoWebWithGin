package base

type DataResultDto struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
