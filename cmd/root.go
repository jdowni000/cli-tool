/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-tool localhost:8080 [PATH]. PATH can be 'root','/','$id'",
	Short: "A simple cli tool that hits a local running image/container on port 8080",
	Long: `A simple cli tool designed to help retrieve information from a local running 
	that is exposed on port 8080. This tool was designed specifically to target 
	https://github.com/jdowni000/gameserver that simply returns json information.
	This can be achieved by running cli-tool localhost:8080 [PATH]. PATH can be 'root','/','$id'`,

	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		u, err := url.Parse(args[0])

		if err != nil {
			panic(err)
		}

		host := u.Hostname()
		port := u.Port()
		path := u.Path

		if host == "" {
			host = "localhost"
		}

		if port == "" {
			port = "8080"
		}

		if path == "" {
			path = "/"
		}

		if args[1] == "root" || args[1] == "/" {
			path = "/"
			fmt.Println("Retrieving information from root")
		} else {
			path = fmt.Sprintf("/game?id=${%s}", args[1])
		}

		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		fmt.Fprintf(conn, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", path, host)

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf[:n]))
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	fmt.Println("*****************************************************************************************")
	fmt.Println("Welcome to cli-tool to run against localhost:8080! Please run --help for more information")
	fmt.Println("*****************************************************************************************")

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
