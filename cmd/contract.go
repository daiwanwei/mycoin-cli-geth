package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"mycoin-cli-geth/business/services"
	"mycoin-cli-geth/pkg/eth"
)

var (
	contractCmd = &cobra.Command{
		Use:   "contract",
		Short: "Gopher CLI in Go",
		Long:  `Gopher CLI application written in Go.`,
	}
)

var deployMyCoinCmd = &cobra.Command{
	Use:   "deployMyCoin",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run:   deployMyCoin,
}

func init() {
	rootCmd.AddCommand(contractCmd)

	contractCmd.AddCommand(deployMyCoinCmd)
	deployMyCoinCmd.Flags().StringP("providerUrl", "u", "wss://{chain}.infura.io/ws/v3/{KEY}", "provider url")
	deployMyCoinCmd.Flags().StringP("privateKey", "k", "", "private key")
}

func deployMyCoin(cmd *cobra.Command, args []string) {
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
	privateKey, err := cmd.Flags().GetString("privateKey")
	if err != nil {
		fmt.Println(err)
		return
	}
	signer, err := eth.NewSinger(provider, privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	service, err := services.NewContractService(provider, signer)
	if err != nil {
		fmt.Println(err)
		return
	}

	contractAddress, err := service.DeployMyCoin(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("contract address is %s\n", contractAddress)
}
