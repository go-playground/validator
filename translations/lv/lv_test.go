package lv

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
		RequiredIf       string
	}

	type Test struct {
		Inner             Inner
		RequiredString    string            `validate:"required"`
		RequiredNumber    int               `validate:"required"`
		RequiredMultiple  []string          `validate:"required"`
		RequiredIf        string            `validate:"required_if=Inner.RequiredIf abcd"`
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
		JWTString         string            `validate:"jwt"`
		LowercaseString   string            `validate:"lowercase"`
		UppercaseString   string            `validate:"uppercase"`
		Datetime          string            `validate:"datetime=2006-01-02"`
		PostCode          string            `validate:"postcode_iso3166_alpha2=SG"`
		PostCodeCountry   string
		PostCodeByField   string `validate:"postcode_iso3166_alpha2_field=PostCodeCountry"`
		BooleanString     string `validate:"boolean"`
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
	test.BooleanString = "A"

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
			expected: "IsColor jābūt derīgai krāsai",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC jābūt derīgai MAC adresei",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr jābūt atrisināmai IP adresei",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 jābūt atrisināmai IPv4 adresei",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 jābūt atrisināmai IPv6 adresei",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr jābūt derīgai UDP adresei",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 jābūt derīgai IPv4 UDP adresei",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 jābūt derīgai IPv6 UDP adresei",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr jābūt derīgai TCP adresei",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 jābūt derīgai IPv4 TCP adresei",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 jābūt derīgai IPv6 TCP adresei",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR jāsatur derīgu CIDR notāciju",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 jāsatur derīgu CIDR notāciju IPv4 adresei",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 jāsatur derīgu CIDR notāciju IPv6 adresei",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN jābūt derīgam SSN numuram",
		},
		{
			ns:       "Test.IP",
			expected: "IP jābūt derīgai IP adresei",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 jābūt derīgai IPv4 adresei",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 jābūt derīgai IPv6 adresei",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI jāsatur derīgs Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude jāsatur derīgus platuma grādus",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude jāsatur derīgus garuma grādus",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte jāsatur multibyte rakstu zīmes",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII jāsatur tikai ascii rakstu zīmes",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII jāsatur tikai drukājamas ascii rakstu zīmes",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID jābūt derīgam UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 jābūt derīgam 3. versijas UUID",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 jābūt derīgam 4. versijas UUID",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 jābūt derīgam 5. versijas UUID",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID jābūt derīgam ULID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN jābūt derīgam ISBN numuram",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 jābūt derīgam ISBN-10 numuram",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 jābūt derīgam ISBN-13 numuram",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes nedrīkst saturēt tekstu 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll nedrīkst saturēt nevienu no sekojošām rakstu zīmēm '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune nedrīkst saturēt sekojošo '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny jāsatur minimums 1 no rakstu zīmēm '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains jāsatur teksts 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 jābūt derīgai Base64 virknei",
		},
		{
			ns:       "Test.Email",
			expected: "Email jābūt derīgai e-pasta adresei",
		},
		{
			ns:       "Test.URL",
			expected: "URL jābūt derīgam URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI jābūt derīgam URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString jābūt derīgai RGB krāsai",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString jābūt derīgai RGBA krāsai",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString jābūt derīgai HSL krāsai",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString jābūt derīgai HSLA krāsai",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString jābūt heksadecimālam skaitlim",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString jābūt derīgai HEX krāsai",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString jāsatur derīgs skaitlis",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString jāsatur tikai cipari",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString jāsatur tikai simboli no alfabēta vai cipari (Alphanumeric)",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString jāsatur tikai simboli no alfabēta",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString jābūt mazākam par MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString jābūt mazākam par MaxString vai vienādam",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString jābūt lielākam par MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString jābūt lielākam par MaxString vai vienādam",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString nedrīkst būt vienāds ar EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString jābūt mazākam par Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString jābūt mazākam par Inner.LteCSFieldString vai vienādam",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString jābūt lielākam par Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString jābūt lielākam par Inner.GteCSFieldString vai vienādam",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString nedrīkst būt vienāds ar Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString jābūt vienādam ar Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString jābūt vienādam ar MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString garumam jābūt minimums 3 rakstu zīmes",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber jābūt 5.56 vai lielākam",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple jāsatur minimums 2 elementi",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime jābūt lielākam par šī brīža Datumu un laiku vai vienādam",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString ir jābūt garākam par 3 rakstu zīmēm",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber jābūt lielākam par 5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple jāsatur vairāk par 2 elementiem",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime jābūt lielākam par šī brīža Datumu un laiku",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString garumam jābūt maksimums 3 rakstu zīmes",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber jābūt 5.56 vai mazākam",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple jāsatur maksimums 2 elementi",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime jābūt mazākam par šī brīža Datumu un laiku vai vienādam",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString garumam jābūt mazākam par 3 rakstu zīmēm",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber jābūt mazākam par 5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple jāsatur mazāk par 2 elementiem",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime jābūt mazākam par šī brīža Datumu un laiku",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString nedrīkst būt vienāds ar ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber nedrīkst būt vienāds ar 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple nedrīkst būt vienāds ar 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString nav vienāds ar 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber nav vienāds ar 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple nav vienāds ar 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString vērtība pārsniedz maksimālo garumu 3 rakstu zīmes",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber vērtībai jābūt 1,113.00 vai mazākai",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple jāsatur maksimums 7 elementi",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString garumam jābūt minimums 1 rakstu zīme",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber vērtībai jābūt 1,113.00 vai lielākai",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple jāsatur minimums 7 elementi",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString garumam jābūt 1 rakstu zīme",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber vērtībai jābūt 1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple vērtībai jāsatur 7 elementi",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString ir obligāts lauks",
		},
		{
			ns:       "Test.RequiredIf",
			expected: "RequiredIf ir obligāts lauks",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber ir obligāts lauks",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple ir obligāts lauks",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen garumam jābūt minimums 10 rakstu zīmes",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen vērtība pārsniedz maksimālo garumu 1 rakstu zīme",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen garumam jābūt 2 rakstu zīmes",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt garumam jābūt mazākam par 1 rakstu zīmi",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte garumam jābūt maksimums 1 rakstu zīme",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt ir jābūt garākam par 10 rakstu zīmēm",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte garumam jābūt minimums 10 rakstu zīmes",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString jābūt vienam no [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt jābūt vienam no [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice jāsatur unikālas vērtības",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray jāsatur unikālas vērtības",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap jāsatur unikālas vērtības",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString jābūt derīgai json virknei",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString jābūt derīgai jwt virknei",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString jābūt mazo burtu virknei",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString jābūt lielo burtu virknei",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime neatbilst formātam 2006-01-02",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode neatbilst pasta indeksa formātam valstī SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField neatbilst pasta indeksa formātam valstī, kura norādīta laukā PostCodeCountry",
		},
		{
			ns:       "Test.BooleanString",
			expected: "BooleanString jābūt derīgai boolean vērtībai",
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
