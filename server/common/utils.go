package common

import (
	"GitHelper/server/services"
	"context"
	"errors"
)

func GetServiceDS(ctx context.Context) (services.DS, error) {
	ds, ok := ctx.Value("DS").(services.DS)
	if !ok {
		return services.DS{}, errors.New("unable to access required implementor")
	}
	return ds, nil
}
