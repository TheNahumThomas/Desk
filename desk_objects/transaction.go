package objects

type Transaction struct {
	TransactionID string
	LoginLocation string
}

func (t *Transaction) GetTransactionID() string {
	return t.TransactionID
}
