package main

import ()

type Message struct {
	Mid  string `json:"mid"`
	Seq  int    `json:"seq"`
	Text string `json:"text"`
}

type Messaging struct {
	Sender struct {
		ID string `json:"id"`
	} `json:"sender"`
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Timestamp int64   `json:"timestamp"`
	Message   Message `json:"message"`
}

type ResponseData struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string      `json:"id"`
		Time      int64       `json:"time"`
		Messaging []Messaging `json:"messaging"`
	} `json:"entry"`
}

type MessageRequestData struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}
