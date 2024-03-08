package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/dstotijn/go-notion"
)

type NotionClient struct {
	client *notion.Client
}

func NewNotionClient() *NotionClient {
	notionToken := os.Getenv("NOTION_TOKEN")
	notionClient := notion.NewClient(notionToken)

	return &NotionClient{
		client: notionClient,
	}
}

func (n *NotionClient) GetUsers() {
	users, err := n.client.ListUsers(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func (n *NotionClient) AddPage() {

	page := notion.CreatePageParams{
		ParentType: notion.ParentTypeDatabase,
		ParentID:   "cc3b974b957a41f69dc02a0e665f0f08",
		DatabasePageProperties: &notion.DatabasePageProperties{
			"Name": notion.DatabasePageProperty{
				Title: []notion.RichText{
					{
						Text: &notion.Text{
							Content: "Test Page",
						},
					},
				},
			},
		},
	}

	_, err := n.client.CreatePage(context.Background(), page)

	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

	fmt.Println("Page created succesfully")
}

//func markdownToChildren(givenMarkdownText string) {
//	markdownBuffer := []byte(givenMarkdownText)
//
//	err := goldmark.Convert(markdownBuffer, os.Stdout)
//
//	if err != nil {
//		fmt.Fprintln(os.Stdout, err)
//	}
//}
