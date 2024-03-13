// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/Godx1an/gp_ent/pkg/ent_work"
)

// The AdminFunc type is an adapter to allow the use of ordinary
// function as Admin mutator.
type AdminFunc func(context.Context, *ent_work.AdminMutation) (ent_work.Value, error)

// Mutate calls f(ctx, m).
func (f AdminFunc) Mutate(ctx context.Context, m ent_work.Mutation) (ent_work.Value, error) {
	if mv, ok := m.(*ent_work.AdminMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent_work.AdminMutation", m)
}

// The SchoolFunc type is an adapter to allow the use of ordinary
// function as School mutator.
type SchoolFunc func(context.Context, *ent_work.SchoolMutation) (ent_work.Value, error)

// Mutate calls f(ctx, m).
func (f SchoolFunc) Mutate(ctx context.Context, m ent_work.Mutation) (ent_work.Value, error) {
	if mv, ok := m.(*ent_work.SchoolMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent_work.SchoolMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent_work.UserMutation) (ent_work.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent_work.Mutation) (ent_work.Value, error) {
	if mv, ok := m.(*ent_work.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent_work.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent_work.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent_work.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent_work.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent_work.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent_work.Op) Condition {
	return func(_ context.Context, m ent_work.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent_work.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent_work.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent_work.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent_work.Hook, cond Condition) ent_work.Hook {
	return func(next ent_work.Mutator) ent_work.Mutator {
		return ent_work.MutateFunc(func(ctx context.Context, m ent_work.Mutation) (ent_work.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent_work.Delete|ent_work.Create)
func On(hk ent_work.Hook, op ent_work.Op) ent_work.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent_work.Update|ent_work.UpdateOne)
func Unless(hk ent_work.Hook, op ent_work.Op) ent_work.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent_work.Hook {
	return func(ent_work.Mutator) ent_work.Mutator {
		return ent_work.MutateFunc(func(context.Context, ent_work.Mutation) (ent_work.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent_work.Hook {
//		return []ent_work.Hook{
//			Reject(ent_work.Delete|ent_work.Update),
//		}
//	}
func Reject(op ent_work.Op) ent_work.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent_work.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent_work.Hook) Chain {
	return Chain{append([]ent_work.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent_work.Hook {
	return func(mutator ent_work.Mutator) ent_work.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent_work.Hook) Chain {
	newHooks := make([]ent_work.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
