package logger

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
)

type TransactionLogger interface {
	WritePost(key, value string)
	WriteDelete(key string)
	RemoveRedundantLogs()
	ReadEvents() (chan Event, chan error)
	Run()

	Err() <-chan error
}

type EventType byte

const (
	_ = iota                   // 0
	EventPost EventType = iota // 1
	EventDelete 			   // 2
)

type Event struct {
	Sequence int64
	EventType EventType
	Key string
	Value string
}

type FileTransactionLogger struct {
	lastSeq int64
	readLogFile *os.File
	writeLogFile *os.File
	events chan<- Event
	errs  <-chan error
}

func NewFileTransactionLogger(filename string) (TransactionLogger, error) {
	writeFile, err := os.OpenFile("write.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}

	l := &FileTransactionLogger{
		writeLogFile: writeFile,
	}

	return l, nil
}

func (l *FileTransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errs := make(chan error, 1)
	l.errs = errs

	go func() {
		for e := range events {
			l.lastSeq++

			_, err := fmt.Fprintf(l.writeLogFile, "%d\t%d\t%s\t%s\n", l.lastSeq, e.EventType, e.Key, e.Value)
		
			if err != nil {
				errs <- err
				return	
			}
		}
	}()
}

func (l *FileTransactionLogger) ReadEvents() (chan Event, chan error) {
	file, _ := os.Open("read.log")
	scanner := bufio.NewScanner(file)
	outEvents := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		var e Event

		defer file.Close()
		defer close(outEvents)
		defer close(outError)

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Sscanf(line, "%d\t%d\t%s\t%s", &e.Sequence, &e.EventType, &e.Key, &e.Value)
			
			if e.EventType == EventDelete {
				e.Value = ""
			}

			l.lastSeq = e.Sequence

			outEvents <- e
		}

		if err := scanner.Err(); err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
			return
		}

	}()
	
	return outEvents, outError
}

func (l *FileTransactionLogger) WritePost(key, value string) {
	l.events <- Event{
		EventType: EventPost,
		Key: key,
		Value: value,
	}
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{
		EventType: EventDelete,
		Key: key,
	}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errs
}

func (l *FileTransactionLogger) RemoveRedundantLogs() {
	
	scanner := bufio.NewScanner(l.writeLogFile)	
	scanner.Split(bufio.ScanLines)
	
	var text []string
	keys := make([]string, 0)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	res := make([]string, 0) 
	for i:=len(text)-1; i>=0; i-- {
		line := strings.Split(text[i], "\t")

		if !contains(keys, line[2]) {
			keys = append(keys, line[2])
			res = append(res, text[i])
		}
	}

	log.Println(res)
	readFile, _ := os.OpenFile("read.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	defer readFile.Close()
	l.readLogFile = readFile
	for _, v := range res {
		fmt.Fprintf(l.readLogFile, "%s\n", v)
	} 

	log.Println("file is cleaned")	
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}