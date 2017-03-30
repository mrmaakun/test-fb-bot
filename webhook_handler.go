package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func receiveMessage(event Messaging) {

	if &event.Message != nil {
		log.Println("Text: " + event.Message.Text)
		log.Println("MID: " + event.Message.Mid)
		log.Println("Seq: " + strconv.Itoa(event.Message.Seq))
	}
	sendTextMessage(event.Sender.ID, event.Message.Text)

}

func sendTextMessage(senderId string, message string) {
	messageData := &MessageRequestData{}
	messageData.Recipient.ID = senderId
	messageData.Message.Text = message
	callSendApi(*messageData)
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered Handler")

	// read body

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	log.Println("Request: " + string(dump))

	//decode json
	var data ResponseData

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	if data.Object == "page" {

		for _, entry := range data.Entry {

			log.Println("pageId: " + entry.ID)
			log.Println("timeofevent: " + strconv.FormatInt(entry.Time, 10))

			for _, event := range entry.Messaging {

				receiveMessage(event)
			}
		}
	}

	challenge_value := r.URL.Query().Get("hub.challenge")
	verify_token := r.URL.Query().Get("hub.verify_token")
	log.Println(challenge_value)
	log.Println(verify_token)
	w.Write([]byte(challenge_value))
}

func registerRouteHandlers() {
	var endpoint_port = os.Getenv("PORT")
	http.HandleFunc("/webhook", webhookHandler)

	log.Println("FBBOT: Listen on port " + endpoint_port)
	err := http.ListenAndServe(":"+endpoint_port, nil) // set listen port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	registerRouteHandlers()
}
