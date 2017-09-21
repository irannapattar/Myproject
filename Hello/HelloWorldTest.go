package Hello

import (
	"fmt"
	"testing"
	//	"github.com/stretchr/testify/assert"
)

func TestHelloworld(t *testing.T) {
	total := hello()
	fmt.Println(total)
	asset.Equal(t, 15, total, "Should be eqaul")
}
