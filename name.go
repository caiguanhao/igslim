package igslim

import (
	"regexp"
	"strings"
)

const (
	reUserNameStr = `[A-Za-z0-9_][A-Za-z0-9_.]{0,28}[A-Za-z0-9_]`
)

var (
	reIg         = regexp.MustCompile(`(?i)(?:^|[^\w])(ig|instagram|insta|ins)(?:[^\w]|$)`)
	reUserNameOn = regexp.MustCompile(`(?i)(^|[^\w])(@?)(` + reUserNameStr + `)(\s+on\s+)$`)
	reUserNameAt = regexp.MustCompile(`(?:^|[^\w])@(` + reUserNameStr + `)`)
	reUserName   = regexp.MustCompile(reUserNameStr)
)

// GetUserNameFromText extracts Instagram user name from "text" (usually the
// biography of a social media account) using regular expressions. Because of
// too many edge cases, this function does not guarantee the accuracy of the
// extracted user names. The "username" argument is returned if Instagram is
// mentioned but no user name is found. Empty string is returned if no user
// name is found.
func GetUserNameFromText(text, username string) string {
	loc := reIg.FindStringSubmatchIndex(text)
	if len(loc) < 4 {
		return ""
	}

	// follow * on Instagram
	if m := reUserNameOn.FindStringSubmatch(text[:loc[2]]); len(m) > 0 {
		// CAN'T DECIDE: this is very likely an English word instead of
		// username if it does not start with @ and its length is less than 8
		if m[2] == "@" || len(m[3]) >= 8 {
			if m[1] != "@" {
				return m[3]
			}
		}
	}

	text = text[loc[3]:]

	var hasColon bool
	if i := strings.IndexAny(text, ":-=：>👉"); i > -1 {
		text = text[i:]
		hasColon = true
	}

	if m := reUserNameAt.FindStringSubmatch(text); len(m) > 0 {
		return m[1]
	}

	if hasColon {
		loc := reUserName.FindStringIndex(text)
		if len(loc) == 2 {
			if loc[1]+1 <= len(text) && text[loc[1]:loc[1]+1] == "@" {
				// more like an email address
				return username
			}
			// CAN'T DECIDE: this is very likely an English word instead of
			// username if it does not start with @ and its length is less than 8
			if name := text[loc[0]:loc[1]]; len(name) >= 8 {
				return name
			}
		}
	}

	return username
}
