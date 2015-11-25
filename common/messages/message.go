// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package messages

import (
	"bytes"
	"fmt"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

//A placeholder structure for messages
type Message struct {
	Timestamp interfaces.Timestamp

	//Not marshalled
	hash interfaces.IHash
}

var _ interfaces.IMsg = (*Message)(nil)

func (m *Message) Process(interfaces.IState) {}

func (m *Message) GetHash() interfaces.IHash {
	if m.hash == nil {
		data, err := m.MarshalForSignature()
		if err != nil {
			panic(fmt.Sprintf("Error in CommitChain.GetHash(): %s", err.Error()))
		}
		m.hash = primitives.Sha(data)
	}
	return m.hash
}

func (m *Message) GetTimestamp() interfaces.Timestamp {
	return m.Timestamp
}

func (m *Message) Type() int {
	return -1
}

func (m *Message) Int() int {
	return -1
}

func (m *Message) Bytes() []byte {
	return nil
}

func (m *Message) UnmarshalBinaryData(data []byte) (newdata []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
		}
	}()

	return nil, nil
}

func (m *Message) UnmarshalBinary(data []byte) error {
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *Message) MarshalBinary() (data []byte, err error) {
	return nil, nil
}

func (m *Message) MarshalForSignature() (data []byte, err error) {
	return nil, nil
}

func (m *Message) String() string {
	return ""
}

func (m *Message) DBHeight() int {
	return 0
}

func (m *Message) ChainID() []byte {
	return nil
}

func (m *Message) ListHeight() int {
	return 0
}

func (m *Message) SerialHash() []byte {
	return nil
}

func (m *Message) Signature() []byte {
	return nil
}

// Validate the message, given the state.  Three possible results:
//  < 0 -- Message is invalid.  Discard
//  0   -- Cannot tell if message is Valid
//  1   -- Message is valid
func (m *Message) Validate(interfaces.IState) int {
	return 0
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *Message) Leader(state interfaces.IState) bool {
	switch state.GetNetworkNumber() {
	case 0: // Main Network
		panic("Not implemented yet")
	case 1: // Test Network
		panic("Not implemented yet")
	case 2: // Local Network
		panic("Not implemented yet")
	default:
		panic("Not implemented yet")
	}

}

// Execute the leader functions of the given message
func (m *Message) LeaderExecute(state interfaces.IState) error {
	return nil
}

// Returns true if this is a message for this server to execute as a follower
func (m *Message) Follower(interfaces.IState) bool {
	return true
}

func (m *Message) FollowerExecute(interfaces.IState) error {
	return nil
}

func (e *Message) JSONByte() ([]byte, error) {
	return primitives.EncodeJSON(e)
}

func (e *Message) JSONString() (string, error) {
	return primitives.EncodeJSONString(e)
}

func (e *Message) JSONBuffer(b *bytes.Buffer) error {
	return primitives.EncodeJSONToBuffer(e, b)
}
