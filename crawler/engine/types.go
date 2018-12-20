package engine

//请求链接 和 解释器
type Request struct {
	Url string
	ParserFunc func( []byte ) ParseResult
}

//解释器返回结果
type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Url 	string
	Type 	string //数据库表名
	Id 		string
	Payload interface{} //接口即什么类型都可以
}


func NilParser( []byte ) ParseResult {
		return ParseResult{}
}