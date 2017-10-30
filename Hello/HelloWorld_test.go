package Hello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloworld(t *testing.T) {
	total := hello()
	fmt.Println("total", total)
	assert.Equal(t, 10, total, "Should be eqaul")
}
