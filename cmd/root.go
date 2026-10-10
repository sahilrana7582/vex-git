package cmd

import (
	"fmt"

	"github.com/sahilrana.com/vex/internal/repo"
	"github.com/sahilrana.com/vex/internal/ui"
	"github.com/spf13/cobra"
)

func printBanner() {
	fmt.Println(ui.Banner.Render(`
    ██╗   ██╗███████╗██╗  ██╗
    ██║   ██║██╔════╝╚██╗██╔╝
    ██║   ██║█████╗   ╚███╔╝
    ╚██╗ ██╔╝██╔══╝   ██╔██╗
     ╚████╔╝ ███████╗██╔╝ ██╗
      ╚═══╝  ╚══════╝╚═╝  ╚═╝
`))
}

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "vex",
		Short: "vex — a minimal version control system",
		Long: ui.Banner.Render(`
    ██╗   ██╗███████╗██╗  ██╗
    ██║   ██║██╔════╝╚██╗██╔╝
    ██║   ██║█████╗   ╚███╔╝
    ╚██╗ ██╔╝██╔══╝   ██╔██╗
     ╚████╔╝ ███████╗██╔╝ ██╗
      ╚═══╝  ╚══════╝╚═╝  ╚═╝

vex — a minimal version control system built to learn how Git works.
`),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			switch cmd.Name() {
			case "init", "help", "completion":
				return nil
			}
			return nil
		},
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a repository",
		Run: func(cmd *cobra.Command, args []string) {
			fresh, err := repo.Init()
			if err != nil {
				fmt.Println(ui.Error.Render("Error:"), err)
				return
			}

			printBanner()

			if fresh {
				fmt.Println(ui.Success.Render("Initialized empty Vex repository in .vex/"))
			} else {
				fmt.Println(ui.Success.Render("Reinitialized existing Vex repository in .vex/"))
			}
		},
	}

	rootCmd.AddCommand(initCmd)

	_ = rootCmd.Execute()
}
