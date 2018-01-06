package cli

import (
	gcli "github.com/urfave/cli"

	"github.com/spaco/spo/src/cipher"
)

func verifyAddressCmd() gcli.Command {
	name := "verifyAddress"
	return gcli.Command{
		Name:         name,
		Usage:        "Verify a spo address",
		ArgsUsage:    "[spo address]",
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) error {
			skyAddr := c.Args().First()
			_, err := cipher.DecodeBase58Address(skyAddr)
			return err
		},
	}
}
