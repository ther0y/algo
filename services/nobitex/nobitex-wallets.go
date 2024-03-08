package nobitex

import (
	"fmt"
	"strings"
)

func GetWallets() ([]Wallet, error) {
	var result GetWalletsResponse

	resp, err := client.R().
		SetSuccessResult(&result).
		Get("/users/wallets/list")

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return result.Wallets, nil
}

func GetWalletsByCurrency(currency string) ([]WalletByCurrency, error) {
	var result GetWalletsByCurrencyResponse

	resp, err := client.R().
		SetSuccessResult(&result).
		SetQueryParam("currencies", strings.ToLower(currency)).
		Get("/v2/wallets")

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	if !resp.IsSuccessState() {
		fmt.Println(resp.String())

		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	var wallets []WalletByCurrency
	for _, wallet := range result.Wallets {
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

func GetWalletsByCurrencies(currencies []string) ([]WalletByCurrency, error) {
	var result GetWalletsByCurrencyResponse

	resp, err := client.R().
		SetSuccessResult(&result).
		SetQueryParam("currency", strings.Join(currencies, ",")).
		Get("/users/wallets")

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccessState() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	var wallets []WalletByCurrency
	for _, wallet := range result.Wallets {
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}
