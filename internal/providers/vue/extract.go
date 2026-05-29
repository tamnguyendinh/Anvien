package vue

import (
	"github.com/tamnguyendinh/anvien/internal/providers/sfc"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

type Request = sfc.Request

func Extract(request Request) (scopeir.ScopeIR, error) {
	return sfc.Extract(request, sfc.Options{
		Name:            "vue",
		Language:        scanner.Vue,
		ScriptExtractor: sfc.ExtractHTMLScript,
	})
}
