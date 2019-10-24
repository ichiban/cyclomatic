package cyclomatic

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestFromFileSystem(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
}
