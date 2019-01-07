package reporter

import (
	"fmt"
	"os"
)

type OSReporter struct {
	infoObject *os.File
	errObject  *os.File
}

func NewOSReporter() *OSReporter {
	return &OSReporter{
		infoObject: os.Stdout,
		errObject:  os.Stderr,
	}
}

func (rp *OSReporter) InfoFormat(format string, args ...interface{}) error {
	_, err := rp.infoObject.WriteString(fmt.Sprintf(format, args...))
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

func (rp *OSReporter) ErrorFormat(format string, args ...interface{}) error {
	_, err := rp.errObject.WriteString(fmt.Sprintf(format+"\n", args...))
	if err != nil {
		return err
	}

	return nil
}
