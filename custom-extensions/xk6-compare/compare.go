package compare

import (
	"fmt"
	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/lib"
)

// init is called by the Go runtime at application startup.
func init() {
	modules.Register("k6/x/compare", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// comparator is the exported type
		comparator *Compare
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu:         vu,
		comparator: &Compare{vu: vu},
	}
}

// Compare is the type for our custom API.
type Compare struct {
	vu               modules.VU // provides methods for accessing internal k6 objects
	ComparisonResult string     // textual description of the most recent comparison
}

// IsGreater returns true if a is greater than b, or false otherwise, setting textual result message.
func (c *Compare) IsGreater(a, b int) bool {
	if a > b {
		c.ComparisonResult = fmt.Sprintf("%d is greater than %d", a, b)
		return true
	} else {
		c.ComparisonResult = fmt.Sprintf("%d is NOT greater than %d", a, b)
		return false
	}
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.comparator,
	}
}

// InternalState holds basic metadata from the runtime state.
type InternalState struct {
	ActiveVUs       int64 `js:"activeVUs"`
	Iteration       int64
	VUID            uint64     `js:"vuID"`
	VUIDFromRuntime goja.Value `js:"vuIDFromRuntime"`
}

// GetInternalState interrogates the current virtual user for state information.
func (c *Compare) GetInternalState() *InternalState {
	state := c.vu.State()
	ctx := c.vu.Context()
	es := lib.GetExecutionState(ctx)
	rt := c.vu.Runtime()

	return &InternalState{
		VUID:            state.VUID,
		VUIDFromRuntime: rt.Get("__VU"),
		Iteration:       state.Iteration,
		ActiveVUs:       es.GetCurrentlyActiveVUsCount(),
	}
}
