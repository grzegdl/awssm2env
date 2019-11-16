package env

import (
	"os"
	"testing"
)

func TestSetEnvs(t *testing.T) {
	testKey := "DATABASE_PASSWORD"
	testValue := "replaceme"

	env := map[string]string{
		testKey: testValue,
	}

	defer os.Unsetenv(testKey)

	err := SetEnvs(env)

	if err != nil {
		t.Fatalf("Unable to SetEnvs")
	}

	if _, ok := os.LookupEnv(testKey); !ok {
		t.Errorf("SetEnv(%q, %q) didn't set $%s", testValue, testKey, testKey)
	}
}

// func TestEnvs2Map(t *testing.T) {
// 	env := []string{
// 		"DATABASE_PASSWORD=replaceme",
// 		"secret_name=replaceme",
// 		"XMODIFIERS=@im=ibus=sss",
// 	}

// 	got := Envs2Map(env) // map

// 	want := map[string]string{
// 		"DATABASE_PASSWORD": "replaceme",
// 		"secret_name":       "replaceme",
// 	}

// 	if (reflect.DeepEqual(want, got)) == false {
// 		t.Errorf("Envs2Map(env) = %v; want *%v", got, want)
// 	}
// }
