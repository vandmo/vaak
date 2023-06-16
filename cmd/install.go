package cmd

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/martinlindhe/base36"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install <blob>",
	Short: "Installs from installation manifest blob",
	Args:  cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		blob := args[0]
		var data Data
		b := base36.DecodeToBytes(blob)
		cbor.Unmarshal(b, &data)
		fmt.Println(data.Name)
		fmt.Println(data.Url)
		fmt.Println(data.Key)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
