package host

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"math"
	"os"
	"sync"
)

type Host struct {
	mu        sync.Mutex
	lastMsgId uint32
}

type Address uint8

const (
	Webpage Address = iota
	ContentScript
	ServiceWorker
	NativeHost
)

type Command uint8

type MessageHeader struct {
	Id        uint32  `json:"id"`
	Source    Address `json:"source"`
	Target    Address `json:"target"`
	Cmd       Command `json:"cmd"`
	Webpage   uint16  `json:"webpage,omitempty"`
	RequestId uint32  `json:"requestId,omitempty"`
}

type Message struct {
	Header  MessageHeader   `json:"header"`
	Payload json.RawMessage `json:"payload"`
}

const MaxMsgSize = 1048576 // 1 MiB
const ReaderBufSize = 4096 // 4 KiB

func NewHost() (*Host, error) {
	return &Host{
		mu:        sync.Mutex{},
		lastMsgId: 0,
	}, nil
}

func (host *Host) Run() error {
	var msgSize uint32
	reader := bufio.NewReaderSize(os.Stdin, ReaderBufSize)
	readMsgSizeErr := binary.Read(reader, binary.NativeEndian, &msgSize)
	for readMsgSizeErr == nil {
		msgBytes := make([]byte, msgSize)
		_, err := io.ReadFull(reader, msgBytes)
		if err != nil {
			return errors.New("failed to read message: " + err.Error())
		}
		var msg Message
		err = json.Unmarshal(msgBytes, &msg)
		if err != nil {
			return errors.New("failed to unmarshal message: " + err.Error())
		}
		host.handleMessage(&msg)
		readMsgSizeErr = binary.Read(reader, binary.NativeEndian, &msgSize)
	}
	if readMsgSizeErr != io.EOF {
		return errors.New("failed to read message size: " + readMsgSizeErr.Error())
	}
	return nil
}

func (host *Host) Send(target Address, cmd Command, webpage uint16, requestId uint32, payload *json.RawMessage) (uint32, error) {
	msg := Message{
		Header: MessageHeader{
			Source:    NativeHost,
			Target:    target,
			Cmd:       cmd,
			Webpage:   webpage,
			RequestId: requestId,
		},
		Payload: *payload,
	}
	msgBytes, err := json.Marshal(&msg)
	if err != nil {
		return 0, errors.New("failed to marshal message: " + err.Error())
	}
	msgSize := len(msgBytes)
	if msgSize > MaxMsgSize {
		return 0, errors.New("message too large")
	}
	host.mu.Lock()
	if host.lastMsgId < math.MaxUint32 {
		msg.Header.Id = host.lastMsgId + 1
	} else {
		msg.Header.Id = 1
	}
	err = os.Stdout.Sync()
	if err != nil {
		host.mu.Unlock()
		return 0, errors.New("failed to sync stdout: " + err.Error())
	}
	err = binary.Write(os.Stdout, binary.NativeEndian, uint32(msgSize))
	if err != nil {
		host.mu.Unlock()
		return 0, errors.New("failed to write message size: " + err.Error())
	}
	_, err = os.Stdout.Write(msgBytes)
	if err != nil {
		host.mu.Unlock()
		return 0, errors.New("failed to write message: " + err.Error())
	}
	err = os.Stdout.Sync()
	if err != nil {
		host.mu.Unlock()
		return 0, errors.New("failed to sync stdout: " + err.Error())
	}
	host.lastMsgId = msg.Header.Id
	host.mu.Unlock()
	return msg.Header.Id, nil
}

func (host *Host) handleMessage(msg *Message) {
	// TODO
}
