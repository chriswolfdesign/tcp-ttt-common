package tcp_ttt_common_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTcpTttCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TcpTttCommon Suite")
}
