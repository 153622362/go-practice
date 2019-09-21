package rootest

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"it's a simpale application")
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "hello",
		Short: "Print the version number of Hugo",
		Long:  `All software has versions. This is Hugo's`,
		Run:   HelloWorld,
	})
}

func HelloWorld(cmd *cobra.Command, args []string) {
	fmt.Println("Hello World!")
	fmt.Println(args)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
