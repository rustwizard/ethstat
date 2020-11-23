package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run ethstat server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		// Routes
		e.GET("/", hello)

		// Start server
		e.Logger.Fatal(e.Start(":9000"))
	},
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVar(&cfgFile, "config", "", "-- config config.yaml")
}
