package zh

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
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
		Inner                 Inner
		RequiredString        string    `validate:"required"`
		RequiredNumber        int       `validate:"required"`
		RequiredMultiple      []string  `validate:"required"`
		RequiredIf            string    `validate:"required_if=Inner.RequiredIf abcd"`
		RequiredUnless        string    `validate:"required_unless=Inner.RequiredUnless abcd"`
		RequiredWith          string    `validate:"required_with=Inner.RequiredWith"`
		RequiredWithAll       string    `validate:"required_with_all=Inner.RequiredWith Inner.RequiredWithAll"`
		RequiredWithout       string    `validate:"required_without=Inner.RequiredWithout"`
		RequiredWithoutAll    string    `validate:"required_without_all=Inner.RequiredWithout Inner.RequiredWithoutAll"`
		ExcludedIf            string    `validate:"excluded_if=Inner.ExcludedIf abcd"`
		ExcludedUnless        string    `validate:"excluded_unless=Inner.ExcludedUnless abcd"`
		ExcludedWith          string    `validate:"excluded_with=Inner.ExcludedWith"`
		ExcludedWithout       string    `validate:"excluded_with_all=Inner.ExcludedWithAll"`
		ExcludedWithAll       string    `validate:"excluded_without=Inner.ExcludedWithout"`
		ExcludedWithoutAll    string    `validate:"excluded_without_all=Inner.ExcludedWithoutAll"`
		IsDefault             string    `validate:"isdefault"`
		LenString             string    `validate:"len=1"`
		LenNumber             float64   `validate:"len=1113.00"`
		LenMultiple           []string  `validate:"len=7"`
		MinString             string    `validate:"min=1"`
		MinNumber             float64   `validate:"min=1113.00"`
		MinMultiple           []string  `validate:"min=7"`
		MaxString             string    `validate:"max=3"`
		MaxNumber             float64   `validate:"max=1113.00"`
		MaxMultiple           []string  `validate:"max=7"`
		EqString              string    `validate:"eq=3"`
		EqNumber              float64   `validate:"eq=2.33"`
		EqMultiple            []string  `validate:"eq=7"`
		NeString              string    `validate:"ne="`
		NeNumber              float64   `validate:"ne=0.00"`
		NeMultiple            []string  `validate:"ne=0"`
		LtString              string    `validate:"lt=3"`
		LtNumber              float64   `validate:"lt=5.56"`
		LtMultiple            []string  `validate:"lt=2"`
		LtTime                time.Time `validate:"lt"`
		LteString             string    `validate:"lte=3"`
		LteNumber             float64   `validate:"lte=5.56"`
		LteMultiple           []string  `validate:"lte=2"`
		LteTime               time.Time `validate:"lte"`
		GtString              string    `validate:"gt=3"`
		GtNumber              float64   `validate:"gt=5.56"`
		GtMultiple            []string  `validate:"gt=2"`
		GtTime                time.Time `validate:"gt"`
		GteString             string    `validate:"gte=3"`
		GteNumber             float64   `validate:"gte=5.56"`
		GteMultiple           []string  `validate:"gte=2"`
		GteTime               time.Time `validate:"gte"`
		EqFieldString         string    `validate:"eqfield=MaxString"`
		EqCSFieldString       string    `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString       string    `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString       string    `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString      string    `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString       string    `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString      string    `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString         string    `validate:"nefield=EqFieldString"`
		GtFieldString         string    `validate:"gtfield=MaxString"`
		GteFieldString        string    `validate:"gtefield=MaxString"`
		LtFieldString         string    `validate:"ltfield=MaxString"`
		LteFieldString        string    `validate:"ltefield=MaxString"`
		AlphaString           string    `validate:"alpha"`
		AlphanumString        string    `validate:"alphanum"`
		AlphanumUnicodeString string    `validate:"alphanumunicode"`
		AlphaUnicodeString    string    `validate:"alphaunicode"`
		NumericString         string    `validate:"numeric"`
		NumberString          string    `validate:"number"`
		HexadecimalString     string    `validate:"hexadecimal"`
		HexColorString        string    `validate:"hexcolor"`
		RGBColorString        string    `validate:"rgb"`
		RGBAColorString       string    `validate:"rgba"`
		HSLColorString        string    `validate:"hsl"`
		HSLAColorString       string    `validate:"hsla"`
		Email                 string    `validate:"email"`
		URL                   string    `validate:"url"`
		URI                   string    `validate:"uri"`
		Base64                string    `validate:"base64"`
		Contains              string    `validate:"contains=purpose"`
		ContainsAny           string    `validate:"containsany=!@#$"`
		ContainsRune          string    `validate:"containsrune=☻"`
		Excludes              string    `validate:"excludes=text"`
		ExcludesAll           string    `validate:"excludesall=!@#$"`
		ExcludesRune          string    `validate:"excludesrune=☻"`
		EndsWith              string    `validate:"endswith=end"`
		StartsWith            string    `validate:"startswith=start"`
		ISBN                  string    `validate:"isbn"`
		ISBN10                string    `validate:"isbn10"`
		ISBN13                string    `validate:"isbn13"`
		ISSN                  string    `validate:"issn"`
		UUID                  string    `validate:"uuid"`
		UUID3                 string    `validate:"uuid3"`
		UUID4                 string    `validate:"uuid4"`
		UUID5                 string    `validate:"uuid5"`
		ULID                  string    `validate:"ulid"`
		ASCII                 string    `validate:"ascii"`
		PrintableASCII        string    `validate:"printascii"`
		MultiByte             string    `validate:"multibyte"`
		DataURI               string    `validate:"datauri"`
		Latitude              string    `validate:"latitude"`
		Longitude             string    `validate:"longitude"`
		SSN                   string    `validate:"ssn"`
		IP                    string    `validate:"ip"`
		IPv4                  string    `validate:"ipv4"`
		IPv6                  string    `validate:"ipv6"`
		CIDR                  string    `validate:"cidr"`
		CIDRv4                string    `validate:"cidrv4"`
		CIDRv6                string    `validate:"cidrv6"`
		TCPAddr               string    `validate:"tcp_addr"`
		TCPAddrv4             string    `validate:"tcp4_addr"`
		TCPAddrv6             string    `validate:"tcp6_addr"`
		UDPAddr               string    `validate:"udp_addr"`
		UDPAddrv4             string    `validate:"udp4_addr"`
		UDPAddrv6             string    `validate:"udp6_addr"`
		IPAddr                string    `validate:"ip_addr"`
		IPAddrv4              string    `validate:"ip4_addr"`
		IPAddrv6              string    `validate:"ip6_addr"`
		UinxAddr              string    `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC                   string    `validate:"mac"`
		IsColor               string    `validate:"iscolor"`
		StrPtrMinLen          *string   `validate:"min=10"`
		StrPtrMaxLen          *string   `validate:"max=1"`
		StrPtrLen             *string   `validate:"len=2"`
		StrPtrLt              *string   `validate:"lt=1"`
		StrPtrLte             *string   `validate:"lte=1"`
		StrPtrGt              *string   `validate:"gt=10"`
		StrPtrGte             *string   `validate:"gte=10"`
		OneOfString           string    `validate:"oneof=red green"`
		OneOfInt              int       `validate:"oneof=5 63"`
		JsonString            string    `validate:"json"`
		LowercaseString       string    `validate:"lowercase"`
		UppercaseString       string    `validate:"uppercase"`
		Datetime              string    `validate:"datetime=2006-01-02"`
		Image                 string    `validate:"image"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"
	test.Inner.RequiredIf = "abcd"
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
	test.AlphanumUnicodeString = "abc3啊!"
	test.AlphaUnicodeString = "abc3啊"
	test.NumericString = "12E.00"
	test.NumberString = "12E"

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ☻"

	test.EndsWith = "this is some test text"
	test.StartsWith = "this is some test text"

	test.ASCII = "ｶﾀｶﾅ"
	test.PrintableASCII = "ｶﾀｶﾅ"

	test.MultiByte = "1234feerf"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.JsonString = "{\"foo\":\"bar\",}"

	test.LowercaseString = "ABCDEFG"
	test.UppercaseString = "abcdefg"

	test.Datetime = "20060102"

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
			ns:       "Test.ULID",
			expected: "ULID必须是一个有效的ULID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN必须是一个有效的ISBN编号",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10必须是一个有效的ISBN-10编号",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13必须是一个有效的ISBN-13编号",
		},
		{
			ns:       "Test.ISSN",
			expected: "ISSN必须是一个有效的ISSN编号",
		},
		{
			ns:       "Test.EndsWith",
			expected: "EndsWith必须以文本'end'结尾",
		},
		{
			ns:       "Test.StartsWith",
			expected: "StartsWith必须以文本'start'开头",
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
			ns:       "Test.ContainsRune",
			expected: "ContainsRune必须包含字符'☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny必须包含至少一个以下字符'!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains必须包含文本'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64必须是一个有效的Base64字符串",
		},
		{
			ns:       "Test.Email",
			expected: "Email必须是一个有效的邮箱",
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
			expected: "HSLAColorString必须是一个有效的HSLA颜色",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString必须是一个有效的十六进制",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString必须是一个有效的十六进制颜色",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString必须是一个有效的数字",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString必须是一个有效的数值",
		},
		{
			ns:       "Test.AlphaUnicodeString",
			expected: "AlphaUnicodeString只能包含字母和Unicode字符",
		},
		{
			ns:       "Test.AlphanumUnicodeString",
			expected: "AlphanumUnicodeString只能包含字母数字和Unicode字符",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString只能包含字母和数字",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString只能包含字母",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString必须小于MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString必须小于或等于MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString必须大于MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString必须大于或等于MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString不能等于EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString必须小于Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString必须小于或等于Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString必须大于Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString必须大于或等于Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString不能等于Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString必须等于Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString必须等于MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString长度必须至少为3个字符",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber必须大于或等于5.56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple必须至少包含2项",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime必须大于或等于当前日期和时间",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString长度必须大于3个字符",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber必须大于5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple必须大于2项",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime必须大于当前日期和时间",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString长度不能超过3个字符",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber必须小于或等于5.56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple最多只能包含2项",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime必须小于或等于当前日期和时间",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString长度必须小于3个字符",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber必须小于5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple必须包含少于2项",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime必须小于当前日期和时间",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString不能等于",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber不能等于0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple不能等于0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString不等于3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber不等于2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple不等于7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString长度不能超过3个字符",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber必须小于或等于1,113.00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple最多只能包含7项",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString长度必须至少为1个字符",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber最小只能为1,113.00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple必须至少包含7项",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString长度必须是1个字符",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber必须等于1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple必须包含7项",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString为必填字段",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber为必填字段",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple为必填字段",
		},
		{
			ns:       "Test.RequiredUnless",
			expected: "RequiredUnless为必填字段",
		},
		{
			ns:       "Test.RequiredWith",
			expected: "RequiredWith为必填字段",
		},
		{
			ns:       "Test.RequiredWithAll",
			expected: "RequiredWithAll为必填字段",
		},
		{
			ns:       "Test.RequiredWithout",
			expected: "RequiredWithout为必填字段",
		},
		{
			ns:       "Test.RequiredWithoutAll",
			expected: "RequiredWithoutAll为必填字段",
		},
		{
			ns:       "Test.ExcludedIf",
			expected: "ExcludedIf为禁填字段",
		},
		{
			ns:       "Test.ExcludedUnless",
			expected: "ExcludedUnless为禁填字段",
		},
		{
			ns:       "Test.ExcludedWith",
			expected: "ExcludedWith为禁填字段",
		},
		{
			ns:       "Test.ExcludedWithAll",
			expected: "ExcludedWithAll为禁填字段",
		},
		{
			ns:       "Test.ExcludedWithout",
			expected: "ExcludedWithout为禁填字段",
		},
		{
			ns:       "Test.ExcludedWithoutAll",
			expected: "ExcludedWithoutAll为禁填字段",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen长度必须至少为10个字符",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen长度不能超过1个字符",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen长度必须是2个字符",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt长度必须小于1个字符",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte长度不能超过1个字符",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt长度必须大于10个字符",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte长度必须至少为10个字符",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString必须是[red green]中的一个",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt必须是[5 63]中的一个",
		},
		{
			ns:       "Test.JsonString",
			expected: "JsonString必须是一个JSON字符串",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString必须是小写字母",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString必须是大写字母",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime的格式必须是2006-01-02",
		},
		{
			ns:       "Test.Image",
			expected: "Image 必须是有效图像",
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
