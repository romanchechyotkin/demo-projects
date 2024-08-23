package redis

import "fmt"

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Log() {
	fmt.Println("redis log")
}
