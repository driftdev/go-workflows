package workflow

import (
	"github.com/cschleiden/go-workflows/internal/core"
)

type (
	Instance = core.WorkflowInstance
	Metadata = core.WorkflowMetadata
	Workflow = interface{}
)
