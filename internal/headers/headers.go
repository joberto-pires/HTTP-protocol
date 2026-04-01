package headers

import (
	"bytes"
	"fmt"
)

type Headers struct {
	header map[string]string
}

var endLine = []byte("\r\n")
var badfdline = fmt.Errorf("malformed field-line!")
var badfdname = fmt.Errorf("malformed field-name!")

func ParseHeaders(data []byte) (string, string, error) {
	head  :=bytes.SplitN(data, []byte(":"), 2)
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

func newHeaders() {
}

func (h Headers) Parse(data []byte) (int, bool, error) {
	idx := bytes.Index(data, endLine)
	if idx == -1 {
		return 0, false, badfdline

	}
	if idx == 0 {
		return idx, true, nil
	}
	name, value, err := ParseHeaders(data[:idx])
	if err != nil {
		return 0, false, badfdline
	}
	n := idx + len(endLine)
	return n, false, nil


}







