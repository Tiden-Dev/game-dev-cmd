/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"path"

	"github.com/Tiden-Dev/game-dev-cmd/ue"
	"github.com/spf13/cobra"
)

// ueCmd represents the ue command
var ueCmd = &cobra.Command{
	Use:   "ue",
	Short: "A set to commands for Unreal Engine",
	Long:  `A set to commands for Unreal Engine`,
	Run: func(cmd *cobra.Command, args []string) {
		fp, err := ue.FindUProject()
		if err != nil {
			logger.Fatal().Err(err).Msg("no unreal project in the working directory")
			return
		}
		project, err := ue.ReadUProject(fp)
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to read project file")
			return
		}

		fp = path.Join(ue.ConfigDir, ue.DefaultGameIniFileName)
		defaultGame, err := ue.LoadDefaultGame(fp, ue.EngineVersion(project.EngineAssociation))
		if err != nil {
			logger.Error().Err(err).Msg("failed to load default game")
			return
		}

		name, _ := defaultGame.GetProjectName()
		logger.Info().Str("name", name).Str("engine", project.EngineAssociation).Msg("unreal project found!")
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
