package validator

import (
	"encoding/json"
	"testing"
)

type benchmarkSignupRequest struct {
	Name         string   `validate:"required,min=2,max=80"`
	Email        string   `validate:"required,email"`
	Age          int      `validate:"gte=18,lte=120"`
	Password     string   `validate:"required,min=12"`
	Confirmation string   `validate:"eqfield=Password"`
	Country      string   `validate:"oneof=US GB DE IN CA"`
	Marketing    *bool    `validate:"omitempty"`
	Topics       []string `validate:"max=5,dive,oneof=news releases offers"`
}

type benchmarkOrderRequest struct {
	CustomerID string `validate:"required,uuid4"`
	Currency   string `validate:"oneof=USD EUR GBP"`
	Shipping   struct {
		Street  string `validate:"required,max=120"`
		City    string `validate:"required,max=80"`
		Postal  string `validate:"required,max=12"`
		Country string `validate:"len=2"`
	}
	Items []struct {
		SKU      string  `validate:"required,alphanum"`
		Quantity int     `validate:"gte=1,lte=100"`
		Price    float64 `validate:"gt=0"`
	} `validate:"required,min=1,max=100,dive"`
	Metadata    map[string]string `validate:"max=10,dive,keys,required,endkeys,required,max=80"`
	GiftMessage *string           `validate:"omitempty,max=200"`
}

func BenchmarkRequestSignup(b *testing.B) {
	benchmarkRequest[benchmarkSignupRequest](b, []byte(`{
		"Name":"Sample User","Email":"sample@example.com","Age":32,
		"Password":"sample-password","Confirmation":"sample-password",
		"Country":"US","Marketing":true,"Topics":["news","releases"]
	}`), func(r *benchmarkSignupRequest) {
		r.Email = "invalid"
		r.Age = 15
		r.Confirmation = "different"
	})
}

func BenchmarkRequestOrder(b *testing.B) {
	benchmarkRequest[benchmarkOrderRequest](b, []byte(`{
		"CustomerID":"550e8400-e29b-41d4-a716-446655440000","Currency":"USD",
		"Shipping":{"Street":"123 Sample Street","City":"London","Postal":"SW1A 1AA","Country":"GB"},
		"Items":[
			{"SKU":"SKU001","Quantity":2,"Price":19.95},
			{"SKU":"SKU002","Quantity":1,"Price":49.50},
			{"SKU":"SKU003","Quantity":3,"Price":5.25},
			{"SKU":"SKU004","Quantity":1,"Price":99.00},
			{"SKU":"SKU005","Quantity":2,"Price":12.00},
			{"SKU":"SKU006","Quantity":1,"Price":8.75},
			{"SKU":"SKU007","Quantity":4,"Price":3.50},
			{"SKU":"SKU008","Quantity":1,"Price":24.99}
		],
		"Metadata":{"source":"web","campaign":"summer","locale":"en-GB"},
		"GiftMessage":"Happy birthday"
	}`), func(r *benchmarkOrderRequest) {
		r.Shipping.Postal = ""
		r.Items[0].Quantity = 0
		r.Items[3].Price = -1
		r.Metadata["source"] = ""
	})
}

// benchmarkRequest models a reused validator with 90% valid requests and 10%
// invalid requests. DecodeValidate includes JSON decoding, but no HTTP or I/O.
func benchmarkRequest[T any](b *testing.B, validJSON []byte, invalidate func(*T)) {
	b.Helper()
	v := New(WithRequiredStructEnabled())
	var requests [10]T
	var payloads [10][]byte
	for i := range requests {
		if err := json.Unmarshal(validJSON, &requests[i]); err != nil {
			b.Fatal(err)
		}
		if i == len(requests)-1 {
			invalidate(&requests[i])
		}
		payload, err := json.Marshal(&requests[i])
		if err != nil {
			b.Fatal(err)
		}
		payloads[i] = payload
		if err := v.Struct(&requests[i]); (err != nil) != (i == len(requests)-1) {
			b.Fatalf("request %d has unexpected validation result: %v", i, err)
		}
	}

	for _, mode := range []string{"Validate", "DecodeValidate"} {
		b.Run(mode, func(b *testing.B) {
			run := func(i int) {
				if mode == "DecodeValidate" {
					var request T
					if err := json.Unmarshal(payloads[i], &request); err != nil {
						b.Error(err)
						return
					}
					_ = v.Struct(&request)
					return
				}
				_ = v.Struct(&requests[i])
			}
			b.Run("Serial", func(b *testing.B) {
				b.ReportAllocs()
				i := 0
				for b.Loop() {
					run(i % len(requests))
					i++
				}
			})
			b.Run("Parallel", func(b *testing.B) {
				b.ReportAllocs()
				b.RunParallel(func(pb *testing.PB) {
					i := 0
					for pb.Next() {
						run(i % len(requests))
						i++
					}
				})
			})
		})
	}
}
