package reporter

import (
	"fmt"
	"os"
)

type OSReporter struct {
	infoObject *os.File
	errObject  *os.File
}

func NewOSReporter() IReporter {
	return &OSReporter{
		infoObject: os.Stdout,
		errObject:  os.Stderr,
	}
}

func (rp *OSReporter) InfoFormat(format string, args ...interface{}) {
	_, _ = rp.infoObject.WriteString(fmt.Sprintf(format, args...))

	fmt.Println()
}

func (rp *OSReporter) ErrorFormat(format string, args ...interface{}) {
	_, _ = rp.errObject.WriteString(fmt.Sprintf(format+"\n", args...))
}
