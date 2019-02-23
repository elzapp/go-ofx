package ofx

import (
	"fmt"
	"testing"
	"time"
)

func TestSingleTransaction(t *testing.T) {
	var btx BankTransaction
	btx.PostedDate = time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	ofx := btx.ToOfx()
	if ofx.DateUser != "20190101" {
		fmt.Println(ofx, "incorrect date")
		t.Fail()
	}
}
