package headers

import (
	"bytes"
	"fmt"
)

// currently the headers are stored in a map.
type Headers map[string]string

var endLine = []byte("\r\n")
var badfdline = fmt.Errorf("malformed field-line!")
var badfdname = fmt.Errorf("malformed field-name!")

// parses the headers.
// given the RFC's definition of a header:
// header = field-name ":" OWS field-value OWS
// field-name = token
// field-value = *( field-content | LWS )
// field-content = <the OCTETs making up the field-value
//                and consisting of either *TEXT or combinations
//                of token, separators, and quoted-string>
// token = 1*<any CHAR except CTLs or separators>
// separators = LWS | ":" | "<" | ">" | "@"
//           | "," | ";" | "\" | <">
//           | "/" | "[" | "]" | "?" | "="
//           | "{" | "}" | SP | HT
// quoted-string = <"> *(qdtext | quoted-pair ) <">
// qdtext = <any TEXT except <">>
// quoted-pair = "\" CHAR
// OWS = < SP | HT >
// CTL = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
func ParseHeaders(data []byte) (string, string, error) {
	head:=bytes.SplitN(data, []byte(":"), 2)
	if len(head) != 2 {

		return "", "", badfdline
	}
	if bytes.HasSuffix(head[0], []byte(" ")) {
		return "", "", badfdname
	}
	// Header = head[0]: head[1] \r\n \r\n
	fieldName := head[0]
	fieldValue := bytes.TrimSpace(head[1])
	return string(fieldName), string(fieldValue), nil
	


	
}

// creates a new headers map.
func NewHeaders() Headers {
	return map[string]string{}
}

//receives a given data and will parse it until the end of the headers.
// if the headers are not finished, it will return the number of bytes parsed and a boolean indicating if the headers are finished.
func (h Headers) Parse(data []byte) (int, bool, error) {
	n := 0
	done := false
free :
	for {
		idx := bytes.Index(data[n:], endLine)
		if idx == -1 {
			break free
		}
		if idx == 0 {
			done = true
			n += len(endLine)
			break free
		}
		name, value, err := ParseHeaders(data[n:+idx+n])
		if err != nil {
			return 0, false, err 
		}
		n += idx + len(endLine)
		h[name] = value
	}
	return n, done, nil


}







