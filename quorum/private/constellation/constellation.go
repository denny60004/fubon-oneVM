package constellation

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

type Constellation struct {
	node                    *Client
	c                       *cache.Cache
	isConstellationNotInUse bool
}

var (
	ErrConstellationIsntInit = errors.New("Constellation not in use")
)

func (g *Constellation) Send(data []byte, from string, to []string) (out []byte, err error) {
	if g.isConstellationNotInUse {
		return nil, ErrConstellationIsntInit
	}
	if g.node.usegrpc{
		out, err = g.node.SendPayloadGrpc(data, from, to)
	} else {
		out, err = g.node.SendPayload(data, from, to)
	}
	if err != nil {
		return nil, err
	}
	g.c.Set(string(out), data, cache.DefaultExpiration)
	return out, nil
}

func (g *Constellation) Receive(data []byte) ([]byte, error) {
	if g.isConstellationNotInUse {
		return nil, nil
	}
	if len(data) == 0 {
		return data, nil
	}
	// Ignore this error since not being a recipient of
	// a payload isn't an error.
	// TODO: Return an error if it's anything OTHER than
	// 'you are not a recipient.'
	dataStr := string(data)
	x, found := g.c.Get(dataStr)
	if found {
		return x.([]byte), nil
	}
	var pl []byte
	if g.node.usegrpc{
		pl, _ = g.node.ReceivePayloadGrpc(data)
	} else {
		pl, _ = g.node.ReceivePayload(data)
	}
	g.c.Set(dataStr, pl, cache.DefaultExpiration)
	return pl, nil
}

func New(path string) (*Constellation, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	// We accept either the socket or a configuration file that points to
	// a socket.
	isSocket := info.Mode()&os.ModeSocket != 0
	var cfg *Config
	if !isSocket {
		cfg, err = LoadConfig(path)
		if err != nil {
			return nil, err
		}
	} else {
		cfg = &Config{
			Socket: path,
			Grpc: true,
		}
	}
	n, err := NewClient(cfg, cfg.Grpc)
	if err != nil {
		return nil, err
	}
	err = UpCheck(n)
	if err != nil {
		return nil, err
	}
	return &Constellation{
		node: n,
		c:    cache.New(5*time.Minute, 5*time.Minute),
		isConstellationNotInUse: false,
	}, nil
}

func MustNew(path string) *Constellation {
	if strings.EqualFold(path, "ignore") {
		return &Constellation{
			node: nil,
			c:    nil,
			isConstellationNotInUse: true,
		}
	}
	g, err := New(path)
	if err != nil {
		panic(fmt.Sprintf("MustNew: Failed to connect to Constellation (%s): %v", path, err))
	}
	return g
}

func MaybeNew(path string) *Constellation {
	if path == "" {
		return nil
	}
	return MustNew(path)
}
