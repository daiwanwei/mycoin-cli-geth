package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mycoin-cli-geth/contractevents"
	"mycoin-cli-geth/pkg/eth"
)

var (
	contractEventCmd = &cobra.Command{
		Use:   "contractEvent",
		Short: "Gopher CLI in Go",
		Long:  `Gopher CLI application written in Go.`,
	}
)

var runMyCoinManagerCmd = &cobra.Command{
	Use:   "runMyCoinManager",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run:   runMyCoinManager,
}

func init() {
	rootCmd.AddCommand(contractEventCmd)

	contractEventCmd.AddCommand(runMyCoinManagerCmd)
	runMyCoinManagerCmd.Flags().StringP("providerUrl", "u", "wss://{chain}.infura.io/ws/v3/{KEY}", "provider url")
}

func runMyCoinManager(cmd *cobra.Command, args []string) {
	providerUrl, err := cmd.Flags().GetString("providerUrl")
	if err != nil {
		fmt.Println(err)
		return
	}
	provider, err := eth.NewProvider(providerUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	service, err := contractevents.NewMyCoinSubManager(provider)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("run mycoin manager")
	err = service.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
