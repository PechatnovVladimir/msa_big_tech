package userservice

import "context"

func (c *Client) CreateUser(ctx context.Context, nickname string, email string, password string) (userID string, err error) {
	return "999", nil
}
