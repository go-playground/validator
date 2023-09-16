package uk

import (
	"log"
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	ukrainian "github.com/go-playground/locales/uk"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	uk := ukrainian.New()
	uni := ut.New(uk, uk)
	trans, _ := uni.GetTranslator("uk")

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
		Inner                   Inner
		RequiredString          string            `validate:"required"`
		RequiredNumber          int               `validate:"required"`
		RequiredMultiple        []string          `validate:"required"`
		LenString               string            `validate:"len=1"`
		LenNumber               float64           `validate:"len=1113.00"`
		LenMultiple             []string          `validate:"len=7"`
		LenMultipleSecond       []string          `validate:"len=2"`
		MinString               string            `validate:"min=1"`
		MinStringMultiple       string            `validate:"min=2"`
		MinStringMultipleSecond string            `validate:"min=7"`
		MinNumber               float64           `validate:"min=1113.00"`
		MinMultiple             []string          `validate:"min=7"`
		MinMultipleSecond       []string          `validate:"min=2"`
		MaxString               string            `validate:"max=3"`
		MaxStringSecond         string            `validate:"max=7"`
		MaxNumber               float64           `validate:"max=1113.00"`
		MaxMultiple             []string          `validate:"max=7"`
		MaxMultipleSecond       []string          `validate:"max=2"`
		EqString                string            `validate:"eq=3"`
		EqNumber                float64           `validate:"eq=2.33"`
		EqMultiple              []string          `validate:"eq=7"`
		NeString                string            `validate:"ne="`
		NeNumber                float64           `validate:"ne=0.00"`
		NeMultiple              []string          `validate:"ne=0"`
		LtString                string            `validate:"lt=3"`
		LtStringSecond          string            `validate:"lt=7"`
		LtNumber                float64           `validate:"lt=5.56"`
		LtMultiple              []string          `validate:"lt=2"`
		LtMultipleSecond        []string          `validate:"lt=7"`
		LtTime                  time.Time         `validate:"lt"`
		LteString               string            `validate:"lte=3"`
		LteStringSecond         string            `validate:"lte=7"`
		LteNumber               float64           `validate:"lte=5.56"`
		LteMultiple             []string          `validate:"lte=2"`
		LteMultipleSecond       []string          `validate:"lte=7"`
		LteTime                 time.Time         `validate:"lte"`
		GtString                string            `validate:"gt=3"`
		GtStringSecond          string            `validate:"gt=7"`
		GtNumber                float64           `validate:"gt=5.56"`
		GtMultiple              []string          `validate:"gt=2"`
		GtMultipleSecond        []string          `validate:"gt=7"`
		GtTime                  time.Time         `validate:"gt"`
		GteString               string            `validate:"gte=3"`
		GteStringSecond         string            `validate:"gte=7"`
		GteNumber               float64           `validate:"gte=5.56"`
		GteMultiple             []string          `validate:"gte=2"`
		GteMultipleSecond       []string          `validate:"gte=7"`
		GteTime                 time.Time         `validate:"gte"`
		EqFieldString           string            `validate:"eqfield=MaxString"`
		EqCSFieldString         string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString         string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString         string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString        string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString         string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString        string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString           string            `validate:"nefield=EqFieldString"`
		GtFieldString           string            `validate:"gtfield=MaxString"`
		GteFieldString          string            `validate:"gtefield=MaxString"`
		LtFieldString           string            `validate:"ltfield=MaxString"`
		LteFieldString          string            `validate:"ltefield=MaxString"`
		AlphaString             string            `validate:"alpha"`
		AlphanumString          string            `validate:"alphanum"`
		NumericString           string            `validate:"numeric"`
		NumberString            string            `validate:"number"`
		HexadecimalString       string            `validate:"hexadecimal"`
		HexColorString          string            `validate:"hexcolor"`
		RGBColorString          string            `validate:"rgb"`
		RGBAColorString         string            `validate:"rgba"`
		HSLColorString          string            `validate:"hsl"`
		HSLAColorString         string            `validate:"hsla"`
		Email                   string            `validate:"email"`
		URL                     string            `validate:"url"`
		URI                     string            `validate:"uri"`
		Base64                  string            `validate:"base64"`
		Contains                string            `validate:"contains=purpose"`
		ContainsAny             string            `validate:"containsany=!@#$"`
		Excludes                string            `validate:"excludes=text"`
		ExcludesAll             string            `validate:"excludesall=!@#$"`
		ExcludesRune            string            `validate:"excludesrune=☻"`
		ISBN                    string            `validate:"isbn"`
		ISBN10                  string            `validate:"isbn10"`
		ISBN13                  string            `validate:"isbn13"`
		UUID                    string            `validate:"uuid"`
		UUID3                   string            `validate:"uuid3"`
		UUID4                   string            `validate:"uuid4"`
		UUID5                   string            `validate:"uuid5"`
		ULID                    string            `validate:"ulid"`
		ASCII                   string            `validate:"ascii"`
		PrintableASCII          string            `validate:"printascii"`
		MultiByte               string            `validate:"multibyte"`
		DataURI                 string            `validate:"datauri"`
		Latitude                string            `validate:"latitude"`
		Longitude               string            `validate:"longitude"`
		SSN                     string            `validate:"ssn"`
		IP                      string            `validate:"ip"`
		IPv4                    string            `validate:"ipv4"`
		IPv6                    string            `validate:"ipv6"`
		CIDR                    string            `validate:"cidr"`
		CIDRv4                  string            `validate:"cidrv4"`
		CIDRv6                  string            `validate:"cidrv6"`
		TCPAddr                 string            `validate:"tcp_addr"`
		TCPAddrv4               string            `validate:"tcp4_addr"`
		TCPAddrv6               string            `validate:"tcp6_addr"`
		UDPAddr                 string            `validate:"udp_addr"`
		UDPAddrv4               string            `validate:"udp4_addr"`
		UDPAddrv6               string            `validate:"udp6_addr"`
		IPAddr                  string            `validate:"ip_addr"`
		IPAddrv4                string            `validate:"ip4_addr"`
		IPAddrv6                string            `validate:"ip6_addr"`
		UinxAddr                string            `validate:"unix_addr"`
		MAC                     string            `validate:"mac"`
		IsColor                 string            `validate:"iscolor"`
		StrPtrMinLen            *string           `validate:"min=10"`
		StrPtrMinLenSecond      *string           `validate:"min=2"`
		StrPtrMaxLen            *string           `validate:"max=1"`
		StrPtrLen               *string           `validate:"len=2"`
		StrPtrLenSecond         *string           `validate:"len=7"`
		StrPtrLt                *string           `validate:"lt=1"`
		StrPtrLte               *string           `validate:"lte=1"`
		StrPtrLteMultiple       *string           `validate:"lte=2"`
		StrPtrLteMultipleSecond *string           `validate:"lte=7"`
		StrPtrGt                *string           `validate:"gt=10"`
		StrPtrGte               *string           `validate:"gte=10"`
		StrPtrGtSecond          *string           `validate:"gt=2"`
		StrPtrGteSecond         *string           `validate:"gte=2"`
		OneOfString             string            `validate:"oneof=red green"`
		OneOfInt                int               `validate:"oneof=5 63"`
		UniqueSlice             []string          `validate:"unique"`
		UniqueArray             [3]string         `validate:"unique"`
		UniqueMap               map[string]string `validate:"unique"`
		Image                   string            `validate:"image"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"

	test.MaxString = "1234"
	test.MaxStringSecond = "12345678"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)
	test.MaxMultipleSecond = make([]string, 3)

	test.LtString = "1234"
	test.LtStringSecond = "12345678"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtMultipleSecond = make([]string, 8)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteStringSecond = "12345678"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteMultipleSecond = make([]string, 8)
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

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}

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
			expected: "IsColor має бути кольором",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC має містити MAC адресу",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr має бути IP адресою, що розпізнається",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 має бути IPv4 адресою, що розпізнається",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 має бути IPv6 адресою, що розпізнається",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr має бути UDP адресою",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 має бути IPv4 UDP адресою",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 має бути IPv6 UDP адресою",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr має бути TCP адресою",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 має бути IPv4 TCP адресою",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 має бути IPv6 TCP адресою",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR має містити CIDR позначення",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 має містити CIDR позначення для IPv4 адреси",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 має містити CIDR позначення для IPv6 адреси",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN має бути SSN номером",
		},
		{
			ns:       "Test.IP",
			expected: "IP має бути IP адресою",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 має бути IPv4 адресою",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 має бути IPv6 адресою",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI має містити Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude має містити координати широти",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude має містити координати довготи",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte має містити мультибайтні символи",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII має містити тільки ascii символи",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII має містити лише доступні для друку ascii символи",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID має бути UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 має бути UUID 3 версії",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 має бути UUID 4 версії",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 має бути UUID 5 версії",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID має бути ULID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN має бути ISBN номером",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 має бути ISBN-10 номером",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 має бути ISBN-13 номером",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes не має містити текст 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll не має містити символи '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune не має містити '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny має містити мінімум один із символів '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains має містити текст 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 має бути Base64 рядком",
		},
		{
			ns:       "Test.Email",
			expected: "Email має бути email адресою",
		},
		{
			ns:       "Test.URL",
			expected: "URL має бути URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI має бути URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString має бути RGB кольором",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString має бути RGBA кольором",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString має бути HSL кольором",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString має бути HSLA кольором",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString має бути шістнадцятковим рядком",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString має бути HEX кольором",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString має бути цифрою",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString має бути цифровим значенням",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString має містити лише літери та цифри",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString має містити лише літери",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString має бути менше MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString має бути меншим або дорівнювати MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString має бути більше MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString має бути більше або дорівнювати MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString не має дорівнювати EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString має бути менше Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString має бути меншим або дорівнювати Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString має бути більше Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString має бути більше або дорівнювати Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString не має дорівнювати Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString має дорівнювати Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString має дорівнювати MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString має містити мінімум 3 символи",
		},
		{
			ns:       "Test.GteStringSecond",
			expected: "GteStringSecond має містити мінімум 7 символів",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber має бути більшим або рівним 5,56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple має містити мінімум 2 елементи",
		},
		{
			ns:       "Test.GteMultipleSecond",
			expected: "GteMultipleSecond має містити мінімум 7 елементів",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime має бути пізніше або дорівнювати поточному моменту",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString має бути довшим за 3 символи",
		},
		{
			ns:       "Test.GtStringSecond",
			expected: "GtStringSecond має бути довшим за 7 символів",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber має бути більше 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple має містити більше 2 елементів",
		},
		{
			ns:       "Test.GtMultipleSecond",
			expected: "GtMultipleSecond має містити більше 7 елементів",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime має бути пізніше поточного моменту",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString має містити максимум 3 символи",
		},
		{
			ns:       "Test.LteStringSecond",
			expected: "LteStringSecond має містити максимум 7 символів",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber має бути меншим або рівним 5,56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple має містити максимум 2 елементи",
		},
		{
			ns:       "Test.LteMultipleSecond",
			expected: "LteMultipleSecond має містити максимум 7 елементів",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime має бути менше або дорівнювати поточній даті та часу",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString має мати менше 3 символів",
		},
		{
			ns:       "Test.LtStringSecond",
			expected: "LtStringSecond має мати менше 7 символів",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber має бути менше 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple має містити не менше 2 елементів",
		},
		{
			ns:       "Test.LtMultipleSecond",
			expected: "LtMultipleSecond має містити не менше 7 елементів",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime має бути менше поточної дати й часу",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString має бути не рівний ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber має бути не рівний 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple має бути не рівний 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString не рівний 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber не рівний 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple не рівний 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString має містити максимум 3 символи",
		},
		{
			ns:       "Test.MaxStringSecond",
			expected: "MaxStringSecond має містити максимум 7 символів",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber має бути меншим або рівним 1 113,00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple має містити максимум 7 елементів",
		},
		{
			ns:       "Test.MaxMultipleSecond",
			expected: "MaxMultipleSecond має містити максимум 2 елементи",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString має містити мінімум 1 символ",
		},
		{
			ns:       "Test.MinStringMultiple",
			expected: "MinStringMultiple має містити мінімум 2 символи",
		},
		{
			ns:       "Test.MinStringMultipleSecond",
			expected: "MinStringMultipleSecond має містити мінімум 7 символів",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber має бути більшим або рівним 1 113,00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple має містити мінімум 7 елементів",
		},
		{
			ns:       "Test.MinMultipleSecond",
			expected: "MinMultipleSecond має містити мінімум 2 елементи",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString має бути довжиною в 1 символ",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber має дорівнювати 1 113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple має містити 7 елементів",
		},
		{
			ns:       "Test.LenMultipleSecond",
			expected: "LenMultipleSecond має містити 2 елементи",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString обов'язкове поле",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber обов'язкове поле",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple обов'язкове поле",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen має містити мінімум 10 символів",
		},
		{
			ns:       "Test.StrPtrMinLenSecond",
			expected: "StrPtrMinLenSecond має містити мінімум 2 символи",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen має містити максимум 1 символ",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen має бути довжиною в 2 символи",
		},
		{
			ns:       "Test.StrPtrLenSecond",
			expected: "StrPtrLenSecond має бути довжиною в 7 символів",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt має мати менше 1 символ",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte має містити максимум 1 символ",
		},
		{
			ns:       "Test.StrPtrLteMultiple",
			expected: "StrPtrLteMultiple має містити максимум 2 символи",
		},
		{
			ns:       "Test.StrPtrLteMultipleSecond",
			expected: "StrPtrLteMultipleSecond має містити максимум 7 символів",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt має бути довшим за 10 символів",
		},
		{
			ns:       "Test.StrPtrGtSecond",
			expected: "StrPtrGtSecond має бути довшим за 2 символи",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte має містити мінімум 10 символів",
		},
		{
			ns:       "Test.StrPtrGteSecond",
			expected: "StrPtrGteSecond має містити мінімум 2 символи",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString має бути одним з [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt має бути одним з [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice має містити унікальні значення",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray має містити унікальні значення",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap має містити унікальні значення",
		},
		{
			ns:       "Test.Image",
			expected: "Image має бути допустимим зображенням",
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

		log.Println(fe)

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}

}
