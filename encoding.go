package u

import (
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
