package textarea

import (
	"reflect"
	"testing"
)

func TestAddCurrent(t *testing.T) {

	t.Run("add a correct input", func(t *testing.T) {
		m := Model{target: "a", correct: []int{}, incorrect: []int{}, cursor: 0}

		m.addCurrent("a")

		assertArrayEquals(t, m.correct, []int{0})
		assertArrayEquals(t, m.incorrect, []int{})
	})

	t.Run("add an incorrect input", func(t *testing.T) {
		m := Model{target: "a", correct: []int{}, incorrect: []int{}, cursor: 0}

		m.addCurrent("b")

		assertArrayEquals(t, m.correct, []int{})
		assertArrayEquals(t, m.incorrect, []int{0})

	})
}

func TestRemovePrevious(t *testing.T) {

	t.Run("remove a correct entry", func(t *testing.T) {
		m := Model{correct: []int{0}, incorrect: []int{}, cursor: 0}

		m.removePrevious()

		assertArrayEquals(t, m.correct, []int{})
	})

	t.Run("remove an incorrect entry", func(t *testing.T) {
		m := Model{correct: []int{}, incorrect: []int{0}, cursor: 0}

		m.removePrevious()

		assertArrayEquals(t, m.incorrect, []int{})
	})

}

func assertArrayEquals(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
