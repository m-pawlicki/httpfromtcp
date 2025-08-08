package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	r, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	rl, err := parseRequestLine(string(r))
	if err != nil {
		return nil, err
	}

	req := &Request{
		RequestLine: *rl,
	}
	return req, nil
}

func parseRequestLine(s string) (*RequestLine, error) {
	split := strings.Split(s, "\r\n")
	requestLine := strings.Split(split[0], " ")
	rl := &RequestLine{}

	if len(requestLine) != 3 {
		return nil, errors.New("not enough parameters")
	}

	if !isUpper(requestLine[0]) {
		return nil, errors.New("improper method")
	}
	rl.Method = requestLine[0]

	rl.RequestTarget = requestLine[1]

	version := strings.Split(requestLine[2], "/")
	if version[1] != "1.1" {
		return nil, errors.New("wrong version")
	}
	rl.HttpVersion = version[1]

	return rl, nil
}

func isUpper(s string) bool {
	for _, s := range s {
		if s < 'A' || s > 'Z' {
			return false
		}
	}
	return true
}
