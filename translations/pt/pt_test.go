package pt

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
)

func TestTranslations(t *testing.T) {

	pt := pt.New()
	uni := ut.New(pt, pt)
	trans, _ := uni.GetTranslator("pt")

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
			expected: "IsColor deve ser uma cor válida",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC deve conter um endereço MAC válido",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr deve ser um endereço IP resolvível",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 deve ser um endereço IPv4 resolvível",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 deve ser um endereço IPv6 resolvível",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr deve ser um endereço UDP válido",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 deve ser um endereço UDP IPv4 válido",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 deve ser um endereço UDP IPv6 válido",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr deve ser um endereço TCP válido",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 deve ser um endereço TCP IPv4 válido",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 deve ser um endereço TCP IPv6 válido",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR não respeita a notação CIDR",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 não respeita a notação CIDR para um endereço IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 não respeita a notação CIDR para um endereço IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN deve ser um número SSN válido",
		},
		{
			ns:       "Test.IP",
			expected: "IP deve ser um endereço IP válido",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 deve ser um endereço IPv4 válido",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 deve ser um endereço IPv6 válido",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI deve conter um Data URI válido",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude deve conter uma coordenada de latitude válida",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude deve conter uma coordenada de longitude válida",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte deve conter caracteres multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII deve conter apenas caracteres ascii",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII deve conter apenas caracteres ascii imprimíveis",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID deve ser um UUID válido",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 deve ser um UUID versão 3 válido",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 deve ser um UUID versão 4 válido",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 deve ser um UUID versão 5 válido",
		},
		{
			ns:       "Test.ULID",
			expected: "ULID deve ser um ULID válido",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN deve ser um número de ISBN válido",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 deve ser um número ISBN-10 válido",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 deve ser um número ISBN-13 válido",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes não deve conter o texto 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll não deve conter os seguintes caracteres '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune não pode conter o seguinte '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny deve conter pelo menos um dos seguintes caracteres '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains deve conter o texto 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 deve ser uma string Base64 válida",
		},
		{
			ns:       "Test.Email",
			expected: "Email deve ser um endereço de e-mail válido",
		},
		{
			ns:       "Test.URL",
			expected: "URL deve ser um URL válido",
		},
		{
			ns:       "Test.URI",
			expected: "URI deve ser um URI válido",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString deve ser uma cor RGB válida",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString deve ser uma cor RGBA válida",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString deve ser uma cor HSL válida",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString deve ser uma cor HSLA válida",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString deve ser um hexadecimal válido",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString deve ser uma cor HEX válida",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString deve ser um número válido",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString deve ser um valor numérico válido",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString deve conter apenas caracteres alfanuméricos",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString deve conter apenas caracteres alfabéticos",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString deve ser menor que MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString deve ser menor ou igual que MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString deve ser maior que MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString deve ser maior ou igual que MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString não deve ser igual a EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString deve ser menor que Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString deve ser menor ou igual que Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString deve ser maior que Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString deve ser maior ou igual que Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString não deve ser igual a Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString deve ser igual a Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString deve ser igual a MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString deve ter pelo menos 3 caracteres",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber deve ser maior ou igual a 5,56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple deve conter pelo menos 2 items",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime deve ser posterior ou igual à data/hora atual",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString deve conter mais de 3 caracteres",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber deve ser maior que 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple deve conter mais de 2 items",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime deve ser posterior à data/hora atual",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString deve ter no máximo 3 caracteres",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber deve ser menor ou igual a 5,56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple deve conter no máximo 2 items",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime deve ser anterior ou igual à data/hora atual",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString deve ter menos de 3 caracteres",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber deve ser menor que 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple deve conter menos de 2 items",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime deve ser anterior à data / hora atual",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString não deve ser igual a ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber não deve ser igual a 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple não deve ser igual a 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString não é igual a 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber não é igual a 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple não é igual a 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString deve ter no máximo 3 caracteres",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber deve ser 1.113,00 ou menos",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple deve conter no máximo 7 items",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString deve ter pelo menos 1 caractere",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber deve ser 1.113,00 ou superior",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple deve conter pelo menos 7 items",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString deve ter 1 caractere",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber deve ser igual a 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple deve conter 7 items",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString é obrigatório",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber é obrigatório",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple é obrigatório",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen deve ter pelo menos 10 caracteres",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen deve ter no máximo 1 caractere",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen deve ter 2 caracteres",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt deve ter menos de 1 caractere",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte deve ter no máximo 1 caractere",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt deve conter mais de 10 caracteres",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte deve ter pelo menos 10 caracteres",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString deve ser um de [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt deve ser um de [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice deve conter valores únicos",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray deve conter valores únicos",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap deve conter valores únicos",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString deve ser uma string json válida",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString deve estar em minuscúlas",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString deve estar em maiúsculas",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime não está no formato 2006-01-02",
		},
		{
			ns: "Test.Image",
			expected: "Image deve ser uma imagem válida",
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
