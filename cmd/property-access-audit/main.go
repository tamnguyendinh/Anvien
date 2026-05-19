package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tamnguyendinh/avmatrix-go/internal/graphaccuracy"
)

func main() {
	repo := flag.String("repo", ".", "repository root used to read source lines")
	graphPath := flag.String("graph", "", "AVmatrix graph snapshot JSON")
	outPath := flag.String("out", "", "output JSON path")
	maxExamples := flag.Int("max-examples", 50, "maximum examples to include per bucket")
	flag.Parse()

	if *graphPath == "" || *outPath == "" {
		exitf("missing -graph or -out")
	}

	result, err := graphaccuracy.RunPropertyAccessAudit(graphaccuracy.PropertyAccessAuditOptions{
		Repo:        *repo,
		GraphPath:   *graphPath,
		OutPath:     *outPath,
		MaxExamples: *maxExamples,
	})
	if err != nil {
		exitf("%v", err)
	}

	fmt.Printf("wrote %s\n", *outPath)
	for _, line := range graphaccuracy.PropertyAccessAuditSummaryLines(result) {
		fmt.Println(line)
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
