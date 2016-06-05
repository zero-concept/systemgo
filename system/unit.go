package system

import (
	"bytes"
	"errors"
	"io"
	"log"
	"sync"

	"github.com/b1101/systemgo/lib/state"
)

type Unit struct {
	Supervisable

	log UnitLog

	name string

	stats struct {
		path   string
		loaded state.Load
	}

	listeners listeners
	rdy       chan interface{}
}

type listeners struct {
	ch []chan interface{}
	sync.Mutex
}

type UnitLog struct { // WIP
	*log.Logger
	out io.Writer
}

func NewUnit() (u *Unit) {
	var b bytes.Buffer
	u = &Unit{}
	u.log.Logger = log.New(&b, "", log.LstdFlags)
	u.log.out = &b
	u.rdy = make(chan interface{})
	go u.readyNotifier()
	return
}

func (u *Unit) readyNotifier() {
	for {
		<-u.rdy
		for _, c := range u.listeners.ch {
			c <- struct{}{}
		}
		u.listeners.ch = []chan interface{}{}
	}
}

func (u *Unit) ready() {
	u.rdy <- struct{}{}
}

func (u *Unit) waitFor() <-chan interface{} {
	u.listeners.Lock()
	c := make(chan interface{})
	u.listeners.ch = append(u.listeners.ch, c)
	u.listeners.Unlock()
	return c
}

func (u *Unit) Log(v ...interface{}) {
	u.log.Logger.Println(v)
}
func (u *Unit) SetOutput(w io.Writer) {
	u.log.Logger.SetOutput(w)
}

func (u Unit) Read(b []byte) (int, error) {
	if reader, ok := u.log.out.(io.Reader); ok {
		return reader.Read(b)
	}
	return 0, errors.New("unreadable")
}

func (u Unit) Name() string {
	return u.name
}
func (u Unit) Description() string {
	if u.Supervisable != nil {
		return u.Supervisable.Description()
	} else {
		return ""
	}
}

func (u Unit) Path() string {
	return u.stats.path
}
func (u Unit) Loaded() state.Load {
	return u.stats.loaded
}

//func (u Unit) Enabled() state.Enable {
//return u.stats.enabled
//}
func (u Unit) Active() state.Activation {
	if u.Supervisable != nil {
		return u.Supervisable.Active()
	} else {
		return state.Inactive
	}
}
func (u Unit) Sub() state.Sub {
	if u.Supervisable != nil {
		return u.Supervisable.Sub()
	} else {
		return state.Unavailable
	}
}