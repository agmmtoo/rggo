package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{
			"NoFilter",
			"testdata",
			config{ext: "", size: 0, list: true},
			"testdata/dir.log\ntestdata/dir2/script.sh\n",
		},
		{
			"FilterExtensionMatch",
			"testdata",
			config{ext: ".log", size: 0, list: true},
			"testdata/dir.log\n",
		},
		{
			"FilterExtensionSizeMatch",
			"testdata",
			config{ext: ".log", size: 10, list: true},
			"testdata/dir.log\n",
		},
		{
			"FilterExtensionSizeNoMatch",
			"testdata",
			config{ext: ".gz", size: 0, list: true},
			"",
		},
		{
			"FilterExtensionNoMatch",
			"testdata",
			config{".gz", 0, true},
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buffer.String()
			if tc.expected != res {
				t.Errorf("Expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}
