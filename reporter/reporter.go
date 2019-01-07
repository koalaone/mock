package reporter

type Reporter interface {
	InfoFormat(fmt string, args ...interface{}) error
	ErrorFormat(fmt string, args ...interface{}) error
}
