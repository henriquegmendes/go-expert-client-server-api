package client

import (
	"fmt"
	"log"
	"os"
)

func SaveLatestQuoteInFile(quote float64) error {
	err := os.MkdirAll("./result", os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create("./result/cotacao.txt")
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("error closing file: %s", err.Error())
		}
	}()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %f", quote))
	if err != nil {
		return err
	}

	return nil
}
