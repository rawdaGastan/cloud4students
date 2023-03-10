// Package cmd to make it cmd app
/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/rawdaGastan/cloud4students/routes"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloud4Students",
	Short: "Cloud helps students to deploy their projects",
	Long: `Cloud for students helps them to deploy their projects with 
	applying for a voucher, For example :
		They can deploy virtual machine that can be small, medium or large.
		They can deploy Kubernetes that can be small, medium, or large.
		The Amount of resources available will depend on their voucher.
		`,

	Run: func(cmd *cobra.Command, args []string) {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatal().Err(err)
		}
		server, err := routes.NewServer(config)
		if err != nil {
			log.Fatal().Err(err)
		}

		err = server.Start()
		if err != nil {
			log.Fatal().Err(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("config", "c", "./config.json", "Enter your configurations path")
}
