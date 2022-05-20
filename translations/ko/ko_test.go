package ko

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	korean "github.com/go-playground/locales/ko"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	kor := korean.New()
	uni := ut.New(kor, kor)
	trans, _ := uni.GetTranslator("ko")

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
		RequiredWithout    string
		RequiredWithoutAll string
	}

	type Test struct {
		Inner                 Inner
		RequiredString        string            `validate:"required"`
		RequiredNumber        int               `validate:"required"`
		RequiredMultiple      []string          `validate:"required"`
		RequiredIf            string            `validate:"required_if=Inner.RequiredIf abcd"`
		RequiredUnless        string            `validate:"required_unless=Inner.RequiredUnless abcd"`
		RequiredWith          string            `validate:"required_with=Inner.RequiredWith"`
		RequiredWithAll       string            `validate:"required_with_all=Inner.GtCSFieldString Inner.GteCSFieldString"`
		RequiredWithout       string            `validate:"required_without=Inner.RequiredWithout"`
		RequiredWithoutAll    string            `validate:"required_without_all=Inner.RequiredUnless Inner.RequiredWithout"`
		LenString             string            `validate:"len=1"`
		LenNumber             float64           `validate:"len=1113.00"`
		LenMultiple           []string          `validate:"len=7"`
		MinString             string            `validate:"min=1"`
		MinNumber             float64           `validate:"min=1113.00"`
		MinMultiple           []string          `validate:"min=7"`
		MaxString             string            `validate:"max=3"`
		MaxNumber             float64           `validate:"max=1113.00"`
		MaxMultiple           []string          `validate:"max=7"`
		EqString              string            `validate:"eq=3"`
		EqNumber              float64           `validate:"eq=2.33"`
		EqMultiple            []string          `validate:"eq=7"`
		NeString              string            `validate:"ne="`
		NeNumber              float64           `validate:"ne=0.00"`
		NeMultiple            []string          `validate:"ne=0"`
		LtString              string            `validate:"lt=3"`
		LtNumber              float64           `validate:"lt=5.56"`
		LtMultiple            []string          `validate:"lt=2"`
		LtTime                time.Time         `validate:"lt"`
		LteString             string            `validate:"lte=3"`
		LteNumber             float64           `validate:"lte=5.56"`
		LteMultiple           []string          `validate:"lte=2"`
		LteTime               time.Time         `validate:"lte"`
		GtString              string            `validate:"gt=3"`
		GtNumber              float64           `validate:"gt=5.56"`
		GtMultiple            []string          `validate:"gt=2"`
		GtTime                time.Time         `validate:"gt"`
		GteString             string            `validate:"gte=3"`
		GteNumber             float64           `validate:"gte=5.56"`
		GteMultiple           []string          `validate:"gte=2"`
		GteTime               time.Time         `validate:"gte"`
		EqFieldString         string            `validate:"eqfield=MaxString"`
		EqCSFieldString       string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString       string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString       string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString      string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString       string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString      string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString         string            `validate:"nefield=EqFieldString"`
		GtFieldString         string            `validate:"gtfield=MaxString"`
		GteFieldString        string            `validate:"gtefield=MaxString"`
		LtFieldString         string            `validate:"ltfield=MaxString"`
		LteFieldString        string            `validate:"ltefield=MaxString"`
		AlphaString           string            `validate:"alpha"`
		AlphanumString        string            `validate:"alphanum"`
		AlphanumUnicodeString string            `validate:"alphanumunicode"`
		AlphaUnicodeString    string            `validate:"alphaunicode"`
		NumericString         string            `validate:"numeric"`
		NumberString          string            `validate:"number"`
		HexadecimalString     string            `validate:"hexadecimal"`
		HexColorString        string            `validate:"hexcolor"`
		RGBColorString        string            `validate:"rgb"`
		RGBAColorString       string            `validate:"rgba"`
		HSLColorString        string            `validate:"hsl"`
		HSLAColorString       string            `validate:"hsla"`
		E164                  string            `validate:"e164"`
		Email                 string            `validate:"email"`
		URL                   string            `validate:"url"`
		URI                   string            `validate:"uri"`
		Base64                string            `validate:"base64"`
		Contains              string            `validate:"contains=purpose"`
		ContainsAny           string            `validate:"containsany=!@#$"`
		ContainsRune          string            `validate:"containsrune=☻"`
		Excludes              string            `validate:"excludes=text"`
		ExcludesAll           string            `validate:"excludesall=!@#$"`
		ExcludesRune          string            `validate:"excludesrune=☻"`
		EndsWith              string            `validate:"endswith=end"`
		StartsWith            string            `validate:"startswith=start"`
		ISBN                  string            `validate:"isbn"`
		ISBN10                string            `validate:"isbn10"`
		ISBN13                string            `validate:"isbn13"`
		UUID                  string            `validate:"uuid"`
		UUID3                 string            `validate:"uuid3"`
		UUID4                 string            `validate:"uuid4"`
		UUID5                 string            `validate:"uuid5"`
		ULID                  string            `validate:"ulid"`
		ASCII                 string            `validate:"ascii"`
		PrintableASCII        string            `validate:"printascii"`
		MultiByte             string            `validate:"multibyte"`
		DataURI               string            `validate:"datauri"`
		Latitude              string            `validate:"latitude"`
		Longitude             string            `validate:"longitude"`
		SSN                   string            `validate:"ssn"`
		IP                    string            `validate:"ip"`
		IPv4                  string            `validate:"ipv4"`
		IPv6                  string            `validate:"ipv6"`
		CIDR                  string            `validate:"cidr"`
		CIDRv4                string            `validate:"cidrv4"`
		CIDRv6                string            `validate:"cidrv6"`
		TCPAddr               string            `validate:"tcp_addr"`
		TCPAddrv4             string            `validate:"tcp4_addr"`
		TCPAddrv6             string            `validate:"tcp6_addr"`
		UDPAddr               string            `validate:"udp_addr"`
		UDPAddrv4             string            `validate:"udp4_addr"`
		UDPAddrv6             string            `validate:"udp6_addr"`
		IPAddr                string            `validate:"ip_addr"`
		IPAddrv4              string            `validate:"ip4_addr"`
		IPAddrv6              string            `validate:"ip6_addr"`
		UinxAddr              string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC                   string            `validate:"mac"`
		IsColor               string            `validate:"iscolor"`
		StrPtrMinLen          *string           `validate:"min=10"`
		StrPtrMaxLen          *string           `validate:"max=1"`
		StrPtrLen             *string           `validate:"len=2"`
		StrPtrLt              *string           `validate:"lt=1"`
		StrPtrLte             *string           `validate:"lte=1"`
		StrPtrGt              *string           `validate:"gt=10"`
		StrPtrGte             *string           `validate:"gte=10"`
		OneOfString           string            `validate:"oneof=red green"`
		OneOfInt              int               `validate:"oneof=5 63"`
		UniqueSlice           []string          `validate:"unique"`
		UniqueArray           [3]string         `validate:"unique"`
		UniqueMap             map[string]string `validate:"unique"`
		JSONString            string            `validate:"json"`
		JWTString             string            `validate:"jwt"`
		LowercaseString       string            `validate:"lowercase"`
		UppercaseString       string            `validate:"uppercase"`
		Datetime              string            `validate:"datetime=2006-01-02"`
		BooleanString         string            `validate:"boolean"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"
	test.Inner.RequiredIf = "abcd"
	test.Inner.RequiredUnless = ""
	test.Inner.RequiredWith = "abcd"

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

	test.JSONString = "{\"foo\":\"bar\",}"

	test.LowercaseString = "ABCDEFG"
	test.UppercaseString = "abcdefg"

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}

	test.Datetime = "20060102"
	test.BooleanString = "A"

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
			expected: "IsColor은효과적인색상이어야합니다",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC은효과적인MAC주소여야합니다",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr은효과적인IP주소여야합니다",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4은효과적인IPv4주소여야합니다",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6은효과적인IPv6주소여야합니다",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr은효과적인UDP주소여야합니다",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4은효과적인IPv4 UDP주소여야합니다",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6은효과적인IPv6 UDP주소여야합니다",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr은효과적인TCP주소여야합니다",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4은효과적인IPv4 TCP주소여야합니다",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6은효과적인IPv6 TCP주소여야합니다",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR은효과적인CIDR이어야합니다",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4은IPv4를포함하는CIDR이어야합니다",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6은IPv6를포함하는CIDR이어야합니다",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN은효과적인사회보장번호(SSN)여야합니다",
		},
		{
			ns:       "Test.IP",
			expected: "IP은효과적인IP주소여야합니다",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4은효과적인IPv4주소여야합니다",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6은효과적인IPv6주소여야합니다",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI에는효과적인데이터URI가포함되어야합니다",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude에는효과적인위도좌표가포함되어야합니다",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude에는효과적인종방향좌표가포함되어야합니다",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte에는멀티바이트문자가포함되어야합니다",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII에는ASCII문자만포함해야합니다",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII에는인쇄가능한ASCII문자만포함해야합니다",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID은효과적인UUID여야합니다",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3은효과적인V3 UUID여야합니다",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4은효과적인V4 UUID여야합니다",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5은효과적인V5 UUID여야합니다",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID은효과적인ULID여야합니다",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN은유효ISBN번호여야합니다",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10은효과적인ISBN-10번호여야합니다",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13은효과적인ISBN-13번호여야합니다",
		},
		{
			ns:       "Test.EndsWith",
			expected: "EndsWith텍스트'end'으로끝나야합니다",
		},
		{
			ns:       "Test.StartsWith",
			expected: "StartsWith텍스트'start'으로시작해야합니다",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes은텍스트를포함할수없습니다'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll은다음문자중하나를포함할수없습니다'!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune은'☻'을포함할수없습니다",
		},
		{
			ns:       "Test.ContainsRune",
			expected: "ContainsRune은문자를포함해야합니다'☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny은하나이상의문자를포함해야합니다'!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains은텍스트를포함해야합니다'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64은효과적인Base64문자열이어야합니다",
		},
		{
			ns:       "Test.E164",
			expected: "E164은효과적인E.164휴대폰번호여야합니다",
		},
		{
			ns:       "Test.Email",
			expected: "Email은효과적인사서함이어야합니다",
		},
		{
			ns:       "Test.URL",
			expected: "URL은효과적인URL이어야합니다",
		},
		{
			ns:       "Test.URI",
			expected: "URI은효과적인URI여야합니다",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString은효과적인RGB색상이어야합니다",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString은효과적인RGBA색상이어야합니다",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString은효과적인HSL색상이어야합니다",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString은효과적인HSLA색상이어야합니다",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString은효과적인16진수여야합니다",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString은효과적인16진수색상이어야합니다",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString은유효숫자여야합니다",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString은유효숫자값이어야합니다",
		},
		{
			ns:       "Test.AlphaUnicodeString",
			expected: "AlphaUnicodeString은문자와Unicode문자만포함할수있습니다",
		},
		{
			ns:       "Test.AlphanumUnicodeString",
			expected: "AlphanumUnicodeString은문자,숫자및Unicode문자만포함할수있습니다",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString에는문자와숫자만포함할수있습니다",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString에는문자만포함할수있습니다",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString은MaxString보다작아야합니다",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString은MaxString보다작거나같아야합니다",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString은MaxString보다커야합니다",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString은MaxString보다크거나같아야합니다",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString은EqFieldString과같지않아야합니다",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString은Inner.LtCSFieldString보다작아야합니다",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString은Inner.LteCSFieldString보다작거나같아야합니다",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString은Inner.GtCSFieldString보다커야합니다",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString은Inner.GteCSFieldString보다크거나같아야합니다",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString은Inner.NeCSFieldString과같지않아야합니다",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString은Inner.EqCSFieldString과같아야합니다",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString은MaxString과같아야합니다",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString길이는3자이상이어야합니다",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber은5.56보다크거나같아야합니다",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple은적어도2항목을포함해야합니다",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime은현재날짜및시간보다크거나동일해야합니다",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString길이는3자보다커야합니다",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber은5.56보다커야합니다",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple은2항목보다커야합니다",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime은현재날짜와시간보다커야합니다",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString길이는3자을초과할수없습니다",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber은5.56보다작거나같아야합니다",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple은최대2항목만포함할수있습니다",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime은현재날짜및시간보다작거나동일해야합니다",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString길이는3자보다작아야합니다",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber은5.56보다작아야합니다",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple은2항목미만을포함해야합니다",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime은현재날짜와시간보다작아야합니다",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString은과같지않아야합니다",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber은0.00과같지않아야합니다",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple은0과같지않아야합니다",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString은3과같지않습니다",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber은2.33과같지않습니다",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple은7과같지않습니다",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString길이는3자을초과할수없습니다",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber은1,113.00보다작거나같아야합니다",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple은최대7항목만포함할수있습니다",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString길이는1자이상이어야합니다",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber는1,113.00이상이어야합니다",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple은적어도7항목을포함해야합니다",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString길이는1자이어야합니다",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber은1,113.00과같아야합니다",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple은7항목을포함해야합니다",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString필요한필드입니다",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber필요한필드입니다",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple필요한필드입니다",
		},
		{
			ns:       "Test.RequiredIf",
			expected: "RequiredIf필요한필드입니다",
		},
		{
			ns:       "Test.RequiredUnless",
			expected: "RequiredUnless필요한필드입니다",
		},
		{
			ns:       "Test.RequiredWith",
			expected: "RequiredWith필요한필드입니다",
		},
		{
			ns:       "Test.RequiredWithAll",
			expected: "RequiredWithAll필요한필드입니다",
		},
		{
			ns:       "Test.RequiredWithout",
			expected: "RequiredWithout필요한필드입니다",
		},
		{
			ns:       "Test.RequiredWithoutAll",
			expected: "RequiredWithoutAll필요한필드입니다",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen길이는10자이상이어야합니다",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen길이는1자을초과할수없습니다",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen길이는2자이어야합니다",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt길이는1자보다작아야합니다",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte길이는1자을초과할수없습니다",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt길이는10자보다커야합니다",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte길이는10자이상이어야합니다",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString은[red green]중하나여야합니다",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt은[5 63]중하나여야합니다",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice필드의값은독특해야합니다",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray필드의값은독특해야합니다",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap필드의값은독특해야합니다",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString은효과적인JSON문자열이어야합니다",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString은효과적인JWT문자열이어야합니다",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString은소문자여야합니다",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString은대문자여야합니다",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime의형식은2006-01-02이어야합니다",
		},
		{
			ns:       "Test.BooleanString",
			expected: "BooleanString은효과적인부울값이어야합니다",
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
