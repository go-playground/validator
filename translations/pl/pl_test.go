package pl

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	polish "github.com/go-playground/locales/pl"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {
	pl := polish.New()
	uni := ut.New(pl, pl)
	trans, ok := uni.GetTranslator("pl")
	if !ok {
		t.Fatalf("didn't found translator")
	}

	validate := validator.New()

	err := RegisterDefaultTranslations(validate, trans)
	Equal(t, err, nil)

	type Inner struct {
		EqCSFieldString    string
		NeCSFieldString    string
		GtCSFieldString    string
		GteCSFieldString   string
		LtCSFieldString    string
		LteCSFieldString   string
		RequiredIf         string
		RequiredUnless     string
		RequiredWith       string
		RequiredWithAll    string
		RequiredWithout    string
		RequiredWithoutAll string
		ExcludedIf         string
		ExcludedUnless     string
		ExcludedWith       string
		ExcludedWithAll    string
		ExcludedWithout    string
		ExcludedWithoutAll string
	}

	type Test struct {
		Inner                     Inner
		RequiredString            string            `validate:"required"`
		RequiredNumber            int               `validate:"required"`
		RequiredMultiple          []string          `validate:"required"`
		RequiredIf                string            `validate:"required_if=Inner.RequiredIf abcd"`
		RequiredUnless            string            `validate:"required_unless=Inner.RequiredUnless abcd"`
		RequiredWith              string            `validate:"required_with=Inner.RequiredWith"`
		RequiredWithAll           string            `validate:"required_with_all=Inner.RequiredWith Inner.RequiredWithAll"`
		RequiredWithout           string            `validate:"required_without=Inner.RequiredWithout"`
		RequiredWithoutAll        string            `validate:"required_without_all=Inner.RequiredWithout Inner.RequiredWithoutAll"`
		ExcludedIf                string            `validate:"excluded_if=Inner.ExcludedIf abcd"`
		ExcludedUnless            string            `validate:"excluded_unless=Inner.ExcludedUnless abcd"`
		ExcludedWith              string            `validate:"excluded_with=Inner.ExcludedWith"`
		ExcludedWithout           string            `validate:"excluded_with_all=Inner.ExcludedWithAll"`
		ExcludedWithAll           string            `validate:"excluded_without=Inner.ExcludedWithout"`
		ExcludedWithoutAll        string            `validate:"excluded_without_all=Inner.ExcludedWithoutAll"`
		IsDefault                 string            `validate:"isdefault"`
		LenString                 string            `validate:"len=1"`
		LenNumber                 float64           `validate:"len=1113.00"`
		LenMultiple               []string          `validate:"len=7"`
		MinString                 string            `validate:"min=1"`
		MinNumber                 float64           `validate:"min=1113.00"`
		MinMultiple               []string          `validate:"min=7"`
		MaxString                 string            `validate:"max=3"`
		MaxNumber                 float64           `validate:"max=1113.00"`
		MaxMultiple               []string          `validate:"max=7"`
		EqString                  string            `validate:"eq=3"`
		EqNumber                  float64           `validate:"eq=2.33"`
		EqMultiple                []string          `validate:"eq=7"`
		NeString                  string            `validate:"ne="`
		NeNumber                  float64           `validate:"ne=0.00"`
		NeMultiple                []string          `validate:"ne=0"`
		LtString                  string            `validate:"lt=3"`
		LtNumber                  float64           `validate:"lt=5.56"`
		LtMultiple                []string          `validate:"lt=2"`
		LtTime                    time.Time         `validate:"lt"`
		LteString                 string            `validate:"lte=3"`
		LteNumber                 float64           `validate:"lte=5.56"`
		LteMultiple               []string          `validate:"lte=2"`
		LteTime                   time.Time         `validate:"lte"`
		GtString                  string            `validate:"gt=3"`
		GtNumber                  float64           `validate:"gt=5.56"`
		GtMultiple                []string          `validate:"gt=2"`
		GtTime                    time.Time         `validate:"gt"`
		GteString                 string            `validate:"gte=3"`
		GteStringSingle           string            `validate:"gte=1"`
		GteStringFive             string            `validate:"gte=5"`
		GteStringTwentyOne        string            `validate:"gte=21"`
		GteStringBigNumber        string            `validate:"gte=852"`
		GteStringAnotherBigNumber string            `validate:"gte=2137"`
		GteNumber                 float64           `validate:"gte=5.56"`
		GteMultiple               []string          `validate:"gte=2"`
		GteTime                   time.Time         `validate:"gte"`
		EqFieldString             string            `validate:"eqfield=MaxString"`
		EqCSFieldString           string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString           string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString           string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString          string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString           string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString          string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString             string            `validate:"nefield=EqFieldString"`
		GtFieldString             string            `validate:"gtfield=MaxString"`
		GteFieldString            string            `validate:"gtefield=MaxString"`
		LtFieldString             string            `validate:"ltfield=MaxString"`
		LteFieldString            string            `validate:"ltefield=MaxString"`
		AlphaString               string            `validate:"alpha"`
		AlphanumString            string            `validate:"alphanum"`
		NumericString             string            `validate:"numeric"`
		NumberString              string            `validate:"number"`
		HexadecimalString         string            `validate:"hexadecimal"`
		HexColorString            string            `validate:"hexcolor"`
		RGBColorString            string            `validate:"rgb"`
		RGBAColorString           string            `validate:"rgba"`
		HSLColorString            string            `validate:"hsl"`
		HSLAColorString           string            `validate:"hsla"`
		Email                     string            `validate:"email"`
		URL                       string            `validate:"url"`
		URI                       string            `validate:"uri"`
		Base64                    string            `validate:"base64"`
		Contains                  string            `validate:"contains=purpose"`
		ContainsAny               string            `validate:"containsany=!@#$"`
		Excludes                  string            `validate:"excludes=text"`
		ExcludesAll               string            `validate:"excludesall=!@#$"`
		ExcludesRune              string            `validate:"excludesrune=☻"`
		ISBN                      string            `validate:"isbn"`
		ISBN10                    string            `validate:"isbn10"`
		ISBN13                    string            `validate:"isbn13"`
		ISSN                      string            `validate:"issn"`
		UUID                      string            `validate:"uuid"`
		UUID3                     string            `validate:"uuid3"`
		UUID4                     string            `validate:"uuid4"`
		UUID5                     string            `validate:"uuid5"`
		ULID                      string            `validate:"ulid"`
		ASCII                     string            `validate:"ascii"`
		PrintableASCII            string            `validate:"printascii"`
		MultiByte                 string            `validate:"multibyte"`
		DataURI                   string            `validate:"datauri"`
		Latitude                  string            `validate:"latitude"`
		Longitude                 string            `validate:"longitude"`
		SSN                       string            `validate:"ssn"`
		IP                        string            `validate:"ip"`
		IPv4                      string            `validate:"ipv4"`
		IPv6                      string            `validate:"ipv6"`
		CIDR                      string            `validate:"cidr"`
		CIDRv4                    string            `validate:"cidrv4"`
		CIDRv6                    string            `validate:"cidrv6"`
		TCPAddr                   string            `validate:"tcp_addr"`
		TCPAddrv4                 string            `validate:"tcp4_addr"`
		TCPAddrv6                 string            `validate:"tcp6_addr"`
		UDPAddr                   string            `validate:"udp_addr"`
		UDPAddrv4                 string            `validate:"udp4_addr"`
		UDPAddrv6                 string            `validate:"udp6_addr"`
		IPAddr                    string            `validate:"ip_addr"`
		IPAddrv4                  string            `validate:"ip4_addr"`
		IPAddrv6                  string            `validate:"ip6_addr"`
		UinxAddr                  string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC                       string            `validate:"mac"`
		FQDN                      string            `validate:"fqdn"`
		IsColor                   string            `validate:"iscolor"`
		StrPtrMinLen              *string           `validate:"min=10"`
		StrPtrMaxLen              *string           `validate:"max=1"`
		StrPtrLen                 *string           `validate:"len=2"`
		StrPtrLt                  *string           `validate:"lt=1"`
		StrPtrLte                 *string           `validate:"lte=1"`
		StrPtrGt                  *string           `validate:"gt=10"`
		StrPtrGte                 *string           `validate:"gte=10"`
		OneOfString               string            `validate:"oneof=red green"`
		OneOfInt                  int               `validate:"oneof=5 63"`
		UniqueSlice               []string          `validate:"unique"`
		UniqueArray               [3]string         `validate:"unique"`
		UniqueMap                 map[string]string `validate:"unique"`
		JSONString                string            `validate:"json"`
		JWTString                 string            `validate:"jwt"`
		LowercaseString           string            `validate:"lowercase"`
		UppercaseString           string            `validate:"uppercase"`
		Datetime                  string            `validate:"datetime=2006-01-02"`
		PostCode                  string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry           string
		PostCodeByField           string `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
		BooleanString             string `validate:"boolean"`
		Image                     string `validate:"image"`
		CveString                 string `validate:"cve"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"
	test.Inner.RequiredUnless = "1234"
	test.Inner.RequiredWith = "1234"
	test.Inner.RequiredWithAll = "1234"
	test.Inner.ExcludedIf = "abcd"
	test.Inner.ExcludedUnless = "1234"
	test.Inner.ExcludedWith = "1234"
	test.Inner.ExcludedWithAll = "1234"

	test.ExcludedIf = "1234"
	test.ExcludedUnless = "1234"
	test.ExcludedWith = "1234"
	test.ExcludedWithAll = "1234"
	test.ExcludedWithout = "1234"
	test.ExcludedWithoutAll = "1234"

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

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}
	test.Datetime = "2008-Feb-01"
	test.BooleanString = "A"
	test.CveString = "A"

	test.Inner.RequiredIf = "abcd"

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
			expected: "IsColor musi być prawdziwym kolorem",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC musi zawierać poprawny MAC adres",
		},
		{
			ns:       "Test.FQDN",
			expected: "FQDN musi być poprawnym FQDN",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr musi być rozpoznawalnym adresem IP",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 musi być rozpoznawalnym adresem IPv4",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 musi być rozpoznawalnym adresem IPv6",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr musi być poprawnym adresem UDP",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 musi być poprawnym adresem IPv4 UDP",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 musi być poprawnym adresem IPv6 UDP",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr musi być poprawnym adresem TCP",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 musi być poprawnym adresem IPv4 TCP",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 musi być poprawnym adresem IPv6 TCP",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR musi zawierać adres zapisany metodą CIDR",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 musi zawierać adres IPv4 zapisany metodą CIDR",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 musi zawierać adres IPv6 zapisany metodą CIDR",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN musi zawierać poprawny numer SSN",
		},
		{
			ns:       "Test.IP",
			expected: "IP musi zawierać poprawny adres IP",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 musi zawierać poprawny adres IPv4",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 musi zawierać poprawny adres IPv6",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI musi zawierać poprawnie zakodowane dane w formie URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude musi zawierać poprawną szerokość geograficzną",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude musi zawierać poprawną długość geograficzną",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte musi zawierać znaki wielobajtowe",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII może zawierać wyłącznie znaki ASCII",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII może zawierać wyłącznie drukowalne znaki ASCII",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID musi być poprawnym identyfikatorem UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 musi być poprawnym identyfikatorem UUID w wersji 3",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 musi być poprawnym identyfikatorem UUID w wersji 4",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 musi być poprawnym identyfikatorem UUID w wersji 5",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID musi być poprawnym identyfikatorem ULID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN musi być poprawnym numerem ISBN",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 musi być poprawnym numerem ISBN-10",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 musi być poprawnym numerem ISBN-13",
		},
		{
			ns:       "Test.ISSN",
			expected: "ISSN musi być poprawnym numerem ISSN",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes nie może zawierać tekstu 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll nie może zawierać żadnych z następujących znaków '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune nie może zawierać następujących znaków '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny musi zawierać przynajmniej jeden z następujących znaków '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains musi zawierać tekst 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 musi być ciągiem znaków zakodowanym w formacie Base64",
		},
		{
			ns:       "Test.Email",
			expected: "Email musi być poprawnym adresem email",
		},
		{
			ns:       "Test.URL",
			expected: "URL musi być poprawnym adresem URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI musi być poprawnym adresem URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString musi być poprawnym kolorem w formacie RGB",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString musi być poprawnym kolorem w formacie RGBA",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString musi być poprawnym kolorem w formacie HSL",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString musi być poprawnym kolorem w formacie HSLA",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString musi być poprawną wartością heksadecymalną",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString musi być poprawnym kolorem w formacie HEX",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString musi być poprawną liczbą",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString musi być poprawną wartością numeryczną",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString może zawierać wyłącznie znaki alfanumeryczne",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString może zawierać wyłącznie znaki alfabetu",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString musi być mniejsze niż MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString musi być mniejsze lub równe MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString musi być większe niż MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString musi być większe lub równe MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString nie może być równe EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString musi być mniejsze niż Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString musi być mniejsze lub równe Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString musi być większe niż Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString musi być większe lub równe niż Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString nie może być równe Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString musi być równe Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString musi być równe MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString musi mieć długość przynajmniej na 3 znaki",
		},
		{
			ns:       "Test.GteStringSingle",
			expected: "GteStringSingle musi mieć długość przynajmniej na 1 znak",
		},
		{
			ns:       "Test.GteStringFive",
			expected: "GteStringFive musi mieć długość przynajmniej na 5 znaków",
		},
		{
			ns:       "Test.GteStringTwentyOne",
			expected: "GteStringTwentyOne musi mieć długość przynajmniej na 21 znaków",
		},
		{
			ns:       "Test.GteStringBigNumber",
			expected: "GteStringBigNumber musi mieć długość przynajmniej na 852 znaki",
		},
		{
			ns:       "Test.GteStringAnotherBigNumber",
			expected: "GteStringAnotherBigNumber musi mieć długość przynajmniej na 2 137 znaków",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber musi być równe 5,56 lub większe",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple musi zawierać co najmniej 2 elementy",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime musi być większe lub równe niż obecny dzień i godzina",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString musi mieć długość większą niż 3 znaki",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber musi być większe niż 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple musi zawierać więcej niż 2 elementy",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime musi być większe niż obecny dzień i godzina",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString musi mieć długość maksymalnie na 3 znaki",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber musi być równe 5,56 lub mniej",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple musi zawierać maksymalnie 2 elementy",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime musi być mniejsze lub równe niż obecny dzień i godzina",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString musi mieć długość mniejszą niż 3 znaki",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber musi być mniejsze niż 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple musi zawierać mniej niż 2 elementy",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime musi być mniejsze niż obecny dzień i godzina",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString nie powinien być równy ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber nie powinien być równy 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple nie powinien być równy 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString nie równa się 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber nie równa się 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple nie równa się 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString musi mieć długość maksymalnie na 3 znaki",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber musi być równe 1 113,00 lub mniej",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple musi zawierać maksymalnie 7 elementów",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString musi mieć długość przynajmniej na 1 znak",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber musi być równe 1 113,00 lub więcej",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple musi zawierać przynajmniej 7 elementów",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString musi mieć długość na 1 znak",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber musi być równe 1 113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple musi zawierać 7 elementów",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredIf",
			expected: "RequiredIf jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredUnless",
			expected: "RequiredUnless jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredWith",
			expected: "RequiredWith jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredWithAll",
			expected: "RequiredWithAll jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredWithout",
			expected: "RequiredWithout jest wymaganym polem",
		},
		{
			ns:       "Test.RequiredWithoutAll",
			expected: "RequiredWithoutAll jest wymaganym polem",
		},
		{
			ns:       "Test.ExcludedIf",
			expected: "ExcludedIf jest wykluczonym polem",
		},
		{
			ns:       "Test.ExcludedUnless",
			expected: "ExcludedUnless jest wykluczonym polem",
		},
		{
			ns:       "Test.ExcludedWith",
			expected: "ExcludedWith jest wykluczonym polem",
		},
		{
			ns:       "Test.ExcludedWithAll",
			expected: "ExcludedWithAll jest wykluczonym polem",
		},
		{
			ns:       "Test.ExcludedWithout",
			expected: "ExcludedWithout jest wykluczonym polem",
		},
		{
			ns:       "Test.ExcludedWithoutAll",
			expected: "ExcludedWithoutAll jest wykluczonym polem",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen musi mieć długość przynajmniej na 10 znaków",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen musi mieć długość maksymalnie na 1 znak",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen musi mieć długość na 2 znaki",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt musi mieć długość mniejszą niż 1 znak",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte musi mieć długość maksymalnie na 1 znak",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt musi mieć długość większą niż 10 znaków",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte musi mieć długość przynajmniej na 10 znaków",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString musi być jednym z [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt musi być jednym z [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice musi zawierać unikalne wartości",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray musi zawierać unikalne wartości",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap musi zawierać unikalne wartości",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString musi być ciągiem znaków w formacie JSON",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString musi być ciągiem znaków w formacie JWT",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString musi zawierać wyłącznie małe litery",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString musi zawierać wyłącznie duże litery",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime nie spełnia formatu 2006-01-02",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode nie spełnia formatu kodu pocztowego kraju SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField nie spełnia formatu kodu pocztowego kraju z pola PostCodeCountry",
		},
		{
			ns:       "Test.BooleanString",
			expected: "BooleanString musi być wartością logiczną",
		},
		{
			ns:       "Test.Image",
			expected: "Image musi być obrazem",
		},
		{
			ns:       "Test.CveString",
			expected: "CveString musi być poprawnym identyfikatorem CVE",
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
