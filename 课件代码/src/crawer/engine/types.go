package engine

type ParserFunc  func(contents []byte, url string) ParseResult
type Request struct {
	Url string
	ParserFunc ParserFunc
}



type ParseResult struct {
	Requests []Request
	Items []Item
}

type Item struct {
	Id 	string //存储时去重。
	Url string
	Type string //存储的配置
	Payload interface{}
}

func NilParser([]byte) ParseResult  {
	return ParseResult{}
}
