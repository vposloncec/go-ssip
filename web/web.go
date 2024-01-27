package web

import (
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"github.com/vposloncec/go-ssip/export"
	"go.uber.org/zap"
	"net/http"
)

func Serve(graph *base.Graph) {
	http.HandleFunc("/nodes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=nodes.csv")
		_, err := export.NodesToCSV(graph.Nodes).WriteTo(w)
		if err != nil {
			return
		}
	})
	http.HandleFunc("/edges", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=edges.csv")
		b := export.EdgesToCSV(graph.Connections)
		var err error
		_, err = b.WriteTo(w)
		if err != nil {
			zap.Error(err)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
