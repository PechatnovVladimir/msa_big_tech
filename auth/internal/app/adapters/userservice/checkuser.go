package userservice

import "context"

func (c *Client) CheckUser(ctx context.Context, nickname string, password string) (userID string, err error) {
	return "999", nil
}
