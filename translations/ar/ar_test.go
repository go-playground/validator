package ar

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
		ISSN              string            `validate:"issn"`
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
		JWTString         string            `validate:"jwt"`
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
			expected: "يجب أن يكون IsColor لون صالح",
		},
		{
			ns:       "Test.MAC",
			expected: "يجب أن يحتوي MAC على عنوان MAC صالح",
		},
		{
			ns:       "Test.IPAddr",
			expected: "يجب أن يكون IPAddr عنوان IP قابل للحل",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "يجب أن يكون IPAddrv4 عنوان IP قابل للحل",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "يجب أن يكون IPAddrv6 عنوان IPv6 قابل للحل",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "يجب أن يكون UDPAddr عنوان UDP صالح",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "يجب أن يكون UDPAddrv4 عنوان IPv4 UDP صالح",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "يجب أن يكون UDPAddrv6 عنوان IPv6 UDP صالح",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "يجب أن يكون TCPAddr عنوان TCP صالح",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "يجب أن يكون TCPAddrv4 عنوان IPv4 TCP صالح",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "يجب أن يكون TCPAddrv6 عنوان IPv6 TCP صالح",
		},
		{
			ns:       "Test.CIDR",
			expected: "يجب أن يحتوي CIDR على علامة CIDR صالحة",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "يجب أن يحتوي CIDRv4 على علامة CIDR صالحة لعنوان IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "يجب أن يحتوي CIDRv6 على علامة CIDR صالحة لعنوان IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "يجب أن يكون SSN رقم SSN صالح",
		},
		{
			ns:       "Test.IP",
			expected: "يجب أن يكون IP عنوان IP صالح",
		},
		{
			ns:       "Test.IPv4",
			expected: "يجب أن يكون IPv4 عنوان IPv4 صالح",
		},
		{
			ns:       "Test.IPv6",
			expected: "يجب أن يكون IPv6 عنوان IPv6 صالح",
		},
		{
			ns:       "Test.DataURI",
			expected: "يجب أن يحتوي DataURI على URI صالح للبيانات",
		},
		{
			ns:       "Test.Latitude",
			expected: "يجب أن يحتوي Latitude على إحداثيات خط عرض صالحة",
		},
		{
			ns:       "Test.Longitude",
			expected: "يجب أن يحتوي Longitude على إحداثيات خط طول صالحة",
		},
		{
			ns:       "Test.MultiByte",
			expected: "يجب أن يحتوي MultiByte على أحرف متعددة البايت",
		},
		{
			ns:       "Test.ASCII",
			expected: "يجب أن يحتوي ASCII على أحرف ascii فقط",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "يجب أن يحتوي PrintableASCII على أحرف ascii قابلة للطباعة فقط",
		},
		{
			ns:       "Test.UUID",
			expected: "يجب أن يكون UUID UUID صالح",
		},
		{
			ns:       "Test.UUID3",
			expected: "يجب أن يكون UUID3 UUID صالح من النسخة 3",
		},
		{
			ns:       "Test.UUID4",
			expected: "يجب أن يكون UUID4 UUID صالح من النسخة 4",
		},
		{
			ns:       "Test.UUID5",
			expected: "يجب أن يكون UUID5 UUID صالح من النسخة 5",
		},
		{
			ns:       "Test.ULID",
			expected: "يجب أن يكون ULID ULID صالح من نسخة",
		},
		{
			ns:       "Test.ISBN",
			expected: "يجب أن يكون ISBN رقم ISBN صالح",
		},
		{
			ns:       "Test.ISBN10",
			expected: "يجب أن يكون ISBN10 رقم ISBN-10 صالح",
		},
		{
			ns:       "Test.ISBN13",
			expected: "يجب أن يكون ISBN13 رقم ISBN-13 صالح",
		},
		{
			ns:       "Test.ISSN",
			expected: "يجب أن يكون ISSN رقم ISSN صالح",
		},
		{
			ns:       "Test.Excludes",
			expected: "لا يمكن أن يحتوي Excludes على النص 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "لا يمكن أن يحتوي ExcludesAll على أي من الأحرف التالية '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "لا يمكن أن يحتوي ExcludesRune على التالي '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "يجب أن يحتوي ContainsAny على حرف واحد على الأقل من الأحرف التالية '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "يجب أن يحتوي Contains على النص 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "يجب أن يكون Base64 سلسلة Base64 صالحة",
		},
		{
			ns:       "Test.Email",
			expected: "يجب أن يكون Email عنوان بريد إلكتروني صالح",
		},
		{
			ns:       "Test.URL",
			expected: "يجب أن يكون URL رابط إنترنت صالح",
		},
		{
			ns:       "Test.URI",
			expected: "يجب أن يكون URI URI صالح",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "يجب أن يكون RGBColorString لون RGB صالح",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "يجب أن يكون RGBAColorString لون RGBA صالح",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "يجب أن يكون HSLColorString لون HSL صالح",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "يجب أن يكون HSLAColorString لون HSLA صالح",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "يجب أن يكون HexadecimalString عددًا سداسيًا عشريًا صالحاً",
		},
		{
			ns:       "Test.HexColorString",
			expected: "يجب أن يكون HexColorString لون HEX صالح",
		},
		{
			ns:       "Test.NumberString",
			expected: "يجب أن يكون NumberString رقم صالح",
		},
		{
			ns:       "Test.NumericString",
			expected: "يجب أن يكون NumericString قيمة رقمية صالحة",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "يمكن أن يحتوي AlphanumString على أحرف أبجدية رقمية فقط",
		},
		{
			ns:       "Test.AlphaString",
			expected: "يمكن أن يحتوي AlphaString على أحرف أبجدية فقط",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "يجب أن يكون LtFieldString أصغر من MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "يجب أن يكون LteFieldString أصغر من أو يساوي MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "يجب أن يكون GtFieldString أكبر من MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "يجب أن يكون GteFieldString أكبر من أو يساوي MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString لا يمكن أن يساوي EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "يجب أن يكون LtCSFieldString أصغر من Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "يجب أن يكون LteCSFieldString أصغر من أو يساوي Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "يجب أن يكون GtCSFieldString أكبر من Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "يجب أن يكون GteCSFieldString أكبر من أو يساوي Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString لا يمكن أن يساوي Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "يجب أن يكون EqCSFieldString مساويا ل Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "يجب أن يكون EqFieldString مساويا ل MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "يجب أن يكون طول GteString على الأقل 3 أحرف",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber يجب أن يكون 5.56 أو أكبر",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "يجب أن يحتوي GteMultiple على 2 عناصر على الأقل",
		},
		{
			ns:       "Test.GteTime",
			expected: "يجب أن يكون GteTime أكبر من أو يساوي التاريخ والوقت الحاليين",
		},
		{
			ns:       "Test.GtString",
			expected: "يجب أن يكون طول GtString أكبر من 3 أحرف",
		},
		{
			ns:       "Test.GtNumber",
			expected: "يجب أن يكون GtNumber أكبر من 5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "يجب أن يحتوي GtMultiple على أكثر من 2 عناصر",
		},
		{
			ns:       "Test.GtTime",
			expected: "يجب أن يكون GtTime أكبر من التاريخ والوقت الحاليين",
		},
		{
			ns:       "Test.LteString",
			expected: "يجب أن يكون طول LteString كحد أقصى 3 أحرف",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber يجب أن يكون 5.56 أو اقل",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "يجب أن يحتوي LteMultiple على 2 عناصر كحد أقصى",
		},
		{
			ns:       "Test.LteTime",
			expected: "يجب أن يكون LteTime أقل من أو يساوي التاريخ والوقت الحاليين",
		},
		{
			ns:       "Test.LtString",
			expected: "يجب أن يكون طول LtString أقل من 3 أحرف",
		},
		{
			ns:       "Test.LtNumber",
			expected: "يجب أن يكون LtNumber أقل من 5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "يجب أن يحتوي LtMultiple على أقل من 2 عناصر",
		},
		{
			ns:       "Test.LtTime",
			expected: "يجب أن يكون LtTime أقل من التاريخ والوقت الحاليين",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString يجب ألا يساوي ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber يجب ألا يساوي 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple يجب ألا يساوي 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString لا يساوي 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber لا يساوي 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple لا يساوي 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "يجب أن يكون طول MaxString بحد أقصى 3 أحرف",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber يجب أن يكون 1,113.00 أو اقل",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "يجب أن يحتوي MaxMultiple على 7 عناصر كحد أقصى",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString يجب أن يكون 1 حرف أو اقل",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber يجب أن يكون 1,113.00 أو اقل",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "يجب أن يحتوي MinMultiple على 7 عناصر على الأقل",
		},
		{
			ns:       "Test.LenString",
			expected: "يجب أن يكون طول LenString مساويا ل 1 حرف",
		},
		{
			ns:       "Test.LenNumber",
			expected: "يجب أن يكون LenNumber مساويا ل 1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "يجب أن يحتوي LenMultiple على 7 عناصر",
		},
		{
			ns:       "Test.RequiredString",
			expected: "حقل RequiredString مطلوب",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "حقل RequiredNumber مطلوب",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "حقل RequiredMultiple مطلوب",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen يجب أن يكون 10 أحرف أو اقل",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "يجب أن يكون طول StrPtrMaxLen بحد أقصى 1 حرف",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "يجب أن يكون طول StrPtrLen مساويا ل 2 أحرف",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "يجب أن يكون طول StrPtrLt أقل من 1 حرف",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "يجب أن يكون طول StrPtrLte كحد أقصى 1 حرف",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "يجب أن يكون طول StrPtrGt أكبر من 10 أحرف",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "يجب أن يكون طول StrPtrGte على الأقل 10 أحرف",
		},
		{
			ns:       "Test.OneOfString",
			expected: "يجب أن يكون OneOfString واحدا من [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "يجب أن يكون OneOfInt واحدا من [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "يجب أن يحتوي UniqueSlice على قيم فريدة",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "يجب أن يحتوي UniqueArray على قيم فريدة",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "يجب أن يحتوي UniqueMap على قيم فريدة",
		},
		{
			ns:       "Test.JSONString",
			expected: "يجب أن يكون JSONString نص json صالح",
		},
		{
			ns:       "Test.JWTString",
			expected: "يجب أن يكون JWTString نص jwt صالح",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "يجب أن يكون LowercaseString نص حروف صغيرة",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "يجب أن يكون UppercaseString نص حروف كبيرة",
		},
		{
			ns:       "Test.Datetime",
			expected: "لا يتطابق Datetime مع تنسيق 2006-01-02",
		},
		{
			ns:       "Test.PostCode",
			expected: "لا يتطابق PostCode مع تنسيق الرمز البريدي للبلد SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "لا يتطابق PostCodeByField مع تنسيق الرمز البريدي للبلد في حقل PostCodeCountry",
		},
		{
			ns: "Test.Image",
			expected: "يجب أن تكون Image صورة صالحة",
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
