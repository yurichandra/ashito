package message

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
)

var CbsSpec = iso8583.MessageSpec{
	Fields: map[int]field.Field{
		0: field.NewString(&field.Spec{
			Length:      4,
			Description: "Message Type Indicator",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		1: field.NewBitmap(&field.Spec{
			Description: "Secondary Bit-Map",
			Enc:         encoding.Binary,
			Pref:        prefix.ASCII.Fixed,
		}),

		// Message fields:
		2: field.NewString(&field.Spec{
			Length:      24,
			Description: "Primary Account Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		3: field.NewString(&field.Spec{
			Length:      6,
			Description: "Processing Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			// Pad:         padding.Left('0'),
		}),
		4: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Transaction",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		5: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Account",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		6: field.NewString(&field.Spec{
			Length:      12,
			Description: "Amount, Cardholder Billing",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		7: field.NewString(&field.Spec{
			Length:      10,
			Description: "Transmission Date and Time",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		// Unused
		// 8: field.NewString(&field.Spec{
		// 	Length:      8,
		// 	Description: "Billing Fee Amount",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		9: field.NewString(&field.Spec{
			Length:      8,
			Description: "Conversion Rate, Cardholder Billing",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		10: field.NewString(&field.Spec{
			Length:      8,
			Description: "Conversion Rate, Account",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		11: field.NewString(&field.Spec{
			Length: 6,
			// Description: "Systems Trace Audit Number (STAN)",
			Description: "Systems Trace Audit Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		12: field.NewString(&field.Spec{
			Length:      12,
			Description: "Time, Local Transaction",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// 13: field.NewString(&field.Spec{
		// 	Length:      4,
		// 	Description: "Local Transaction Date",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		14: field.NewString(&field.Spec{
			Length: 4,
			// Description: "Expiration Date",
			Description: "Date, Expiration",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		15: field.NewString(&field.Spec{
			Length:      6,
			Description: "Settlement Date",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// 16: field.NewString(&field.Spec{
		// 	Length:      4,
		// 	Description: "Currency Conversion Date",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// 17: field.NewString(&field.Spec{
		// 	Length:      4,
		// 	Description: "Capture Date",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		18: field.NewString(&field.Spec{
			Length:      4,
			Description: "Merchant Type",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// Unused
		// 19: field.NewString(&field.Spec{
		// 	Length:      3,
		// 	Description: "Acquiring Institution Country Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// Unused
		// 20: field.NewString(&field.Spec{
		// 	Length:      3,
		// 	Description: "PAN Extended Country Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// 21: field.NewString(&field.Spec{
		// 	Length:      3,
		// 	Description: "Forwarding Institution Country Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		22: field.NewString(&field.Spec{
			Length: 12,
			// Description: "Point of Sale (POS) Entry Mode",
			Description: "Point of Service Date Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// Unused
		// 23: field.NewString(&field.Spec{
		// 	Length:      3,
		// 	Description: "Card Sequence Number (CSN)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		24: field.NewString(&field.Spec{
			Length:      3,
			Description: "Function Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// 25: field.NewString(&field.Spec{
		// 	Length:      2,
		// 	Description: "Point of Service Condition Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// 26: field.NewString(&field.Spec{
		// 	Length:      2,
		// 	Description: "Point of Service PIN Capture Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// Unused
		// 27: field.NewString(&field.Spec{
		// 	Length:      1,
		// 	Description: "Authorizing Identification Response Length",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// // Unused
		// 28: field.NewString(&field.Spec{
		// 	Length:      9,
		// 	Description: "Transaction Fee Amount",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// 29: field.NewString(&field.Spec{
		// 	Length:      9,
		// 	Description: "Settlement Fee Amount",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// 30: field.NewString(&field.Spec{
		// 	Length:      9,
		// 	Description: "Transaction Processing Fee Amount",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		// 31: field.NewString(&field.Spec{
		// 	Length:      9,
		// 	Description: "Settlement Processing Fee Amount",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		32: field.NewString(&field.Spec{
			Length:      11,
			Description: "Acquiring Institution Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		// 33: field.NewString(&field.Spec{
		// 	Length:      11,
		// 	Description: "Forwarding Institution Identification Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LL,
		// }),
		// Unused
		// 34: field.NewString(&field.Spec{
		// 	Length:      30,
		// 	Description: "Extended Primary Account Number",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LL,
		// }),
		// // Unused
		// 35: field.NewString(&field.Spec{
		// 	Length:      37,
		// 	Description: "Track 2 Data",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LL,
		// }),
		// 36: field.NewString(&field.Spec{
		// 	Length:      104,
		// 	Description: "Track 3 Data",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		37: field.NewString(&field.Spec{
			Length:      12,
			Description: "Retrieval Reference Number",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		38: field.NewString(&field.Spec{
			Length:      6,
			Description: "Authorization Identification Response",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		39: field.NewString(&field.Spec{
			Length:      3,
			Description: "Response Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// 40: field.NewString(&field.Spec{
		// 	Length:      3,
		// 	Description: "Service Restriction Code",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		41: field.NewString(&field.Spec{
			Length:      8,
			Description: "Card Acceptor Terminal Identification",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		42: field.NewString(&field.Spec{
			Length:      15,
			Description: "Card Acceptor Identification Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		43: field.NewString(&field.Spec{
			Length: 40,
			// Description: "Card Acceptor Name/Location",
			Description: "Card Acceptor Name and Location",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// 44: field.NewString(&field.Spec{
		// 	Length:      99,
		// 	Description: "Additional Data",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LL,
		// }),
		// 45: field.NewString(&field.Spec{
		// 	Length:      76,
		// 	Description: "Track 1 Data",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LL,
		// }),
		// 46: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Additional data (ISO)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		// Unused
		// 47: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Additional data (National)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		48: field.NewString(&field.Spec{
			Length: 999,
			// Description: "Additional data (Private)",
			Description: "Additional Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		49: field.NewString(&field.Spec{
			Length: 3,
			// Description: "Transaction Currency Code",
			Description: "Currency Code, Transaction",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		50: field.NewString(&field.Spec{
			Length: 3,
			// Description: "Settlement Currency Code",
			Description: "Currency Code, Account",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		51: field.NewString(&field.Spec{
			Length: 3,
			// Description: "Cardholder Billing Currency Code",
			Description: "Currency Code, Cardholder Billing",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
		}),
		// 52: field.NewString(&field.Spec{
		// 	Length:      8,
		// 	Description: "PIN Data",
		// 	Enc:         encoding.Hex,
		// 	Pref:        prefix.Hex.Fixed,
		// }),
		// 53: field.NewString(&field.Spec{
		// 	Length:      16,
		// 	Description: "Security Related Control Information",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		54: field.NewString(&field.Spec{
			Length:      999,
			Description: "Additional Amounts",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		55: field.NewString(&field.Spec{
			Length: 255,
			// Description: "ICC Data – EMV Having Multiple Tags",
			Description: "EMV Data",
			// HMMMM encoding is: llvar b...2555
			Enc:  encoding.ASCII,
			Pref: prefix.ASCII.LLL,
		}),
		// 56: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Reserved (ISO)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		// 57: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Reserved (National)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		// 58: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Reserved (National)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		// 59: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Reserved (National)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		// 60: field.NewString(&field.Spec{
		// 	Length:      999,
		// 	Description: "Reserved (National)",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.LLL,
		// }),
		61: field.NewString(&field.Spec{
			Length: 999,
			// Description: "Reserved (Private)",
			Description: "Additional Amounts and Counters",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		62: field.NewString(&field.Spec{
			// TODO: hmm how do i represent this, because the length is dependant on protocol version
			// https://app.clickup.com/t/f2a7wm
			Length:      999,
			Description: "Mini-statement data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		63: field.NewString(&field.Spec{
			Length:      12,
			Description: "Account Balance",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.Fixed,
			Pad:         padding.Left('0'),
		}),
		64: field.NewString(&field.Spec{
			Length: 8,
			// Description: "Message Authentication Code (MAC)",
			Description: "Primary MAC Data",
			Enc:         encoding.BytesToASCIIHex,
			Pref:        prefix.Hex.Fixed,
		}),
		// 90: field.NewString(&field.Spec{
		// 	Length:      42,
		// 	Description: "Original Data Elements",
		// 	Enc:         encoding.ASCII,
		// 	Pref:        prefix.ASCII.Fixed,
		// }),
		91: field.NewString(&field.Spec{
			Length:      1,
			Description: "Action Code",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		95: field.NewString(&field.Spec{
			Length:      999,
			Description: "Replacement Amounts",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLL,
		}),
		100: field.NewString(&field.Spec{
			Length:      4,
			Description: "SVFE Issuer Institution Identifier",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		102: field.NewString(&field.Spec{
			Length:      32,
			Description: "Account Identification",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		103: field.NewString(&field.Spec{
			Length:      32,
			Description: "Account Identification - 2",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LL,
		}),
		124: field.NewString(&field.Spec{
			Length:      99999,
			Description: "Private Data",
			Enc:         encoding.ASCII,
			Pref:        prefix.ASCII.LLLL,
		}),
		128: field.NewString(&field.Spec{
			Length:      8,
			Description: "Secondary MAC Data",
			Enc:         encoding.BytesToASCIIHex,
			Pref:        prefix.Hex.Fixed,
		}),
	},
}
