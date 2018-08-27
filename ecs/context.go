package ecs

import (
	"reflect"
	"sort"
)

// Context contains a bunch of Entities, and a bunch of Systems. It is the
// recommended way to run ecs.
type Context struct {
	systems      systems
	sysIn, sysEx map[reflect.Type]reflect.Type
}

// AddSystem adds the given System to the Context, sorted by priority.
func (w *Context) AddSystem(system System) {
	if initializer, ok := system.(Initializer); ok {
		initializer.New(w)
	}

	w.systems = append(w.systems, system)
	sort.Sort(w.systems)
}

// AddSystemInterface adds a system to the Context, but also adds a filter that allows
// automatic adding of entities that match the provided in interface, and excludes any
// that match the provided ex interface, even if they also match in. in and ex must be
// pointers to the interface or else this panics.
func (w *Context) AddSystemInterface(sys SystemAddByInterfacer, in interface{}, ex interface{}) {
	w.AddSystem(sys)

	if w.sysIn == nil {
		w.sysIn = make(map[reflect.Type]reflect.Type)
	}

	w.sysIn[reflect.TypeOf(sys)] = reflect.TypeOf(in).Elem()

	if ex == nil {
		return
	}

	if w.sysEx == nil {
		w.sysEx = make(map[reflect.Type]reflect.Type)
	}

	w.sysEx[reflect.TypeOf(sys)] = reflect.TypeOf(ex).Elem()
}

// AddEntity adds the entity to all systems that have been added via
// AddSystemInterface. If the system was added via AddSystem the entity will not be
// added to it.
func (w *Context) AddEntity(e Identifier) {
	if w.sysIn == nil {
		w.sysIn = make(map[reflect.Type]reflect.Type)
	}
	if w.sysEx == nil {
		w.sysEx = make(map[reflect.Type]reflect.Type)
	}
	for _, system := range w.systems {
		sys, ok := system.(SystemAddByInterfacer)
		if !ok {
			continue
		}
		if ex, not := w.sysEx[reflect.TypeOf(sys)]; not {
			if reflect.TypeOf(e).Implements(ex) {
				continue
			}
		}
		if in, ok := w.sysIn[reflect.TypeOf(sys)]; ok {
			if reflect.TypeOf(e).Implements(in) {
				sys.AddByInterface(e)
			}
		}
	}
}

// Systems returns the list of Systems managed by the Context.
func (w *Context) Systems() []System {
	return w.systems
}

// Update updates each System managed by the Context. It is invoked by the engine
// once every frame, with dt being the duration since the previous update.
func (w *Context) Update(dt float32) {
	for _, system := range w.Systems() {
		system.Update(dt)
	}
}

// RemoveEntity removes the entity across all systems.
func (w *Context) RemoveEntity(e BasicEntity) {
	for _, sys := range w.systems {
		sys.Remove(e)
	}
}