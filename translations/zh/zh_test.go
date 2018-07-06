package en

import (
	"testing"
	"time"

	zhongwen "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	. "gopkg.in/go-playground/assert.v1"
	"gopkg.in/go-playground/validator.v9"
)

func TestTranslations(t *testing.T) {

	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, ok := uni.GetTranslator("zh")

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
			expected: "IsColor必须是一个有效的颜色",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC必须是一个有效的MAC地址",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr必须是一个有效的IP地址",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4必须是一个有效的IPv4地址",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6必须是一个有效的IPv6地址",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr必须是一个有效的UDP地址",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4必须是一个有效的IPv4 UDP地址",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6必须是一个有效的IPv6 UDP地址",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr必须是一个有效的TCP地址",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4必须是一个有效的IPv4 TCP地址",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6必须是一个有效的IPv6 TCP地址",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR必须是一个有效的无类别域间路由(CIDR)",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4必须是一个包含IPv4地址的有效无类别域间路由(CIDR)",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6必须是一个包含IPv6地址的有效无类别域间路由(CIDR)",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN必须是一个有效的社会安全号码(SSN)",
		},
		{
			ns:       "Test.IP",
			expected: "IP必须是一个有效的IP地址",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4必须是一个有效的IPv4地址",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6必须是一个有效的IPv6地址",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI必须包含有效的数据URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude必须包含有效的纬度坐标",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude必须包含有效的经度坐标",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte必须包含多字节字符",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII必须只包含ascii字符",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII必须只包含可打印的ascii字符",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID必须是一个有效的UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3必须是一个有效的V3 UUID",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4必须是一个有效的V4 UUID",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5必须是一个有效的V5 UUID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN必须是一个有效的ISBN号码",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10必须是一个有效的ISBN-10号码",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13必须是一个有效的ISBN-13号码",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes不能包含文本'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll不能包含以下任何字符'!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune不能包含'☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny必须包含至少一个以下字符'!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains必须包含的文本'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64必须是一个有效的Base64字符串",
		},
		{
			ns:       "Test.Email",
			expected: "Email必须是一个有效的邮箱地址",
		},
		{
			ns:       "Test.URL",
			expected: "URL必须是一个有效的URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI必须是一个有效的URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString必须是一个有效的RGB颜色",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString必须是一个有效的RGBA颜色",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString必须是一个有效的HSL颜色",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString必须是有效的HSLA颜色",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString must be a valid hexadecimal",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString must be a valid HEX color",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString must be a valid number",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString must be a valid numeric value",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString can only contain alphanumeric characters",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString can only contain alphabetic characters",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString must be less than MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString must be less than or equal to MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString must be greater than MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString must be greater than or equal to MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString cannot be equal to EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString must be less than Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString must be less than or equal to Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString must be greater than Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString must be greater than or equal to Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString cannot be equal to Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString must be equal to Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString must be equal to MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString must be at least 3 characters in length",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber must be 5.56 or greater",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple must contain at least 2 items",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime must be greater than or equal to the current Date & Time",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString must be greater than 3 characters in length",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber must be greater than 5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple must contain more than 2 items",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime must be greater than the current Date & Time",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString must be at maximum 3 characters in length",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber must be 5.56 or less",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple must contain at maximum 2 items",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime must be less than or equal to the current Date & Time",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString must be less than 3 characters in length",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber must be less than 5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple must contain less than 2 items",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime must be less than the current Date & Time",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString should not be equal to ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber should not be equal to 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple should not be equal to 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString is not equal to 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber is not equal to 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple is not equal to 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString must be a maximum of 3 characters in length",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber must be 1,113.00 or less",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple must contain at maximum 7 items",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString must be at least 1 character in length",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber must be 1,113.00 or greater",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple must contain at least 7 items",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString长度必须是1个字符",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber must be equal to 1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple must contain 7 items",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString is a required field",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber is a required field",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple is a required field",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen must be at least 10 characters in length",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen must be a maximum of 1 character in length",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen must be 2 characters in length",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt must be less than 1 character in length",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte must be at maximum 1 character in length",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt must be greater than 10 characters in length",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte must be at least 10 characters in length",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString must be one of [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt must be one of [5 63]",
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
