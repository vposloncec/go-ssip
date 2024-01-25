package export

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"strconv"
)

var nodeHeaders = []string{"id", "packages_received", "reliability"}

func NodesToCSV(nodes []*base.Node) *bytes.Buffer {
	var b bytes.Buffer

	writer := csv.NewWriter(&b)
	defer writer.Flush()

	writer.Write(nodeHeaders)
	for _, node := range nodes {
		writer.Write(nodeToRow(node))
	}

	return &b
}

func nodeToRow(node *base.Node) []string {
	row := make([]string, len(nodeHeaders))
	row[0] = node.ID.String()
	row[1] = strconv.Itoa(node.PackagesReceived)
	row[2] = fmt.Sprint(node.Reliability)

	return row
}
