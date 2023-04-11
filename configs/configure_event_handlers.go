package configs

import (
	"context"

	"github.com/OnLab-Clinical/onlab-clinical-services/contexts/shared/sharedapp"
)

// Configure Event Handlers
func ConfigureEventHandlers(ctx context.Context, handlers ...sharedapp.EventHandler) {
	for _, handler := range handlers {
		go handler.Handle(ctx)
	}
}
