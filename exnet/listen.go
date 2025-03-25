package exnet

import (
	"context"
	"net"
)

// Listen likes net.Listen, but it will close the listener when ctx is done.
func Listen(ctx context.Context, network, address string) (net.Listener, error) {
	var lc net.ListenConfig

	l, err := lc.Listen(ctx, network, address)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()

		_ = l.Close()
	}()

	return l, nil
}
