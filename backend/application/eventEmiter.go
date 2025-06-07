package application

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type EventEmitterImp struct {
	Ctx context.Context
}

func NewEventEmitter(ctx context.Context) *EventEmitterImp {
	return &EventEmitterImp{
		Ctx: ctx,
	}
}

func (p *EventEmitterImp) Emit(event string, args ...any) {
	runtime.EventsEmit(p.Ctx, event, args...)
}

func (p *EventEmitterImp) EmitLoading(loading bool) {
	p.Emit("Loading", loading)
}

func (p *EventEmitterImp) EmitRoute(route string) {
	p.Emit("Route", route)
}
