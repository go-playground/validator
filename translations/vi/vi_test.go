package vi

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	vietnamese "github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {
	vie := vietnamese.New()
	uni := ut.New(vie, vie)
	trans, _ := uni.GetTranslator("vi")

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
			expected: "IsColor phải là màu sắc hợp lệ",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC chỉ được chứa địa chỉ MAC",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr phải là địa chỉ IP có thể phân giải",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 phải là địa chỉ IPv4 có thể phân giải",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 phải là địa chỉ IPv6 có thể phân giải",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr phải là địa chỉ UDP",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 phải là địa chỉ IPv4 UDP",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 phải là địa chỉ IPv6 UDP",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr phải là địa chỉ TCP",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 phải là địa chỉ IPv4 TCP",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 phải là địa chỉ IPv6 TCP",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR chỉ được chứa CIDR notation",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 chỉ được chứa CIDR notation của một địa chỉ IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 chỉ được chứa CIDR notation của một địa chỉ IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN phải là SSN number",
		},
		{
			ns:       "Test.IP",
			expected: "IP phải là địa chỉ IP",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 phải là địa chỉ IPv4",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 phải là địa chỉ IPv6",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI chỉ được chứa Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude chỉ được chứa latitude (vỹ độ)",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude chỉ được chứa longitude (kinh độ)",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte chỉ được chứa ký tự multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII chỉ được chứa ký tự ASCII",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII chỉ được chứa ký tự ASCII có thể in ấn",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID phải là giá trị UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 phải là giá trị UUID phiên bản 3",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 phải là giá trị UUID phiên bản 4",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 phải là giá trị UUID phiên bản 5",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN phải là số ISBN",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 phải là số ISBN-10",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 phải là số ISBN-13",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes không được chứa chuỗi 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll không được chứa bất kỳ ký tự nào trong nhóm ký tự '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune không được chứa '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny phải chứa ít nhất 1 trong cách ký tự sau '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains phải chứa chuỗi 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 phải là giá trị chuỗi Base64",
		},
		{
			ns:       "Test.Email",
			expected: "Email phải là giá trị email address",
		},
		{
			ns:       "Test.URL",
			expected: "URL phải là giá trị URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI phải là giá trị URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString phải là giá trị RGB color",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString phải là giá trị RGBA color",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString phải là giá trị HSL color",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString phải là giá trị HSLA color",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString phải là giá trị hexadecimal",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString phải là giá trị HEX color",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString chỉ được chứa giá trị số",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString chỉ được chứa giá trị số hoặc số dưới dạng chữ",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString chỉ được chứa ký tự dạng alphanumeric",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString chỉ được chứa ký tự dạng alphabetic",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString chỉ được nhỏ hơn MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString chỉ được nhỏ hơn hoặc bằng MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString phải lớn hơn MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString phải lớn hơn hoặc bằng MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString không được phép bằng EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString chỉ được nhỏ hơn Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString chỉ được nhỏ hơn hoặc bằng Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString phải lớn hơn Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString phải lớn hơn hoặc bằng Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString không được phép bằng Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString phải bằng Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString phải bằng MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString phải có độ dài ít nhất 3 ký tự",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber phải là 5,56 hoặc lớn hơn",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple phải chứa ít nhất 2 phần tử",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime phải lớn hơn hoặc bằng Ngày & Giờ hiện tại",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString phải có độ dài lớn hơn 3 ký tự",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber phải lớn hơn 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple phải chứa nhiều hơn 2 phần tử",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime phải lớn hơn Ngày & Giờ hiện tại",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString chỉ được có độ dài tối đa là 3 ký tự",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber phải là 5,56 hoặc nhỏ hơn",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple chỉ được chứa nhiều nhất 2 phần tử",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime chỉ được nhỏ hơn hoặc bằng Ngày & Giờ hiện tại",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString phải có độ dài nhỏ hơn 3 ký tự",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber phải nhỏ hơn 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple chỉ được chứa ít hơn 2 phần tử",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime phải nhỏ hơn Ngày & Giờ hiện tại",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString không được bằng ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber không được bằng 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple không được bằng 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString không bằng 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber không bằng 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple không bằng 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString chỉ được chứa tối đa 3 ký tự",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber phải là 1.113,00 hoặc nhỏ hơn",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple chỉ được chứa tối đa 7 phần tử",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString phải chứa ít nhất 1 ký tự",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber phải bằng 1.113,00 hoặc lớn hơn",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple phải chứa ít nhất 7 phần tử",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString phải có độ dài là 1 ký tự",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber phải bằng 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple phải chứa 7 phần tử",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString không được bỏ trống",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber không được bỏ trống",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple không được bỏ trống",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen phải chứa ít nhất 10 ký tự",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen chỉ được chứa tối đa 1 ký tự",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen phải có độ dài là 2 ký tự",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt phải có độ dài nhỏ hơn 1 ký tự",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte chỉ được có độ dài tối đa là 1 ký tự",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt phải có độ dài lớn hơn 10 ký tự",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte phải có độ dài ít nhất 10 ký tự",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString phải là trong những giá trị [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt phải là trong những giá trị [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice chỉ được chứa những giá trị không trùng lặp",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray chỉ được chứa những giá trị không trùng lặp",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap chỉ được chứa những giá trị không trùng lặp",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString phải là một chuỗi json hợp lệ",
		},
		{
			ns:       "Test.JWTString",
			expected: "JWTString phải là một chuỗi jwt hợp lệ",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString phải được viết thường",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString phải được viết hoa",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime không trùng định dạng ngày tháng 2006-01-02",
		},
		{
			ns:       "Test.PostCode",
			expected: "PostCode sai định dạng postcode của quốc gia SG",
		},
		{
			ns:       "Test.PostCodeByField",
			expected: "PostCodeByField sai định dạng postcode của quốc gia tương ứng thuộc trường PostCodeCountry",
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
