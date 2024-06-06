package uk

import (
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
	trans, ok := uni.GetTranslator("uk")
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
		Inner                   Inner
		RequiredString          string            `validate:"required"`
		RequiredNumber          int               `validate:"required"`
		RequiredMultiple        []string          `validate:"required"`
		RequiredIf              string            `validate:"required_if=Inner.RequiredIf abcd"`
		RequiredUnless          string            `validate:"required_unless=Inner.RequiredUnless abcd"`
		RequiredWith            string            `validate:"required_with=Inner.RequiredWith"`
		RequiredWithAll         string            `validate:"required_with_all=Inner.RequiredWith Inner.RequiredWithAll"`
		RequiredWithout         string            `validate:"required_without=Inner.RequiredWithout"`
		RequiredWithoutAll      string            `validate:"required_without_all=Inner.RequiredWithout Inner.RequiredWithoutAll"`
		ExcludedIf              string            `validate:"excluded_if=Inner.ExcludedIf abcd"`
		ExcludedUnless          string            `validate:"excluded_unless=Inner.ExcludedUnless abcd"`
		ExcludedWith            string            `validate:"excluded_with=Inner.ExcludedWith"`
		ExcludedWithout         string            `validate:"excluded_with_all=Inner.ExcludedWithAll"`
		ExcludedWithAll         string            `validate:"excluded_without=Inner.ExcludedWithout"`
		ExcludedWithoutAll      string            `validate:"excluded_without_all=Inner.ExcludedWithoutAll"`
		IsDefault               string            `validate:"isdefault"`
		LenString               string            `validate:"len=1"`
		LenNumber               float64           `validate:"len=1113.00"`
		LenMultiple             []string          `validate:"len=7"`
		LenMultipleSecond       []string          `validate:"len=2"`
		LenMultipleThird        []string          `validate:"len=1"`
		MinString               string            `validate:"min=1"`
		MinStringMultiple       string            `validate:"min=2"`
		MinStringMultipleSecond string            `validate:"min=7"`
		MinStringMultipleThird  string            `validate:"min=1"`
		MinNumber               float64           `validate:"min=1113.00"`
		MinMultiple             []string          `validate:"min=7"`
		MinMultipleSecond       []string          `validate:"min=2"`
		MinMultipleThird        []string          `validate:"min=1"`
		MaxString               string            `validate:"max=3"`
		MaxStringSecond         string            `validate:"max=7"`
		MaxStringThird          string            `validate:"max=1"`
		MaxNumber               float64           `validate:"max=1113.00"`
		MaxMultiple             []string          `validate:"max=7"`
		MaxMultipleSecond       []string          `validate:"max=2"`
		MaxMultipleThird        []string          `validate:"max=1"`
		EqString                string            `validate:"eq=3"`
		EqNumber                float64           `validate:"eq=2.33"`
		EqMultiple              []string          `validate:"eq=7"`
		NeString                string            `validate:"ne="`
		NeNumber                float64           `validate:"ne=0.00"`
		NeMultiple              []string          `validate:"ne=0"`
		LtString                string            `validate:"lt=3"`
		LtStringSecond          string            `validate:"lt=7"`
		LtStringThird           string            `validate:"lt=1"`
		LtNumber                float64           `validate:"lt=5.56"`
		LtMultiple              []string          `validate:"lt=2"`
		LtMultipleSecond        []string          `validate:"lt=7"`
		LtMultipleThird         []string          `validate:"lt=1"`
		LtTime                  time.Time         `validate:"lt"`
		LteString               string            `validate:"lte=3"`
		LteStringSecond         string            `validate:"lte=7"`
		LteStringThird          string            `validate:"lte=1"`
		LteNumber               float64           `validate:"lte=5.56"`
		LteMultiple             []string          `validate:"lte=2"`
		LteMultipleSecond       []string          `validate:"lte=7"`
		LteMultipleThird        []string          `validate:"lte=1"`
		LteTime                 time.Time         `validate:"lte"`
		GtString                string            `validate:"gt=3"`
		GtStringSecond          string            `validate:"gt=7"`
		GtStringThird           string            `validate:"gt=1"`
		GtNumber                float64           `validate:"gt=5.56"`
		GtMultiple              []string          `validate:"gt=2"`
		GtMultipleSecond        []string          `validate:"gt=7"`
		GtMultipleThird         []string          `validate:"gt=1"`
		GtTime                  time.Time         `validate:"gt"`
		GteString               string            `validate:"gte=3"`
		GteStringSecond         string            `validate:"gte=7"`
		GteStringThird          string            `validate:"gte=1"`
		GteNumber               float64           `validate:"gte=5.56"`
		GteMultiple             []string          `validate:"gte=2"`
		GteMultipleSecond       []string          `validate:"gte=7"`
		GteMultipleThird        []string          `validate:"gte=1"`
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
		ISSN                    string            `validate:"issn"`
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
		UinxAddr                string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC                     string            `validate:"mac"`
		FQDN                    string            `validate:"fqdn"`
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
		JSONString              string            `validate:"json"`
		JWTString               string            `validate:"jwt"`
		LowercaseString         string            `validate:"lowercase"`
		UppercaseString         string            `validate:"uppercase"`
		Datetime                string            `validate:"datetime=2006-01-02"`
		PostCode                string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry         string
		PostCodeByField         string `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
		BooleanString           string `validate:"boolean"`
		Image                   string `validate:"image"`
		CveString               string `validate:"cve"`
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
	test.MaxStringSecond = "12345678"
	test.MaxStringThird = "12"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)
	test.MaxMultipleSecond = make([]string, 3)
	test.MaxMultipleThird = make([]string, 2)

	test.LtString = "1234"
	test.LtStringSecond = "12345678"
	test.LtStringThird = "12"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtMultipleSecond = make([]string, 8)
	test.LtMultipleThird = make([]string, 2)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteStringSecond = "12345678"
	test.LteStringThird = "12"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteMultipleSecond = make([]string, 8)
	test.LteMultipleThird = make([]string, 2)
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
			expected: "IsColor має бути кольором",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC має містити MAC адресу",
		},
		{
			ns:       "Test.FQDN",
			expected: "FQDN має містити FQDN",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr має бути розпізнаваною IP адресою",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 має бути розпізнаваною IPv4 адресою",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 має бути розпізнаваною IPv6 адресою",
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
			expected: "MultiByte має містити мультібайтні символи",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII має містити тільки ascii символи",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII має містити тільки доступні для друку ascii символи",
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
			ns:       "Test.ISSN",
			expected: "ISSN має бути ISSN номером",
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
			expected: "ContainsAny має містити щонайменше один із символів '!@#$'",
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
			expected: "AlphanumString має містити тільки літери та цифри",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString має містити тільки літери",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString має бути менше MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString має бути менше чи дорівнювати MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString має бути більше MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString має бути більше чи дорівнювати MaxString",
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
			expected: "LteCSFieldString має бути менше чи дорівнювати Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString має бути більше Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString має бути більше чи дорівнювати Inner.GteCSFieldString",
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
			expected: "GteString має містити щонайменше 3 символи",
		},
		{
			ns:       "Test.GteStringSecond",
			expected: "GteStringSecond має містити щонайменше 7 символів",
		},
		{
			ns:       "Test.GteStringThird",
			expected: "GteStringThird має містити щонайменше 1 символ",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber має бути більше чи дорівнювати 5,56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple має містити щонайменше 2 елементи",
		},
		{
			ns:       "Test.GteMultipleSecond",
			expected: "GteMultipleSecond має містити щонайменше 7 елементів",
		},
		{
			ns:       "Test.GteMultipleThird",
			expected: "GteMultipleThird має містити щонайменше 1 елемент",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime має бути пізніше чи дорівнювати теперішньому моменту",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString має бути довше за 3 символи",
		},
		{
			ns:       "Test.GtStringSecond",
			expected: "GtStringSecond має бути довше за 7 символів",
		},
		{
			ns:       "Test.GtStringThird",
			expected: "GtStringThird має бути довше за 1 символ",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber має бути більше 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple має містити більше ніж 2 елементи",
		},
		{
			ns:       "Test.GtMultipleSecond",
			expected: "GtMultipleSecond має містити більше ніж 7 елементів",
		},
		{
			ns:       "Test.GtMultipleThird",
			expected: "GtMultipleThird має містити більше ніж 1 елемент",
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
			ns:       "Test.LteStringThird",
			expected: "LteStringThird має містити максимум 1 символ",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber має бути менше чи дорівнювати 5,56",
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
			ns:       "Test.LteMultipleThird",
			expected: "LteMultipleThird має містити максимум 1 елемент",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime має бути менше чи дорівнювати поточній даті та часу",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString має мати менше за 3 символи",
		},
		{
			ns:       "Test.LtStringSecond",
			expected: "LtStringSecond має мати менше за 7 символів",
		},
		{
			ns:       "Test.LtStringThird",
			expected: "LtStringThird має мати менше за 1 символ",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber має бути менше 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple має містити менше ніж 2 елементи",
		},
		{
			ns:       "Test.LtMultipleSecond",
			expected: "LtMultipleSecond має містити менше ніж 7 елементів",
		},
		{
			ns:       "Test.LtMultipleThird",
			expected: "LtMultipleThird має містити менше ніж 1 елемент",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime має бути менше поточної дати й часу",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString має не дорівнювати ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber має не дорівнювати 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple має не дорівнювати 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString не дорівнює 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber не дорівнює 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple не дорівнює 7",
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
			expected: "MaxNumber має бути менше чи дорівнювати 1 113,00",
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
			expected: "MinString має містити щонайменше 1 символ",
		},
		{
			ns:       "Test.MinStringMultiple",
			expected: "MinStringMultiple має містити щонайменше 2 символи",
		},
		{
			ns:       "Test.MinStringMultipleSecond",
			expected: "MinStringMultipleSecond має містити щонайменше 7 символів",
		},
		{
			ns:       "Test.MinStringMultipleThird",
			expected: "MinStringMultipleThird має містити щонайменше 1 символ",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber має бути більше чи дорівнювати 1 113,00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple має містити щонайменше 7 елементів",
		},
		{
			ns:       "Test.MinMultipleSecond",
			expected: "MinMultipleSecond має містити щонайменше 2 елементи",
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
			ns:       "Test.LenMultipleThird",
			expected: "LenMultipleThird має містити 1 елемент",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString обов'язкове поле",
		},
		{
			ns:       "Test.RequiredIf",
			expected: "RequiredIf обов'язкове поле",
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
			ns:       "Test.RequiredUnless",
			expected: "RequiredUnless обов'язкове поле",
		},
		{
			ns:       "Test.RequiredWith",
			expected: "RequiredWith обов'язкове поле",
		},
		{
			ns:       "Test.RequiredWithAll",
			expected: "RequiredWithAll обов'язкове поле",
		},
		{
			ns:       "Test.RequiredWithout",
			expected: "RequiredWithout обов'язкове поле",
		},
		{
			ns:       "Test.RequiredWithoutAll",
			expected: "RequiredWithoutAll обов'язкове поле",
		},
		{
			ns:       "Test.ExcludedIf",
			expected: "ExcludedIf є виключеним полем",
		},
		{
			ns:       "Test.ExcludedUnless",
			expected: "ExcludedUnless є виключеним полем",
		},
		{
			ns:       "Test.ExcludedWith",
			expected: "ExcludedWith є виключеним полем",
		},
		{
			ns:       "Test.ExcludedWithAll",
			expected: "ExcludedWithAll є виключеним полем",
		},
		{
			ns:       "Test.ExcludedWithout",
			expected: "ExcludedWithout є виключеним полем",
		},
		{
			ns:       "Test.ExcludedWithoutAll",
			expected: "ExcludedWithoutAll є виключеним полем",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen має містити щонайменше 10 символів",
		},
		{
			ns:       "Test.StrPtrMinLenSecond",
			expected: "StrPtrMinLenSecond має містити щонайменше 2 символи",
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
			expected: "StrPtrLt має мати менше за 1 символ",
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
			expected: "StrPtrGt має бути довше за 10 символів",
		},
		{
			ns:       "Test.StrPtrGtSecond",
			expected: "StrPtrGtSecond має бути довше за 2 символи",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte має містити щонайменше 10 символів",
		},
		{
			ns:       "Test.StrPtrGteSecond",
			expected: "StrPtrGteSecond має містити щонайменше 2 символи",
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
			ns:       "Test.JSONString",
			expected: "JSONString має бути json рядком",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString має бути jwt рядком",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString має бути рядком у нижньому регістрі",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString має бути рядком у верхньому регістрі",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime не відповідає 2006-01-02 формату",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode не відповідає формату поштового індексу країни SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField не відповідає формату поштового індексу країни в PostCodeCountry полі",
		},
		{
			ns:       "Test.BooleanString",
			expected: "BooleanString має бути булевим значенням",
		},
		{
			ns:       "Test.Image",
			expected: "Image має бути допустимим зображенням",
		},
		{
			ns:       "Test.CveString",
			expected: "CveString має бути cve ідентифікатором",
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
