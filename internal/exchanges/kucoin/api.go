package kucoin

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/boises-finest-dao/investmentdao-backend/internal/models"
)

type KucoinServer struct {
	*kucoin.ApiService
}

type KuCoinTradingBalance struct {
	Timestamp           time.Time                              `json:"timestamp"`
	TradingBalances     map[string]KuCoinTradinCurrencyBalance `json:"trading_balances"`
	TotalTradingBalance float64                                `json:"total_trading_balance"`
}

type KuCoinTradinCurrencyBalance struct {
	CurrencyName    string  `json:"currency_name"`
	FiatValue       float64 `json:"fiat_value"`
	AmountAvailable float64 `json:"amount_available"`
	AmountOnHold    float64 `json:"amount_on_hold"`
	TotalAmount     float64 `json:"total_amount"`
	UsdAmount       float64 `json:"usd_amount"`
	UsdValue        string  `json:"usd_value"`
	TotalValue      string  `json:"total_value"`
}

func ConfigureConnection(key string, secret string, pass string) *KucoinServer {
	// Setup KuCoin API
	ks := &KucoinServer{
		kucoin.NewApiService(
			// kucoin.ApiBaseURIOption("https://api.kucoin.com"),
			kucoin.ApiKeyOption(key),
			kucoin.ApiSecretOption(secret),
			kucoin.ApiPassPhraseOption(pass),
			kucoin.ApiKeyVersionOption(kucoin.ApiKeyVersionV2)),
	}

	return ks
}

func (ks *KucoinServer) GetTradingBalances(tradingBalanceID uint) *models.ExchangeBalance {
	// Without pagination
	rsp, err := ks.Accounts("", "")
	if err != nil {
		// Handle error
		return nil
	}

	as := kucoin.AccountsModel{}
	if err := rsp.ReadData(&as); err != nil {
		// Handle error
		return nil
	}

	var coinsArray []string
	for _, a := range as {
		coinsArray = append(coinsArray, a.Currency)
	}

	coinsString := strings.Join(coinsArray, ",")

	c, err := ks.Prices("USD", coinsString)
	if err != nil {
		return nil
	}

	cas := kucoin.PricesModel{}
	if err := c.ReadData(&cas); err != nil {
		fmt.Println(err)
		return nil
	}

	// tradingBalance := &KuCoinTradingBalance{
	// 	Timestamp:           time.Now(),
	// 	TradingBalances:     make(map[string]KuCoinTradinCurrencyBalance),
	// 	TotalTradingBalance: 0,
	// }

	tradingBalance := &models.ExchangeBalance{
		TradingBalanceID: tradingBalanceID,
	}

	for _, a := range as {
		if a.Type == "trade" {
			fiatValue, _ := strconv.ParseFloat(cas[a.Currency], 8)
			amountAvailable, _ := strconv.ParseFloat(a.Available, 8)
			amountOnHold, _ := strconv.ParseFloat(a.Holds, 8)
			totalAmount := amountAvailable + amountOnHold
			usdAmount := fiatValue * totalAmount
			usdValue := strconv.FormatFloat(usdAmount, 'f', 2, 64)
			totalValue := strconv.FormatFloat(totalAmount, 'f', 8, 64)

			if totalAmount > 0 {
				tradingBalance.ExchangeCurrencyBalances = append(tradingBalance.ExchangeCurrencyBalances, models.ExchangeCurrencyBalance{
					CurrencyName:    a.Currency,
					FiatValue:       fiatValue,
					AmountAvailable: amountAvailable,
					AmountOnHold:    amountOnHold,
					TotalAmount:     totalAmount,
					UsdAmount:       usdAmount,
					UsdValue:        usdValue,
					TotalValue:      totalValue,
				})

				tradingBalance.TotalExchangeBalance += usdAmount
			}
		}
	}

	totalCoinsString := strconv.FormatFloat(tradingBalance.TotalExchangeBalance, 'f', 2, 64)
	log.Printf("Total Trading Value: $%s", totalCoinsString)

	return tradingBalance
}
