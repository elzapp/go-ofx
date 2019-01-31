package ofx

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

type BankTransaction struct {
	PostedDate         time.Time
	InterestDate       time.Time
	Ref                string
	DestinationAccount string
	TxType             string
	Memo               string
	Amount             float64
}

type OfxTransactionList struct {
	XMLName          xml.Name         `xml:"OFX"`
	CurDef           string           `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>CURDEF"`
	PayerBank        string           `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKACCTFROM>BANKID"`
	PayerAccount     string           `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKACCTFROM>ACCTID"`
	PayerAccountType string           `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKACCTFROM>ACCTTYPE"`
	Transactions     []OfxTransaction `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS>BANKTRANLIST>STMTTRN"`
}

func (txs *OfxTransactionList) WriteOFX(w io.Writer) {
	header := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<?OFX OFXHEADER="200" VERSION="211" SECURITY="NONE" OLDFILEUID="NONE" NEWFILEUID="NONE"?>
`
	w.Write([]byte(header))
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(txs); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

type OfxTransaction struct {
	XMLName          xml.Name `xml:"STMTTRN"`
	Type             string   `xml:"TRNTYPE"`
	DatePosted       string   `xml:"DTPOSTED"`
	DateUser         string   `xml:"DTUSER"`
	Amount           float64  `xml:"TRNAMT"`
	ID               string   `xml:"FITID"`
	Payee            string   `xml:"NAME"`
	Memo             string   `xml:"MEMO"`
	PayeeBank        string   `xml:"BANKACCTTO>BANKID"`
	PayeeAccount     string   `xml:"BANKACCTTO>ACCTID"`
	PayeeAccountType string   `xml:"BANKACCTTO>ACCTTYPE"`
}

func (btx *BankTransaction) ToOfx() OfxTransaction {
	var otx OfxTransaction
	otx.Type = btx.TxType
	otx.DatePosted = btx.InterestDate.Format("20060102")
	otx.DateUser = btx.PostedDate.Format("20060102")
	otx.Amount = btx.Amount
	otx.Memo = btx.Memo
	return otx
}
