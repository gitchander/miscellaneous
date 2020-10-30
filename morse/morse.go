package morse

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// short and long signals called "dots" and "dashes"
// samples: {'•', '—'}

const (
	dash = iota
	dot
)

var charMap = charMapEnglish

var charMapEnglish = map[rune][]int{
	'a':  {dot, dash},
	'b':  {dash, dot, dot, dot},
	'c':  {dash, dot, dash, dot},
	'd':  {dash, dot, dot},
	'e':  {dot},
	'f':  {dot, dot, dash, dot},
	'g':  {dash, dash, dot},
	'h':  {dot, dot, dot, dot},
	'i':  {dot, dot},
	'j':  {dot, dash, dash, dash},
	'k':  {dash, dot, dash},
	'l':  {dot, dash, dot, dot},
	'm':  {dash, dash},
	'n':  {dash, dot},
	'o':  {dash, dash, dash},
	'p':  {dot, dash, dash, dot},
	'q':  {dash, dash, dot, dash},
	'r':  {dot, dash, dot},
	's':  {dot, dot, dot},
	't':  {dash},
	'u':  {dot, dot, dash},
	'v':  {dot, dot, dot, dash},
	'w':  {dot, dash, dash},
	'x':  {dash, dot, dot, dash},
	'y':  {dash, dot, dash, dash},
	'z':  {dash, dash, dot, dot},
	'1':  {dot, dash, dash, dash, dash},
	'2':  {dot, dot, dash, dash, dash},
	'3':  {dot, dot, dot, dash, dash},
	'4':  {dot, dot, dot, dot, dash},
	'5':  {dot, dot, dot, dot, dot},
	'6':  {dash, dot, dot, dot, dot},
	'7':  {dash, dash, dot, dot, dot},
	'8':  {dash, dash, dash, dot, dot},
	'9':  {dash, dash, dash, dash, dot},
	'0':  {dash, dash, dash, dash, dash},
	'.':  {dot, dash, dot, dash, dot, dash},
	',':  {dash, dash, dot, dot, dash, dash},
	':':  {dash, dash, dash, dot, dot, dot},
	'?':  {dot, dot, dash, dash, dot, dot},
	'\'': {dot, dash, dash, dash, dash, dot},
	'-':  {dash, dot, dot, dot, dot, dash},
	'/':  {dash, dot, dot, dash, dot},
	'(':  {dash, dot, dash, dash, dot},
	')':  {dash, dot, dash, dash, dot, dash},
	'+':  {dot, dash, dot, dash, dot},
	'×':  {dash, dot, dot, dash},
	'@':  {dot, dash, dash, dot, dash, dot},
}

var charMapRussian = map[rune][]int{
	'а':  {dot, dash},
	'б':  {dash, dot, dot, dot},
	'в':  {dot, dash, dash},
	'г':  {dash, dash, dot},
	'д':  {dash, dot, dot},
	'е':  {dot},
	'ё':  {dot},
	'ж':  {dot, dot, dot, dash},
	'з':  {dash, dash, dot, dot},
	'и':  {dot, dot},
	'й':  {dot, dash, dash, dash},
	'к':  {dash, dot, dash},
	'л':  {dot, dash, dot, dot},
	'м':  {dash, dash},
	'н':  {dash, dot},
	'о':  {dash, dash, dash},
	'п':  {dot, dash, dash, dot},
	'р':  {dot, dash, dot},
	'с':  {dot, dot, dot},
	'т':  {dash},
	'у':  {dot, dot, dash},
	'ф':  {dot, dot, dash, dot},
	'х':  {dot, dot, dot, dot},
	'ц':  {dash, dot, dash, dot},
	'ч':  {dash, dash, dash, dot},
	'ш':  {dash, dash, dash, dash},
	'щ':  {dash, dash, dot, dash},
	'ъ':  {dash, dash, dot, dash, dash},
	'ы':  {dash, dot, dash, dash},
	'ь':  {dash, dot, dot, dash},
	'э':  {dot, dot, dash, dot, dot},
	'ю':  {dot, dot, dash, dash},
	'я':  {dot, dash, dot, dash},
	'0':  {dash, dash, dash, dash, dash},
	'1':  {dot, dash, dash, dash, dash},
	'2':  {dot, dot, dash, dash, dash},
	'3':  {dot, dot, dot, dash, dash},
	'4':  {dot, dot, dot, dot, dash},
	'5':  {dot, dot, dot, dot, dot},
	'6':  {dash, dot, dot, dot, dot},
	'7':  {dash, dash, dot, dot, dot},
	'8':  {dash, dash, dash, dot, dot},
	'9':  {dash, dash, dash, dash, dot},
	'.':  {dot, dot, dot, dot, dot, dot},
	',':  {dot, dash, dot, dash, dot, dash},
	':':  {dash, dash, dash, dot, dot, dot},
	';':  {dash, dot, dash, dot, dash},
	'(':  {dash, dot, dash, dash, dot, dash},
	')':  {dash, dot, dash, dash, dot, dash},
	'\'': {dot, dash, dash, dash, dash, dot},
	'"':  {dot, dash, dot, dot, dash, dot},
	'—':  {dash, dot, dot, dot, dot, dash},
	'/':  {dash, dot, dot, dash, dot},
	'?':  {dot, dot, dash, dash, dot, dot},
	'!':  {dash, dash, dot, dot, dash, dash},
	'@':  {dot, dash, dash, dot, dash, dot},
}

// Durations
const (
	durDash = 3
	durDot  = 1

	durSymbolSpace = 1
	durLetterSpace = 3
	durWordSpace   = 7
)

const (
	unitSignalON  = '='
	unitSignalOFF = '.'
)

func writeSignal(n int, b byte, buffer *bytes.Buffer) {
	for i := 0; i < n; i++ {
		buffer.WriteByte(b)
	}
}

func writeWord(word string, buffer *bytes.Buffer) {

	data := []byte(word)

	firstLetter := true

	for {
		r, size := utf8.DecodeRune(data)
		if size == 0 {
			break
		}
		data = data[size:]

		ps, ok := charMap[unicode.ToLower(r)]
		if !ok {
			continue
		}

		if firstLetter {
			firstLetter = false
		} else {
			writeSignal(durLetterSpace, unitSignalOFF, buffer)
		}

		for i, p := range ps {
			if i > 0 {
				writeSignal(durSymbolSpace, unitSignalOFF, buffer)
			}
			switch p {
			case dash:
				writeSignal(durDash, unitSignalON, buffer)
			case dot:
				writeSignal(durDot, unitSignalON, buffer)
			}
		}
	}
}

func Units(s string) ([]byte, error) {

	var buffer bytes.Buffer

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	firstWord := true

	for scanner.Scan() {
		word := scanner.Text()
		if firstWord {
			firstWord = false
		} else {
			writeSignal(durWordSpace, unitSignalOFF, &buffer)
		}
		writeWord(word, &buffer)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
