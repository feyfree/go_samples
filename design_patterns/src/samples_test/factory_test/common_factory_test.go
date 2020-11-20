package factory

import "testing"
import "../../samples/factory/common"

func TestCommonFactory(t *testing.T) {
	factory := &common.FileLoggerFactory{}
	logger := factory.CreateLogger()
	logger.WriteLog()
}
