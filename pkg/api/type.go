package api

import (
	"sync"
)

var (
	once sync.Once
)

const (
	DefaultUserAPIFlag Type = 1
)

type (
	Type int
)
