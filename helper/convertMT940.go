package helper

import (
	"strings"
)

type MT940Response struct {
	BankSender                 string              `json:"bank_sender,omitempty"`
	BankReceiver               string              `json:"bank_receiver,omitempty"`
	TransactionReferenceNumber string              `json:"transaction_reference_number,omitempty"`
	AccountIdentification      string              `json:"account_identification,omitempty"`
	StatementNumber            string              `json:"statement_number,omitempty"`
	SequenceNumber             string              `json:"sequence_number,omitempty"`
	Kind                       string              `json:"kind,omitempty"`
	Date                       string              `json:"date,omitempty"`
	Currency                   string              `json:"currency,omitempty"`
	OpeningBalance             string              `json:"opening_balance,omitempty"`
	History                    []map[string]string `json:"history,omitempty"`
	ClosingBalance             map[string]string   `json:"closing_balance,omitempty"`
}

func ConvertMT940(path string) MT940Response {
	data, _ := ReadFile(path)
	dataArray := strings.Split(data, "\n")
	dataArrayLineOne := strings.Split(dataArray[0], "}")
	bankSender := dataArrayLineOne[0][6:]
	bankReceiver := dataArrayLineOne[1][7:]
	transactionReferenceNumber := dataArray[1][4:]
	accountIdentification := dataArray[2][4:]
	dataArrayLineFour := strings.Split(dataArray[3], "/")
	statementNumber := dataArrayLineFour[0][5:]
	sequenceNumber := "0"
	if v := len(dataArrayLineFour); v > 1 {
		sequenceNumber = dataArrayLineFour[1]
	}
	openingBalance := dataArray[4][5:]
	kindOfOpeningBalance := IdentifyTransaction(string(openingBalance[0]))
	balanceDate := ReverseDateYYMMDDtoDDMMYY(openingBalance[1:7])
	currencyCode := openingBalance[7:10]
	finalOpeningBalanceAmount := openingBalance[10:]
	history := []map[string]string{
		{
			"name": "transaction one",
			"date": ReverseDateYYMMDDtoDDMMYY(dataArray[5][4:10]),
			"kind": IdentifyTransaction(string(dataArray[5][10])),
		},
		{
			"name": "transaction two",
			"date": ReverseDateYYMMDDtoDDMMYY(dataArray[7][4:10]),
			"kind": IdentifyTransaction(string(dataArray[7][10])),
		},
	}
	closingBalance := map[string]string{
		"kind":     IdentifyTransaction(string(dataArray[9][5])),
		"date":     ReverseDateYYMMDDtoDDMMYY(dataArray[9][6:12]),
		"currency": dataArray[9][12:15],
		"amount":   dataArray[9][15:],
	}

	return MT940Response{
		BankSender:                 bankSender,
		BankReceiver:               bankReceiver,
		TransactionReferenceNumber: transactionReferenceNumber,
		AccountIdentification:      accountIdentification,
		StatementNumber:            statementNumber,
		SequenceNumber:             sequenceNumber,
		Kind:                       kindOfOpeningBalance,
		Date:                       balanceDate,
		Currency:                   currencyCode,
		OpeningBalance:             finalOpeningBalanceAmount,
		History:                    history,
		ClosingBalance:             closingBalance,
	}
}
