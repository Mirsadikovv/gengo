package sharedutil

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type validate = validator.Validate

type Validator struct {
	*validate
}

func NewValidator() *Validator {

	validate := validator.New()
	validate.RegisterValidation("iso639_2_alpha3", Iso6392Alpha3)

	return &Validator{
		validate: validate,
	}
}

func (cv *Validator) Validate(t any) error {
	return cv.Struct(t)
}

//	Validator Langs

var validLanguages = map[string]bool{
	"eng": true, // English
	"fra": true, // French
	"spa": true, // Spanish
	"deu": true, // German
	"rus": true, // Russian
	"ita": true, // Italian
	"zho": true, // Chinese
	"jpn": true, // Japanese
	"ara": true, // Arabic
	"por": true, // Portuguese
	"hin": true, // Hindi
	"ben": true, // Bengali
	"urd": true, // Urdu
	"kor": true, // Korean
	"tur": true, // Turkish
	"pol": true, // Polish
	"nld": true, // Dutch
	"dan": true, // Danish
	"nor": true, // Norwegian
	"fin": true, // Finnish
	"swe": true, // Swedish
	"ukr": true, // Ukrainian
	"uzb": true, // Uzbek
	"kaz": true, // Kazakh
	"tam": true, // Tamil
	"tel": true, // Telugu
	"mal": true, // Malayalam
	"vie": true, // Vietnamese
	"tha": true, // Thai
	"gre": true, // Greek
	"heb": true, // Hebrew
	"cze": true, // Czech
	"hun": true, // Hungarian
	"rom": true, // Romanian
	"slk": true, // Slovak
	"slv": true, // Slovenian
	"srp": true, // Serbian
	"hrv": true, // Croatian
	"bul": true, // Bulgarian
	"lit": true, // Lithuanian
	"lav": true, // Latvian
	"est": true, // Estonian
	"aze": true, // Azerbaijani
	"arm": true, // Armenian
	"fas": true, // Persian
	"bos": true, // Bosnian
	"cat": true, // Catalan
	"ind": true, // Indonesian
	"msa": true, // Malay
	"fil": true, // Filipino
	"isl": true, // Icelandic
	"gle": true, // Irish
	"tgl": true, // Tagalog
	"sqi": true, // Albanian
	"bel": true, // Belarusian
	"mkd": true, // Macedonian
	"ces": true, // Czech (alternative code for Cze)
	"tir": true, // Tigrinya
	"amh": true, // Amharic
	"som": true, // Somali
	"kin": true, // Kinyarwanda
	"lug": true, // Luganda
	"swa": true, // Swahili
	"afr": true, // Afrikaans
	"zul": true, // Zulu
	"xho": true, // Xhosa
	"yor": true, // Yoruba
	"ibo": true, // Igbo
	"hau": true, // Hausa
}

func Iso6392Alpha3(fl validator.FieldLevel) bool {
	code := fl.Field().String()
	_, exists := validLanguages[strings.ToLower(code)]
	return exists
}
