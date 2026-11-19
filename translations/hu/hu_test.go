package hu

import (
	"reflect"
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	hungarian "github.com/go-playground/locales/hu"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Foo struct{}

func (Foo) IsBar() bool { return false }

func TestTranslations(t *testing.T) {
	hun := hungarian.New()
	uni := ut.New(hun, hun)
	trans, _ := uni.GetTranslator("hu")

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
		RequiredString        string            `validate:"required"`
		RequiredNumber        int               `validate:"required"`
		RequiredMultiple      []string          `validate:"required"`
		RequiredIf            string            `validate:"required_if=Inner.RequiredIf abcd"`
		RequiredUnless        string            `validate:"required_unless=Inner.RequiredUnless abcd"`
		RequiredWith          string            `validate:"required_with=Inner.RequiredWith"`
		RequiredWithAll       string            `validate:"required_with_all=Inner.RequiredWith Inner.RequiredWithAll"`
		RequiredWithout       string            `validate:"required_without=Inner.RequiredWithout"`
		RequiredWithoutAll    string            `validate:"required_without_all=Inner.RequiredWithout Inner.RequiredWithoutAll"`
		ExcludedIf            string            `validate:"excluded_if=Inner.ExcludedIf abcd"`
		ExcludedUnless        string            `validate:"excluded_unless=Inner.ExcludedUnless abcd"`
		ExcludedWith          string            `validate:"excluded_with=Inner.ExcludedWith"`
		ExcludedWithout       string            `validate:"excluded_with_all=Inner.ExcludedWithAll"`
		ExcludedWithAll       string            `validate:"excluded_without=Inner.ExcludedWithout"`
		ExcludedWithoutAll    string            `validate:"excluded_without_all=Inner.ExcludedWithoutAll"`
		IsDefault             string            `validate:"isdefault"`
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
		AlphaSpaceString      string            `validate:"alphaspace"`
		AlphaNumSpaceString   string            `validate:"alphanumspace"`
		AlphaUnicodeString    string            `validate:"alphaunicode"`
		AlphaNumUnicodeString string            `validate:"alphanumunicode"`
		NumericString         string            `validate:"numeric"`
		NumberString          string            `validate:"number"`
		HexadecimalString     string            `validate:"hexadecimal"`
		HexColorString        string            `validate:"hexcolor"`
		RGBColorString        string            `validate:"rgb"`
		RGBAColorString       string            `validate:"rgba"`
		HSLColorString        string            `validate:"hsl"`
		HSLAColorString       string            `validate:"hsla"`
		Email                 string            `validate:"email"`
		URL                   string            `validate:"url"`
		URI                   string            `validate:"uri"`
		Base64                string            `validate:"base64"`
		Contains              string            `validate:"contains=purpose"`
		ContainsAny           string            `validate:"containsany=!@#$"`
		StartsWith            string            `validate:"startswith=start"`
		EndsWith              string            `validate:"endswith=end"`
		StartsNotWith         string            `validate:"startsnotwith=start"`
		EndsNotWith           string            `validate:"endsnotwith=end"`
		Excludes              string            `validate:"excludes=text"`
		ExcludesAll           string            `validate:"excludesall=!@#$"`
		ExcludesRune          string            `validate:"excludesrune=☻"`
		ISBN                  string            `validate:"isbn"`
		ISBN10                string            `validate:"isbn10"`
		ISBN13                string            `validate:"isbn13"`
		ISSN                  string            `validate:"issn"`
		URN                   string            `validate:"urn_rfc2141"`
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
		FQDN                  string            `validate:"fqdn"`
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
		Timezone              string            `validate:"timezone"`
		PostCode              string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry       string
		PostCodeByField       string        `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
		BooleanString         string        `validate:"boolean"`
		Image                 string        `validate:"image"`
		MIMEType              string        `validate:"mimetype=image/png"`
		CveString             string        `validate:"cve"`
		MinDuration           time.Duration `validate:"min=1h30m,max=2h"`
		MaxDuration           time.Duration `validate:"min=1h30m,max=2h"`
		ValidateFn            Foo           `validate:"validateFn=IsBar"`
		EqInt                 int           `validate:"eq=1000"`
		NeIntPtr              *int          `validate:"ne=1000"`
		EqDuration            time.Duration `validate:"eq=1h30m"`
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
	test.StartsWith = "end"
	test.EndsWith = "start"
	test.StartsNotWith = "start value"
	test.EndsNotWith = "value end"

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
	test.Timezone = "abc"
	test.BooleanString = "A"
	test.CveString = "A"

	test.MinDuration = 1 * time.Hour
	test.MaxDuration = 3 * time.Hour

	neInt := 1000
	test.NeIntPtr = &neInt

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
			expected: "IsColor értékének érvényes színkódnak kell lennie",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC értékének érvényes MAC-címet kell tartalmaznia",
		},
		{
			ns:       "Test.FQDN",
			expected: "FQDN értékének érvényes teljes tartománynévnek (FQDN) kell lennie",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr értékének feloldható IP-címnek kell lennie",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 értékének feloldható IPv4-címnek kell lennie",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 értékének feloldható IPv6-címnek kell lennie",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr értékének érvényes UDP-címnek kell lennie",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 értékének érvényes IPv4-es UDP-címnek kell lennie",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 értékének érvényes IPv6-os UDP-címnek kell lennie",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr értékének érvényes TCP-címnek kell lennie",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 értékének érvényes IPv4-es TCP-címnek kell lennie",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 értékének érvényes IPv6-os TCP-címnek kell lennie",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR értékének érvényes CIDR-jelölést kell tartalmaznia",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 értékének érvényes, IPv4-címhez tartozó CIDR-jelölést kell tartalmaznia",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 értékének érvényes, IPv6-címhez tartozó CIDR-jelölést kell tartalmaznia",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN értékének érvényes társadalombiztosítási számnak (SSN) kell lennie",
		},
		{
			ns:       "Test.IP",
			expected: "IP értékének érvényes IP-címnek kell lennie",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 értékének érvényes IPv4-címnek kell lennie",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 értékének érvényes IPv6-címnek kell lennie",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI értékének érvényes Data URI-t kell tartalmaznia",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude értékének érvényes földrajzi szélességi koordinátát kell tartalmaznia",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude értékének érvényes földrajzi hosszúsági koordinátát kell tartalmaznia",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte értékének több bájtból álló karaktereket kell tartalmaznia",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII értéke csak ASCII-karaktereket tartalmazhat",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII értéke csak nyomtatható ASCII-karaktereket tartalmazhat",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID értékének érvényes UUID-azonosítónak kell lennie",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 értékének érvényes, 3-as verziójú UUID-azonosítónak kell lennie",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 értékének érvényes, 4-es verziójú UUID-azonosítónak kell lennie",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 értékének érvényes, 5-ös verziójú UUID-azonosítónak kell lennie",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID értékének érvényes ULID-azonosítónak kell lennie",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN értékének érvényes ISBN-számnak kell lennie",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 értékének érvényes ISBN-10 számnak kell lennie",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 értékének érvényes ISBN-13 számnak kell lennie",
		},
		{
			ns:       "Test.ISSN",
			expected: "ISSN értékének érvényes ISSN-számnak kell lennie",
		},
		{
			ns:       "Test.URN",
			expected: "URN értékének érvényes, RFC 2141 szerinti URN-azonosítónak kell lennie",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes értéke nem tartalmazhatja a következő szöveget: 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll értéke nem tartalmazhatja a következő karakterek egyikét sem: '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune értéke nem tartalmazhatja a következőt: '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny értékének tartalmaznia kell legalább egyet a következő karakterek közül: '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains értékének tartalmaznia kell a következő szöveget: 'purpose'",
		},
		{
			ns:       "Test.StartsWith",
			expected: "StartsWith értékének a következővel kell kezdődnie: 'start'",
		},
		{
			ns:       "Test.EndsWith",
			expected: "EndsWith értékének a következővel kell végződnie: 'end'",
		},
		{
			ns:       "Test.StartsNotWith",
			expected: "StartsNotWith értéke nem kezdődhet a következővel: 'start'",
		},
		{
			ns:       "Test.EndsNotWith",
			expected: "EndsNotWith értéke nem végződhet a következővel: 'end'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 értékének érvényes Base64-karakterláncnak kell lennie",
		},
		{
			ns:       "Test.Email",
			expected: "Email értékének érvényes e-mail-címnek kell lennie",
		},
		{
			ns:       "Test.URL",
			expected: "URL értékének érvényes URL-címnek kell lennie",
		},
		{
			ns:       "Test.URI",
			expected: "URI értékének érvényes URI-azonosítónak kell lennie",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString értékének érvényes RGB-színkódnak kell lennie",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString értékének érvényes RGBA-színkódnak kell lennie",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString értékének érvényes HSL-színkódnak kell lennie",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString értékének érvényes HSLA-színkódnak kell lennie",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString értékének érvényes hexadecimális számnak kell lennie",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString értékének érvényes HEX-színkódnak kell lennie",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString értékének érvényes számnak kell lennie",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString értékének érvényes numerikus értéknek kell lennie",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString értéke csak alfanumerikus karaktereket tartalmazhat",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString értéke csak betűket tartalmazhat",
		},
		{
			ns:       "Test.AlphaSpaceString",
			expected: "AlphaSpaceString értéke csak betűket és szóközöket tartalmazhat",
		},
		{
			ns:       "Test.AlphaNumSpaceString",
			expected: "AlphaNumSpaceString értéke csak alfanumerikus karaktereket és szóközöket tartalmazhat",
		},
		{
			ns:       "Test.AlphaUnicodeString",
			expected: "AlphaUnicodeString értéke csak Unicode-betűket tartalmazhat",
		},
		{
			ns:       "Test.AlphaNumUnicodeString",
			expected: "AlphaNumUnicodeString értéke csak alfanumerikus Unicode-karaktereket tartalmazhat",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString értékének kisebbnek kell lennie, mint MaxString értéke",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString értéke nem lehet nagyobb, mint MaxString értéke",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString értékének nagyobbnak kell lennie, mint MaxString értéke",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString értéke nem lehet kisebb, mint MaxString értéke",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString értéke nem egyezhet meg EqFieldString értékével",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString értékének kisebbnek kell lennie, mint Inner.LtCSFieldString értéke",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString értéke nem lehet nagyobb, mint Inner.LteCSFieldString értéke",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString értékének nagyobbnak kell lennie, mint Inner.GtCSFieldString értéke",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString értéke nem lehet kisebb, mint Inner.GteCSFieldString értéke",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString értéke nem egyezhet meg Inner.NeCSFieldString értékével",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString értékének meg kell egyeznie Inner.EqCSFieldString értékével",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString értékének meg kell egyeznie MaxString értékével",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString értékének legalább 3 karakter hosszúnak kell lennie",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber értéke nem lehet kisebb, mint 5,56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple értékének legalább 2 elemet kell tartalmaznia",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime értéke nem lehet korábbi, mint az aktuális dátum és idő",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString értékének hosszabbnak kell lennie, mint 3 karakter",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber értékének nagyobbnak kell lennie, mint 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple értékének több mint 2 elemet kell tartalmaznia",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime értékének későbbinek kell lennie, mint az aktuális dátum és idő",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString értéke legfeljebb 3 karakter hosszú lehet",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber értéke nem lehet nagyobb, mint 5,56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple értéke legfeljebb 2 elemet tartalmazhat",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime értéke nem lehet későbbi, mint az aktuális dátum és idő",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString értékének rövidebbnek kell lennie, mint 3 karakter",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber értékének kisebbnek kell lennie, mint 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple értékének kevesebb mint 2 elemet kell tartalmaznia",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime értékének korábbinak kell lennie, mint az aktuális dátum és idő",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString értéke nem lehet egyenlő a következővel: ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber értéke nem lehet egyenlő a következővel: 0,00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple értéke nem lehet egyenlő a következővel: 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString értékének egyenlőnek kell lennie ezzel: 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber értékének egyenlőnek kell lennie ezzel: 2,33",
		},
		{
			ns:       "Test.EqInt",
			expected: "EqInt értékének egyenlőnek kell lennie ezzel: 1\u00a0000",
		},
		{
			ns:       "Test.NeIntPtr",
			expected: "NeIntPtr értéke nem lehet egyenlő a következővel: 1\u00a0000",
		},
		{
			ns:       "Test.EqDuration",
			expected: "EqDuration értékének egyenlőnek kell lennie ezzel: 1h30m",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple értékének egyenlőnek kell lennie ezzel: 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString értéke legfeljebb 3 karakter hosszú lehet",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber értéke nem lehet nagyobb, mint 1\u00a0113,00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple értéke legfeljebb 7 elemet tartalmazhat",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString értékének legalább 1 karakter hosszúnak kell lennie",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber értéke nem lehet kisebb, mint 1\u00a0113,00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple értékének legalább 7 elemet kell tartalmaznia",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString értékének 1 karakter hosszúnak kell lennie",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber értékének egyenlőnek kell lennie a következővel: 1\u00a0113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple értékének 7 elemet kell tartalmaznia",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString kötelező mező",
		},
		{
			ns:       "Test.RequiredIf",
			expected: "RequiredIf kötelező mező",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber kötelező mező",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple kötelező mező",
		},
		{
			ns:       "Test.RequiredUnless",
			expected: "RequiredUnless kötelező mező",
		},
		{
			ns:       "Test.RequiredWith",
			expected: "RequiredWith kötelező mező",
		},
		{
			ns:       "Test.RequiredWithAll",
			expected: "RequiredWithAll kötelező mező",
		},
		{
			ns:       "Test.RequiredWithout",
			expected: "RequiredWithout kötelező mező",
		},
		{
			ns:       "Test.RequiredWithoutAll",
			expected: "RequiredWithoutAll kötelező mező",
		},
		{
			ns:       "Test.ExcludedIf",
			expected: "ExcludedIf kizárt mező",
		},
		{
			ns:       "Test.ExcludedUnless",
			expected: "ExcludedUnless kizárt mező",
		},
		{
			ns:       "Test.ExcludedWith",
			expected: "ExcludedWith kizárt mező",
		},
		{
			ns:       "Test.ExcludedWithAll",
			expected: "ExcludedWithAll kizárt mező",
		},
		{
			ns:       "Test.ExcludedWithout",
			expected: "ExcludedWithout kizárt mező",
		},
		{
			ns:       "Test.ExcludedWithoutAll",
			expected: "ExcludedWithoutAll kizárt mező",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen értékének legalább 10 karakter hosszúnak kell lennie",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen értéke legfeljebb 1 karakter hosszú lehet",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen értékének 2 karakter hosszúnak kell lennie",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt értékének rövidebbnek kell lennie, mint 1 karakter",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte értéke legfeljebb 1 karakter hosszú lehet",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt értékének hosszabbnak kell lennie, mint 10 karakter",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte értékének legalább 10 karakter hosszúnak kell lennie",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString értékének a következők egyikének kell lennie: [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt értékének a következők egyikének kell lennie: [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice értéke csak egyedi elemeket tartalmazhat",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray értéke csak egyedi elemeket tartalmazhat",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap értéke csak egyedi elemeket tartalmazhat",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString értékének érvényes JSON-karakterláncnak kell lennie",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString értékének érvényes JWT-karakterláncnak kell lennie",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString értékének kisbetűs karakterláncnak kell lennie",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString értékének nagybetűs karakterláncnak kell lennie",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime értéke nem felel meg a következő formátumnak: 2006-01-02",
		},
		{
			ns:       "Test.Timezone",
			expected: "Timezone értékének érvényes időzónának kell lennie",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode értéke nem felel meg a következő ország irányítószám-formátumának: SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField értéke nem felel meg a következő mezőben megadott ország irányítószám-formátumának: PostCodeCountry",
		},
		{
			ns:       "Test.BooleanString",
			expected: "BooleanString értékének érvényes logikai értéknek kell lennie",
		},
		{
			ns:       "Test.Image",
			expected: "Image értékének érvényes képnek kell lennie",
		},
		{
			ns:       "Test.MIMEType",
			expected: "MIMEType értékének érvényes MIME-típusnak kell lennie",
		},
		{
			ns:       "Test.CveString",
			expected: "CveString értékének érvényes CVE-azonosítónak kell lennie",
		},
		{
			ns:       "Test.MinDuration",
			expected: "MinDuration értéke nem lehet kisebb, mint 1h30m",
		},
		{
			ns:       "Test.MaxDuration",
			expected: "MaxDuration értéke nem lehet nagyobb, mint 2h",
		},
		{
			ns:       "Test.ValidateFn",
			expected: "ValidateFn értékének érvényes objektumnak kell lennie",
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

// stubFieldError implements only the methods fmtNumberParam calls; any other
// method panics on the nil embedded interface.
type stubFieldError struct {
	validator.FieldError

	typ   reflect.Type
	param string
}

func (f stubFieldError) Type() reflect.Type { return f.typ }
func (f stubFieldError) Param() string      { return f.param }

func TestFmtNumberParam(t *testing.T) {
	hun := hungarian.New()
	uni := ut.New(hun, hun)
	trans, _ := uni.GetTranslator("hu")

	tests := []struct {
		name     string
		typ      reflect.Type
		param    string
		expected string
	}{
		{"int", reflect.TypeOf(0), "1000", "1\u00a0000"},
		{"negative int", reflect.TypeOf(0), "-1000", "-1\u00a0000"},
		{"zero", reflect.TypeOf(0), "0", "0"},
		{"uint", reflect.TypeOf(uint64(0)), "1000000", "1\u00a0000\u00a0000"},
		{"float", reflect.TypeOf(float64(0)), "2.33", "2,33"},
		{"float with trailing zeros", reflect.TypeOf(float32(0)), "1113.00", "1\u00a0113,00"},

		// The validator reports a pointer type only for nil pointers, other
		// pointers are dereferenced before validation.
		{"nil pointer to int", reflect.TypeOf((*int)(nil)), "1000", "1\u00a0000"},
		{"nil pointer to float", reflect.TypeOf((*float64)(nil)), "2.5", "2,5"},

		// Everything else is returned as is.
		{"string", reflect.TypeOf(""), "1000", "1000"},
		{"string with decimal point", reflect.TypeOf(""), "2.33", "2.33"},
		{"duration", reflect.TypeOf(time.Duration(0)), "1h30m", "1h30m"},
		{"duration without unit", reflect.TypeOf(time.Duration(0)), "1000", "1000"},
		{"nil pointer to duration", reflect.TypeOf((*time.Duration)(nil)), "1000", "1000"},
		{"slice", reflect.TypeOf([]int(nil)), "3", "3"},
		{"no type", nil, "1000", "1000"},
		{"unparsable param", reflect.TypeOf(0), "abc", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Equal(t, tt.expected, fmtNumberParam(trans, stubFieldError{typ: tt.typ, param: tt.param}))
		})
	}
}
