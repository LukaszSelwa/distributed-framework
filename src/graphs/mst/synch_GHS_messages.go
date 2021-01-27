package mst

import (
	"encoding/json"
	"lib"
)

type messageType int

const (
	nilMessage messageType = iota
	msgVerifyEdge
	msgProposeEdge
	msgChooseEdge
	msgConnectEdge
	msgElectNewRoot
	msgCompleated
)

type messageSynchronizedGHS struct {
	Type  messageType
	MWOE  *edge // MWOE - Minimal Weight Outgoing Edge
	Index int
	Root  int
}

func sendMessageSynchGHS(v lib.WeightedGraphNode, index int, message *messageSynchronizedGHS) {
	outMessage, _ := json.Marshal(message)
	v.SendMessage(index, outMessage)
}

func receiveMessageSynchGHS(v lib.WeightedGraphNode, index int) *messageSynchronizedGHS {
	var inMessage messageSynchronizedGHS
	if data := v.ReceiveMessage(index); data != nil {
		if err := json.Unmarshal(data, &inMessage); err == nil {
			return &inMessage
		}
	}
	return nil
}
