package headers

import (
	"bytes"
	"fmt"
)

//type Headers struct {
	type Headers map[string]string
//}

var endLine = []byte("\r\n")
var badfdline = fmt.Errorf("malformed field-line!")
var badfdname = fmt.Errorf("malformed field-name!")

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

func NewHeaders() Headers {
	return map[string]string{}
}

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







