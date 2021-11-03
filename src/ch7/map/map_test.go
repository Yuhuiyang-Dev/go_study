package _map

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Log("len m1=", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len m2=", len(m2))

	m3 := make(map[int]int, 10)
	t.Log("len m3=", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 8: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
