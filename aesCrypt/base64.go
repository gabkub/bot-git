package aesCrypt

import "encoding/base64"

func decodeBytes(encoded string) ([]byte, error) {
	v, e := base64.StdEncoding.DecodeString(encoded)
	if e != nil {
		return nil, e
	}
	return v, nil
}

func encodeStringFromBytes(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

