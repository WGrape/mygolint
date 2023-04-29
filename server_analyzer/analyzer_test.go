package server_analyzer

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestRedisAnalyzer(t *testing.T) {
	var testdataDir = analysistest.TestData()
	fmt.Println(testdataDir)

	analysistest.Run(t, testdataDir, ServerAnalyzer)
}
