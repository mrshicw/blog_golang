package tests

import (
	"blog_golang/testutil"
	"fmt"
	"testing"
)

func TestSetupDemo(t *testing.T) {
	fmt.Println("TestSetupDemo started")
	testutil.TestDatabaseSetup(t)
	fmt.Println("TestSetupDemo completed")
}
