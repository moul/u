package u

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

// B64Encode returns a base64 encoded string of input bytes.
func B64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// B64Decode try to decode an input string and returns bytes if success.
func B64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

// IsASCII checks whether a buffer only contains ASCII characters.
func IsASCII(buf []byte) bool {
	for _, b := range buf {
		if b > 0x7F {
			return false
		}
	}
	return true
}

// JSON returns a JSON representation of the passed input.
func JSON(input interface{}) string {
	out, _ := json.Marshal(input)
	return string(out)
}

// PrettyJSON returns an indented JSON representation of the passed input.
func PrettyJSON(input interface{}) string {
	out, _ := json.MarshalIndent(input, "", "  ")
	return string(out)
}

// IsBinary returns whether the provided buffer looks like binary or human-readable.
//
// It is inspired by the implementation made in the Git project.
// https://github.com/git/git/blob/49f38e2de47a401fc2b0f4cce38e9f07fb63df48/xdiff-interface.c#L188.
func IsBinary(buf []byte) bool {
	const prefixLen = 8000
	if len(buf) > prefixLen {
		buf = buf[0:prefixLen]
	}
	pos := bytes.IndexByte(buf, byte(0))
	return pos != -1
}
