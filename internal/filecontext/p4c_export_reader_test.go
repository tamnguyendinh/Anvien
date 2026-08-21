package filecontext

import (
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
)

func TestP4CExportedSymbolUsesCanonicalExportFieldBeforeAccessVisibility(t *testing.T) {
	tests := []struct {
		name string
		node graph.Node
		want bool
	}{
		{
			name: "canonical false does not inherit public visibility",
			node: graph.Node{Properties: graph.NodeProperties{"isExported": false, "visibility": "public"}},
			want: false,
		},
		{
			name: "canonical true is independent of private visibility",
			node: graph.Node{Properties: graph.NodeProperties{"isExported": true, "visibility": "private"}},
			want: true,
		},
		{
			name: "legacy graph without canonical field keeps visibility fallback",
			node: graph.Node{Properties: graph.NodeProperties{"visibility": "public"}},
			want: true,
		},
		{
			name: "legacy exported compatibility field remains supported",
			node: graph.Node{Properties: graph.NodeProperties{"exported": true}},
			want: true,
		},
		{
			name: "malformed canonical field fails closed",
			node: graph.Node{Properties: graph.NodeProperties{"isExported": "true", "visibility": "public"}},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := exportedSymbol(test.node); got != test.want {
				t.Fatalf("exportedSymbol() = %v, want %v for %#v", got, test.want, test.node.Properties)
			}
		})
	}
}
