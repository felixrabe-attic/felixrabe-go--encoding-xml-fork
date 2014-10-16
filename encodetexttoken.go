package xml

import (
	"io"
	"unicode/utf8"
)

// encodeTextToken writes to w the properly escaped XML equivalent of the
// plain text data s.
//
// This works like EscapeText, but does not escape whitespace.
func encodeTextToken(w io.Writer, s []byte) error {
	var esc []byte
	last := 0
	for i := 0; i < len(s); {
		r, width := utf8.DecodeRune(s[i:])
		i += width
		switch r {
		case '"':
			esc = esc_quot
		case '\'':
			esc = esc_apos
		case '&':
			esc = esc_amp
		case '<':
			esc = esc_lt
		case '>':
			esc = esc_gt
		// case '\t':
		// 	esc = esc_tab
		// case '\n':
		// 	esc = esc_nl
		// case '\r':
		// 	esc = esc_cr
		default:
			if !isInCharacterRange(r) || (r == 0xFFFD && width == 1) {
				esc = esc_fffd
				break
			}
			continue
		}
		if _, err := w.Write(s[last : i-width]); err != nil {
			return err
		}
		if _, err := w.Write(esc); err != nil {
			return err
		}
		last = i
	}
	if _, err := w.Write(s[last:]); err != nil {
		return err
	}
	return nil
}
