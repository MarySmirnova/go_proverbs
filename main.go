package main

import (
	"math/rand"
	"net"
	"time"
)

const addr = "0.0.0.0:12345"
const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleProverbs(conn)
	}
}

func handleProverbs(conn net.Conn) {
	for {
		conn.Write([]byte(chooseProverb() + "\n"))
		time.Sleep(3 * time.Second)
	}
}

func chooseProverb() string {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"Make the zero value useful.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
	r := rand.Intn(len(proverbs))

	return proverbs[r]
}
