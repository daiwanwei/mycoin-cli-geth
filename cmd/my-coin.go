package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"mycoin-cli-geth/business/services"
	"mycoin-cli-geth/pkg/eth"
)

var (
	myCoinCmd = &cobra.Command{
		Use:   "myCoin",
		Short: "Gopher CLI in Go",
		Long:  `Gopher CLI application written in Go.`,
	}
)

var balanceOfCmd = &cobra.Command{
	Use:   "balanceOf",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run:   balanceOf,
}

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "This command will get the desired Gopher",
	Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
	Run:   transfer,
}

func init() {
	rootCmd.AddCommand(myCoinCmd)

	myCoinCmd.AddCommand(balanceOfCmd)
	balanceOfCmd.Flags().StringP("providerUrl", "u", "wss://{chain}.infura.io/ws/v3/{KEY}", "provider url")
	balanceOfCmd.Flags().StringP("privateKey", "k", "", "private key")
	balanceOfCmd.Flags().StringP("accountAddress", "a", "", "account address")
	balanceOfCmd.Flags().StringP("contractAddress", "c", "", "contract address")

	myCoinCmd.AddCommand(transferCmd)
	transferCmd.Flags().StringP("providerUrl", "u", "wss://{chain}.infura.io/ws/v3/{KEY}", "provider url")
	transferCmd.Flags().StringP("privateKey", "k", "", "private key")
	transferCmd.Flags().StringP("toAddress", "t", "", "to address")
	transferCmd.Flags().StringP("contractAddress", "c", "", "contract address")
	transferCmd.Flags().Int64P("amount", "a", 1000000, "amount")
}

func balanceOf(cmd *cobra.Command, args []string) {
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
	service, err := services.NewMyCoinService(provider, signer)
	if err != nil {
		fmt.Println(err)
		return
	}

	contractAddress, err := cmd.Flags().GetString("contractAddress")
	if err != nil {
		fmt.Println(err)
		return
	}
	accountAddress, err := cmd.Flags().GetString("accountAddress")
	if err != nil {
		fmt.Println(err)
		return
	}
	balance, err := service.BalanceOf(context.Background(), accountAddress, contractAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("contract address is %s\n", balance)
}

func transfer(cmd *cobra.Command, args []string) {
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
	service, err := services.NewMyCoinService(provider, signer)
	if err != nil {
		fmt.Println(err)
		return
	}

	contractAddress, err := cmd.Flags().GetString("contractAddress")
	if err != nil {
		fmt.Println(err)
		return
	}
	toAddress, err := cmd.Flags().GetString("toAddress")
	if err != nil {
		fmt.Println(err)
		return
	}
	amount, err := cmd.Flags().GetInt64("amount")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = service.Transfer(context.Background(), toAddress, contractAddress, amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("transfer successfully!\n")
}
