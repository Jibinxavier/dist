package cmd

import (
	"dist/pkg/api"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
)

var (
	port string
)

func init() {
	rootCmd.AddCommand(startSimpleLeader)
	startSimpleLeader.Flags().StringVarP(&port, "port", "p", "", "port")
}

var startSimpleLeader = &cobra.Command{
	Use:   "smleader start",
	Short: "To run the leader ",
	Long:  `Start leader work`,
	Run: func(cmd *cobra.Command, args []string) {

		router := httprouter.New()
		router.GET("/", api.Index)
		router.GET("/health", api.Health)
		router.GET("/hello/:name", api.Hello)

		log.Fatal(http.ListenAndServe(":"+port, router))
	},
}
