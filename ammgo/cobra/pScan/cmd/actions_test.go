package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"agmmtoo.me/ammgo/cobra/pScan/scan"
)

func setup(t *testing.T, hosts []string, initList bool) (string, func()) {
	tf, err := os.CreateTemp("", "pScan")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	if initList {
		hl := &scan.HostsList{}
		for _, h := range hosts {
			hl.Add(h)
		}
		if err := hl.Save(tf.Name()); err != nil {
			t.Fatal(err)
		}
	}
	return tf.Name(), func() {
		os.Remove(tf.Name())
	}
}

func TestHostActions(t *testing.T) {
	hosts := []string{
		"host1",
		"host2",
		"host3",
	}

	testCases := []struct {
		name           string
		args           []string
		expectedOut    string
		initList       bool
		actionFunction func(io.Writer, string, []string) error
	}{
		{
			name:           "AddAction",
			args:           hosts,
			expectedOut:    "Added host: host1\nAdded host: host2\nAdded host: host3\n",
			initList:       false,
			actionFunction: addAction,
		},
		{
			name:           "ListAction",
			args:           hosts,
			expectedOut:    "host1\nhost2\nhost3\n",
			initList:       true,
			actionFunction: listAction,
		},
		{
			name:           "DeleteAction",
			args:           []string{"host1", "host2"},
			expectedOut:    "Deleted host: host1\nDeleted host: host2\n",
			initList:       true,
			actionFunction: deleteAction,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tf, cleanup := setup(t, tc.args, tc.initList)
			defer cleanup()
			var out bytes.Buffer
			if err := tc.actionFunction(&out, tf, tc.args); err != nil {
				t.Fatal(err)
			}
			if out.String() != tc.expectedOut {
				t.Errorf("Expected output: %q, got: %q\n", tc.expectedOut, out.String())
			}
		})
	}
}
