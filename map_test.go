package p_map

import (
	"fmt"
	"testing"
)

func TestMapClass_Keys(t *testing.T) {
	fmt.Println(Map.Keys(map[string]interface{}{
		`test`: `1213`,
		`test1`: 11,
	}))
}
