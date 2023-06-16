package cmd

import (
    "encoding/json"
	"fmt"
    "io/ioutil"
	
	"github.com/martinlindhe/base36"
	"github.com/spf13/cobra"
	"github.com/fxamacker/cbor/v2"
)

type Data struct {
    Name string `json:"name"`
    Url string `json:"url"`
    Key []byte `json:"public-key"`
}

var cimCmd = &cobra.Command{
	Use:   "cim <file>",
	Short: "Compiles an installation manifest",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		byteValue, _ := ioutil.ReadFile(filename)
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