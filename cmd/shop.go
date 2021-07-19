package cmd

import (
	"encoding/json"
	"fiftyfive/service/checkout"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Execute function will receive and create the root commands
func Execute(logger *zap.Logger, svcName, ver string) {
	var (
		fileName = "price_rules.json"
		rootCmd  = &cobra.Command{
			Use:     "Checkout",
			Short:   "CheckoutTweet Test",
			Long:    `This is supermarket checkout page`,
			Version: ver,
			// Pre Run will set PriceRules so we can all priceRule set before the execution
			PreRun: func(cmd *cobra.Command, args []string) {
				ruleFile, err := os.Open(fileName)
				if err != nil {
					logger.Fatal("Failed to Open Price Rules file : ", zap.Error(err))
				}
				priceRulesData, _ := ioutil.ReadAll(ruleFile)
				err = json.Unmarshal(priceRulesData, &checkout.PriceRules)
				if err != nil {
					logger.Fatal("Failed to decode Json Data : ", zap.Error(err))
				}
			},
			Run: func(cmd *cobra.Command, args []string) {
				// Scan the product input
				// This can be params/ Here using scanf to enter the value from user
				for {
					var skuName string
					fmt.Println("Please enter Product SKU / exit if done")
					fmt.Scanf("%s", &skuName)
					if skuName == "exit" {
						total, err := checkout.CalculateTotal()
						if err != nil {
							logger.Error("Error to calculate total")
						}
						fmt.Println("Total Amount : ", total)
						break
					}
					_, checkSku := checkout.IsSkuExist(skuName)
					if !checkSku {
						fmt.Println("Please enter Valid SKU")
						continue
					}
					checkout.Scan(logger, skuName)
				}

			},
		}
	)

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("failed to execute command", zap.Error(err))
	}

}
