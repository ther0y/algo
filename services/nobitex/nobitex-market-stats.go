package nobitex

import "fmt"

func GetMarketStats(srcCurrency string, dstCurrency string) (MarketStats, error) {
	var result MarketStatsResponse

	resp, err := client.R().
		SetSuccessResult(&result).
		SetQueryParam("srcCurrency", srcCurrency).
		SetQueryParam("dstCurrency", dstCurrency).
		Get("/market/stats")

	if err != nil {
		return MarketStats{}, err
	}

	if !resp.IsSuccessState() {
		return MarketStats{}, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	return result.Stats[srcCurrency+"-"+dstCurrency], nil
}
