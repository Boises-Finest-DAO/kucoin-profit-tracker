package models

import "gorm.io/gorm"

type Fund struct {
	gorm.Model
	BotID           uint             `json:"bot_id"`
	Name            string           `json:"name"`
	Description     string           `json:"host"`
	Bot             Bot              `json:"bots"`
	TradingBalances []TradingBalance `json:"trading_balances"`
}

type TradingBalance struct {
	gorm.Model
	FundID              uint              `json:"fund_id"`
	ExchangeBalances    []ExchangeBalance `json:"exchange_balances"`
	TotalTradingBalance float64           `json:"total_trading_balance"`
}

type ExchangeBalance struct {
	gorm.Model
	TradingBalanceID         uint                      `json:"trading_balance"`
	ExchangeCurrencyBalances []ExchangeCurrencyBalance `json:"exchange_currency_balances"`
	TotalExchangeBalance     float64                   `json:"total_exchange_balance"`
}

type ExchangeCurrencyBalance struct {
	gorm.Model
	ExchangeBalanceID uint    `json:"exchange_balance_id"`
	CurrencyName      string  `json:"currency_name"`
	FiatValue         float64 `json:"fiat_value"`
	AmountAvailable   float64 `json:"amount_available"`
	AmountOnHold      float64 `json:"amount_on_hold"`
	TotalAmount       float64 `json:"total_amount"`
	UsdAmount         float64 `json:"usd_amount"`
	UsdValue          string  `json:"usd_value"`
	TotalValue        string  `json:"total_value"`
}
