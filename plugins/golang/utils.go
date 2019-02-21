package sbgo

import (
	"bytes"
	"fmt"
	"strings"
)

// comment lines prefixes each line in lines with "// ".
func commentLines(lines string) string {
	lines = "// " + strings.Replace(lines, "\n", "\n// ", -1)
	return lines
}

// FprintZipData converts zip binary contents to a string literal.
func FprintZipData(dest *bytes.Buffer, zipData []byte) {
	for _, b := range zipData {
		if b == '\n' {
			dest.WriteString(`\n`)
			continue
		}
		if b == '\\' {
			dest.WriteString(`\\`)
			continue
		}
		if b == '"' {
			dest.WriteString(`\"`)
			continue
		}
		if (b >= 32 && b <= 126) || b == '\t' {
			dest.WriteByte(b)
			continue
		}
		fmt.Fprintf(dest, "\\x%02x", b)
	}
}
