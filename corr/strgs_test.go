package corr

import (
	"strings"
	"testing"
)

func TestStrsOne(t *testing.T) {
	t.Log(">>> Initiating test <<<")
	str := "Hello"
	res := strings.Split(str, "")
	println(res)
}
