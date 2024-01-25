package export

import (
	"bytes"
	"encoding/csv"
	"github.com/vposloncec/go-ssip/base"
	"os"
)

type edgeExport struct {
}

var edgeHeaders = []string{"node1", "node2"}

func EdgesToCSV(edges []base.ConnectionPair) *bytes.Buffer {
	var b bytes.Buffer

	writer := csv.NewWriter(&b)
	writer.Flush()

	writer.Write(edgeHeaders)
	for _, edge := range edges {
		writer.Write(edgeToRow(&edge))
	}

	return &b
}

func EdgeToCSV(c base.ConnectionPair) {
	var b bytes.Buffer

	writer := csv.NewWriter(&b)
	writer.Write(edgeHeaders)
	writer.Write(edgeToRow(&c))
	writer.Flush()
	b.WriteTo(os.Stdout)
}

func edgeToRow(c *base.ConnectionPair) []string {
	return []string{c[0].String(), c[1].String()}
}
