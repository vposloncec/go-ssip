package export

import (
	"bytes"
	"encoding/csv"
	"github.com/vposloncec/go-ssip/base"
	"go.uber.org/zap"
	"strconv"
)

var nodeHeaders = []string{"id", "packages_received", "reliability"}

func NodesToCSV(nodes []*base.Node) *bytes.Buffer {
	var b bytes.Buffer

	writer := csv.NewWriter(&b)
	defer writer.Flush()

	err := writer.Write(nodeHeaders)
	if err != nil {
		zap.Error(err)
	}

	for _, node := range nodes {
		err := writer.Write(nodeToRow(node))
		if err != nil {
			zap.Error(err)
		}
	}

	return &b
}

func nodeToRow(node *base.Node) []string {
	row := make([]string, len(nodeHeaders))
	row[0] = node.ID.String()
	row[1] = strconv.Itoa(node.PackagesReceived)
	row[2] = node.Reliability.String()

	return row
}
