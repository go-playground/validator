package fa

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
		RequiredString    string            `validate:"required"`
		RequiredNumber    int               `validate:"required"`
		RequiredMultiple  []string          `validate:"required"`
		LenString         string            `validate:"len=1"`
		LenNumber         float64           `validate:"len=1113.00"`
		LenMultiple       []string          `validate:"len=7"`
		MinString         string            `validate:"min=1"`
		MinNumber         float64           `validate:"min=1113.00"`
		MinMultiple       []string          `validate:"min=7"`
		MaxString         string            `validate:"max=3"`
		MaxNumber         float64           `validate:"max=1113.00"`
		MaxMultiple       []string          `validate:"max=7"`
		EqString          string            `validate:"eq=3"`
		EqNumber          float64           `validate:"eq=2.33"`
		EqMultiple        []string          `validate:"eq=7"`
		NeString          string            `validate:"ne="`
		NeNumber          float64           `validate:"ne=0.00"`
		NeMultiple        []string          `validate:"ne=0"`
		LtString          string            `validate:"lt=3"`
		LtNumber          float64           `validate:"lt=5.56"`
		LtMultiple        []string          `validate:"lt=2"`
		LtTime            time.Time         `validate:"lt"`
		LteString         string            `validate:"lte=3"`
		LteNumber         float64           `validate:"lte=5.56"`
		LteMultiple       []string          `validate:"lte=2"`
		LteTime           time.Time         `validate:"lte"`
		GtString          string            `validate:"gt=3"`
		GtNumber          float64           `validate:"gt=5.56"`
		GtMultiple        []string          `validate:"gt=2"`
		GtTime            time.Time         `validate:"gt"`
		GteString         string            `validate:"gte=3"`
		GteNumber         float64           `validate:"gte=5.56"`
		GteMultiple       []string          `validate:"gte=2"`
		GteTime           time.Time         `validate:"gte"`
		EqFieldString     string            `validate:"eqfield=MaxString"`
		EqCSFieldString   string            `validate:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string            `validate:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string            `validate:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string            `validate:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string            `validate:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string            `validate:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string            `validate:"nefield=EqFieldString"`
		GtFieldString     string            `validate:"gtfield=MaxString"`
		GteFieldString    string            `validate:"gtefield=MaxString"`
		LtFieldString     string            `validate:"ltfield=MaxString"`
		LteFieldString    string            `validate:"ltefield=MaxString"`
		AlphaString       string            `validate:"alpha"`
		AlphanumString    string            `validate:"alphanum"`
		NumericString     string            `validate:"numeric"`
		NumberString      string            `validate:"number"`
		HexadecimalString string            `validate:"hexadecimal"`
		HexColorString    string            `validate:"hexcolor"`
		RGBColorString    string            `validate:"rgb"`
		RGBAColorString   string            `validate:"rgba"`
		HSLColorString    string            `validate:"hsl"`
		HSLAColorString   string            `validate:"hsla"`
		Email             string            `validate:"email"`
		URL               string            `validate:"url"`
		URI               string            `validate:"uri"`
		Base64            string            `validate:"base64"`
		Contains          string            `validate:"contains=purpose"`
		ContainsAny       string            `validate:"containsany=!@#$"`
		Excludes          string            `validate:"excludes=text"`
		ExcludesAll       string            `validate:"excludesall=!@#$"`
		ExcludesRune      string            `validate:"excludesrune=☻"`
		ISBN              string            `validate:"isbn"`
		ISBN10            string            `validate:"isbn10"`
		ISBN13            string            `validate:"isbn13"`
		UUID              string            `validate:"uuid"`
		UUID3             string            `validate:"uuid3"`
		UUID4             string            `validate:"uuid4"`
		UUID5             string            `validate:"uuid5"`
		ULID              string            `validate:"ulid"`
		ASCII             string            `validate:"ascii"`
		PrintableASCII    string            `validate:"printascii"`
		MultiByte         string            `validate:"multibyte"`
		DataURI           string            `validate:"datauri"`
		Latitude          string            `validate:"latitude"`
		Longitude         string            `validate:"longitude"`
		SSN               string            `validate:"ssn"`
		IP                string            `validate:"ip"`
		IPv4              string            `validate:"ipv4"`
		IPv6              string            `validate:"ipv6"`
		CIDR              string            `validate:"cidr"`
		CIDRv4            string            `validate:"cidrv4"`
		CIDRv6            string            `validate:"cidrv6"`
		TCPAddr           string            `validate:"tcp_addr"`
		TCPAddrv4         string            `validate:"tcp4_addr"`
		TCPAddrv6         string            `validate:"tcp6_addr"`
		UDPAddr           string            `validate:"udp_addr"`
		UDPAddrv4         string            `validate:"udp4_addr"`
		UDPAddrv6         string            `validate:"udp6_addr"`
		IPAddr            string            `validate:"ip_addr"`
		IPAddrv4          string            `validate:"ip4_addr"`
		IPAddrv6          string            `validate:"ip6_addr"`
		UinxAddr          string            `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string            `validate:"mac"`
		IsColor           string            `validate:"iscolor"`
		StrPtrMinLen      *string           `validate:"min=10"`
		StrPtrMaxLen      *string           `validate:"max=1"`
		StrPtrLen         *string           `validate:"len=2"`
		StrPtrLt          *string           `validate:"lt=1"`
		StrPtrLte         *string           `validate:"lte=1"`
		StrPtrGt          *string           `validate:"gt=10"`
		StrPtrGte         *string           `validate:"gte=10"`
		OneOfString       string            `validate:"oneof=red green"`
		OneOfInt          int               `validate:"oneof=5 63"`
		UniqueSlice       []string          `validate:"unique"`
		UniqueArray       [3]string         `validate:"unique"`
		UniqueMap         map[string]string `validate:"unique"`
		JSONString        string            `validate:"json"`
		LowercaseString   string            `validate:"lowercase"`
		UppercaseString   string            `validate:"uppercase"`
		Datetime          string            `validate:"datetime=2006-01-02"`
		PostCode          string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry   string
		PostCodeByField   string `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
		Image			  string			`validate:"image"`
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
			expected: "IsColor باید یک رنگ معتبر باشد",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC باید یک مک‌آدرس معتبر باشد",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr باید یک آدرس آی‌پی قابل دسترس باشد",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 باید یک آدرس آی‌پی IPv4 قابل دسترس باشد",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 باید یک آدرس آی‌پی IPv6 قابل دسترس باشد",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr باید یک آدرس UDP معتبر باشد",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 باید یک آدرس UDP IPv4 معتبر باشد",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 باید یک آدرس UDP IPv6 معتبر باشد",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr باید یک آدرس TCP معتبر باشد",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 باید یک آدرس TCP IPv4 معتبر باشد",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 باید یک آدرس TCP IPv6 معتبر باشد",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR باید یک نشانه‌گذاری CIDR معتبر باشد",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 باید یک نشانه‌گذاری CIDR معتبر برای آدرس آی‌پی IPv4 باشد",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 باید یک نشانه‌گذاری CIDR معتبر برای آدرس آی‌پی IPv6 باشد",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN باید یک شماره SSN معتبر باشد",
		},
		{
			ns:       "Test.IP",
			expected: "IP باید یک آدرس آی‌پی معتبر باشد",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 باید یک آدرس آی‌پی IPv4 معتبر باشد",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 باید یک آدرس آی‌پی IPv6 معتبر باشد",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI باید یک Data URI معتبر باشد",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude باید یک عرض جغرافیایی معتبر باشد",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude باید یک طول جغرافیایی معتبر باشد",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte باید شامل کاراکترهای چندبایته باشد",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII باید فقط شامل کاراکترهای اسکی باشد",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII باید فقط شامل کاراکترهای اسکی قابل چاپ باشد",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID باید یک UUID معتبر باشد",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 باید یک UUID نسخه 3 معتبر باشد",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 باید یک UUID نسخه 4 معتبر باشد",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 باید یک UUID نسخه 5 معتبر باشد",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID باید یک ULID معتبر باشد",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN باید یک شابک معتبر باشد",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 باید یک شابک(ISBN-10) معتبر باشد",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 باید یک شابک(ISBN-13) معتبر باشد",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes نمیتواند شامل 'text' باشد",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll نمیتواند شامل کاراکترهای '!@#$' باشد",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune نمیتواند شامل '☻' باشد",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny باید شامل کاراکترهای '!@#$' باشد",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains باید شامل 'purpose' باشد",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 باید یک متن درمبنای64 معتبر باشد",
		},
		{
			ns:       "Test.Email",
			expected: "Email باید یک ایمیل معتبر باشد",
		},
		{
			ns:       "Test.URL",
			expected: "URL باید یک آدرس اینترنتی معتبر باشد",
		},
		{
			ns:       "Test.URI",
			expected: "URI باید یک URI معتبر باشد",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString باید یک کد رنگ RGB باشد",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString باید یک کد رنگ RGBA باشد",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString باید یک کد رنگ HSL باشد",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString باید یک کد رنگ HSLA باشد",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString باید یک عدد درمبنای16 باشد",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString باید یک کد رنگ HEX باشد",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString باید یک عدد معتبر باشد",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString باید یک عدد معتبر باشد",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString میتواند فقط شامل حروف و اعداد باشد",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString میتواند فقط شامل حروف باشد",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "طول LtFieldString باید کمتر از MaxString باشد",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "طول LteFieldString باید کمتر یا برابر MaxString باشد",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "طول GtFieldString باید بیشتر از MaxString باشد",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "طول GteFieldString باید بیشتر یا برابر MaxString باشد",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString نمیتواند برابر EqFieldString باشد",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "طول LtCSFieldString باید کمتر از Inner.LtCSFieldString باشد",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "طول LteCSFieldString باید کمتر یا برابر Inner.LteCSFieldString باشد",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "طول GtCSFieldString باید بیشتر از Inner.GtCSFieldString باشد",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "طول GteCSFieldString باید بیشتر یا برابر Inner.GteCSFieldString باشد",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString نمیتواند برابر Inner.NeCSFieldString باشد",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString باید برابر Inner.EqCSFieldString باشد",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString باید برابر MaxString باشد",
		},
		{
			ns:       "Test.GteString",
			expected: "طول GteString باید حداقل 3 کاراکتر باشد",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber باید بیشتر یا برابر 5.56 باشد",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple باید شامل حداقل 2 آیتم باشد",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime باید بعد یا برابر تاریخ و زمان کنونی باشد",
		},
		{
			ns:       "Test.GtString",
			expected: "طول GtString باید بیشتر از 3 کاراکتر باشد",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber باید بیشتر از 5.56 باشد",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple باید دارای بیشتر از 2 آیتم باشد",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime باید بعد از تاریخ و زمان کنونی باشد",
		},
		{
			ns:       "Test.LteString",
			expected: "طول LteString باید حداکثر 3 کاراکتر باشد",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber باید کمتر یا برابر 5.56 باشد",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple باید حداکثر شامل 2 آیتم باشد",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime باید قبل یا برابر تاریخ و زمان کنونی باشد",
		},
		{
			ns:       "Test.LtString",
			expected: "طول LtString باید کمتر از 3 کاراکتر باشد",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber باید کمتر از 5.56 باشد",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple باید دارای کمتر از 2 آیتم باشد",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime باید قبل از تاریخ و زمان کنونی باشد",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString نباید برابر  باشد",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber نباید برابر 0.00 باشد",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple نباید برابر 0 باشد",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString برابر 3 نمیباشد",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber برابر 2.33 نمیباشد",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple برابر 7 نمیباشد",
		},
		{
			ns:       "Test.MaxString",
			expected: "طول MaxString باید حداکثر 3 کاراکتر باشد",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber باید کمتر یا برابر 1,113.00 باشد",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple باید شامل حداکثر 7 آیتم باشد",
		},
		{
			ns:       "Test.MinString",
			expected: "طول MinString باید حداقل 1 کاراکتر باشد",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber باید بزرگتر یا برابر 1,113.00 باشد",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple باید شامل حداقل 7 آیتم باشد",
		},
		{
			ns:       "Test.LenString",
			expected: "طول LenString باید 1 کاراکتر باشد",
		},
		{
			ns:       "Test.LenNumber",
			expected: "طول LenNumber باید برابر 1,113.00 باشد",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "تعداد LenMultiple باید برابر 7 آیتم باشد",
		},
		{
			ns:       "Test.RequiredString",
			expected: "فیلد RequiredString اجباری میباشد",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "فیلد RequiredNumber اجباری میباشد",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "فیلد RequiredMultiple اجباری میباشد",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "طول StrPtrMinLen باید حداقل 10 کاراکتر باشد",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "طول StrPtrMaxLen باید حداکثر 1 کاراکتر باشد",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "طول StrPtrLen باید 2 کاراکتر باشد",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "طول StrPtrLt باید کمتر از 1 کاراکتر باشد",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "طول StrPtrLte باید حداکثر 1 کاراکتر باشد",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "طول StrPtrGt باید بیشتر از 10 کاراکتر باشد",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "طول StrPtrGte باید حداقل 10 کاراکتر باشد",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString باید یکی از مقادیر [red green] باشد",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt باید یکی از مقادیر [5 63] باشد",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice باید شامل مقادیر منحصربفرد باشد",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray باید شامل مقادیر منحصربفرد باشد",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap باید شامل مقادیر منحصربفرد باشد",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString باید یک json معتبر باشد",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString باید یک متن با حروف کوچک باشد",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString باید یک متن با حروف بزرگ باشد",
		},
		{
			ns:       "Test.Datetime",
			expected: "فرمت Datetime با 2006-01-02 سازگار نیست",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode یک کدپستی معتبر کشور SG نیست",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField یک کدپستی معتبر کشور فیلد PostCodeCountry نیست",
		},
		{
			ns:         "Test.Image",
			expected: "Image باید یک تصویر معتبر باشد",
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
