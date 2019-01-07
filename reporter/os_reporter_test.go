package reporter

import "testing"

func TestOSReporter_ErrorFormat(t *testing.T) {
	rp := NewOSReporter()
	err := rp.ErrorFormat("report error value:%v", "test")
	if err != nil {
		t.Errorf("ErrorFormat error:%v", err.Error())
		return
	}

	err = rp.ErrorFormat("report error value:%v", "test123")
	if err != nil {
		t.Errorf("ErrorFormat error:%v", err.Error())
		return
	}
}

func TestOSReporter_InfoFormat(t *testing.T) {
	rp := NewOSReporter()
	err := rp.InfoFormat("report info value:%v", "test")
	if err != nil {
		t.Errorf("ErrorFormat error:%v", err.Error())
		return
	}

	err = rp.InfoFormat("report info value:%v", "test123")
	if err != nil {
		t.Errorf("ErrorFormat error:%v", err.Error())
		return
	}
}
