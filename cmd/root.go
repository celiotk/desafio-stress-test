/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"stress_test/internal/infra"
	"stress_test/internal/usecase"

	"github.com/spf13/cobra"
)

type RunFunc func(cmd *cobra.Command, args []string)

var rootCmd *cobra.Command

func NewRootCmd(uc *usecase.StressTestUsecase) *cobra.Command {
	return &cobra.Command{
		Use:   "stress_test",
		Short: "Testes de carga em um serviço web",
		Long: `stress_test é um sistema CLI em Go que realiza testes de carga em um serviço web.
	O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.
	Ao final do teste o sistema gerará um relatório contendo:
	- Tempo total gasto na execução.
	- Quantidade total de requests realizados.
	- Quantidade de requests com status HTTP 200.
	- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: runCreate(uc),
	}
}

func runCreate(uc *usecase.StressTestUsecase) RunFunc {
	return func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		usecaseInput := usecase.StressTestInputDTO{
			URL:         url,
			Requests:    requests,
			Concurrency: concurrency,
		}
		res, err := uc.Run(ctx, usecaseInput)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
		fmt.Printf("Result:\n%+v\n", string(data))
		// cmd.Help()
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	req := infra.NewServiceRequest()
	uc := usecase.NewStressTest(req)
	rootCmd = NewRootCmd(uc)
	rootCmd.Flags().StringP("url", "u", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntP("requests", "r", 100, " Número total de requests")
	rootCmd.Flags().IntP("concurrency", "c", 10, "Número de chamadas simultâneas")
	rootCmd.MarkFlagRequired("url")
}
