package headers

import "bytes"

type Headers struct {
	header map[string]string
}


func (h Headers) Parse(data []byte) (int, bool, error) {
	head  :=bytes.SplitN(data, []byte(":"), 2)
	// Header = head[0] : head[1] \r\n \r\n
	value := bytes.TrimSpace(head[1])

}









