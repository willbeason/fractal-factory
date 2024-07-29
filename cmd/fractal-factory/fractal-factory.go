package main

import (
	"context"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	ctx := context.Background()

	err := mainCmd().ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func mainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `fractal-factory`,
		Short: `Run the game Fractal Factory`,
		RunE:  runCmd,
	}

	return cmd
}

func runCmd(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true


	return nil
}