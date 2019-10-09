package metcli

import (
	"encoding/base64"
	"strings"
)

// EncodePayload will base64 the payload and prepend/append magic chars
func EncodePayload(data string, mode string, aid string, magicstr string, magictermstr string) string {
	preEnc := mode + "||" + aid + "||" + data
	encStr := base64.StdEncoding.EncodeToString([]byte(preEnc))
	fin := magicstr + encStr + magictermstr
	return fin
}

//DecodePayload will decode from MAD into plaintext string
func DecodePayload(payload string, magicstr string, magictermstr string) string {
	encodedPayload := strings.Replace(payload, magicstr, "", -1) //trim magic chars from payload
	encodedPayload = strings.Replace(encodedPayload, magictermstr, "", -1)
	data, err := base64.StdEncoding.DecodeString(encodedPayload)
	if err != nil {
		return ""
	}
	return string(data)
}
