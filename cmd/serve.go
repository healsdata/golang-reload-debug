package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

const (
	readTimeout  = 5
	writeTimeout = 10
	idleTimeout  = 120
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	w.WriteHeader(returnStatus)
	message := fmt.Sprintf("Hello World! %s", r.UserAgent())
	w.Write([]byte(message))
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Hello World web server",
	Long:  `Hello World web server`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddress := ":8080"
		l := log.New(os.Stdout, "sample-srv ", log.LstdFlags|log.Lshortfile)
		m := mux.NewRouter()

		m.HandleFunc("/", IndexHandler)

		srv := &http.Server{
			Addr:         serverAddress,
			ReadTimeout:  readTimeout * time.Second,
			WriteTimeout: writeTimeout * time.Second,
			IdleTimeout:  idleTimeout * time.Second,
			Handler:      m,
		}

		l.Printf("server started at http://localhost%s", serverAddress)
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
