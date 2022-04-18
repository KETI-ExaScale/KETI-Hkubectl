/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	userpb "hkubectl/protos"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("delete called")
		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("add needs a name for the yaml file"))
		}
		data, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("read file error", err)
		}
		grpcdeletework(data)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	var yamlflag bool
	var namespaceflag string
	deleteCmd.Flags().BoolVarP(&yamlflag, "file", "f", false, "Delete to yaml")
	deleteCmd.Flags().StringVarP(&namespaceflag, "namespace", "n", "default", "Delete to Workname and namespace")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func grpcdeletework(data []byte) {
	conn, err := grpc.Dial("10.0.5.24:30036", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserClient(conn)

	// Contact the server and print out its response.
	//name := "gpu"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = c.DeleteWork(ctx, &userpb.DeleteWorkRequest{Yamldata: string(data)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Delete Work\n")
	//log.Printf("Greeting: %s", r.GetPodMessage())

	time.Sleep(1 * time.Second)
	cancel()
}
