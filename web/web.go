package web

import (
	"bytes"
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"github.com/vposloncec/go-ssip/export"
	"net/http"
	"time"
)

func Serve(graph *base.Graph) {
	http.HandleFunc("/nodes", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "text/csv")
		//w.Header().Set("Content-Disposition", "attachment; filename=nodes.csv")
		b := export.NodesToCSV(graph.Nodes)
		reader := bytes.NewReader(b.Bytes())
		http.ServeContent(w, r, "nodes", time.Now(), reader)
	})
	http.HandleFunc("/edges", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "text/csv")
		//w.Header().Set("Content-Disposition", "attachment; filename=edges.csv")
		b := export.EdgesToCSV(graph.Connections)
		reader := bytes.NewReader(b.Bytes())
		http.ServeContent(w, r, "edges", time.Now(), reader)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
