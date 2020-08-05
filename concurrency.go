package u

import "context"

// UniqueChild is a goroutine manager (parent) that can only have one child at a time.
// When you call UniqueChild.SetChild(), UniqueChild cancels the previous child context (if any), then run a new child.
// The child needs to auto-kill itself when its context is done.
type UniqueChild interface {
	SetChild(childFn func(context.Context))
	CloseChild()
}

func NewUniqueChild(ctx context.Context) UniqueChild { return &uniqueChild{ctx: ctx} }

type uniqueChild struct {
	ctx               context.Context
	lastChildCancelFn func()
}

func (parent *uniqueChild) SetChild(childFn func(context.Context)) {
	if parent.lastChildCancelFn != nil {
		parent.lastChildCancelFn()
	}

	var childCtx context.Context
	childCtx, parent.lastChildCancelFn = context.WithCancel(parent.ctx)

	go childFn(childCtx)
}

func (parent *uniqueChild) CloseChild() {
	if parent.lastChildCancelFn != nil {
		parent.lastChildCancelFn()
	}
}
