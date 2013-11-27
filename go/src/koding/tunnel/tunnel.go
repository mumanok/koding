package tunnel

import (
	"bufio"
	"fmt"
	"io"
	"koding/tunnel/conn"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// tunnel is implementing the net.Conn interface and defines a tcp connection
// between the client and the public requester conn. tunnel is created via the
// client.
type tunnel struct {
	// underlying tcp connection
	*conn.Conn

	// start time of the tunnel connection
	start time.Time

	// protects single http requests
	sync.Mutex
}

func newTunnel(nc net.Conn) *tunnel {
	t := &tunnel{
		start: time.Now(),
	}

	t.Conn = conn.New(nc, false)
	return t
}

func newTunnelDial(addr string, serverMsg *ServerMsg) *tunnel {
	t := &tunnel{}
	t.Conn = conn.Dial(addr, false)

	err := t.connect(serverMsg)
	if err != nil {
		log.Fatalln("newTunnelConn", err)
	}

	return t
}

func (t *tunnel) connect(serverMsg *ServerMsg) error {
	remoteAddr := fmt.Sprintf("http://%s%s", t.RemoteAddr(), TunnelPath)
	req, err := http.NewRequest("CONNECT", remoteAddr, nil)
	if err != nil {
		return fmt.Errorf("CONNECT", err)
	}

	req.Header.Set("protocol", serverMsg.Protocol)
	req.Header.Set("tunnelID", serverMsg.TunnelID)
	req.Header.Set("identififer", serverMsg.Identifier)
	req.Write(t)

	resp, err := http.ReadResponse(bufio.NewReader(t), req)
	if err != nil {
		return fmt.Errorf("read response", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.Status != Connected {
		return fmt.Errorf("tunnel server response: %s", resp.Status)
	}

	return nil
}

func (t *tunnel) proxy(w http.ResponseWriter, r *http.Request) error {
	t.Lock()
	defer t.Unlock()

	err := r.Write(t)
	if err != nil {
		return fmt.Errorf("write to tunnel: %s", err)
	}

	resp, err := http.ReadResponse(bufio.NewReader(t), r)
	if err != nil {
		return fmt.Errorf("read from tunnel: %s", err.Error())
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)

	return nil
}
