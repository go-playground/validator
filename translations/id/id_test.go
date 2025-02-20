package id

import (
	"testing"
	"time"

	. "github.com/go-playground/assert/v2"
	indonesia "github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// InitValidator initializes a new validator with Indonesian translations
func InitValidator() (*validator.Validate, ut.Translator, error) {

	// setup translator
	idn := indonesia.New()
	uni := ut.New(idn, idn)
	trans, _ := uni.GetTranslator("id")

	// initialize validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	// register translations
	err := RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, nil, err
	}

	return validate, trans, nil
}

// TestFieldTagsTranslations tests all field tags registered translations for Indonesian language
func TestFieldTagsTranslations(t *testing.T) {

	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// Inner struct for cross-field validations
	type Inner struct {
		EqCSFieldString  string
		NeCSFieldString  string
		GtCSFieldString  string
		GteCSFieldString string
		LtCSFieldString  string
		LteCSFieldString string
		FieldContains    string
		FieldExcludes    string
	}

	// TestFieldTags for field validations
	type TestFieldTags struct {
		Inner Inner

		// Equal field validations
		EqCSFieldString string `validate:"eqcsfield=Inner.EqCSFieldString"`
		EqFieldString   string `validate:"eqfield=EqString"`
		EqString        string

		// Not equal field validations
		NeCSFieldString string `validate:"necsfield=Inner.NeCSFieldString"`
		NeFieldString   string `validate:"nefield=NeString"`
		NeString        string

		// Greater than field validations
		GtCSFieldString string `validate:"gtcsfield=Inner.GtCSFieldString"`
		GtFieldString   string `validate:"gtfield=GtString"`
		GtString        string

		// Greater than or equal field validations
		GteCSFieldString string `validate:"gtecsfield=Inner.GteCSFieldString"`
		GteFieldString   string `validate:"gtefield=GteString"`
		GteString        string

		// Less than field validations
		LtCSFieldString string `validate:"ltcsfield=Inner.LtCSFieldString"`
		LtFieldString   string `validate:"ltfield=LtString"`
		LtString        string

		// Less than or equal field validations
		LteCSFieldString string `validate:"ltecsfield=Inner.LteCSFieldString"`
		LteFieldString   string `validate:"ltefield=LteString"`
		LteString        string

		// Field contains/excludes validations
		ContainsField string `validate:"fieldcontains=Inner.FieldContains"`
		ExcludesField string `validate:"fieldexcludes=Inner.FieldExcludes"`
	}

	// init test struct with values
	test := TestFieldTags{
		Inner: Inner{
			EqCSFieldString:  "equal",
			NeCSFieldString:  "not-equal",
			GtCSFieldString:  "zzz",
			GteCSFieldString: "zzz",
			LtCSFieldString:  "aaa",
			LteCSFieldString: "aaa",
			FieldContains:    "contains",
			FieldExcludes:    "excludes",
		},
		EqCSFieldString: "not-equal", // should equal Inner.EqCSFieldString
		EqFieldString:   "not-equal", // should equal EqString
		EqString:        "equal",

		NeCSFieldString: "not-equal", // should not equal Inner.NeCSFieldString
		NeFieldString:   "same",      // should not equal NeString
		NeString:        "same",

		GtCSFieldString: "aaa", // should be greater than Inner.GtCSFieldString
		GtFieldString:   "aaa", // should be greater than GtString
		GtString:        "zzz",

		GteCSFieldString: "aaa",  // should be greater than or equal to Inner.GteCSFieldString
		GteFieldString:   "aaa",  // should be greater than or equal to GteString
		GteString:        "bbbb", // reference value for GteFieldString

		LtCSFieldString: "zzz", // should be less than Inner.LtCSFieldString
		LtFieldString:   "zzz", // should be less than LtString
		LtString:        "aaa",

		LteCSFieldString: "zzzzz", // should be less than or equal to Inner.LteCSFieldString
		LteFieldString:   "zzzz",  // should be less than or equal to LteString
		LteString:        "yyy",   // reference value for LteFieldString

		ContainsField: "xyz",          // should contain Inner.FieldContains
		ExcludesField: "has-excludes", // should not contain Inner.FieldExcludes
	}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestFieldTags.EqCSFieldString",
			expected: "EqCSFieldString harus sama dengan Inner.EqCSFieldString",
		},
		{
			ns:       "TestFieldTags.EqFieldString",
			expected: "EqFieldString harus sama dengan EqString",
		},
		{
			ns:       "TestFieldTags.NeCSFieldString",
			expected: "NeCSFieldString tidak sama dengan Inner.NeCSFieldString",
		},
		{
			ns:       "TestFieldTags.NeFieldString",
			expected: "NeFieldString tidak sama dengan NeString",
		},
		{
			ns:       "TestFieldTags.GtCSFieldString",
			expected: "GtCSFieldString harus lebih besar dari Inner.GtCSFieldString",
		},
		{
			ns:       "TestFieldTags.GtFieldString",
			expected: "GtFieldString harus lebih besar dari GtString",
		},
		{
			ns:       "TestFieldTags.GteCSFieldString",
			expected: "GteCSFieldString harus lebih besar dari atau sama dengan Inner.GteCSFieldString",
		},
		{
			ns:       "TestFieldTags.GteFieldString",
			expected: "GteFieldString harus lebih besar dari atau sama dengan GteString",
		},
		{
			ns:       "TestFieldTags.LtCSFieldString",
			expected: "LtCSFieldString harus kurang dari Inner.LtCSFieldString",
		},
		{
			ns:       "TestFieldTags.LtFieldString",
			expected: "LtFieldString harus kurang dari LtString",
		},
		{
			ns:       "TestFieldTags.LteCSFieldString",
			expected: "LteCSFieldString harus kurang dari atau sama dengan Inner.LteCSFieldString",
		},
		{
			ns:       "TestFieldTags.LteFieldString",
			expected: "LteFieldString harus kurang dari atau sama dengan LteString",
		},
		{
			ns:       "TestFieldTags.ContainsField",
			expected: "ContainsField harus berisi nilai dari field Inner.FieldContains",
		},
		{
			ns:       "TestFieldTags.ExcludesField",
			expected: "ExcludesField tidak boleh berisi nilai dari field Inner.FieldExcludes",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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

// TestNetworkTagsTranslations tests all network tags registered translations for Indonesian language
func TestNetworkTagsTranslations(t *testing.T) {
	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// TestNetworkTags for network validations
	type TestNetworkTags struct {
		CIDR            string `validate:"cidr"`
		CIDRv4          string `validate:"cidrv4"`
		CIDRv6          string `validate:"cidrv6"`
		DataURI         string `validate:"datauri"`
		FQDN            string `validate:"fqdn"`
		Hostname        string `validate:"hostname"`
		HostnamePort    string `validate:"hostname_port"`
		HostnameRFC1123 string `validate:"hostname_rfc1123"`
		IP              string `validate:"ip"`
		IP4Addr         string `validate:"ip4_addr"`
		IP6Addr         string `validate:"ip6_addr"`
		IPAddr          string `validate:"ip_addr"`
		IPv4            string `validate:"ipv4"`
		IPv6            string `validate:"ipv6"`
		MAC             string `validate:"mac"`
		TCP4Addr        string `validate:"tcp4_addr"`
		TCP6Addr        string `validate:"tcp6_addr"`
		TCPAddr         string `validate:"tcp_addr"`
		UDP4Addr        string `validate:"udp4_addr"`
		UDP6Addr        string `validate:"udp6_addr"`
		UDPAddr         string `validate:"udp_addr"`
		UnixAddr        string `validate:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		URI             string `validate:"uri"`
		URL             string `validate:"url"`
		HTTPURL         string `validate:"http_url"`
		URLEncoded      string `validate:"url_encoded"`
		URN             string `validate:"urn_rfc2141"`
	}

	// init test struct
	test := TestNetworkTags{
		URLEncoded: "<%az", // invalid URL encoded string
	}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestNetworkTags.CIDR",
			expected: "CIDR harus berupa notasi CIDR yang valid",
		},
		{
			ns:       "TestNetworkTags.CIDRv4",
			expected: "CIDRv4 harus berupa notasi CIDR IPv4 yang valid",
		},
		{
			ns:       "TestNetworkTags.CIDRv6",
			expected: "CIDRv6 harus berupa notasi CIDR IPv6 yang valid",
		},
		{
			ns:       "TestNetworkTags.DataURI",
			expected: "DataURI harus berisi URI Data yang valid",
		},
		{
			ns:       "TestNetworkTags.FQDN",
			expected: "FQDN harus berupa FQDN yang valid",
		},
		{
			ns:       "TestNetworkTags.Hostname",
			expected: "Hostname harus berupa hostname sesuai RFC 952 yang valid",
		},
		{
			ns:       "TestNetworkTags.HostnamePort",
			expected: "HostnamePort harus berupa hostname dan port yang valid",
		},
		{
			ns:       "TestNetworkTags.HostnameRFC1123",
			expected: "HostnameRFC1123 harus berupa hostname sesuai RFC 1123 yang valid",
		},
		{
			ns:       "TestNetworkTags.IP",
			expected: "IP harus berupa alamat IP yang valid",
		},
		{
			ns:       "TestNetworkTags.IP4Addr",
			expected: "IP4Addr harus berupa alamat IPv4 yang valid",
		},
		{
			ns:       "TestNetworkTags.IP6Addr",
			expected: "IP6Addr harus berupa alamat IPv6 yang valid",
		},
		{
			ns:       "TestNetworkTags.IPAddr",
			expected: "IPAddr harus berupa alamat IP yang valid",
		},
		{
			ns:       "TestNetworkTags.IPv4",
			expected: "IPv4 harus berupa alamat IPv4 yang valid",
		},
		{
			ns:       "TestNetworkTags.IPv6",
			expected: "IPv6 harus berupa alamat IPv6 yang valid",
		},
		{
			ns:       "TestNetworkTags.MAC",
			expected: "MAC harus berisi alamat MAC yang valid",
		},
		{
			ns:       "TestNetworkTags.TCP4Addr",
			expected: "TCP4Addr harus berupa alamat TCP IPv4 yang valid",
		},
		{
			ns:       "TestNetworkTags.TCP6Addr",
			expected: "TCP6Addr harus berupa alamat TCP IPv6 yang valid",
		},
		{
			ns:       "TestNetworkTags.TCPAddr",
			expected: "TCPAddr harus berupa alamat TCP yang valid",
		},
		{
			ns:       "TestNetworkTags.UDP4Addr",
			expected: "UDP4Addr harus berupa alamat IPv4 UDP yang valid",
		},
		{
			ns:       "TestNetworkTags.UDP6Addr",
			expected: "UDP6Addr harus berupa alamat IPv6 UDP yang valid",
		},
		{
			ns:       "TestNetworkTags.UDPAddr",
			expected: "UDPAddr harus berupa alamat UDP yang valid",
		},
		{
			ns:       "TestNetworkTags.URI",
			expected: "URI harus berupa URI yang valid",
		},
		{
			ns:       "TestNetworkTags.URL",
			expected: "URL harus berupa URL yang valid",
		},
		{
			ns:       "TestNetworkTags.HTTPURL",
			expected: "HTTPURL harus berupa URL HTTP/HTTPS yang valid",
		},
		{
			ns:       "TestNetworkTags.URLEncoded",
			expected: "URLEncoded harus berupa string URL yang terenkode",
		},
		{
			ns:       "TestNetworkTags.URN",
			expected: "URN harus berupa URN sesuai RFC 2141 yang valid",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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

// TestStringTagsTranslations tests all string tags registered translations for Indonesian language
func TestStringTagsTranslations(t *testing.T) {

	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// TestStringTags for string validations
	type TestStringTags struct {
		Alpha         string `validate:"alpha"`
		Alphanum      string `validate:"alphanum"`
		AlphanumUni   string `validate:"alphanumunicode"`
		AlphaUni      string `validate:"alphaunicode"`
		ASCII         string `validate:"ascii"`
		Boolean       string `validate:"boolean"`
		Contains      string `validate:"contains=test"`
		ContainsAny   string `validate:"containsany=!@#"`
		ContainsRune  string `validate:"containsrune=☺"`
		EndsNotWith   string `validate:"endsnotwith=end"`
		EndsWith      string `validate:"endswith=end"`
		Excludes      string `validate:"excludes=exclude"`
		ExcludesAll   string `validate:"excludesall=!@#"`
		ExcludesRune  string `validate:"excludesrune=☺"`
		Lowercase     string `validate:"lowercase"`
		Multibyte     string `validate:"multibyte"`
		Number        string `validate:"number"`
		Numeric       string `validate:"numeric"`
		PrintASCII    string `validate:"printascii"`
		StartsNotWith string `validate:"startsnotwith=start"`
		StartsWith    string `validate:"startswith=start"`
		Uppercase     string `validate:"uppercase"`
	}

	// init test struct with invalid values
	test := TestStringTags{
		Alpha:         "123",                // should only contain letters
		Alphanum:      "!@#",                // should only contain letters and numbers
		AlphanumUni:   "!@#",                // should only contain unicode letters and numbers
		AlphaUni:      "123",                // should only contain unicode letters
		ASCII:         "ñ",                  // should only contain ASCII characters
		Boolean:       "invalid",            // should be a valid boolean
		Contains:      "invalid",            // should contain "test"
		ContainsAny:   "abc",                // should contain any of "!@#"
		ContainsRune:  "abc",                // should contain "☺"
		EndsNotWith:   "test-end",           // should not end with "end"
		EndsWith:      "test-no-start",      // should end with "end"
		Excludes:      "has-exclude-here",   // should not contain "exclude"
		ExcludesAll:   "test!@#",            // should not contain any of "!@#"
		ExcludesRune:  "test☺here",          // should not contain "☺"
		Lowercase:     "TEST",               // should be lowercase
		Multibyte:     "abc",                // should contain multibyte characters
		Number:        "abc",                // should be a valid number
		Numeric:       "abc",                // should be numeric
		PrintASCII:    string([]byte{0x7f}), // should only contain printable ASCII
		StartsNotWith: "start-test",         // should not start with "start"
		StartsWith:    "test-no-start",      // should start with "start"
		Uppercase:     "test",               // should be uppercase
	}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestStringTags.Alpha",
			expected: "Alpha hanya dapat berisi karakter alfanumerik",
		},
		{
			ns:       "TestStringTags.Alphanum",
			expected: "Alphanum hanya dapat berisi karakter alfanumerik",
		},
		{
			ns:       "TestStringTags.AlphanumUni",
			expected: "AlphanumUni hanya boleh berisi karakter alfanumerik unicode",
		},
		{
			ns:       "TestStringTags.AlphaUni",
			expected: "AlphaUni hanya boleh berisi karakter alfanumerik unicode",
		},
		{
			ns:       "TestStringTags.ASCII",
			expected: "ASCII hanya boleh berisi karakter ASCII",
		},
		{
			ns:       "TestStringTags.Boolean",
			expected: "Boolean harus berupa nilai boolean yang valid",
		},
		{
			ns:       "TestStringTags.Contains",
			expected: "Contains harus berisi teks 'test'",
		},
		{
			ns:       "TestStringTags.ContainsAny",
			expected: "ContainsAny harus berisi setidaknya salah satu karakter berikut '!@#'",
		},
		{
			ns:       "TestStringTags.ContainsRune",
			expected: "ContainsRune harus berisi setidaknya salah satu karakter berikut '☺'",
		},
		{
			ns:       "TestStringTags.EndsNotWith",
			expected: "EndsNotWith tidak boleh diakhiri dengan 'end'",
		},
		{
			ns:       "TestStringTags.EndsWith",
			expected: "EndsWith harus diakhiri dengan 'end'",
		},
		{
			ns:       "TestStringTags.Excludes",
			expected: "Excludes tidak boleh berisi teks 'exclude'",
		},
		{
			ns:       "TestStringTags.ExcludesAll",
			expected: "ExcludesAll tidak boleh berisi salah satu karakter berikut '!@#'",
		},
		{
			ns:       "TestStringTags.ExcludesRune",
			expected: "ExcludesRune tidak boleh berisi '☺'",
		},
		{
			ns:       "TestStringTags.Lowercase",
			expected: "Lowercase harus berupa string huruf kecil",
		},
		{
			ns:       "TestStringTags.Multibyte",
			expected: "Multibyte harus berisi karakter multibyte",
		},
		{
			ns:       "TestStringTags.Number",
			expected: "Number harus berupa angka yang valid",
		},
		{
			ns:       "TestStringTags.Numeric",
			expected: "Numeric harus berupa nilai numerik yang valid",
		},
		{
			ns:       "TestStringTags.PrintASCII",
			expected: "PrintASCII hanya boleh berisi karakter ASCII yang dapat dicetak",
		},
		{
			ns:       "TestStringTags.StartsNotWith",
			expected: "StartsNotWith tidak boleh diawali dengan 'start'",
		},
		{
			ns:       "TestStringTags.StartsWith",
			expected: "StartsWith harus diawali dengan 'start'",
		},
		{
			ns:       "TestStringTags.Uppercase",
			expected: "Uppercase harus berupa string huruf besar",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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

// TestFormatTagsTranslations tests all format tags registered translations for Indonesian language
func TestFormatTagsTranslations(t *testing.T) {

	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// TestFormatTags for format validations
	type TestFormatTags struct {
		Hexadecimal                string `validate:"hexadecimal"`
		Base64                     string `validate:"base64"`
		Base64URL                  string `validate:"base64url"`
		Base64RawURL               string `validate:"base64rawurl"`
		BIC                        string `validate:"bic"`
		BCP47Lang                  string `validate:"bcp47_language_tag"`
		BTCAddr                    string `validate:"btc_addr"`
		BTCAddrBech32              string `validate:"btc_addr_bech32"`
		CreditCard                 string `validate:"credit_card"`
		MongoDB                    string `validate:"mongodb"`
		MongoDBConn                string `validate:"mongodb_connection_string"`
		Cron                       string `validate:"cron"`
		SpiceDB                    string `validate:"spicedb"`
		DateTime                   string `validate:"datetime=2006-01-02"`
		E164                       string `validate:"e164"`
		Email                      string `validate:"email"`
		EthAddr                    string `validate:"eth_addr"`
		HexColor                   string `validate:"hexcolor"`
		HSL                        string `validate:"hsl"`
		HSLA                       string `validate:"hsla"`
		HTML                       string `validate:"html"`
		HTMLEncoded                string `validate:"html_encoded"`
		ISBN                       string `validate:"isbn"`
		ISBN10                     string `validate:"isbn10"`
		ISBN13                     string `validate:"isbn13"`
		ISSN                       string `validate:"issn"`
		ISO3166Alpha2              string `validate:"iso3166_1_alpha2"`
		ISO3166Alpha3              string `validate:"iso3166_1_alpha3"`
		ISO3166AlphaNumeric        string `validate:"iso3166_1_alpha_numeric"`
		ISO31662                   string `validate:"iso3166_2"`
		ISO4217                    string `validate:"iso4217"`
		JSON                       string `validate:"json"`
		JWT                        string `validate:"jwt"`
		Latitude                   string `validate:"latitude"`
		Longitude                  string `validate:"longitude"`
		LuhnChecksum               string `validate:"luhn_checksum"`
		PostcodeISO3166Alpha2      string `validate:"postcode_iso3166_alpha2"`
		PostcodeISO3166Alpha2Field string `validate:"postcode_iso3166_alpha2_field"`
		RGB                        string `validate:"rgb"`
		RGBA                       string `validate:"rgba"`
		SSN                        string `validate:"ssn"`
		Timezone                   string `validate:"timezone"`
		UUID                       string `validate:"uuid"`
		UUID3                      string `validate:"uuid3"`
		UUID3RFC4122               string `validate:"uuid3_rfc4122"`
		UUID4                      string `validate:"uuid4"`
		UUID4RFC4122               string `validate:"uuid4_rfc4122"`
		UUID5                      string `validate:"uuid5"`
		UUID5RFC4122               string `validate:"uuid5_rfc4122"`
		UUIDRFC4122                string `validate:"uuid_rfc4122"`
		MD4                        string `validate:"md4"`
		MD5                        string `validate:"md5"`
		SHA256                     string `validate:"sha256"`
		SHA384                     string `validate:"sha384"`
		SHA512                     string `validate:"sha512"`
		RIPEMD128                  string `validate:"ripemd128"`
		RIPEMD160                  string `validate:"ripemd160"`
		Tiger128                   string `validate:"tiger128"`
		Tiger160                   string `validate:"tiger160"`
		Tiger192                   string `validate:"tiger192"`
		Semver                     string `validate:"semver"`
		ULID                       string `validate:"ulid"`
		CVE                        string `validate:"cve"`
	}

	// init test struct with invalid values
	test := TestFormatTags{}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestFormatTags.Hexadecimal",
			expected: "Hexadecimal harus berupa heksadesimal yang valid",
		},
		{
			ns:       "TestFormatTags.Base64",
			expected: "Base64 harus berupa string Base64 yang valid",
		},
		{
			ns:       "TestFormatTags.Base64URL",
			expected: "Base64URL harus berupa string Base64 URL yang valid",
		},
		{
			ns:       "TestFormatTags.Base64RawURL",
			expected: "Base64RawURL harus berupa string Base64 Raw URL yang valid",
		},
		{
			ns:       "TestFormatTags.BIC",
			expected: "BIC harus berupa kode BIC (SWIFT) yang valid sesuai ISO 9362",
		},
		{
			ns:       "TestFormatTags.BCP47Lang",
			expected: "BCP47Lang harus berupa tag bahasa BCP 47 yang valid",
		},
		{
			ns:       "TestFormatTags.BTCAddr",
			expected: "BTCAddr harus berupa alamat Bitcoin yang valid",
		},
		{
			ns:       "TestFormatTags.BTCAddrBech32",
			expected: "BTCAddrBech32 harus berupa alamat Bitcoin Bech32 yang valid",
		},
		{
			ns:       "TestFormatTags.CreditCard",
			expected: "CreditCard harus berupa nomor kartu kredit yang valid",
		},
		{
			ns:       "TestFormatTags.MongoDB",
			expected: "MongoDB harus berupa ObjectID MongoDB yang valid",
		},
		{
			ns:       "TestFormatTags.MongoDBConn",
			expected: "MongoDBConn harus berupa string koneksi MongoDB yang valid",
		},
		{
			ns:       "TestFormatTags.Cron",
			expected: "Cron harus berupa ekspresi cron yang valid",
		},
		{
			ns:       "TestFormatTags.SpiceDB",
			expected: "SpiceDB harus berupa format SpiceDB yang valid",
		},
		{
			ns:       "TestFormatTags.DateTime",
			expected: "DateTime tidak sesuai dengan format 2006-01-02",
		},
		{
			ns:       "TestFormatTags.E164",
			expected: "E164 harus berupa nomor telepon format E.164 yang valid",
		},
		{
			ns:       "TestFormatTags.Email",
			expected: "Email harus berupa alamat email yang valid",
		},
		{
			ns:       "TestFormatTags.EthAddr",
			expected: "EthAddr harus berupa alamat Ethereum yang valid",
		},
		{
			ns:       "TestFormatTags.HexColor",
			expected: "HexColor harus berupa warna HEX yang valid",
		},
		{
			ns:       "TestFormatTags.HSL",
			expected: "HSL harus berupa warna HSL yang valid",
		},
		{
			ns:       "TestFormatTags.HSLA",
			expected: "HSLA harus berupa warna HSLA yang valid",
		},
		{
			ns:       "TestFormatTags.HTML",
			expected: "HTML harus berupa HTML yang valid",
		},
		{
			ns:       "TestFormatTags.HTMLEncoded",
			expected: "HTMLEncoded harus berupa HTML terenkode yang valid",
		},
		{
			ns:       "TestFormatTags.ISBN",
			expected: "ISBN harus berupa nomor ISBN yang valid",
		},
		{
			ns:       "TestFormatTags.ISBN10",
			expected: "ISBN10 harus berupa nomor ISBN-10 yang valid",
		},
		{
			ns:       "TestFormatTags.ISBN13",
			expected: "ISBN13 harus berupa nomor ISBN-13 yang valid",
		},
		{
			ns:       "TestFormatTags.ISSN",
			expected: "ISSN harus berupa nomor ISSN yang valid",
		},
		{
			ns:       "TestFormatTags.ISO3166Alpha2",
			expected: "ISO3166Alpha2 harus berupa kode negara ISO 3166-1 alpha-2 yang valid",
		},
		{
			ns:       "TestFormatTags.ISO3166Alpha3",
			expected: "ISO3166Alpha3 harus berupa kode negara ISO 3166-1 alpha-3 yang valid",
		},
		{
			ns:       "TestFormatTags.ISO3166AlphaNumeric",
			expected: "ISO3166AlphaNumeric harus berupa kode negara numerik ISO 3166-1 yang valid",
		},
		{
			ns:       "TestFormatTags.ISO31662",
			expected: "ISO31662 harus berupa kode subdivisi negara ISO 3166-2 yang valid",
		},
		{
			ns:       "TestFormatTags.ISO4217",
			expected: "ISO4217 harus berupa kode mata uang ISO 4217 yang valid",
		},
		{
			ns:       "TestFormatTags.JSON",
			expected: "JSON harus berupa string JSON yang valid",
		},
		{
			ns:       "TestFormatTags.JWT",
			expected: "JWT harus berupa JSON Web Token (JWT) yang valid",
		},
		{
			ns:       "TestFormatTags.Latitude",
			expected: "Latitude harus berisi koordinat lintang yang valid",
		},
		{
			ns:       "TestFormatTags.Longitude",
			expected: "Longitude harus berisi koordinat bujur yang valid",
		},
		{
			ns:       "TestFormatTags.LuhnChecksum",
			expected: "LuhnChecksum harus memiliki checksum Luhn yang valid",
		},
		{
			ns:       "TestFormatTags.PostcodeISO3166Alpha2",
			expected: "PostcodeISO3166Alpha2 tidak sesuai dengan format kode pos negara ",
		},
		{
			ns:       "TestFormatTags.PostcodeISO3166Alpha2Field",
			expected: "PostcodeISO3166Alpha2Field tidak sesuai dengan format kode pos negara dalam field ",
		},
		{
			ns:       "TestFormatTags.RGB",
			expected: "RGB harus berupa warna RGB yang valid",
		},
		{
			ns:       "TestFormatTags.RGBA",
			expected: "RGBA harus berupa warna RGBA yang valid",
		},
		{
			ns:       "TestFormatTags.SSN",
			expected: "SSN harus berupa nomor SSN (Social Security Number) yang valid",
		},
		{
			ns:       "TestFormatTags.Timezone",
			expected: "Timezone harus berupa zona waktu yang valid",
		},
		{
			ns:       "TestFormatTags.UUID",
			expected: "UUID harus berupa UUID yang valid",
		},
		{
			ns:       "TestFormatTags.UUID3",
			expected: "UUID3 harus berupa UUID versi 3 yang valid",
		},
		{
			ns:       "TestFormatTags.UUID3RFC4122",
			expected: "UUID3RFC4122 harus berupa UUID versi 3 RFC4122 yang valid",
		},
		{
			ns:       "TestFormatTags.UUID4",
			expected: "UUID4 harus berupa UUID versi 4 yang valid",
		},
		{
			ns:       "TestFormatTags.UUID4RFC4122",
			expected: "UUID4RFC4122 harus berupa UUID versi 4 RFC4122 yang valid",
		},
		{
			ns:       "TestFormatTags.UUID5",
			expected: "UUID5 harus berupa UUID versi 5 yang valid",
		},
		{
			ns:       "TestFormatTags.UUID5RFC4122",
			expected: "UUID5RFC4122 harus berupa UUID versi 5 RFC4122 yang valid",
		},
		{
			ns:       "TestFormatTags.UUIDRFC4122",
			expected: "UUIDRFC4122 harus berupa UUID RFC4122 yang valid",
		},
		{
			ns:       "TestFormatTags.MD4",
			expected: "MD4 harus berupa hash MD4 yang valid",
		},
		{
			ns:       "TestFormatTags.MD5",
			expected: "MD5 harus berupa hash MD5 yang valid",
		},
		{
			ns:       "TestFormatTags.SHA256",
			expected: "SHA256 harus berupa hash SHA256 yang valid",
		},
		{
			ns:       "TestFormatTags.SHA384",
			expected: "SHA384 harus berupa hash SHA384 yang valid",
		},
		{
			ns:       "TestFormatTags.SHA512",
			expected: "SHA512 harus berupa hash SHA512 yang valid",
		},
		{
			ns:       "TestFormatTags.RIPEMD128",
			expected: "RIPEMD128 harus berupa hash RIPEMD128 yang valid",
		},
		{
			ns:       "TestFormatTags.RIPEMD160",
			expected: "RIPEMD160 harus berupa hash RIPEMD160 yang valid",
		},
		{
			ns:       "TestFormatTags.Tiger128",
			expected: "Tiger128 harus berupa hash TIGER128 yang valid",
		},
		{
			ns:       "TestFormatTags.Tiger160",
			expected: "Tiger160 harus berupa hash TIGER160 yang valid",
		},
		{
			ns:       "TestFormatTags.Tiger192",
			expected: "Tiger192 harus berupa hash TIGER192 yang valid",
		},
		{
			ns:       "TestFormatTags.Semver",
			expected: "Semver harus berupa nomor versi semantik yang valid",
		},
		{
			ns:       "TestFormatTags.ULID",
			expected: "ULID harus berupa ULID yang valid",
		},
		{
			ns:       "TestFormatTags.CVE",
			expected: "CVE harus berupa identifikasi CVE yang valid",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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

// TestComparisonTagsTranslations tests all comparison tags registered translations for Indonesian language
func TestComparisonTagsTranslations(t *testing.T) {

	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// TestComparisonTags for comparison validations
	type TestComparisonTags struct {
		// Equal comparisons
		EqString     string `validate:"eq=test"`
		EqNumber     int    `validate:"eq=10"`
		EqIgnoreCase string `validate:"eq_ignore_case=Test"`

		// Not equal comparisons
		NeString     string `validate:"ne=test"`
		NeNumber     int    `validate:"ne=10"`
		NeIgnoreCase string `validate:"ne_ignore_case=Test"`

		// Greater than comparisons
		GtString string    `validate:"gt=5"` // length > 5
		GtNumber float64   `validate:"gt=10.5"`
		GtTime   time.Time `validate:"gt"`
		GtSlice  []string  `validate:"gt=1"` // length > 1

		// Greater than or equal comparisons
		GteString string    `validate:"gte=5"` // length >= 5
		GteNumber float64   `validate:"gte=10.5"`
		GteTime   time.Time `validate:"gte"`
		GteSlice  []string  `validate:"gte=1"` // length >= 1

		// Less than comparisons
		LtString string    `validate:"lt=5"` // length < 5
		LtNumber float64   `validate:"lt=10.5"`
		LtTime   time.Time `validate:"lt"`
		LtSlice  []string  `validate:"lt=2"` // length < 1

		// Less than or equal comparisons
		LteString string    `validate:"lte=5"` // length <= 5
		LteNumber float64   `validate:"lte=10.5"`
		LteTime   time.Time `validate:"lte"`
		LteSlice  []string  `validate:"lte=1"` // length <= 1
	}

	// init test struct with invalid values
	now := time.Now()
	test := TestComparisonTags{
		EqString:     "not-test",
		EqNumber:     20,
		EqIgnoreCase: "not-test",

		NeString:     "test",
		NeNumber:     10,
		NeIgnoreCase: "Test",

		GtString: "abc",               // length = 3, should be > 5
		GtNumber: 5.5,                 // should be > 10.5
		GtTime:   now.Add(-time.Hour), // should be > now

		GteString: "abc",               // length = 3, should be >= 5
		GteNumber: 5.5,                 // should be >= 10.5
		GteTime:   now.Add(-time.Hour), // should be >= now

		LtString: "toolong",               // length = 7, should be < 5
		LtNumber: 15.5,                    // should be < 10.5
		LtTime:   now.Add(time.Hour),      // should be < now
		LtSlice:  []string{"satu", "dua"}, // should be < 2

		LteString: "toolong",               // length = 7, should be <= 5
		LteNumber: 15.5,                    // should be <= 10.5
		LteTime:   now.Add(time.Hour),      // should be <= now
		LteSlice:  []string{"satu", "dua"}, // should be <= 1
	}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestComparisonTags.EqString",
			expected: "EqString tidak sama dengan test",
		},
		{
			ns:       "TestComparisonTags.EqNumber",
			expected: "EqNumber tidak sama dengan 10",
		},
		{
			ns:       "TestComparisonTags.EqIgnoreCase",
			expected: "EqIgnoreCase harus sama dengan Test (tidak case-sensitive)",
		},
		{
			ns:       "TestComparisonTags.NeString",
			expected: "NeString tidak sama dengan test",
		},
		{
			ns:       "TestComparisonTags.NeNumber",
			expected: "NeNumber tidak sama dengan 10",
		},
		{
			ns:       "TestComparisonTags.NeIgnoreCase",
			expected: "NeIgnoreCase tidak sama dengan Test (tidak case-sensitive)",
		},
		{
			ns:       "TestComparisonTags.GtString",
			expected: "panjang GtString harus lebih dari 5 karakter",
		},
		{
			ns:       "TestComparisonTags.GtNumber",
			expected: "GtNumber harus lebih besar dari 10,5",
		},
		{
			ns:       "TestComparisonTags.GtTime",
			expected: "GtTime harus lebih besar dari tanggal & waktu saat ini",
		},
		{
			ns:       "TestComparisonTags.GtSlice",
			expected: "GtSlice harus berisi lebih dari 1 item",
		},
		{
			ns:       "TestComparisonTags.GteString",
			expected: "panjang minimal GteString adalah 5 karakter",
		},
		{
			ns:       "TestComparisonTags.GteNumber",
			expected: "GteNumber harus 10,5 atau lebih besar",
		},
		{
			ns:       "TestComparisonTags.GteTime",
			expected: "GteTime harus lebih besar dari atau sama dengan tanggal & waktu saat ini",
		},
		{
			ns:       "TestComparisonTags.GteSlice",
			expected: "GteSlice harus berisi setidaknya 1 item",
		},
		{
			ns:       "TestComparisonTags.LtString",
			expected: "panjang LtString harus kurang dari 5 karakter",
		},
		{
			ns:       "TestComparisonTags.LtNumber",
			expected: "LtNumber harus kurang dari 10,5",
		},
		{
			ns:       "TestComparisonTags.LtTime",
			expected: "LtTime harus kurang dari tanggal & waktu saat ini",
		},
		{
			ns:       "TestComparisonTags.LtSlice",
			expected: "LtSlice harus berisi kurang dari 2 item",
		},
		{
			ns:       "TestComparisonTags.LteString",
			expected: "panjang maksimal LteString adalah 5 karakter",
		},
		{
			ns:       "TestComparisonTags.LteNumber",
			expected: "LteNumber harus 10,5 atau kurang",
		},
		{
			ns:       "TestComparisonTags.LteTime",
			expected: "LteTime harus kurang dari atau sama dengan tanggal & waktu saat ini",
		},
		{
			ns:       "TestComparisonTags.LteSlice",
			expected: "LteSlice harus berisi maksimal 1 item",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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

// TestOtherTagsTranslations tests all other tags registered translations for Indonesian language
func TestOtherTagsTranslations(t *testing.T) {

	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// TestOtherTags for other validations
	type Inner struct {
		RequiredWith string
		ExcludedWith string
		Field        string
	}

	type TestOtherTags struct {
		Dir                string   `validate:"dir"`
		DirPath            string   `validate:"dirpath"`
		File               string   `validate:"file"`
		FilePath           string   `validate:"filepath"`
		Image              string   `validate:"image"`
		LenString          string   `validate:"len=5"`
		LenSlice           []string `validate:"len=3"`
		LenNumber          int      `validate:"len=10"`
		MinString          string   `validate:"min=3"`
		MinSlice           []string `validate:"min=1"`
		MaxString          string   `validate:"max=5"`
		MaxSlice           []string `validate:"max=1"`
		IsDefault          string   `validate:"isdefault"`
		Required           string   `validate:"required"`
		RequiredIf         string   `validate:"required_if=Inner.RequiredWith value"`
		RequiredUnless     string   `validate:"required_unless=Inner.RequiredWith values"`
		RequiredWith       string   `validate:"required_with=Inner.RequiredWith"`
		RequiredWithAll    string   `validate:"required_with_all=Inner.RequiredWith"`
		RequiredWithout    string   `validate:"required_without=Inner.ExcludedWith"`
		RequiredWithoutAll string   `validate:"required_without_all=Inner.ExcludedWith"`
		ExcludedIf         string   `validate:"excluded_if=Inner.RequiredWith value"`
		ExcludedUnless     string   `validate:"excluded_unless=Inner.ExcludedWith value"`
		ExcludedWith       string   `validate:"excluded_with=Inner.RequiredWith"`
		ExcludedWithAll    string   `validate:"excluded_with_all=Inner.RequiredWith"`
		ExcludedWithout    string   `validate:"excluded_without=Inner.ExcludedWith"`
		ExcludedWithoutAll string   `validate:"excluded_without_all=Inner.ExcludedWith"`
		OneOf              string   `validate:"oneof=red green blue"`
		Unique             []string `validate:"unique"`
		Inner              Inner
	}

	// init test struct with invalid values
	test := TestOtherTags{
		Dir:     "nonexistent",
		DirPath: "invalid/dir/path",
		File:    "nonexistent.txt",
		Image:   "not-an-image.txt",

		LenString: "toolong",               // should be exactly 5 chars
		LenSlice:  []string{"a", "b"},      // should be exactly 3 items
		LenNumber: 5,                       // should be 10
		MinString: "ab",                    // should be min 3 chars
		MinSlice:  []string{},              // should be min 1 item
		MaxString: "toolong",               // should be max 5 chars
		MaxSlice:  []string{"satu", "dua"}, // should be max 1 item

		IsDefault: "non-default",

		ExcludedIf:         "value", // should fail when Inner.RequiredWith is "value"
		ExcludedUnless:     "value", // should fail unless Inner.ExcludedWith is "value"
		ExcludedWith:       "value", // should fail when Inner.Field is populated
		ExcludedWithAll:    "value", // should fail when both Inner.Field and Inner.ExcludedWith are populated
		ExcludedWithout:    "value", // should fail when Inner.ExcludedWith is not populated
		ExcludedWithoutAll: "value", // should fail when Inner.ExcludedWithoutAll is not populated

		OneOf:  "yellow",           // not in [red green blue]
		Unique: []string{"a", "a"}, // contains duplicate

		Inner: Inner{
			RequiredWith: "value", // triggers required_if validation
			ExcludedWith: "",      // triggers excluded_unless validation
		},
	}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestOtherTags.Dir",
			expected: "Dir harus berupa direktori yang ada",
		},
		{
			ns:       "TestOtherTags.DirPath",
			expected: "DirPath harus berupa path direktori yang valid",
		},
		{
			ns:       "TestOtherTags.File",
			expected: "File harus berupa file yang valid",
		},
		{
			ns:       "TestOtherTags.FilePath",
			expected: "FilePath harus berupa path file yang valid",
		},
		{
			ns:       "TestOtherTags.Image",
			expected: "Image harus berupa gambar yang valid",
		},
		{
			ns:       "TestOtherTags.LenString",
			expected: "panjang LenString harus 5 karakter",
		},
		{
			ns:       "TestOtherTags.LenSlice",
			expected: "LenSlice harus berisi 3 item",
		},
		{
			ns:       "TestOtherTags.LenNumber",
			expected: "LenNumber harus sama dengan 10",
		},
		{
			ns:       "TestOtherTags.MinString",
			expected: "panjang minimal MinString adalah 3 karakter",
		},
		{
			ns:       "TestOtherTags.MinSlice",
			expected: "MinSlice harus berisi minimal 1 item",
		},
		{
			ns:       "TestOtherTags.MaxString",
			expected: "panjang maksimal MaxString adalah 5 karakter",
		},
		{
			ns:       "TestOtherTags.MaxSlice",
			expected: "MaxSlice harus berisi maksimal 1 item",
		},
		{
			ns:       "TestOtherTags.IsDefault",
			expected: "IsDefault harus berupa nilai default",
		},
		{
			ns:       "TestOtherTags.Required",
			expected: "Required wajib diisi",
		},
		{
			ns:       "TestOtherTags.RequiredIf",
			expected: "RequiredIf wajib diisi jika Inner.RequiredWith value",
		},
		{
			ns:       "TestOtherTags.RequiredUnless",
			expected: "RequiredUnless wajib diisi kecuali Inner.RequiredWith values",
		},
		{
			ns:       "TestOtherTags.RequiredWith",
			expected: "RequiredWith wajib diisi jika Inner.RequiredWith telah diisi",
		},
		{
			ns:       "TestOtherTags.RequiredWithAll",
			expected: "RequiredWithAll wajib diisi jika Inner.RequiredWith telah diisi",
		},
		{
			ns:       "TestOtherTags.RequiredWithout",
			expected: "RequiredWithout wajib diisi jika Inner.ExcludedWith tidak diisi",
		},
		{
			ns:       "TestOtherTags.RequiredWithoutAll",
			expected: "RequiredWithoutAll wajib diisi jika Inner.ExcludedWith tidak diisi",
		},
		{
			ns:       "TestOtherTags.ExcludedIf",
			expected: "ExcludedIf tidak boleh diisi jika Inner.RequiredWith value",
		},
		{
			ns:       "TestOtherTags.ExcludedUnless",
			expected: "ExcludedUnless tidak boleh diisi kecuali Inner.ExcludedWith value",
		},
		{
			ns:       "TestOtherTags.ExcludedWith",
			expected: "ExcludedWith tidak boleh diisi jika Inner.RequiredWith telah diisi",
		},
		{
			ns:       "TestOtherTags.ExcludedWithAll",
			expected: "ExcludedWithAll tidak boleh diisi jika semua Inner.RequiredWith telah diisi",
		},
		{
			ns:       "TestOtherTags.ExcludedWithout",
			expected: "ExcludedWithout tidak boleh diisi jika Inner.ExcludedWith tidak diisi",
		},
		{
			ns:       "TestOtherTags.ExcludedWithoutAll",
			expected: "ExcludedWithoutAll tidak boleh diisi jika Inner.ExcludedWith tidak diisi",
		},
		{
			ns:       "TestOtherTags.OneOf",
			expected: "OneOf harus berupa salah satu dari [red green blue]",
		},
		{
			ns:       "TestOtherTags.Unique",
			expected: "Unique harus berisi nilai yang unik",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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

// TestAliasesTagsTranslations tests all aliases tags registered translations for Indonesian language
func TestAliasesTagsTranslations(t *testing.T) {

	// init validator with Indonesian translations
	validate, trans, err := InitValidator()
	Equal(t, err, nil)

	// TestAliasTags for alias validations
	type TestAliasTags struct {
		// Color validations
		Color     string `validate:"iscolor"`
		HexColor  string `validate:"hexcolor"`
		RGBColor  string `validate:"rgb"`
		RGBAColor string `validate:"rgba"`
		HSLColor  string `validate:"hsl"`
		HSLAColor string `validate:"hsla"`

		// Country code validation
		CountryCode string `validate:"country_code"`
		ISO2Code    string `validate:"iso3166_1_alpha2"`
		ISO3Code    string `validate:"iso3166_1_alpha3"`
	}

	// init test struct with invalid values
	test := TestAliasTags{
		Color:     "not-a-color",
		HexColor:  "not-hex",
		RGBColor:  "not-rgb",
		RGBAColor: "not-rgba",
		HSLColor:  "not-hsl",
		HSLAColor: "not-hsla",

		CountryCode: "XX",  // invalid country code
		ISO2Code:    "XX",  // invalid ISO 3166-1 alpha-2
		ISO3Code:    "XXX", // invalid ISO 3166-1 alpha-3
	}

	// validate struct
	err = validate.Struct(test)
	NotEqual(t, err, nil)

	// get validation errors
	errs := err.(validator.ValidationErrors)

	// verify each expected error message
	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "TestAliasTags.Color",
			expected: "Color harus berupa warna yang valid",
		},
		{
			ns:       "TestAliasTags.HexColor",
			expected: "HexColor harus berupa warna HEX yang valid",
		},
		{
			ns:       "TestAliasTags.RGBColor",
			expected: "RGBColor harus berupa warna RGB yang valid",
		},
		{
			ns:       "TestAliasTags.RGBAColor",
			expected: "RGBAColor harus berupa warna RGBA yang valid",
		},
		{
			ns:       "TestAliasTags.HSLColor",
			expected: "HSLColor harus berupa warna HSL yang valid",
		},
		{
			ns:       "TestAliasTags.HSLAColor",
			expected: "HSLAColor harus berupa warna HSLA yang valid",
		},
		{
			ns:       "TestAliasTags.CountryCode",
			expected: "CountryCode harus berupa kode negara yang valid",
		},
	}

	// verify each expected error message
	for _, tt := range tests {
		var fe validator.FieldError

		// find matching field error
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
