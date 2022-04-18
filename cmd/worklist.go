/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	userpb "hkubectl/protos"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var allflag bool
var selectnamespace string

// worklistCmd represents the worklist command
var worklistCmd = &cobra.Command{
	Use:   "worklist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("worklist called")
		grpcgetworklistdata()
	},
}

func init() {
	getCmd.AddCommand(worklistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// worklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// worklistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	worklistCmd.Flags().BoolVarP(&allflag, "all", "A", false, "Get All NameSpace")
	// podCmd.Flags().StringP("namespace", "n", "", "Get All Pod")
	worklistCmd.Flags().StringVarP(&selectnamespace, "namespace", "n", "default", "Get NameSpace Pod")
}

func grpcgetworklistdata() {
	conn, err := grpc.Dial("10.0.5.24:30036", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserClient(conn)

	// Contact the server and print out its response.
	//name := "gpu"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	var r *userpb.GetWorkListResponse
	if allflag {
		r, err = c.GetWorkList(ctx, &userpb.GetWorkListRequest{Namespace: ""})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	} else {
		r, err = c.GetWorkList(ctx, &userpb.GetWorkListRequest{Namespace: selectnamespace})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}

	//log.Printf("Greeting: %s", r.GetPodMessage())
	runtime := strings.Split(r.GetWorklistMessage.Runtime, " ")
	namespace := strings.Split(r.GetWorklistMessage.Namespace, " ")
	name := strings.Split(r.GetWorklistMessage.Name, " ")
	ready := strings.Split(r.GetWorklistMessage.Ready, " ")
	status := strings.Split(r.GetWorklistMessage.Status, " ")
	restart := strings.Split(r.GetWorklistMessage.Restart, " ")
	age := strings.Split(r.GetWorklistMessage.Age, " ")
	fmt.Printf("RunTime       Namespace       Name                                       Ready   Status       Restart  Age\n")
	//fmt.Println(r.GetWorklistMessage)
	for i := 1; i < len(strings.Split(r.GetWorklistMessage.Runtime, " ")); i++ {
		fmt.Printf("%-11s   %-13s   %-40s   %-5s   %-10s   %-6s   %-5s\n", runtime[i], namespace[i], name[i], ready[i], status[i], restart[i], SecondChange(age[i]))
		//fmt.Println(age[i])

	}
	//time.Sleep(1 * time.Second)
	cancel()
}

func SecondChange(second string) string {
	intsecond, _ := strconv.Atoi(second)
	if intsecond >= 864000 {
		return strconv.Itoa(intsecond/86400) + "d"
	} else if intsecond >= 86400 {
		return strconv.Itoa(intsecond/86400) + "d" + strconv.Itoa(intsecond%86400/3600) + "h"
	} else if intsecond >= 3600 {
		return strconv.Itoa(intsecond/3600) + "h"
	} else if intsecond >= 60 {
		return strconv.Itoa(intsecond/60) + "m"
	} else {
		return strconv.Itoa(intsecond) + "s"
	}
}
