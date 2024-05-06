package tool

import (
	"os"
	"reflect"
	"testing"
)

func TestEnsureEnv(t *testing.T) {

	testCases := []struct {
		SetInitialEnvironment func()
		ExpectedEnvironment   []string
	}{
		{
			SetInitialEnvironment: func() {
				os.Clearenv()
			},
			ExpectedEnvironment: []string{"DB_ROOT_PASSWORD=pass"},
		},
		{
			SetInitialEnvironment: func() {
				os.Clearenv()
				_ = os.Setenv("DB_ROOT_PASSWORD", "test1")
			},
			ExpectedEnvironment: []string{"DB_ROOT_PASSWORD=test1"},
		},
	}
	for _, test := range testCases {

		test.SetInitialEnvironment()

		env := EnsureEnv()

		if !reflect.DeepEqual(env, test.ExpectedEnvironment) {
			t.Errorf("got %v want %v", env, test.ExpectedEnvironment)
		}
	}
}
