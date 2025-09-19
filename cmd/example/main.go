package main

import (
	"github.com/Noah-Wagner/example/internal"
)

func GetPolicy() internal.ExampleSum {
	return nil
}

func main() {
	policy := GetPolicy()

	switch policy.(type) {
	// Missing types here.
	}
}
