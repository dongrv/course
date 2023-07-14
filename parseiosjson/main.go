package parseiosjson

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func SplitFile(stream []byte) map[string]string {
	var rowsMap = make(map[string]string)
	err := json.Unmarshal(stream, &rowsMap)
	if err != nil {
		return nil
	}
	return rowsMap
}

func Validate() (int, int) {
	var stream, err = ReadFile("./new_ios_pay_logs.json")
	if err != nil {
		println(err)
		return 0, 0
	}

	var iosResp = struct {
		AppAccountToken    string `json:"appAccountToken,omitempty"`
		BundleId           string `json:"bundleId"`
		Environment        string `json:"environment"`
		InAppOwnershipType string `json:"inAppOwnershipType"`
		ProductId          string `json:"productId"`
		PurchaseDate       int64  `json:"purchaseDate"`
		Quantity           int    `json:"quantity"`
		SignedDate         int64  `json:"signedDate"`
		Storefront         string `json:"storefront"`
		TransactionReason  string `json:"transactionReason"`
		Type               string `json:"type"`
	}{}

	var (
		total    int
		minInt64 int64 = math.MaxInt64
		maxInt64 int64
	)

	rows := SplitFile(stream)

	for k, v := range rows {
		right := strings.Split(v, "->")[1]
		err = json.Unmarshal([]byte(right), &iosResp)
		if err != nil {
			println(err.Error())
			break
		}
		if iosResp.AppAccountToken != "" {
			total++
		}
		bytes, _ := json.Marshal(iosResp)
		fmt.Println("transactionId: ", k, "\t", "response: ", string(bytes))
		//fmt.Printf("账单：%s\n", bytes)
		if iosResp.PurchaseDate > maxInt64 {
			maxInt64 = iosResp.PurchaseDate
		}
		if iosResp.PurchaseDate < minInt64 {
			minInt64 = iosResp.PurchaseDate
		}
	}
	fmt.Printf("订单总数: %d\t透传成功订单数:%d \n", len(rows), total)
	return len(rows), total
}
