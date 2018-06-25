package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 第二个Parser是接口
type Request struct {
	Url string
	Parser Parser
}

//type SerializedParser struct {
//	Name string
//	Args interface{}
//}

// {"ParseCityList", nil}, {"ProfileParser", userName}


type ParseResult struct {
	Requests []Request
	Items   []Item
}

type Item struct {
	Url 		string
	Type  		string
	Id  		string
	Payload  	interface{}
}

type NilParser struct {}

// 实现NilParser的interface
func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (
	name string, args interface{}) {
	return "NilParser", nil
}

//func NilParser([]byte) ParseResult {
//	return ParseResult{}
//}

// 包装ParserFunc
type FuncParser struct {
	parser ParserFunc
	name string
}

// 对接口parser实现FuncParser函数
// FuncParser函数只有一个函数，没有参数
func (f *FuncParser) Parse(
	contents []byte, url string) ParseResult {
	return  f.parser(contents, url)
}

func (f *FuncParser) Serialize() (
	name string, args interface{}) {
	return f.name, nil
}

// 用工厂函数的方法建FuncParser
func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:	name,
	}
}



