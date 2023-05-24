package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	transactionsJson, _ := json.MarshalIndent(transactions, "", "")
	fmt.Println(string(transactionsJson))

}
