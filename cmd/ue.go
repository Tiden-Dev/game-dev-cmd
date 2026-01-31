/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// ueCmd represents the ue command
var ueCmd = &cobra.Command{
	Use:   "ue",
	Short: "A set to commands for Unreal Engine",
	Long:  `A set to commands for Unreal Engine`,
	Run: func(cmd *cobra.Command, args []string) {
		p := promptui.Prompt{
			Label:     "Do you want to disable global illumination?",
			IsConfirm: true,
		}
		result, err := p.Run()
		if err != nil {
			logger.Error().Err(err).Msg("Failed to disable global illumination")
			return
		}
		logger.Info().Msg(result)
	},
}

func init() {
	rootCmd.AddCommand(ueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
