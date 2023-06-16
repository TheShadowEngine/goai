package goai

import "github.com/julienschmidt/httprouter"

type Operatable interface {
	Routes() map[string]map[string]httprouter.Handle
	Templates() map[string][]byte
}
