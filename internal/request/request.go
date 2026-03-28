package request

import (
	"bytes"
	"fmt"
	"io"
)

var BadReqLine 			= fmt.Errorf("Malformed Request-Line!")
var BadHttpVer    	= fmt.Errorf("Unsupported Http Version!")
var BadParseState 	= fmt.Errorf("Unable to Continue! Parsing Terminated!")
var EndLine 				= []byte("\r\n")

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type ParserState string 
const (
   StateInit  ParserState = "init"
	 StateDone  ParserState = "done"
	 StateError ParserState = "error"
)

type Request struct {
	RequestLine RequestLine
	State ParserState
}

func (r *Request) done() bool {
	return r.State == StateDone 
}

func (r *Request) err() bool {
  return r.State == StateError
}

func newRequest() *Request {
 return &Request{
	 State : StateInit,
 }
}


func (r *Request) parse(data []byte) (int, error) {
	idx := 0
	dance :
		for {
			switch (r.State) {
				case StateInit :
					rl, n, err := parseRequestLine(data[idx:])
					if err != nil {
						r.State = StateError
						return 0, err
					}
					if n == 0 {
						break dance
					}
					idx += n
					r.State = StateDone
					r.RequestLine = *rl
				case StateDone :
					break dance

				case StateError :
					return 0, BadParseState


			}

		}
	return idx, nil
} 

func parseRequestLine(data []byte) (*RequestLine, int, error){
	idx := bytes.Index(data, EndLine)
	if idx == -1 {
		return nil, 0, nil
	}
	reqLine := bytes.Split(data[:idx], []byte(" "))
	if len(reqLine) != 3 {
    return nil, idx, BadReqLine
	}
	httpVer := bytes.Split(reqLine[2], []byte("/"))
	if len(httpVer)!= 2 || string(httpVer[0]) != "HTTP" || string(httpVer[1]) != "1.1" {
	  return nil, idx, BadHttpVer	
	}
	httpRequestLine := &RequestLine {
		Method: string(reqLine[0]),
		RequestTarget: string(reqLine[1]),
		HttpVersion: string(httpVer[1]),
	}

  return httpRequestLine, idx+len(EndLine), nil
}


func RequestFromReader(reader io.Reader) (*Request, error) {
	rl := newRequest()
	buf := make([]byte, 1024)
	buflen := 0

	for !rl.done() && !rl.err() {
		data, err := reader.Read(buf[buflen:])
		if err != nil {
			//rl.State = StateError 
			return nil, err
		}
		buflen += data
		n, err := rl.parse(buf[:buflen])
		if err != nil {
			return nil, err
		}
		copy(buf, buf[n:buflen])
    buflen -= n
	}

	return rl, nil

}




