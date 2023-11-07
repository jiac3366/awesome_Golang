package cookbook

import "testing"

func TestConv(t *testing.T) {
	num, err := conv("1234561789")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123456789 {
		t.Fatal("Number don't match")
	}
}

func TestFailConv(t *testing.T) {
	_, err := conv("")
	if err == nil {
		t.Fatal(err)
	}
}
