package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/hexennacht/signme/enigma/code"
	"github.com/hexennacht/signme/enigma/proto"
)

var root = &cobra.Command{
	Use:   "jormungand",
	Short: "Command line interface tool for handling service",
	Long: `Jormungand is a powerful command-line tool designed to streamline service management. 
From generating protobuf files to deploying services, Jormungand simplifies the entire process with its intuitive interface and comprehensive features. 
With Jormungand, developers can effortlessly generate protobuf files, saving time and eliminating errors. 
Furthermore, this versatile tool enables seamless service deployment, handling dependencies and configuration effortlessly. 
Its flexible command-line interface integrates smoothly into existing workflows, making it suitable for projects of any scale. 
Jormungand empowers developers to efficiently manage services, enhancing productivity and enabling them to focus on building exceptional applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.AddCommand(GenProto())
	},
}

func init() {
	cobra.OnInitialize(initJormungandConfig)
	root.AddCommand(proto.GenProto())
	root.AddCommand(code.GenCode())
}

func initJormungandConfig() {
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
