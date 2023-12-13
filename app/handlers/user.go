package handlers

import (
	"github.com/getfider/fider/app/models/cmd"
	"github.com/getfider/fider/app/pkg/bus"
	"github.com/getfider/fider/app/pkg/web"
)

// BlockUser is used to block an existing user from using Fider
func BlockUser() web.HandlerFunc {
	return func(c *web.Context) error {
		userID, err := c.ParamAsInt("userID")
		if err != nil {
			return c.NotFound()
		}

		err = bus.Dispatch(c, &cmd.BlockUser{UserID: userID})
		if err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{})
	}
}

// UnblockUser is used to unblock an existing user so they can use Fider again
func UnblockUser() web.HandlerFunc {
	return func(c *web.Context) error {
		userID, err := c.ParamAsInt("userID")
		if err != nil {
			return c.NotFound()
		}

		err = bus.Dispatch(c, &cmd.UnblockUser{UserID: userID})
		if err != nil {
			return c.Failure(err)
		}

		return c.Ok(web.Map{})
	}
}
