package util

import (
	"github.com/ac-kurniawan/mychat/core"
	"github.com/ac-kurniawan/mychat/library"
)

type Util struct {
	library.AppTrace
	library.AppLog
}

func NewUtil(module Util) core.IUtil {
	return &module
}
