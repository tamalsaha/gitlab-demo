package main

import (
	"fmt"
	"log"
	"os"

	gitlab "github.com/xanzy/go-gitlab"
)

/*
Tamal Saha
Name: appscode FULL_NAME: appscode FullPath: appscode
Name: ghh FULL_NAME: ghh FullPath: ghh
Name: kubedb FULL_NAME: kubedb FullPath: kubedb
Name: kubepack FULL_NAME: kubepack FullPath: kubepack
Name: kubeware FULL_NAME: kubeware FullPath: kubeware
Name: pharmer FULL_NAME: pharmer FullPath: pharmer
Name: subpharm FULL_NAME: pharmer / subpharm FullPath: pharmer/subpharm
*/
func main() {
	token := os.Getenv("GITLAB_TOKEN")

	client := gitlab.NewClient(nil, token)

	user, _, err := client.Users.CurrentUser()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(user.Name)

	var groups []string
	// https://docs.gitlab.com/ee/api/README.html#pagination
	page := 1
	pageSize := 20
	for {
		list, _, err := client.Groups.ListGroups(&gitlab.ListGroupsOptions{
			ListOptions: gitlab.ListOptions{Page: page, PerPage: pageSize},
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, g := range list {
			fmt.Println("Name: " + g.Name + " FULL_NAME: " + g.FullName + " FullPath: " + g.FullPath)
			groups = append(groups, g.Name)
		}
		if len(list) < pageSize {
			break
		}
		page++
	}
}

