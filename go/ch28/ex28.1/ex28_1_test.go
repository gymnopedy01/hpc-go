package main

import "testing"

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9) return %d", rst)
	}
}

// 테스트 실행방법
// $ pwd
// C:\GitWorkSpace\hpc-go\go\ch28\ex28.1
// $ go mod init ch28/ex28.1
// $ go mod tidy
// $ go test
// PASS
// ok      ch28/ex28.1     0.275s
