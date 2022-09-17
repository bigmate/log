package log //nolint // api restrictions

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	const (
		greeting   = "hello there, %v"
		question   = "how are you?"
		outputFile = "output.txt"
	)

	buf := &bytes.Buffer{}
	globalLogger = newLogger(defaultEncoder(), nopSyncer{buf})

	SetLevel(LevelFromString("debug"))

	Debugf(greeting, question)
	Infof(greeting, question)
	Warnf(greeting, question)
	Errorf(greeting, question)

	if err := Close(); err != nil {
		t.Errorf("failed to close logger, %v", err)
	}

	var (
		fields = []Field{
			Namespace("planet"),
			String("name", "Jupiter"),
			String("discovered_by", "Galileo Galilei"),
			Bool("has_moons", true),
			Bool("is_surface_solid", false),
			Err(errors.New("failed to run fusion reaction")),
			Any("moons", []string{"Ganymede", "Europa", "Callisto", "Titan"}),
			Duration("day_length", 9*time.Hour+56*time.Minute),
		}
	)

	Debug(question, fields...)
	Info(question, fields...)
	Warn(question, fields...)
	Error(question, fields...)

	f, err := os.Open(outputFile)
	if err != nil {
		t.Errorf("failed to open file: %v", err)
	}

	defer f.Close()

	actual, err := decodeJSON(buf)
	if err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}

	expected, err := decodeJSON(f)
	if err != nil {
		t.Fatalf("failed to decode json: %v", err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("resulting output is different from what it's supposed to be")
		t.Log("actual", actual)
		t.Log("expected", expected)
	}
}

func decodeJSON(r io.Reader) (res []map[string]interface{}, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		buf := scanner.Bytes()

		if len(buf) == 0 {
			continue
		}

		row := make(map[string]interface{})

		if err := json.Unmarshal(buf, &row); err != nil {
			return nil, err
		}

		delete(row, "@timestamp")

		res = append(res, row)
	}

	return res, scanner.Err()
}
