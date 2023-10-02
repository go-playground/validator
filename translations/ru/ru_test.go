package ru

import (
	"log"
	//"github.com/rustery/validator"
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	russian "github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	ru := russian.New()
	uni := ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

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
		UinxAddr                string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
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
		Image			  string			`validate:"image"`
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
			expected: "IsColor должен быть цветом",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC должен содержать MAC адрес",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr должен быть распознаваемым IP адресом",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 должен быть распознаваемым IPv4 адресом",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 должен быть распознаваемым IPv6 адресом",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr должен быть UDP адресом",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 должен быть IPv4 UDP адресом",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 должен быть IPv6 UDP адресом",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr должен быть TCP адресом",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 должен быть IPv4 TCP адресом",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 должен быть IPv6 TCP адресом",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR должен содержать CIDR обозначения",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 должен содержать CIDR обозначения для IPv4 адреса",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 должен содержать CIDR обозначения для IPv6 адреса",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN должен быть SSN номером",
		},
		{
			ns:       "Test.IP",
			expected: "IP должен быть IP адресом",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 должен быть IPv4 адресом",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 должен быть IPv6 адресом",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI должен содержать Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude должен содержать координаты широты",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude должен содержать координаты долготы",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte должен содержать мультибайтные символы",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII должен содержать только ascii символы",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII должен содержать только доступные для печати ascii символы",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID должен быть UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 должен быть UUID 3 версии",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 должен быть UUID 4 версии",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 должен быть UUID 5 версии",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID должен быть ULID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN должен быть ISBN номером",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 должен быть ISBN-10 номером",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 должен быть ISBN-13 номером",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes не должен содержать текст 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll не должен содержать символы '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune не должен содержать '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny должен содержать минимум один из символов '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains должен содержать текст 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 должен быть Base64 строкой",
		},
		{
			ns:       "Test.Email",
			expected: "Email должен быть email адресом",
		},
		{
			ns:       "Test.URL",
			expected: "URL должен быть URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI должен быть URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString должен быть RGB цветом",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString должен быть RGBA цветом",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString должен быть HSL цветом",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString должен быть HSLA цветом",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString должен быть шестнадцатеричной строкой",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString должен быть HEX цветом",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString должен быть цифрой",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString должен быть цифровым значением",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString должен содержать только буквы и цифры",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString должен содержать только буквы",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString должен быть менее MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString должен быть менее или равен MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString должен быть больше MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString должен быть больше или равен MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString не должен быть равен EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString должен быть менее Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString должен быть менее или равен Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString должен быть больше Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString должен быть больше или равен Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString не должен быть равен Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString должен быть равен Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString должен быть равен MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString должен содержать минимум 3 символа",
		},
		{
			ns:       "Test.GteStringSecond",
			expected: "GteStringSecond должен содержать минимум 7 символов",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber должен быть больше или равно 5,56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple должен содержать минимум 2 элемента",
		},
		{
			ns:       "Test.GteMultipleSecond",
			expected: "GteMultipleSecond должен содержать минимум 7 элементов",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime должна быть позже или равна текущему моменту",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString должен быть длиннее 3 символов",
		},
		{
			ns:       "Test.GtStringSecond",
			expected: "GtStringSecond должен быть длиннее 7 символов",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber должен быть больше 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple должен содержать более 2 элементов",
		},
		{
			ns:       "Test.GtMultipleSecond",
			expected: "GtMultipleSecond должен содержать более 7 элементов",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime должна быть позже текущего момента",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString должен содержать максимум 3 символа",
		},
		{
			ns:       "Test.LteStringSecond",
			expected: "LteStringSecond должен содержать максимум 7 символов",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber должен быть менее или равен 5,56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple должен содержать максимум 2 элемента",
		},
		{
			ns:       "Test.LteMultipleSecond",
			expected: "LteMultipleSecond должен содержать максимум 7 элементов",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime must be less than or equal to the current Date & Time",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString должен иметь менее 3 символов",
		},
		{
			ns:       "Test.LtStringSecond",
			expected: "LtStringSecond должен иметь менее 7 символов",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber должен быть менее 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple должен содержать менее 2 элементов",
		},
		{
			ns:       "Test.LtMultipleSecond",
			expected: "LtMultipleSecond должен содержать менее 7 элементов",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime must be less than the current Date & Time",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString должен быть не равен ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber должен быть не равен 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple должен быть не равен 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString не равен 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber не равен 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple не равен 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString должен содержать максимум 3 символа",
		},
		{
			ns:       "Test.MaxStringSecond",
			expected: "MaxStringSecond должен содержать максимум 7 символов",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber должен быть меньше или равно 1 113,00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple должен содержать максимум 7 элементов",
		},
		{
			ns:       "Test.MaxMultipleSecond",
			expected: "MaxMultipleSecond должен содержать максимум 2 элемента",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString должен содержать минимум 1 символ",
		},
		{
			ns:       "Test.MinStringMultiple",
			expected: "MinStringMultiple должен содержать минимум 2 символа",
		},
		{
			ns:       "Test.MinStringMultipleSecond",
			expected: "MinStringMultipleSecond должен содержать минимум 7 символов",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber должен быть больше или равно 1 113,00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple должен содержать минимум 7 элементов",
		},
		{
			ns:       "Test.MinMultipleSecond",
			expected: "MinMultipleSecond должен содержать минимум 2 элемента",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString должен быть длиной в 1 символ",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber должен быть равен 1 113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple должен содержать 7 элементов",
		},
		{
			ns:       "Test.LenMultipleSecond",
			expected: "LenMultipleSecond должен содержать 2 элемента",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString обязательное поле",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber обязательное поле",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple обязательное поле",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen должен содержать минимум 10 символов",
		},
		{
			ns:       "Test.StrPtrMinLenSecond",
			expected: "StrPtrMinLenSecond должен содержать минимум 2 символа",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen должен содержать максимум 1 символ",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen должен быть длиной в 2 символа",
		},
		{
			ns:       "Test.StrPtrLenSecond",
			expected: "StrPtrLenSecond должен быть длиной в 7 символов",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt должен иметь менее 1 символ",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte должен содержать максимум 1 символ",
		},
		{
			ns:       "Test.StrPtrLteMultiple",
			expected: "StrPtrLteMultiple должен содержать максимум 2 символа",
		},
		{
			ns:       "Test.StrPtrLteMultipleSecond",
			expected: "StrPtrLteMultipleSecond должен содержать максимум 7 символов",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt должен быть длиннее 10 символов",
		},
		{
			ns:       "Test.StrPtrGtSecond",
			expected: "StrPtrGtSecond должен быть длиннее 2 символов",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte должен содержать минимум 10 символов",
		},
		{
			ns:       "Test.StrPtrGteSecond",
			expected: "StrPtrGteSecond должен содержать минимум 2 символа",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString должен быть одним из [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt должен быть одним из [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice должен содержать уникальные значения",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray должен содержать уникальные значения",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap должен содержать уникальные значения",
		},
		{
			ns: "Test.Image",
			expected: "Image должно быть допустимым изображением",
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
