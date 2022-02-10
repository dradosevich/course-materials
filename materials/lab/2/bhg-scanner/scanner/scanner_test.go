package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T) {

	gotopen, gotclosed := PortScanner(0, 1024) // Currently function returns only number of open ports

	want := 2 // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns
	//consider what would happen if you parameterize the portscanner address and ports to scan
	closedwanted := 1022
	total := 1024
	if gotopen != want {
		t.Errorf("got %d, wanted %d", gotopen, want)
	}
	if gotclosed != closedwanted {
		t.Errorf("got %d closed ports, wanted %d", gotclosed, closedwanted)
	}
	if gotopen+gotclosed != total {
		t.Errorf("total number ports scanned: %d, expected %d", gotopen+gotclosed, total)
	}

}

/*
func TestTotalPortsScanned(t *testing.T) {
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

	gotopen, gotclosed := PortScanner(0, 1024) // Currently function returns only number of open ports
	want := 1023                               // default value; consider what would happen if you parameterize the portscanner ports to scan

	if gotopen+gotclosed != want {
		t.Errorf("total number ports scanned: %d, expected %d", gotopen+gotclosed, want)
	}
} */
