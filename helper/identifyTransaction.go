package helper

func IdentifyTransaction(s string) string {
	if s == "C" {
		return "Credit"
	}
	return "Debit"
}