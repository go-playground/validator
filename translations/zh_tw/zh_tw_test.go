package zh_tw

import (
	"testing"
	"time"

	zhongwen "github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	. "github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")

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
			expected: "IsColor必須是一個有效的顏色",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC必須是一個有效的MAC地址",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr必須是一個有效的IP地址",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4必須是一個有效的IPv4地址",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6必須是一個有效的IPv6地址",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr必須是一個有效的UDP地址",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4必須是一個有效的IPv4 UDP地址",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6必須是一個有效的IPv6 UDP地址",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr必須是一個有效的TCP地址",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4必須是一個有效的IPv4 TCP地址",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6必須是一個有效的IPv6 TCP地址",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR必須是一個有效的無類別域間路由(CIDR)",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4必須是一个包含IPv4地址的有效無類別域間路由(CIDR)",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6必須是一个包含IPv6地址的有效無類別域間路由(CIDR)",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN必須是一個有效的社會安全編號(SSN)",
		},
		{
			ns:       "Test.IP",
			expected: "IP必須是一個有效的IP地址",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4必須是一個有效的IPv4地址",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6必須是一個有效的IPv6地址",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI必須包含有效的數據URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude必須包含有效的緯度座標",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude必須包含有效的經度座標",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte必須包含多個字元",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII必須只包含ascii字元",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII必須只包含可輸出的ascii字元",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID必須是一個有效的UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3必須是一個有效的V3 UUID",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4必須是一個有效的V4 UUID",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5必須是一個有效的V5 UUID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN必須是一個有效的ISBN編號",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10必須是一個有效的ISBN-10編號",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13必須是一個有效的ISBN-13編號",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes不能包含文字'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll不能包含以下任何字元'!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune不能包含'☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny必須包含至少一個以下字元'!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains必須包含文字'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64必須是一個有效的Base64字元串",
		},
		{
			ns:       "Test.Email",
			expected: "Email必須是一個有效的信箱",
		},
		{
			ns:       "Test.URL",
			expected: "URL必須是一個有效的URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI必須是一個有效的URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString必須是一個有效的RGB顏色",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString必須是一個有效的RGBA顏色",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString必須是一個有效的HSL顏色",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString必須是一個有效的HSLA顏色",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString必須是一個有效的十六進制",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString必須是一個有效的十六進制顏色",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString必須是一個有效的數字",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString必須是一個有效的數值",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString只能包含字母和數字",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString只能包含字母",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString必須小於MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString必須小於或等於MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString必須大於MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString必須大於或等於MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString不能等於EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString必須小於Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString必須小於或等於Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString必須大於Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString必須大於或等於Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString不能等於Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString必須等於Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString必須等於MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString長度必須至少為3個字元",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber必須大於或等於5.56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple必須至少包含2項",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime必須大於或等於目前日期和時間",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString長度必須大於3個字元",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber必須大於5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple必須大於2項",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime必須大於目前日期和時間",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString長度不能超過3個字元",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber必須小於或等於5.56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple最多只能包含2項",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime必須小於或等於目前日期和時間",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString長度必須小於3個字元",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber必須小於5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple必須包含少於2項",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime必須小於目前日期和時間",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString不能等於",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber不能等於0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple不能等於0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString不等於3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber不等於2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple不等於7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString長度不能超過3個字元",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber必須小於或等於1,113.00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple最多只能包含7項",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString長度必須至少為1個字元",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber最小只能為1,113.00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple必須至少包含7項",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString長度必須為1個字元",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber必須等於1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple必須包含7項",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString為必填欄位",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber為必填欄位",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple為必填欄位",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen長度必須至少為10個字元",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen長度不能超過1個字元",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen長度必須為2個字元",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt長度必須小於1個字元",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte長度不能超過1個字元",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt長度必須大於10個字元",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte長度必須至少為10個字元",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString必須是[red green]中的一個",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt必須是[5 63]中的一個",
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
