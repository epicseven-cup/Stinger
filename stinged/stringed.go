package stinged

import (
	"context"
	"github.com/containers/podman/v5/pkg/bindings"
	"os"
)

type Stringed struct {
	podman       context.Context // Context that can used in other podman function calls
	imageBuilder *Buildah        // image builder instance, can be any kind of oci image builder. Using buildah for now
}

// Just example that I was thinking
func NewStringed() (*Stringed, error) {
	conn, err := bindings.NewConnection(context.Background(), os.Getenv("PODMAND_URI"))
	if err != nil {
		return nil, err
	}

	return &Stringed{
		podman: conn,
	}, nil
}
