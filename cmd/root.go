package cmd

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "main",
	Short: `
	 _______  _______             _______ _________ _______  _______  _______ 
	(  ____ \(  ___  )           (       )\__   __/(  ____ \(  ____ )(  ___  )
	| (    \/| (   ) |           | () () |   ) (   | (    \/| (    )|| (   ) |
	| |      | |   | |   _____   | || || |   | |   | |      | (____)|| |   | |
	| | ____ | |   | |  (_____)  | |(_)| |   | |   | |      |     __)| |   | |
	| | \_  )| |   | |           | |   | |   | |   | |      | (\ (   | |   | |
	| (___) || (___) |           | )   ( |___) (___| (____/\| ) \ \__| (___) |
	(_______)(_______)           |/     \|\_______/(_______/|/   \__/(_______) by Diki Haryadi
    `,
}

func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	rootCmd.AddCommand(ServeCmd())
	//ServeCmd().PersistentFlags().StringVarP(nil, "config", "c", "", "Config URL i.e. file://config.json")
	ServeCmd().Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(LoadDataCmd())
	//LoadDataCmd().PersistentFlags().StringVarP(nil, "config", "c", "", "Config URL i.e. file://config.json")
	LoadDataCmd().Flags().BoolP("toggle", "t", false, "Help message for toggle")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(1)
	}
}
