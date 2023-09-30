package podman

import (
	"context"

	"github.com/containers/common/pkg/config"
	"github.com/containers/podman/v4/pkg/bindings"
)

func ListEngines() (map[string]config.Destination, string, error) {
	cfg, err := config.ReadCustomConfig()
	if err != nil {
		return nil, "", err
	}
	engines := cfg.Engine.ServiceDestinations
	defaultEngineName := cfg.Engine.ActiveService
	return engines, defaultEngineName, err
}

func Connect(ctx context.Context, engine config.Destination) (context.Context, error) {
	conn, err := bindings.NewConnectionWithIdentity(ctx, engine.URI, engine.Identity, engine.IsMachine)
	return conn, err
}
