package nl

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

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
		Contains          string    `validate:"contains=purpose"`
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
		OneOfString       string    `validate:"oneof=red green"`
		OneOfInt          int       `validate:"oneof=5 63"`
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
			expected: "IsColor moet een geldige kleur zijn",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC moet een geldig MAC adres bevatten",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr moet een oplosbaar IP adres zijn",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 moet een oplosbaar IPv4 adres zijn",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 moet een oplosbaar IPv6 adres zijn",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr moet een geldig UDP adres zijn",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 moet een geldig IPv4 UDP adres zijn",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 moet een geldig IPv6 UDP adres zijn",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr moet een geldig TCP adres zijn",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 moet een geldig IPv4 TCP adres zijn",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 moet een geldig IPv6 TCP adres zijn",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR moet een geldige CIDR notatie bevatten",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 moet een geldige CIDR notatie voor een IPv4 adres bevatten",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 moet een geldige CIDR notatie voor een IPv6 adres bevatten",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN moet een geldig SSN nummer zijn",
		},
		{
			ns:       "Test.IP",
			expected: "IP moet een geldig IP adres zijn",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 moet een geldig IPv4 adres zijn",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 moet een geldig IPv6 adres zijn",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI moet een geldige Data URI bevatten",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude moet geldige breedtegraadcoördinaten bevatten",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude moet geldige lengtegraadcoördinaten bevatten",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte moet multibyte karakters bevatten",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII mag alleen ascii karakters bevatten",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII mag alleen afdrukbare ascii karakters bevatten",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID moet een geldige UUID zijn",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 moet een geldige versie 3 UUID zijn",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 moet een geldige versie 4 UUID zijn",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 moet een geldige versie 5 UUID zijn",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID moet een geldige ULID zijn",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN moet een geldig ISBN nummer zijn",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 moet een geldig ISBN-10 nummer zijn",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 moet een geldig ISBN-13 nummer zijn",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes mag niet de tekst 'text' bevatten",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll mag niet een van de volgende karakters bevatten '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune mag niet het volgende bevatten '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny moet tenminste een van de volgende karakters bevatten '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains moet de tekst 'purpose' bevatten",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 moet een geldige Base64 string zijn",
		},
		{
			ns:       "Test.Email",
			expected: "Email moet een geldig email adres zijn",
		},
		{
			ns:       "Test.URL",
			expected: "URL moet een geldige URL zijn",
		},
		{
			ns:       "Test.URI",
			expected: "URI moet een geldige URI zijn",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString moet een geldige RGB kleur zijn",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString moet een geldige RGBA kleur zijn",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString moet een geldige HSL kleur zijn",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString moet een geldige HSLA kleur zijn",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString moet een geldig hexadecimaal getal zijn",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString moet een geldige HEX kleur zijn",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString moet een geldig getal zijn",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString moet een geldige numerieke waarde zijn",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString mag alleen alfanumerieke karakters bevatten",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString mag alleen alfabetische karakters bevatten",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString moet kleiner zijn dan MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString moet kleiner dan of gelijk aan MaxString zijn",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString moet groter zijn dan MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString moet groter dan of gelijk aan MaxString zijn",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString mag niet gelijk zijn aan EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString moet kleiner zijn dan Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString moet kleiner dan of gelijk aan Inner.LteCSFieldString zijn",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString moet groter zijn dan Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString moet groter dan of gelijk aan Inner.GteCSFieldString zijn",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString mag niet gelijk zijn aan Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString moet gelijk zijn aan Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString moet gelijk zijn aan MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString moet tenminste 3 karakters lang zijn",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber moet 5.56 of groter zijn",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple moet tenminste 2 items bevatten",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime moet groter dan of gelijk zijn aan de huidige datum & tijd",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString moet langer dan 3 karakters zijn",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber moet groter zijn dan 5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple moet meer dan 2 items bevatten",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime moet groter zijn dan de huidige datum & tijd",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString mag maximaal 3 karakters lang zijn",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber moet 5.56 of minder zijn",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple mag maximaal 2 items bevatten",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime moet kleiner dan of gelijk aan de huidige datum & tijd zijn",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString moet minder dan 3 karakters lang zijn",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber moet kleiner zijn dan 5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple moet minder dan 2 items bevatten",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime moet kleiner zijn dan de huidige datum & tijd",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString mag niet gelijk zijn aan ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber mag niet gelijk zijn aan 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple mag niet gelijk zijn aan 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString is niet gelijk aan 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber is niet gelijk aan 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple is niet gelijk aan 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString mag maximaal 3 karakters lang zijn",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber moet 1,113.00 of kleiner zijn",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple mag maximaal 7 items bevatten",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString moet tenminste 1 karakter lang zijn",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber moet 1,113.00 of groter zijn",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple moet tenminste 7 items bevatten",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString moet 1 karakter lang zijn",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber moet gelijk zijn aan 1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple moet 7 items bevatten",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString is een verplicht veld",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber is een verplicht veld",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple is een verplicht veld",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen moet tenminste 10 karakters lang zijn",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen mag maximaal 1 karakter lang zijn",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen moet 2 karakters lang zijn",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt moet minder dan 1 karakter lang zijn",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte mag maximaal 1 karakter lang zijn",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt moet langer dan 10 karakters zijn",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte moet tenminste 10 karakters lang zijn",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString moet een van de volgende zijn [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt moet een van de volgende zijn [5 63]",
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
