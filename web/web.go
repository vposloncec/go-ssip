package web

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"github.com/vposloncec/go-ssip/base"
	"github.com/vposloncec/go-ssip/export"
	"net/http"
	_ "net/http/pprof"
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

	port := viper.GetInt("port")
	fmt.Printf("Server is running on http://localhost:%v\n", port)
	fmt.Println("Available endpoints are /nodes and /edges")

	addr := fmt.Sprintf(":%v", port)
	http.ListenAndServe(addr, nil)
}
