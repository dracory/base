package cfmt

import (
	"bytes"
	"testing"
)

func TestColorOutput(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer
	SetOutput(&buf)

	// Test Info output
	Info("test info")
	if buf.String() != BoldBlue+"test info"+Reset {
		t.Errorf("Info() = %q, want %q", buf.String(), BoldBlue+"test info"+Reset)
	}

	// Reset buffer
	buf.Reset()

	// Test Infoln output
	Infoln("test info")
	if buf.String() != BoldBlue+"test info\n"+Reset {
		t.Errorf("Infoln() = %q, want %q", buf.String(), BoldBlue+"test info\n"+Reset)
	}

	// Reset buffer
	buf.Reset()

	// Test Error output
	Error("test error")
	if buf.String() != BoldRed+"test error"+Reset {
		t.Errorf("Error() = %q, want %q", buf.String(), BoldRed+"test error"+Reset)
	}

	// Reset buffer
	buf.Reset()

	// Test Errorln output
	Errorln("test error")
	if buf.String() != BoldRed+"test error\n"+Reset {
		t.Errorf("Errorln() = %q, want %q", buf.String(), BoldRed+"test error\n"+Reset)
	}

	// Reset buffer
	buf.Reset()

	// Test Success output
	Success("test success")
	if buf.String() != BoldGreen+"test success"+Reset {
		t.Errorf("Success() = %q, want %q", buf.String(), BoldGreen+"test success"+Reset)
	}

	// Reset buffer
	buf.Reset()

	// Test Warning output
	Warning("test warning")
	if buf.String() != BoldYellow+"test warning"+Reset {
		t.Errorf("Warning() = %q, want %q", buf.String(), BoldYellow+"test warning"+Reset)
	}
}

func TestFormattedOutput(t *testing.T) {
	// Create a buffer to capture output
	var buf bytes.Buffer
	SetOutput(&buf)

	// Test Infof output
	Infof("test %s", "info")
	if buf.String() != BoldBlue+"test info"+Reset {
		t.Errorf("Infof() = %q, want %q", buf.String(), BoldBlue+"test info"+Reset)
	}

	// Reset buffer
	buf.Reset()

	// Test Errorf output
	Errorf("test %s", "error")
	if buf.String() != BoldRed+"test error"+Reset {
		t.Errorf("Errorf() = %q, want %q", buf.String(), BoldRed+"test error"+Reset)
	}
}
