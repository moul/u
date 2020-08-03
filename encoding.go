package u

import "encoding/base64"

// B64Encode returns a base64 encoded string of input bytes
func B64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// B64Decode try to decode an input string and returns bytes if success
func B64Decode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
