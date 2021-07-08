package engine

type ParseResult struct {
	Items []interface{}
	Requests []Request
}
type Request struct {
	Url string
	ParseFunc func([]byte) ParseResult
}
func NilParse([]byte) ParseResult{
	return ParseResult{}
}
