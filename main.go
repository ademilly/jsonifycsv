package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	filename := flag.String("filename", "", "CSV file to convert")
	sep := flag.String("sep", ",", "CSV separetor to use")

	flag.Parse()

	if *filename == "" {
		log.Fatalf("filename flag is empty")
	}

	if len(*sep) > 1 {
		log.Fatalf("separator is too long, should be one character")
	}

	b, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("could not read file %s: %v\n", *filename, err)
	}
	r := csv.NewReader(bytes.NewReader(b))
	r.Comma = []rune(*sep)[0]

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("could not read all file %s to CSV: %v\n", *filename, err)
	}

	header, rows := records[0], records[1:]

	objects, err := jsonifyAll(header, rows)
	if err != nil {
		log.Fatalf("could not convert to JSON: %v", err)
	}

	fmt.Println(squareBracketify(strings.Join(objects, ",\n  ")))
}

func quote(s string) string {
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return s
	}

	if _, err := strconv.ParseBool(s); err == nil {
		return s
	}

	return fmt.Sprintf("\"%s\"", s)
}

func mapValuesToKeys(keys, values []string) string {
	object := make([]string, 0, len(keys))
	for index, value := range values {
		s := fmt.Sprintf("\"%s\": %s", keys[index], quote(value))
		object = append(object, s)
	}

	return strings.Join(object, ", ")
}

func squareBracketify(s string) string {
	return fmt.Sprintf("%s %s %s", "[", s, "]")
}

func curlyBracketify(s string) string {
	return fmt.Sprintf("%s %s %s", "{", s, "}")
}

func jsonify(header, line []string) (string, error) {
	if len(header) != len(line) {
		return "", fmt.Errorf("%v is not the same length as %v", header, line)
	}

	return curlyBracketify(mapValuesToKeys(header, line)), nil
}

func jsonifyAll(header []string, lines [][]string) ([]string, error) {
	objects := make([]string, 0, len(lines))
	for _, line := range lines {
		object, err := jsonify(header, line)
		if err != nil {
			return []string{}, err
		}
		objects = append(objects, object)
	}

	return objects, nil
}
