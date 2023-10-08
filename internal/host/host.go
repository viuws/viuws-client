package host

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"math"
	"os"
	"sync"
)

type Host struct {
	lastId      uint32
	lastIdMutex sync.Mutex
	stdoutMutex sync.Mutex
}

type Address uint8

const (
	Webpage Address = iota
	ContentScript
	ServiceWorker
	NativeHost
)

type MessageHeader struct {
	Id        uint32  `json:"id"`
	Source    Address `json:"source"`
	Target    Address `json:"target"`
	Webpage   uint16  `json:"webpage,omitempty"`
	RequestId uint32  `json:"requestId,omitempty"`
}

func NewHost() (*Host, error) {
	return &Host{
		lastId:      0,
		lastIdMutex: sync.Mutex{},
		stdoutMutex: sync.Mutex{},
	}, nil
}

func (host *Host) Run() {
	// TODO
}

func (host *Host) Send(target Address, webpage uint16, requestId uint32, payload *interface{}) (uint32, error) {
	header := MessageHeader{
		Id:        host.generateId(),
		Source:    NativeHost,
		Target:    target,
		Webpage:   webpage,
		RequestId: requestId,
	}
	headerJson, err := json.Marshal(&header)
	if err != nil {
		return header.Id, errors.New("failed to marshal header: " + err.Error())
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return header.Id, errors.New("failed to marshal payload: " + err.Error())
	}
	length := len(headerJson) + len(payloadJson)
	if length > 1048576 {
		return header.Id, errors.New("message too long")
	}
	host.stdoutMutex.Lock()
	binary.Write(os.Stdout, binary.NativeEndian, uint32(length))
	os.Stdout.Write(headerJson)
	os.Stdout.Write(payloadJson)
	os.Stdout.Sync()
	host.stdoutMutex.Unlock()
	return header.Id, nil
}

func (host *Host) generateId() uint32 {
	host.lastIdMutex.Lock()
	if host.lastId < math.MaxUint32 {
		host.lastId++
	} else {
		host.lastId = 1
	}
	id := host.lastId
	host.lastIdMutex.Unlock()
	return id
}
