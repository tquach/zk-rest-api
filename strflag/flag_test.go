package strflag

import (
	"flag"
	"testing"
)

var (
	hosts     StringSlice
	testFlags = flag.NewFlagSet("slice-test", flag.ContinueOnError)
)

func TestSetSliceFlag(t *testing.T) {
	if err := testFlags.Parse([]string{"-hosts", "localhost,0.0.0.0,random-host"}); err != nil {
		t.Fatal(err)
	}

	if len(hosts) != 3 {
		t.Fatal("Expected 3 values but got ", len(hosts))
	}

	expected := "[localhost 0.0.0.0 random-host]"
	if actual := hosts.String(); actual != expected {
		t.Fatalf("Expected %s got %s", expected, actual)
	}
}

func TestSetSliceWithOneArg(t *testing.T) {
	if err := testFlags.Parse([]string{"-hosts", "localhost"}); err != nil {
		t.Fatal(err)
	}

	if len(hosts) != 1 {
		t.Fatal("Expected 1 values but got ", len(hosts))
	}

	expected := "[localhost]"
	if actual := hosts.String(); actual != expected {
		t.Fatalf("Expected %s got %s", expected, actual)
	}
}

func init() {
	testFlags.Var(&hosts, "hosts", "A comma-separated list of hostnames.")
}
