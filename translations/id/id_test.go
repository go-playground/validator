package id

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	indonesia "github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	idn := indonesia.New()
	uni := ut.New(idn, idn)
	trans, _ := uni.GetTranslator("id")

	validate := validator.New()

	err := RegisterDefaultTranslations(validate, trans)
	Equal(t, err, nil)

	type Inner struct {
		EqCSFieldString  string
		NeCSFieldString  string
		GtCSFieldString  string
		GteCSFieldString string
		LtCSFieldString  string
		LteCSFieldString string
	}

	type Test struct {
		Inner             Inner
		RequiredString    string    `validate:"required"`
		RequiredNumber    int       `validate:"required"`
		RequiredMultiple  []string  `validate:"required"`
		LenString         string    `validate:"len=1"`
		LenNumber         float64   `validate:"len=1113.00"`
		LenMultiple       []string  `validate:"len=7"`
		MinString         string    `validate:"min=1"`
		MinNumber         float64   `validate:"min=1113.00"`
		MinMultiple       []string  `validate:"min=7"`
		MaxString         string    `validate:"max=3"`
		MaxNumber         float64   `validate:"max=1113.00"`
		MaxMultiple       []string  `validate:"max=7"`
		EqString          string    `validate:"eq=3"`
		EqNumber          float64   `validate:"eq=2.33"`
		EqMultiple        []string  `validate:"eq=7"`
		NeString          string    `validate:"ne="`
		NeNumber          float64   `validate:"ne=0.00"`
		NeMultiple        []string  `validate:"ne=0"`
		LtString          string    `validate:"lt=3"`
		LtNumber          float64   `validate:"lt=5.56"`
		LtMultiple        []string  `validate:"lt=2"`
		LtTime            time.Time `validate:"lt"`
		LteString         string    `validate:"lte=3"`
		LteNumber         float64   `validate:"lte=5.56"`
		LteMultiple       []string  `validate:"lte=2"`
		LteTime           time.Time `validate:"lte"`
		GtString          string    `validate:"gt=3"`
		GtNumber          float64   `validate:"gt=5.56"`
		GtMultiple        []string  `validate:"gt=2"`
		GtTime            time.Time `validate:"gt"`
		GteString         string    `validate:"gte=3"`
		GteNumber         float64   `validate:"gte=5.56"`
		GteMultiple       []string  `validate:"gte=2"`
		GteTime           time.Time `validate:"gte"`
		EqFieldString     string    `validate:"eqfield=MaxString"`
		EqCSFieldString   string    `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string    `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string    `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string    `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string    `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string    `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string    `validate:"nefield=EqFieldString"`
		GtFieldString     string    `validate:"gtfield=MaxString"`
		GteFieldString    string    `validate:"gtefield=MaxString"`
		LtFieldString     string    `validate:"ltfield=MaxString"`
		LteFieldString    string    `validate:"ltefield=MaxString"`
		AlphaString       string    `validate:"alpha"`
		AlphanumString    string    `validate:"alphanum"`
		NumericString     string    `validate:"numeric"`
		NumberString      string    `validate:"number"`
		HexadecimalString string    `validate:"hexadecimal"`
		HexColorString    string    `validate:"hexcolor"`
		RGBColorString    string    `validate:"rgb"`
		RGBAColorString   string    `validate:"rgba"`
		HSLColorString    string    `validate:"hsl"`
		HSLAColorString   string    `validate:"hsla"`
		Email             string    `validate:"email"`
		URL               string    `validate:"url"`
		URI               string    `validate:"uri"`
		Base64            string    `validate:"base64"`
		Contains          string    `validate:"contains=tujuan"`
		ContainsAny       string    `validate:"containsany=!@#$"`
		Excludes          string    `validate:"excludes=text"`
		ExcludesAll       string    `validate:"excludesall=!@#$"`
		ExcludesRune      string    `validate:"excludesrune=☻"`
		ISBN              string    `validate:"isbn"`
		ISBN10            string    `validate:"isbn10"`
		ISBN13            string    `validate:"isbn13"`
		UUID              string    `validate:"uuid"`
		UUID3             string    `validate:"uuid3"`
		UUID4             string    `validate:"uuid4"`
		UUID5             string    `validate:"uuid5"`
		ULID              string    `validate:"ulid"`
		ASCII             string    `validate:"ascii"`
		PrintableASCII    string    `validate:"printascii"`
		MultiByte         string    `validate:"multibyte"`
		DataURI           string    `validate:"datauri"`
		Latitude          string    `validate:"latitude"`
		Longitude         string    `validate:"longitude"`
		SSN               string    `validate:"ssn"`
		IP                string    `validate:"ip"`
		IPv4              string    `validate:"ipv4"`
		IPv6              string    `validate:"ipv6"`
		CIDR              string    `validate:"cidr"`
		CIDRv4            string    `validate:"cidrv4"`
		CIDRv6            string    `validate:"cidrv6"`
		TCPAddr           string    `validate:"tcp_addr"`
		TCPAddrv4         string    `validate:"tcp4_addr"`
		TCPAddrv6         string    `validate:"tcp6_addr"`
		UDPAddr           string    `validate:"udp_addr"`
		UDPAddrv4         string    `validate:"udp4_addr"`
		UDPAddrv6         string    `validate:"udp6_addr"`
		IPAddr            string    `validate:"ip_addr"`
		IPAddrv4          string    `validate:"ip4_addr"`
		IPAddrv6          string    `validate:"ip6_addr"`
		UinxAddr          string    `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string    `validate:"mac"`
		IsColor           string    `validate:"iscolor"`
		StrPtrMinLen      *string   `validate:"min=10"`
		StrPtrMaxLen      *string   `validate:"max=1"`
		StrPtrLen         *string   `validate:"len=2"`
		StrPtrLt          *string   `validate:"lt=1"`
		StrPtrLte         *string   `validate:"lte=1"`
		StrPtrGt          *string   `validate:"gt=10"`
		StrPtrGte         *string   `validate:"gte=10"`
		OneOfString       string    `validate:"oneof=merah hijau"`
		OneOfInt          int       `validate:"oneof=5 63"`
		Image			  string			`validate:"image"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"

	test.MaxString = "1234"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)

	test.LtString = "1234"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteTime = time.Now().Add(time.Hour * 24)

	test.LtFieldString = "12345"
	test.LteFieldString = "12345"

	test.LtCSFieldString = "1234"
	test.LteCSFieldString = "1234"

	test.AlphaString = "abc3"
	test.AlphanumString = "abc3!"
	test.NumericString = "12E.00"
	test.NumberString = "12E"

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ☻"

	test.ASCII = "ｶﾀｶﾅ"
	test.PrintableASCII = "ｶﾀｶﾅ"

	test.MultiByte = "1234feerf"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	err = validate.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(validator.ValidationErrors)
	Equal(t, ok, true)

	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "Test.IsColor",
			expected: "IsColor harus berupa warna yang valid",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC harus berisi alamat MAC yang valid",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr harus berupa alamat IP yang dapat dipecahkan",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 harus berupa alamat IPv4 yang dapat diatasi",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 harus berupa alamat IPv6 yang dapat diatasi",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr harus berupa alamat UDP yang valid",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 harus berupa alamat IPv4 UDP yang valid",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 harus berupa alamat IPv6 UDP yang valid",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr harus berupa alamat TCP yang valid",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 harus berupa alamat TCP IPv4 yang valid",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 harus berupa alamat TCP IPv6 yang valid",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR harus berisi notasi CIDR yang valid",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 harus berisi notasi CIDR yang valid untuk alamat IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 harus berisi notasi CIDR yang valid untuk alamat IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN harus berupa nomor SSN yang valid",
		},
		{
			ns:       "Test.IP",
			expected: "IP harus berupa alamat IP yang valid",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 harus berupa alamat IPv4 yang valid",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 harus berupa alamat IPv6 yang valid",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI harus berisi URI Data yang valid",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude harus berisi koordinat lintang yang valid",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude harus berisi koordinat bujur yang valid",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte harus berisi karakter multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII hanya boleh berisi karakter ascii",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII hanya boleh berisi karakter ascii yang dapat dicetak",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID harus berupa UUID yang valid",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 harus berupa UUID versi 3 yang valid",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 harus berupa UUID versi 4 yang valid",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 harus berupa UUID versi 5 yang valid",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID harus berupa ULID yang valid",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN harus berupa nomor ISBN yang valid",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 harus berupa nomor ISBN-10 yang valid",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 harus berupa nomor ISBN-13 yang valid",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes tidak boleh berisi teks 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll tidak boleh berisi salah satu karakter berikut '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune tidak boleh berisi '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny harus berisi setidaknya salah satu karakter berikut '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains harus berisi teks 'tujuan'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 harus berupa string Base64 yang valid",
		},
		{
			ns:       "Test.Email",
			expected: "Email harus berupa alamat email yang valid",
		},
		{
			ns:       "Test.URL",
			expected: "URL harus berupa URL yang valid",
		},
		{
			ns:       "Test.URI",
			expected: "URI harus berupa URI yang valid",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString harus berupa warna RGB yang valid",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString harus berupa warna RGBA yang valid",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString harus berupa warna HSL yang valid",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString harus berupa warna HSLA yang valid",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString harus berupa heksadesimal yang valid",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString harus berupa warna HEX yang valid",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString harus berupa angka yang valid",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString harus berupa nilai numerik yang valid",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString hanya dapat berisi karakter alfanumerik",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString hanya dapat berisi karakter abjad",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString harus kurang dari MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString harus kurang dari atau sama dengan MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString harus lebih besar dari MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString harus lebih besar dari atau sama dengan MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString tidak sama dengan EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString harus kurang dari Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString harus kurang dari atau sama dengan Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString harus lebih besar dari Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString harus lebih besar dari atau sama dengan Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString tidak sama dengan Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString harus sama dengan Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString harus sama dengan MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "panjang minimal GteString adalah 3 karakter",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber harus 5,56 atau lebih besar",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple harus berisi setidaknya 2 item",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime harus lebih besar dari atau sama dengan tanggal & waktu saat ini",
		},
		{
			ns:       "Test.GtString",
			expected: "panjang GtString harus lebih dari 3 karakter",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber harus lebih besar dari 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple harus berisi lebih dari 2 item",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime harus lebih besar dari tanggal & waktu saat ini",
		},
		{
			ns:       "Test.LteString",
			expected: "panjang maksimal LteString adalah 3 karakter",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber harus 5,56 atau kurang",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple harus berisi maksimal 2 item",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime harus kurang dari atau sama dengan tanggal & waktu saat ini",
		},
		{
			ns:       "Test.LtString",
			expected: "panjang LtString harus kurang dari 3 karakter",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber harus kurang dari 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple harus berisi kurang dari 2 item",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime harus kurang dari tanggal & waktu saat ini",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString tidak sama dengan ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber tidak sama dengan 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple tidak sama dengan 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString tidak sama dengan 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber tidak sama dengan 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple tidak sama dengan 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "panjang maksimal MaxString adalah 3 karakter",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber harus 1.113,00 atau kurang",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple harus berisi maksimal 7 item",
		},
		{
			ns:       "Test.MinString",
			expected: "panjang minimal MinString adalah 1 karakter",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber harus 1.113,00 atau lebih besar",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "panjang minimal MinMultiple adalah 7 item",
		},
		{
			ns:       "Test.LenString",
			expected: "panjang LenString harus 1 karakter",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber harus sama dengan 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple harus berisi 7 item",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString wajib diisi",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber wajib diisi",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple wajib diisi",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "panjang minimal StrPtrMinLen adalah 10 karakter",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "panjang maksimal StrPtrMaxLen adalah 1 karakter",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "panjang StrPtrLen harus 2 karakter",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "panjang StrPtrLt harus kurang dari 1 karakter",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "panjang maksimal StrPtrLte adalah 1 karakter",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "panjang StrPtrGt harus lebih dari 10 karakter",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "panjang minimal StrPtrGte adalah 10 karakter",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString harus berupa salah satu dari [merah hijau]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt harus berupa salah satu dari [5 63]",
		},
		{
			ns: "Test.Image",
			expected: "Image harus berupa gambar yang valid",
		},
	}

	for _, tt := range tests {

		var fe validator.FieldError

		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fe = e
				break
			}
		}

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}

}
