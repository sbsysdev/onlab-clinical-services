package sharedinfra

import (
	"context"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/sharedapp"
)

type ConfigureEventHandlers func(context.Context, ...sharedapp.EventHandler)
