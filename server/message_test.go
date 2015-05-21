package server

import (
	"github.com/koding/kite"
	"golang.org/x/net/context"
	"testing"
)

func TestMessageID(t *testing.T) {
	ctx := context.Background()
	k := kite.New("test", "0.0.1")
	b := NewBroker(ctx, "/tmp/test", k)
	b.Init()
	m := NewMessage(b.NewID(), nil)
	if len(m.ID) <= 0 {
		t.Errorf("messageID is %v", m.ID)
	}

	t.Logf("%s", m.ID)
}
