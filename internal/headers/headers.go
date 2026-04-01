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

func (h Headers) Parse(data []byte) (int, bool, error) {
	idx := bytes.Index(data, endLine)
	if idx == -1 {
		return 0, false, badfdline

	}
	if idx == 0 {
		return idx, true, nil
	}
	head  :=bytes.SplitN(data, []byte(":"), 2)
	if len(head) != 2 {
		return 0, false, badfdline
	}
	// Header = head[0]: head[1] \r\n \r\n
	//fieldName := head[0]
	//fieldValue := bytes.TrimSpace(head[1])

	return 0, false, nil


}







