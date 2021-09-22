package igslim

import "testing"

const originalUserName = "original"

var testcases = map[string]string{
	"Follow Me On IG - @example":                  "example",
	"Follow the IG: EXAMPLE_01":                   "EXAMPLE_01",
	"follow my insta >username_ business":         "username_",
	"INSTA=example__official example@example.com": "example__official",
	"IG: @foo_bar foo@bar.com":                    "foo_bar",
	"Instagram: loremipsum YouTube:LoremIpsum":    "loremipsum",
	"Insta:Some_name":                             "Some_name",
	"IG: hello_world yes":                         "hello_world",
	"Follow our Twitter @ourtwitter":              "",
	"Follow us on Instagram and YouTube":          originalUserName,
	"Insta/YouTube: James002":                     "James002",
	"Insta / YouTube: foobar123":                  "foobar123",
	"Instagram:@example email: example@gmail.com": "example",
	"Get my IG to 100K @foobar!":                  "foobar",
	"Insta • Youtube • Blog":                      originalUserName,
	"alright":                                     "",
	"BIG":                                         "",
	"IG":                                          originalUserName,
	"Insta ↓ Thank you":                           originalUserName,
	"Follow me on Instagram: example@gmail.com":   originalUserName,
	"Follow me on Instagram: example@":            originalUserName,
	"Follow me on Instagram: example003":          "example003",
	"Why not follow helloworld on Instagram?":     "helloworld",
	"Why not follow @example on IG?":              "example",
	"leaked face on IG":                           originalUserName,
	"face reveal on instagram":                    originalUserName,
	"Why not follow foo@example.com on IG?":       originalUserName,
	"helloworld ON IG":                            "helloworld",
	"insta👉hello_world":                           "hello_world",
	"insta -15%":                                  originalUserName,
}

func TestGetUserNameFromText(t *testing.T) {
	for text, expected := range testcases {
		actual := GetUserNameFromText(text, originalUserName)
		if actual != expected {
			t.Errorf("GetUserNameFromText(%s) should return %s instead of %s", text, expected, actual)
		}
	}
}
