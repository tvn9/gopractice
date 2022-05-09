// Testing in Go exercies
package main

import (
	"fmt"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	name := "Chris"
	chris := Player{}
	chris.NewPlayer(name)
	if chris.name != name {
		t.Errorf("fail to create player name, expected %s, but got %s instead\n", name, chris.name)
	} else {
		chris.printPlayerStat()
	}

	testHealth := chris.health
	fmt.Println("Test health", testHealth)

	testHealth -= 50
	fmt.Println("testHealth", testHealth)

	chris.lossHealth(50)
	if chris.health != testHealth {
		t.Errorf("result missmatch, expect %d, got %d instead\n", testHealth, chris.health)
	} else {
		chris.printPlayerStat()
	}

}
