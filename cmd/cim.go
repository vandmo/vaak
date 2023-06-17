package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fxamacker/cbor/v2"
	"github.com/martinlindhe/base36"
	"github.com/spf13/cobra"
)

type Data struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Key  []byte `json:"public-key"`
}

var cimCmd = &cobra.Command{
	Use:   "cim <file>",
	Short: "Compiles an installation manifest",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		byteValue, _ := os.ReadFile(filename)
		var data Data
		json.Unmarshal(byteValue, &data)
		b, _ := cbor.Marshal(data)
		fmt.Println(base36.EncodeBytes(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cimCmd)
}
