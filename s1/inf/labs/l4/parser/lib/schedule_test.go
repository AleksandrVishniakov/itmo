package lib_test

import (
	"io"
	"os"
	"parser/lib"
	"testing"
	"time"
)

func TestParseSchedule(t *testing.T) {
	hclFile, err := os.Open("../test_files/schedule.hcl")
	if err != nil {
		t.Fatal(err)
	}
	defer hclFile.Close()

	yamlFile, err := os.OpenFile("schedule.yaml", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer yamlFile.Close()

	if err := lib.ParseSchedule(hclFile, yamlFile); err != nil {
		t.Fatal(err)
	}
}

func TestParseSchedule100(t *testing.T) {
	hclFile, err := os.Open("../test_files/schedule.hcl")
	if err != nil {
		t.Fatal(err)
	}
	defer hclFile.Close()

	start := time.Now()

	for range 100 {
		_, err = hclFile.Seek(0, io.SeekStart)
		if err != nil {
			panic(err)
		}

		w := io.Discard
		if err := lib.ParseSchedule(hclFile, w); err != nil {
			t.Fatal(err)
		}
	}

	t.Log(time.Since(start))
}
