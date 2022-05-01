package it

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	italian "github.com/go-playground/locales/it"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {
	ita := italian.New()
	uni := ut.New(ita, ita)
	trans, _ := uni.GetTranslator("it")

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
		Inner               Inner
		RequiredString      string            `validate:"required"`
		RequiredNumber      int               `validate:"required"`
		RequiredMultiple    []string          `validate:"required"`
		LenString           string            `validate:"len=1"`
		LenNumber           float64           `validate:"len=1113.00"`
		LenMultiple         []string          `validate:"len=7"`
		MinString           string            `validate:"min=1"`
		MinNumber           float64           `validate:"min=1113.00"`
		MinMultiple         []string          `validate:"min=7"`
		MaxString           string            `validate:"max=3"`
		MaxNumber           float64           `validate:"max=1113.00"`
		MaxMultiple         []string          `validate:"max=7"`
		EqString            string            `validate:"eq=3"`
		EqNumber            float64           `validate:"eq=2.33"`
		EqMultiple          []string          `validate:"eq=7"`
		NeString            string            `validate:"ne="`
		NeNumber            float64           `validate:"ne=0.00"`
		NeMultiple          []string          `validate:"ne=0"`
		LtString            string            `validate:"lt=3"`
		LtNumber            float64           `validate:"lt=5.56"`
		LtMultiple          []string          `validate:"lt=2"`
		LtTime              time.Time         `validate:"lt"`
		LteString           string            `validate:"lte=3"`
		LteNumber           float64           `validate:"lte=5.56"`
		LteMultiple         []string          `validate:"lte=2"`
		LteTime             time.Time         `validate:"lte"`
		GtString            string            `validate:"gt=3"`
		GtNumber            float64           `validate:"gt=5.56"`
		GtMultiple          []string          `validate:"gt=2"`
		GtTime              time.Time         `validate:"gt"`
		GteString           string            `validate:"gte=3"`
		GteNumber           float64           `validate:"gte=5.56"`
		GteMultiple         []string          `validate:"gte=2"`
		GteTime             time.Time         `validate:"gte"`
		EqFieldString       string            `validate:"eqfield=MaxString"`
		EqCSFieldString     string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString     string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString     string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString    string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString     string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString    string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString       string            `validate:"nefield=EqFieldString"`
		GtFieldString       string            `validate:"gtfield=MaxString"`
		GteFieldString      string            `validate:"gtefield=MaxString"`
		LtFieldString       string            `validate:"ltfield=MaxString"`
		LteFieldString      string            `validate:"ltefield=MaxString"`
		AlphaString         string            `validate:"alpha"`
		AlphanumString      string            `validate:"alphanum"`
		NumericString       string            `validate:"numeric"`
		NumberString        string            `validate:"number"`
		HexadecimalString   string            `validate:"hexadecimal"`
		HexColorString      string            `validate:"hexcolor"`
		RGBColorString      string            `validate:"rgb"`
		RGBAColorString     string            `validate:"rgba"`
		HSLColorString      string            `validate:"hsl"`
		HSLAColorString     string            `validate:"hsla"`
		Email               string            `validate:"email"`
		URL                 string            `validate:"url"`
		URI                 string            `validate:"uri"`
		Base64              string            `validate:"base64"`
		Contains            string            `validate:"contains=purpose"`
		ContainsAny         string            `validate:"containsany=!@#$"`
		Excludes            string            `validate:"excludes=text"`
		ExcludesAll         string            `validate:"excludesall=!@#$"`
		ExcludesRune        string            `validate:"excludesrune=☻"`
		ISBN                string            `validate:"isbn"`
		ISBN10              string            `validate:"isbn10"`
		ISBN13              string            `validate:"isbn13"`
		UUID                string            `validate:"uuid"`
		UUID3               string            `validate:"uuid3"`
		UUID4               string            `validate:"uuid4"`
		UUID5               string            `validate:"uuid5"`
		ULID                string            `validate:"ulid"`
		ASCII               string            `validate:"ascii"`
		PrintableASCII      string            `validate:"printascii"`
		MultiByte           string            `validate:"multibyte"`
		DataURI             string            `validate:"datauri"`
		Latitude            string            `validate:"latitude"`
		Longitude           string            `validate:"longitude"`
		SSN                 string            `validate:"ssn"`
		IP                  string            `validate:"ip"`
		IPv4                string            `validate:"ipv4"`
		IPv6                string            `validate:"ipv6"`
		CIDR                string            `validate:"cidr"`
		CIDRv4              string            `validate:"cidrv4"`
		CIDRv6              string            `validate:"cidrv6"`
		TCPAddr             string            `validate:"tcp_addr"`
		TCPAddrv4           string            `validate:"tcp4_addr"`
		TCPAddrv6           string            `validate:"tcp6_addr"`
		UDPAddr             string            `validate:"udp_addr"`
		UDPAddrv4           string            `validate:"udp4_addr"`
		UDPAddrv6           string            `validate:"udp6_addr"`
		IPAddr              string            `validate:"ip_addr"`
		IPAddrv4            string            `validate:"ip4_addr"`
		IPAddrv6            string            `validate:"ip6_addr"`
		UinxAddr            string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC                 string            `validate:"mac"`
		IsColor             string            `validate:"iscolor"`
		StrPtrMinLen        *string           `validate:"min=10"`
		StrPtrMaxLen        *string           `validate:"max=1"`
		StrPtrLen           *string           `validate:"len=2"`
		StrPtrLt            *string           `validate:"lt=1"`
		StrPtrLte           *string           `validate:"lte=1"`
		StrPtrGt            *string           `validate:"gt=10"`
		StrPtrGte           *string           `validate:"gte=10"`
		OneOfString         string            `validate:"oneof=red green"`
		OneOfInt            int               `validate:"oneof=5 63"`
		UniqueSlice         []string          `validate:"unique"`
		UniqueArray         [3]string         `validate:"unique"`
		UniqueMap           map[string]string `validate:"unique"`
		BooleanString       string            `validate:"boolean"`
		JSONString          string            `validate:"json"`
		JWTString           string            `validate:"jwt"`
		LowercaseString     string            `validate:"lowercase"`
		UppercaseString     string            `validate:"uppercase"`
		StartsWithString    string            `validate:"startswith=foo"`
		StartsNotWithString string            `validate:"startsnotwith=foo"`
		EndsWithString      string            `validate:"endswith=foo"`
		EndsNotWithString   string            `validate:"endsnotwith=foo"`
		Datetime            string            `validate:"datetime=2006-01-02"`
		PostCode            string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry     string
		PostCodeByField     string `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
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

	test.LowercaseString = "ABCDEFG"
	test.UppercaseString = "abcdefg"

	test.StartsWithString = "hello"
	test.StartsNotWithString = "foo-hello"
	test.EndsWithString = "hello"
	test.EndsNotWithString = "hello-foo"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}
	test.Datetime = "2008-Feb-01"

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
			expected: "IsColor deve essere un colore valido",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC deve contenere un indirizzo MAC valido",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr deve essere un indirizzo IP risolvibile",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 deve essere un indirizzo IPv4 risolvibile",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 deve essere un indirizzo IPv6 risolvibile",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr deve essere un indirizzo UDP valido",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 deve essere un indirizzo IPv4 UDP valido",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 deve essere un indirizzo IPv6 UDP valido",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr deve essere un indirizzo TCP valido",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 deve essere un indirizzo IPv4 TCP valido",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 deve essere un indirizzo IPv6 TCP valido",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR deve contenere una notazione CIDR valida",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 deve contenere una notazione CIDR per un indirizzo IPv4 valida",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 deve contenere una notazione CIDR per un indirizzo IPv6 valida",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN deve essere un numero SSN valido",
		},
		{
			ns:       "Test.IP",
			expected: "IP deve essere un indirizzo IP valido",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 deve essere un indirizzo IPv4 valido",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 deve essere un indirizzo IPv6 valido",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI deve contenere un Data URI valido",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude deve contenere una latitudine valida",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude deve contenere una longitudine valida",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte deve contenere caratteri multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII deve contenere solo caratteri ascii",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII deve contenere solo caratteri ascii stampabili",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID deve essere un UUID valido",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 deve essere un UUID versione 3 valido",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 deve essere un UUID versione 4 valido",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 deve essere un UUID versione 5 valido",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID deve essere un ULID valido",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN deve essere un numero ISBN valido",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 deve essere un numero ISBN-10 valido",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 deve essere un numero ISBN-13 valido",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes non deve contenere il testo 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll non deve contenere alcuno dei seguenti caratteri '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune non deve contenere '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny deve contenere almeno uno dei seguenti caratteri '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains deve contenere il testo 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 deve essere una stringa Base64 valida",
		},
		{
			ns:       "Test.Email",
			expected: "Email deve essere un indirizzo email valido",
		},
		{
			ns:       "Test.URL",
			expected: "URL deve essere un URL valido",
		},
		{
			ns:       "Test.URI",
			expected: "URI deve essere un URI valido",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString deve essere un colore RGB valido",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString deve essere un colore RGBA valido",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString deve essere un colore HSL valido",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString deve essere un colore HSLA valido",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString deve essere un esadecimale valido",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString deve essere un colore HEX valido",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString deve essere un numero valido",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString deve essere un valore numerico valido",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString può contenere solo caratteri alfanumerici",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString può contenere solo caratteri alfabetici",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString deve essere minore di MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString deve essere minore o uguale a MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString deve essere maggiore di MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString deve essere maggiore o uguale a MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString deve essere diverso da EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString deve essere minore di Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString deve essere minore o uguale a Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString deve essere maggiore di Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString deve essere maggiore o uguale a Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString deve essere diverso da Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString deve essere uguale a Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString deve essere uguale a MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString deve essere lungo almeno 3 caratteri",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber deve essere maggiore o uguale a 5,56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple deve contenere almeno 2 elementi",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime deve essere uguale o successivo alla Data/Ora corrente",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString deve essere lungo più di 3 caratteri",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber deve essere maggiore di 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple deve contenere più di 2 elementi",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime deve essere successivo alla Data/Ora corrente",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString deve essere lungo al massimo 3 caratteri",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber deve essere minore o uguale a 5,56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple deve contenere al massimo 2 elementi",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime deve essere uguale o precedente alla Data/Ora corrente",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString deve essere lungo meno di 3 caratteri",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber deve essere minore di 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple deve contenere meno di 2 elementi",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime deve essere precedente alla Data/Ora corrente",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString deve essere diverso da ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber deve essere diverso da 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple deve essere diverso da 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString non è uguale a 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber non è uguale a 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple non è uguale a 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString deve essere lungo al massimo 3 caratteri",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber deve essere minore o uguale a 1.113,00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple deve contenere al massimo 7 elementi",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString deve essere lungo almeno 1 carattere",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber deve essere maggiore o uguale a 1.113,00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple deve contenere almeno 7 elementi",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString deve essere lungo 1 carattere",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber deve essere uguale a 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple deve contenere 7 elementi",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString è un campo obbligatorio",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber è un campo obbligatorio",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple è un campo obbligatorio",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen deve essere lungo almeno 10 caratteri",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen deve essere lungo al massimo 1 carattere",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen deve essere lungo 2 caratteri",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt deve essere lungo meno di 1 carattere",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte deve essere lungo al massimo 1 carattere",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt deve essere lungo più di 10 caratteri",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte deve essere lungo almeno 10 caratteri",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString deve essere uno di [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt deve essere uno di [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice deve contenere valori unici",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray deve contenere valori unici",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap deve contenere valori unici",
		},
		{
			ns:       "Test.BooleanString",
			expected: "BooleanString deve rappresentare un valore booleano",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString deve essere una stringa json valida",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString deve essere una stringa jwt valida",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString deve essere una stringa minuscola",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString deve essere una stringa maiuscola",
		},
		{
			ns:       "Test.StartsWithString",
			expected: "StartsWithString deve iniziare con foo",
		},
		{
			ns:       "Test.StartsNotWithString",
			expected: "StartsNotWithString non deve iniziare con foo",
		},
		{
			ns:       "Test.EndsWithString",
			expected: "EndsWithString deve terminare con foo",
		},
		{
			ns:       "Test.EndsNotWithString",
			expected: "EndsNotWithString non deve terminare con foo",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime non corrisponde al formato 2006-01-02",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode non corrisponde al formato del codice postale dello stato SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField non corrisponde al formato del codice postale dello stato nel campo PostCodeCountry",
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
