package main

import (
	"reflect"
	"testing"
)

func TestQuote(t *testing.T) {
	testingCases := []struct {
		name     string
		value    string
		expected string
	}{
		{"number", "64", "64"},
		{"boolean", "True", "True"},
		{"string", "hello", "\"hello\""},
	}

	for _, tc := range testingCases {
		t.Run(tc.name, func(t *testing.T) {
			result := quote(tc.value)
			if result != tc.expected {
				t.Fatalf("quoting %s should be %s ; got %s", tc.value, tc.expected, result)
			}
		})
	}
}

func TestJsonify(t *testing.T) {
	testingCases := []struct {
		name     string
		header   []string
		line     []string
		expected string
		err      string
	}{
		{
			name:     "valid case",
			header:   []string{"a", "b", "c"},
			line:     []string{"1", "2", "3"},
			expected: "{ \"a\": 1, \"b\": 2, \"c\": 3 }",
		},
		{
			name:   "invalid case",
			header: []string{"a", "b", "c"},
			line:   []string{"1", "2"},
			err:    "[a b c] is not the same length as [1 2]",
		},
	}

	for _, tc := range testingCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := jsonify(tc.header, tc.line)
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("error should be %v ; got %v", tc.err, err)
				}

				return
			}

			if result != tc.expected {
				value := struct {
					header []string
					line   []string
				}{tc.header, tc.line}
				t.Fatalf("jsonifying %v should be %s ; got %s", value, tc.expected, result)
			}
		})
	}
}

func TestJsonifyAll(t *testing.T) {
	testingCases := []struct {
		name     string
		header   []string
		lines    [][]string
		expected []string
		err      string
	}{
		{
			name:     "valid case",
			header:   []string{"a", "b", "c"},
			lines:    [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			expected: []string{"{ \"a\": 1, \"b\": 2, \"c\": 3 }", "{ \"a\": 4, \"b\": 5, \"c\": 6 }"},
		},
	}

	for _, tc := range testingCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := jsonifyAll(tc.header, tc.lines)
			if err != nil {
				return
			}

			if !reflect.DeepEqual(result, tc.expected) {
				value := struct {
					header []string
					lines  [][]string
				}{tc.header, tc.lines}
				t.Fatalf("jsonifying all %v should be %v ; got %v", value, tc.expected, result)
			}
		})
	}
}
