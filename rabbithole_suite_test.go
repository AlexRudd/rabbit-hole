package rabbithole

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRabbitHole(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rabbithole Suite")
}
