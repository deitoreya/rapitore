/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	pkgscreening "github.com/deitoreya/rapitore/pkg/service/screening"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"
)

// screeningCmd represents the screening command
var screeningCmd = &cobra.Command{
	Use:   "screening",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		execScreening()
		fmt.Println("screening called")
	},
}

func init() {
	rootCmd.AddCommand(screeningCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// screeningCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// screeningCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func execScreening() {
	screening := pkgscreening.NewScreening()
	ctx := context.Background()
	if err := screening.Exec(ctx); err != nil {
		slog.ErrorCtx(ctx, "failed exec screening", zap.Error(err))
	}
}
