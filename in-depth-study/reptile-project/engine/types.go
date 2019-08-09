package engine


type ParserFunc func(contents []byte) ParseResult


type Request struct {
	Url string
	// ParserFunc func([]byte) ParseResult
	ParserFunc ParserFunc
}


type ParseResult struct {
	Requests 	[]Request
	Items 		[]interface{}
}


func NilParser([]byte) ParseResult {
	return ParseResult{}
}