package reporter

type IReporter interface {
	InfoFormat(fmt string, args ...interface{})
	ErrorFormat(fmt string, args ...interface{})
}
