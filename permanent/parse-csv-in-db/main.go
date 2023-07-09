package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Transaction struct {
	TransactionDate string
	PostDate        string
	Description     string
	Category        string
	Type            string
	Amount          float32
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)

	var transactions []Transaction

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		amount, _ := strconv.ParseFloat(line[5], 32)

		transactions = append(transactions,
			Transaction{
				TransactionDate: line[0],
				PostDate:        line[1],
				Description:     line[2],
				Category:        line[3],
				Type:            line[4],
				Amount:          float32(amount),
			})
	}

	fmt.Printf("ALDI: %v \n", getTotalAmountSpent(transactions, "aldi"))
	fmt.Printf("NETFIX: %v \n", getTotalAmountSpent(transactions, "netflix"))
	fmt.Printf("YouTube Premium: %v \n", getTotalAmountSpent(transactions, "youtube"))
	fmt.Printf("Apple: %v \n", getTotalAmountSpent(transactions, "apple"))
	fmt.Printf("Harvest Time: %v \n", getTotalAmountSpent(transactions, "HARVESTIME"))
	fmt.Printf("Spotify: %v \n", getTotalAmountSpent(transactions, "Spotify"))
	fmt.Printf("Laundry: %v \n", getTotalAmountSpent(transactions, "OPS"))

	//By category
	fmt.Printf("Car Gas: %v \n", getTotalAmountSpentBasedOnCategory(transactions, "gas"))
	fmt.Printf("Food: %v \n", getTotalAmountSpentBasedOnCategory(transactions, "food"))

	// for _, v := range transactions {
	// 	if strings.Contains(v.Description, "NETFLIX") {
	// 		fmt.Printf("%s \n", v)
	// 	}
	// }

	// fmt.Println()
	// fmt.Println(len(transactions))

	// transactionsJson, _ := json.MarshalIndent(transactions, "", "")
	// fmt.Println(string(transactionsJson))

}
func getTotalAmountSpent(transactions []Transaction, substring string) float32 {
	var total float32 = 0.0
	for _, v := range transactions {
		desc := strings.ToLower(v.Description)
		match := strings.ToLower(substring)
		if strings.Contains(desc, match) {
			total += v.Amount
		}
	}
	return total
}

func getTotalAmountSpentBasedOnCategory(transactions []Transaction, substring string) float32 {
	var total float32 = 0.0
	for _, v := range transactions {
		desc := strings.ToLower(v.Category)
		match := strings.ToLower(substring)
		if strings.Contains(desc, match) {
			total += v.Amount
		}
	}
	return total
}
