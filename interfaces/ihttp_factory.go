package interfaces

type IHttpFactory interface {
	IReport

	SetDomain(value string) IHttpFactory
	Request(method, path string, pathArgs ...interface{}) IRequest
	POST(path string, pathArgs ...interface{}) IRequest
	GET(path string, pathArgs ...interface{}) IRequest
	DELETE(path string, pathArgs ...interface{}) IRequest
	PATCH(path string, pathArgs ...interface{}) IRequest
	PUT(path string, pathArgs ...interface{}) IRequest
	OPTIONS(path string, pathArgs ...interface{}) IRequest
	HEAD(path string, pathArgs ...interface{}) IRequest
	Value(value interface{}) IValue
	Object(value map[string]interface{}) IObject
	Array(value []interface{}) IArray
	String(value string) IString
	Float(value float64) IFloat
	Integer(value int) IInteger
	Boolean(value bool) IBool
}
