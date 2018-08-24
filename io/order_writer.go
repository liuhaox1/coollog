package io

import "github.com/cool/coollog/common"

type exefunc func(str string)
type operation struct {
	execute map[common.Token]exefunc
}
