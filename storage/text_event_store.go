package storage

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"reflect"
)

type TextEventStore struct {
	filename string
	store    ReadWriter
	registry func(string) interface{}
}

func NewTextEventStore(store ReadWriter, registry func(string) interface{}) *TextEventStore {
	return &TextEventStore{
		filename: defaultFilename,
		store:    store,
		registry: registry,
	}
}

func (this *TextEventStore) Store(messages []interface{}) error {
	buffer := bytes.NewBuffer([]byte{})
	writeToBuffer(buffer, messages)
	return this.store.Write(this.filename, ioutil.NopCloser(buffer))
}
func writeToBuffer(buffer *bytes.Buffer, messages []interface{}) {
	for _, message := range messages {
		buffer.WriteString(reflect.TypeOf(message).Name())
		buffer.WriteString(fieldDelimiter)
		buffer.WriteString(serialize(message))
		buffer.WriteString(lineBreak)
	}
}
func serialize(message interface{}) string {
	if serialized, err := json.Marshal(message); err == nil {
		return string(serialized)
	} else {
		panic(err)
	}
}

func (this *TextEventStore) Load() <-chan interface{} {
	output := make(chan interface{}, 1024)
	go this.load(output)
	return output
}
func (this *TextEventStore) load(channel chan<- interface{}) {
	reader, err := this.store.Read(this.filename)
	if err != nil && err == NotFoundError {
		close(channel)
		return
	} else if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		line := scanner.Bytes()
		index := bytes.Index(line, []byte(fieldDelimiterBytes))
		if index < 0 {
			log.Panic(missingDelimiterError)
		}

		instance := this.registry(string(line[0:index]))
		body := line[index:]
		if err := json.Unmarshal(body, &instance); err != nil {
			panic(err)
		}
		channel <- instance
	}

	close(channel)
}

const (
	fieldDelimiter  = "\t"
	lineBreak       = "\n"
	defaultFilename = "events.txt"
)

var fieldDelimiterBytes = []byte(fieldDelimiter)
var missingDelimiterError = errors.New("missing field delimiter")
