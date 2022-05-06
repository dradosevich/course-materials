// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f") // Currently function returns only number of open ports
	want := "foo"
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}
func TestAddSha(t *testing.T) {
	GenHashMapsImproved("../main/rockyou-75.txt")
	shawant := "1DF1854015E31CA286D015345EAFF29A6C6073F70984A3A746823D4CAC16B075"
	pass := "zxcvbnm"
	pass, err := GetMD5(shawant)
	if pass != "zxcvbnm" {
		t.Errorf("Wrong password received, got %d wanted %d", pass, "zxcvbnm")
	}
}
