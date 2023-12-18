package scan_test

import (
	"net"
	"strconv"
	"testing"

	"agmmtoo.me/ammgo/cobra/pScan/scan"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}
	if ps.Open.String() != "closed" {
		t.Errorf("Expected closed, got %q", ps.Open.String())
	}
	ps.Open = true
	if ps.Open != true {
		t.Errorf("Expected open, got %v", ps.Open.String())
	}
}

func TestRunHostFound(t *testing.T) {
	testCases := []struct {
		name         string
		expecteState string
	}{
		{"OpenPort", "open"},
		{"ClosedPort", "closed"},
	}

	host := "localhost"
	hl := &scan.HostsList{}
	hl.Add(host)

	ports := []int{}
	for _, tc := range testCases {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, "0"))
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal(err)
		}
		ports = append(ports, port)

		if tc.name == "ClosedPort" {
			ln.Close()
		}
	}

	res := scan.Run(hl, ports)

	if len(res) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(res))
	}
	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q", host, res[0].Host)
	}
	if res[0].NotFound {
		t.Errorf("Expected host %q to be found", host)
	}
	if len(res[0].PortStates) != 2 {
		t.Fatalf("Expected 2 port states, got %d", len(res[0].PortStates))
	}

	for i, tc := range testCases {
		if res[0].PortStates[i].Port != ports[i] {
			t.Errorf("Expected port %d, got %d", ports[0], res[0].PortStates[i].Port)
		}
		if res[0].PortStates[i].Open.String() != tc.expecteState {
			t.Errorf("Expected port %d to be %s\n", ports[i], tc.expecteState)
		}
	}
}

func TestRunHostNotFound(t *testing.T) {
	host := "389.389.389.389"
	hl := &scan.HostsList{}
	hl.Add(host)
	res := scan.Run(hl, []int{})
	if len(res) != 1 {
		t.Fatalf("Expected 1 result, got %d instead\n", len(res))
	}
	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q\n", host, res[0].Host)
	}
	if !res[0].NotFound {
		t.Errorf("Expected host %q to not be found\n", host)
	}
	if len(res[0].PortStates) != 0 {
		t.Errorf("Expected no port states, got %d\n", len(res[0].PortStates))
	}
}
