package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniqueString(t *testing.T) {
	test := []string{"123", "123", "asd", "asd"}
	result := UniqueString(test)
	require.Equal(t, result, []string{"123", "asd"})
}

func TestReversSort(t *testing.T) {
	test := []string{"123", "123", "asd", "asd"}
	result := ReversSort(test)
	require.Equal(t, result, []string{"asd", "asd", "123", "123"})
}

func TestSortNumeric(t *testing.T) {
	test := []string{"9 8 2 5 6 7"}
	result := SortNum(test)
	require.Equal(t, result, []string{"2 5 6 7 8 9"})
}

func TestSortByColumn(t *testing.T) {
	test := []string{"dfg sdfgs", "qwerqer "}
	result := SortByColumn(test, 2)
	require.Equal(t, result, []string{"qwerqer ", "dfg sdfgs"})
}
