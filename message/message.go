package message

import (
	// "code.google.com/p/go-uuid/uuid"
	"github.com/garyburd/redigo/redis"
	"github.com/saintienn/go-spamc"
	"os"
)

var (
	redis_url        = os.Getenv("REDIS_URL")
	spamassassin_url = os.Getenv("SPAMASSASSIN_URL")

	redis_connection, err = redis.Dial("tcp", redis_url)
)

type Message struct {
	MessageID string
	Body      string
	Headers   string
	Result    interface{}
}

func Find(uuid string) {
	redis_connection.Send("GET", uuid)
}

func (m *Message) SpamCheck() {
	html := "<html>Hello world. I'm not a Spam, don't kill me SpamAssassin!</html>"

	client := spamc.New(spamassassin_url, 10)

	//the 2nd parameter is optional, you can set who (the unix user) do the call
	reply, _ := client.Report(html)

	m.Result = reply
}
