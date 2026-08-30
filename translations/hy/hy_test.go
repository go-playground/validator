package hy

import (
	"testing"

	hyLocale "github.com/go-playground/locales/hy"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func newArmenianTranslator(t *testing.T) (*validator.Validate, ut.Translator) {
	t.Helper()

	locale := hyLocale.New()
	universalTranslator := ut.New(locale, locale)
	translator, found := universalTranslator.GetTranslator("hy")
	if !found {
		t.Fatal("armenian translator was not found")
	}

	validate := validator.New()
	if err := RegisterDefaultTranslations(validate, translator); err != nil {
		t.Fatalf("failed to register armenian translations: %v", err)
	}

	return validate, translator
}

func firstTranslatedError(t *testing.T, validate *validator.Validate, translator ut.Translator, value any) string {
	t.Helper()

	err := validate.Struct(value)
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok || len(validationErrors) == 0 {
		t.Fatalf("expected validator.ValidationErrors, got %T: %v", err, err)
	}

	return validationErrors[0].Translate(translator)
}

func TestRegisterDefaultTranslations(t *testing.T) {
	newArmenianTranslator(t)
}

func TestBasicTranslations(t *testing.T) {
	validate, translator := newArmenianTranslator(t)

	tests := []struct {
		name     string
		value    any
		expected string
	}{
		{
			name: "required",
			value: struct {
				Email string `validate:"required"`
			}{},
			expected: "«Email» դաշտը պարտադիր է",
		},
		{
			name: "email",
			value: struct {
				Email string `validate:"email"`
			}{Email: "invalid"},
			expected: "«Email» դաշտը պետք է պարունակի վավեր էլեկտրոնային փոստի հասցե",
		},
		{
			name: "minimum string length",
			value: struct {
				Code string `validate:"min=2"`
			}{Code: "A"},
			expected: "«Code» դաշտը պետք է պարունակի առնվազն 2 նիշ",
		},
		{
			name: "exact collection length",
			value: struct {
				Items []string `validate:"len=2"`
			}{Items: []string{"one"}},
			expected: "«Items» դաշտը պետք է պարունակի 2 տարր",
		},
		{
			name: "maximum number",
			value: struct {
				Amount int `validate:"max=10"`
			}{Amount: 11},
			expected: "«Amount» դաշտի արժեքը չպետք է գերազանցի 10",
		},
		{
			name: "equal fields",
			value: struct {
				Password             string
				PasswordConfirmation string `validate:"eqfield=Password"`
			}{Password: "first", PasswordConfirmation: "second"},
			expected: "«PasswordConfirmation» և «Password» դաշտերի արժեքները պետք է համընկնեն",
		},
		{
			name: "one of",
			value: struct {
				Status string `validate:"oneof=draft published"`
			}{Status: "deleted"},
			expected: "«Status» դաշտի արժեքը պետք է լինի հետևյալներից մեկը՝ [draft published]",
		},
		{
			name: "date time format",
			value: struct {
				Date string `validate:"datetime=2006-01-02"`
			}{Date: "23/07/2026"},
			expected: "«Date» դաշտը չի համապատասխանում 2006-01-02 ձևաչափին",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := firstTranslatedError(t, validate, translator, test.value)
			if actual != test.expected {
				t.Fatalf("unexpected translation:\nwant: %q\n got: %q", test.expected, actual)
			}
		})
	}
}

func TestArmenianCardinalFormsAreRegistered(t *testing.T) {
	_, translator := newArmenianTranslator(t)

	tests := []struct {
		name   string
		key    string
		number float64
		digits uint64
		value  string
		want   string
	}{
		{name: "zero characters", key: "len-string-character", number: 0, value: "0", want: "0 նիշ"},
		{name: "one character", key: "len-string-character", number: 1, value: "1", want: "1 նիշ"},
		{name: "two characters", key: "len-string-character", number: 2, value: "2", want: "2 նիշ"},
		{name: "one item", key: "len-items-item", number: 1, value: "1", want: "1 տարր"},
		{name: "two items", key: "len-items-item", number: 2, value: "2", want: "2 տարր"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := translator.C(test.key, test.number, test.digits, test.value)
			if err != nil {
				t.Fatalf("failed to translate cardinal form: %v", err)
			}
			if actual != test.want {
				t.Fatalf("unexpected cardinal translation:\nwant: %q\n got: %q", test.want, actual)
			}
		})
	}
}
