package reporter

import "testing"

func TestOSReporter_ErrorFormat(t *testing.T) {
	rp := NewOSReporter()
	rp.ErrorFormat("report error value:%v", "test")

	rp.ErrorFormat("report error value:%v", "test123")
}

func TestOSReporter_InfoFormat(t *testing.T) {
	rp := NewOSReporter()
	rp.InfoFormat("report info value:%v", "test")

	rp.InfoFormat("report info value:%v", "test123")
}
