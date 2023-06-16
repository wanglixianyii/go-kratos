package encryption

import (
	"fmt"
	"testing"
)

func TestMd5BySalt(t *testing.T) {
	fmt.Println(Md5BySalt("123456", "pGsPfQfKaP"))
}
