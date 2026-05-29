package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/tamnguyendinh/anvien/internal/graphaccuracy"
)

func main() {
	repo := flag.String("repo", ".", "repository root to analyze")
	outPath := flag.String("out", "", "output JSON path")
	maxExamples := flag.Int("max-examples", 50, "maximum examples to include per bucket")
	flag.Parse()

	if *outPath == "" {
		exitf("missing -out")
	}

	result, err := graphaccuracy.RunAccessCandidateAudit(context.Background(), graphaccuracy.AccessCandidateAuditOptions{
		Repo:        *repo,
		OutPath:     *outPath,
		MaxExamples: *maxExamples,
	})
	if err != nil {
		exitf("%v", err)
	}

	fmt.Printf("wrote %s\n", *outPath)
	for _, line := range graphaccuracy.AccessCandidateAuditSummaryLines(result) {
		fmt.Println(line)
	}
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
