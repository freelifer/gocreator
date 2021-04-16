package cmd

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "open web ui",
	Long:  `UI Web Open`,
	Run: func(cmd *cobra.Command, args []string) {
		open()
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var commands = map[string]string{
	"windows": "start ",
	"darwin":  "open ",
	"linux":   "xdg-open ",
}

// Open calls the OS default program for uri
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	_, err := runCommandWithErr(run + uri)
	return err
}

func open() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pong": "ok",
		})
	})

	Open("http://localhost:8000/ping")

	router.Run(":8000")
}
