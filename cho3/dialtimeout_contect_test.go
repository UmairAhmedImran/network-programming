package cho3

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()
	
	var d net.Dialer
	d.Control = func (_ , _ string, _ syscall.RawConn) error  {
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}

    conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
        conn.Close()
        t.Fatal("Connection did not time out")
    }
	nErr, ok := err.(net.Error)
	if !ok {
        t.Fatal(err)
    } else { 
		if !nErr.Timeout() {
        t.Fatal("Error is not timeout")
    }
}
	if ctx.Err() != nil {
		t.Errorf("expected deadline exceeded, got %v", ctx.Err())
	}
}