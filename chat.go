package main

import (
	"container/list"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var (
	subscribe  = make(chan (chan<- Subscription), 10)
	unsubsribe = make(chan (<-chan Event), 10)
	publish    = make(chan Event, 10)
)

// Event
type Event struct {
	EvtType   string
	User      string
	Timestamp int
	Text      string
}

// Subscription
type Subscription struct {
	Archive []Event
	New     <-chan Event
}

// Newevent
func NewEvent(evtType, user, msg string) Event {
	return Event{evtType, user, int(time.Now().Unix()), msg}
}

// for sub
func Subscribe() Subscription {
	c := make(chan Subscription)
	subscribe <- c
	return <-c
}

// Cancel
func (s Subscription) Cancel() {
	unsubsribe <- s.New

	for {
		select {
		case _, ok := <-s.New:
			if !ok {
				return
			}
		default:
			return
		}
	}
}

func Join(user string) {
	publish <- NewEvent("join", user, "")
}

func Say(user, message string) {
	publish <- NewEvent("message", user, message)
}

func Leave(user string) {
	publish <- NewEvent("leave", user, "")
}

func Chatroom() {
	archive := list.New()
	subscribers := list.New()

	for {
		select {
		case c := <-subscribe:
			var events []Event
			for e := archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(Event))
			}

			subscriber := make(chan Event, 10)
			subscribers.PushBack(subscriber)
			c <- Subscription{events, subscriber}

		case c := <-unsubsribe:
			for e := subscribers.Front(); e != nil; e = e.Next() {
				subscriber := e.Value.(chan Event)
				if subscriber == c {
					subscribers.Remove(e)
					break
				}
			}
		}
	}
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	go Chatroom()

	server.OnConnect("/", func(so socketio.Conn) error {
		Join(so.ID())
		log.Println("connected: ", so.ID())
		s := Subscribe()

		for _, event := range s.Archive {
			so.Emit("event", event)
		}

		newMessages := make(chan string)
		server.OnEvent("/", "event", func(so socketio.Conn, msg string) {
			log.Println("Emit Event: ", msg)
			newMessages <- msg
		})

		server.OnDisconnect("/", func(so socketio.Conn, reason string) {
			log.Println("Disconnected: ", so.ID())
			Leave(so.ID())
			s.Cancel()
		})

		go func() {
			for {
				select {
				case event := <-s.New:
					so.Emit("event", event)
				case msg := <-newMessages:
					Say(so.ID(), msg)
				}
			}
		}()

		return nil
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":80", nil)
}
