package global

import "time"

type serverini struct {
	HttpAddress  string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
