package internet

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/lucaskatayama/pigo/internal/core"
	"github.com/lucaskatayama/pigo/internal/ui/display"
)

type Internet struct {
	core.FeatureI

	_conn string
	_ip   net.IP

	_back func()
}

func New(back func()) *Internet {
	return &Internet{
		_back: back,
	}
}

func (i *Internet) IP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal("unable to discover IP address")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func (i *Internet) Check() bool {
	client := &http.Client{Timeout: 5 * time.Second}
	_, err := client.Get("http://clients3.google.com/generate_204")
	return err == nil
}

func (i *Internet) Draw(d display.Display) {
	d.Begin()
	d.Text("Internet", 2, 2)

	if i._conn != "" {
		d.Text(i._conn, 2, 15)
	}

	if i._ip != nil {
		d.Text(i._ip.String(), 2, 28)
	}

	d.Flush()
}

func (i *Internet) OnEnter() {
	f := "%sonnected"
	st := "C"
	if ok := i.Check(); !ok {
		st = "Disc"
	}
	i._conn = fmt.Sprintf(f, st)

	i._ip = i.IP()
}

func (i *Internet) OnDoubleClick() {
	if i._back != nil {
		i._back()
	}
}
