package speedtest

import (
	"context"
	"fmt"
	"time"

	"github.com/lucaskatayama/pigo/internal/core"
	"github.com/lucaskatayama/pigo/internal/ui/display"
	"github.com/showwin/speedtest-go/speedtest"
)

type State int

const (
	Begin State = iota
	Testing
	Result
)

type SpeedTest struct {
	core.FeatureI

	_back   func()
	_state  State
	_cancel context.CancelFunc

	_ping     time.Duration
	_upload   string
	_download string
	_ip       string
}

func New(back func()) *SpeedTest {
	st := &SpeedTest{
		_back:  back,
		_state: Begin,
	}

	return st
}

func (st *SpeedTest) Draw(d display.Display) {
	d.Begin()
	d.Text("SpeedTest", 2, 2)
	if st._state == Begin {
		st.Reset()
		d.Text("Long press", 2, 14)
		d.Text("to start", 2, 27)
	} else if st._state == Testing {
		d.TextFont(st._ip, "pixelmix", 2, 14)
	} else if st._state == Result {
		y := 0
		y += d.TextFont(st._ip, "pixelmix", 2, 15)
		y += d.TextFont(fmt.Sprintf("Ping: %s", st._ping.Round(time.Millisecond)), "pixelmix", 2, 16+y)
		y += d.TextFont(fmt.Sprintf("Down: %s", st._download), "pixelmix", 2, 17+y)
		d.TextFont(fmt.Sprintf("Up__: %s", st._upload), "pixelmix", 2, 18+y)
	}
	d.Flush()
}

func (st *SpeedTest) Cancel() {
	if st._cancel != nil {
		st._cancel()
	}
}

func (st *SpeedTest) Reset() {
	st.Cancel()
	st._ip = ""
	st._ping = 0
	st._download = ""
	st._upload = ""
}

func (st *SpeedTest) Start(_st State) {
	st.Reset()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	st._cancel = cancel
	defer cancel()
	st._state = _st

	st.Test(ctx)
}

func (st *SpeedTest) OnEnter() {
	st.Reset()

}

func (st *SpeedTest) OnDoubleClick() {
	st._back()
}

func (st *SpeedTest) OnPress() {
	st.Cancel()
	switch st._state {
	case Begin:
		st.Start(Testing)
	case Result:
		st.Start(Testing)
	}
}

func (st *SpeedTest) Test(ctx context.Context) {

	var speedtestClient = speedtest.New()

	// Use a proxy for the speedtest. eg: socks://127.0.0.1:7890
	// speedtest.WithUserConfig(&speedtest.UserConfig{Proxy: "socks://127.0.0.1:7890"})(speedtestClient)

	// Select a network card as the data interface.
	// speedtest.WithUserConfig(&speedtest.UserConfig{Source: "192.168.1.101"})(speedtestClient)

	user, _ := speedtestClient.FetchUserInfo()
	// Get a list of servers near a specified location
	// user.SetLocationByCity("Tokyo")
	// user.SetLocation("Osaka", 34.6952, 135.5006)
	st._ip = user.IP

	serverList, _ := speedtestClient.FetchServers()
	s := serverList[0]
	st._state = Result

	s.PingTest(nil)
	st._ping = s.Latency

	ticker := speedtestClient.CallbackDownloadRate(func(downRate float64) {
		st._download = fmt.Sprintf("%.2f Mbps", downRate)
	})
	s.DownloadTestContext(ctx)
	ticker.Stop()
	st._download = fmt.Sprintf("%.2f Mbps", s.DLSpeed)

	ticker = speedtestClient.CallbackUploadRate(func(upRate float64) {
		st._upload = fmt.Sprintf("%.2f Mbps", upRate)
	})
	s.UploadTestContext(ctx)
	ticker.Stop()
	st._upload = fmt.Sprintf("%.2f Mbps", s.ULSpeed)
}
