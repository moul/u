package u

import "context"

// MonoChild is a goroutine manager (parent) that can only have one child at a time.
// When you call MonoChild.SetChild(), MonoChild cancels the previous child context (if any), then run a new child.
// The child needs to auto-kill itself when its context is done.
type MonoChild interface {
	SetChild(childFn func(context.Context))
}

func NewMonoChild(ctx context.Context) MonoChild { return &monoChild{ctx: ctx} }

type monoChild struct {
	ctx               context.Context
	lastChildCancelFn func()
}

func (parent *monoChild) SetChild(childFn func(context.Context)) {
	if parent.ctx == nil {
		panic("requires a parent context")
	}
	if parent.lastChildCancelFn != nil {
		parent.lastChildCancelFn()
	}

	var childCtx context.Context
	childCtx, parent.lastChildCancelFn = context.WithCancel(parent.ctx)

	go childFn(childCtx)
}
