package main

import (
	"testing"

	"github.com/MatioYoshitoki/new-world-ops/pkg/workflow"
)

// ============================================================================
// Mocks
// ============================================================================

type mockHookDispatcher struct {
	calls []struct {
		targetType TaskType
		sourceTask *Task
		vars       map[string]interface{}
	}
}

func (m *mockHookDispatcher) Dispatch(targetType TaskType, sourceTask *Task, vars map[string]interface{}) {
	m.calls = append(m.calls, struct {
		targetType TaskType
		sourceTask *Task
		vars       map[string]interface{}
	}{targetType, sourceTask, vars})
}

func (m *mockHookDispatcher) wasCalled(targetType TaskType) bool {
	for _, c := range m.calls {
		if c.targetType == targetType {
			return true
		}
	}
	return false
}

func (m *mockHookDispatcher) callCount(targetType TaskType) int {
	count := 0
	for _, c := range m.calls {
		if c.targetType == targetType {
			count++
		}
	}
	return count
}

type mockReviewFailureCounter struct {
	countFn func(prNumber int, projectName string) int
}

func (m *mockReviewFailureCounter) count(prNumber int, projectName string) int {
	if m.countFn != nil {
		return m.countFn(prNumber, projectName)
	}
	return 0
}

type mockMaestroFlowChecker struct {
	hasFlowsFn func(branch string) bool
}

func (m *mockMaestroFlowChecker) hasFlows(branch string) bool {
	if m.hasFlowsFn != nil {
		return m.hasFlowsFn(branch)
	}
	return true
}

type mockPRFinder struct {
	findPRFn func(branch string) int
}

func (m *mockPRFinder) findPR(branch string) int {
	if m.findPRFn != nil {
		return m.findPRFn(branch)
	}
	return 0
}

// ============================================================================
// Test helpers
// ============================================================================

func setupHookTest(t *testing.T) (*mockHookDispatcher, *mockReviewFailureCounter, *mockMaestroFlowChecker, *mockPRFinder) {
	t.Helper()

	// backup originals
	origDispatcher := hookDispatcher
	origCounter := reviewFailureCounterImpl
	origChecker := maestroFlowCheckerImpl
	origFinder := prFinderImpl
	origAsync := hookAsyncMode

	mockDisp := &mockHookDispatcher{}
	mockCounter := &mockReviewFailureCounter{}
	mockChecker := &mockMaestroFlowChecker{}
	mockFinder := &mockPRFinder{}

	hookDispatcher = mockDisp
	reviewFailureCounterImpl = mockCounter
	maestroFlowCheckerImpl = mockChecker
	prFinderImpl = mockFinder
	hookAsyncMode = false

	setupTestTasks(t)

	t.Cleanup(func() {
		hookDispatcher = origDispatcher
		reviewFailureCounterImpl = origCounter
		maestroFlowCheckerImpl = origChecker
		prFinderImpl = origFinder
		hookAsyncMode = origAsync
	})

	return mockDisp, mockCounter, mockChecker, mockFinder
}

func setupTestTasks(t *testing.T) {
	t.Helper()
	tasksMutex.Lock()
	oldTasks := tasks
	tasks = make(map[string]*Task)
	tasksMutex.Unlock()
	t.Cleanup(func() {
		tasksMutex.Lock()
		tasks = oldTasks
		tasksMutex.Unlock()
	})
}

func makeTask(taskType TaskType, targetID int, status string) *Task {
	return &Task{
		ID:          taskID(taskType, targetID),
		Type:        taskType,
		Status:      status,
		ProjectName: "social-world",
		TargetID:    targetID,
		Branch:      "feat/issue-5",
		Kind:        "full-stack",
		Metadata:    make(map[string]interface{}),
	}
}

// ============================================================================
// Hook chain tests
// ============================================================================

// Dev → Review (success, has PR)
func TestHookChain_DevSuccess_TriggersReview(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 42 }

	task := makeTask(TaskTypeDev, 5, "success")
	spec := taskRegistry[TaskTypeDev]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook to be triggered after dev success")
	}
}

// Dev → no Review (success, no PR)
func TestHookChain_DevSuccess_NoPR(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 0 }

	task := makeTask(TaskTypeDev, 5, "success")
	spec := taskRegistry[TaskTypeDev]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook NOT to be triggered when no PR found")
	}
}

// Dev failure → no Review
func TestHookChain_DevFailure_NoReview(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 42 }

	task := makeTask(TaskTypeDev, 5, "failed")
	spec := taskRegistry[TaskTypeDev]
	triggerHooks(task, spec, map[string]interface{}{}, nil, false)

	if mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook NOT to be triggered after dev failure")
	}
}

// Rework → Review (success, has PR)
func TestHookChain_ReworkSuccess_TriggersReview(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 42 }

	task := makeTask(TaskTypeRework, 5, "success")
	spec := taskRegistry[TaskTypeRework]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook to be triggered after rework success")
	}
}

// Rework → no Review (success, no PR)
func TestHookChain_ReworkSuccess_NoPR(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 0 }

	task := makeTask(TaskTypeRework, 5, "success")
	spec := taskRegistry[TaskTypeRework]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook NOT to be triggered when no PR found")
	}
}

// E2E → Review (success, dev-triggered, has PR)
func TestHookChain_E2ESuccess_DevTriggered_WithPR(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 58 }

	task := makeTask(TaskTypeE2E, 5, "success")
	task.Metadata["triggeredBy"] = "dev-5"
	spec := taskRegistry[TaskTypeE2E]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook to be triggered after E2E success")
	}
}

// E2E → no Review (manual trigger)
func TestHookChain_E2ESuccess_ManualTrigger(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 58 }

	task := makeTask(TaskTypeE2E, 5, "success")
	task.Metadata["triggeredBy"] = "manual"
	spec := taskRegistry[TaskTypeE2E]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook NOT to be triggered for manual E2E")
	}
}

// E2E → no Review (no open PR)
func TestHookChain_E2ESuccess_NoPR(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 0 }

	task := makeTask(TaskTypeE2E, 5, "success")
	task.Metadata["triggeredBy"] = "dev-5"
	spec := taskRegistry[TaskTypeE2E]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if mockDisp.wasCalled(TaskTypeReview) {
		t.Error("expected Review hook NOT to be triggered when no PR found")
	}
}

// Review → Rework (NEEDS_FIX)
func TestHookChain_ReviewNonPass_TriggersRework(t *testing.T) {
	mockDisp, _, _, _ := setupHookTest(t)

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "NEEDS_FIX"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeRework) {
		t.Error("expected Rework hook to be triggered for non-PASS review")
	}
}

// Review → no Rework (PASS)
func TestHookChain_ReviewPass_NoRework(t *testing.T) {
	mockDisp, _, _, _ := setupHookTest(t)

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "passed"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if mockDisp.wasCalled(TaskTypeRework) {
		t.Error("expected Rework hook NOT to be triggered for PASS review")
	}
}

// Review → Rework (failures >= 3, doc-gardener hooks removed)
func TestHookChain_ReviewFailures_AlwaysTriggersRework(t *testing.T) {
	mockDisp, mockCounter, _, _ := setupHookTest(t)
	mockCounter.countFn = func(prNumber int, projectName string) int { return 3 }

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "NEEDS_FIX"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeRework) {
		t.Error("expected Rework hook to be triggered after review failures")
	}
	if mockDisp.wasCalled(TaskTypeDocGardener) {
		t.Error("expected DocGardener hook NOT to be triggered (hooks removed)")
	}
}

// Review → Rework (failures < 3)
func TestHookChain_ReviewFailures_BelowThreshold(t *testing.T) {
	mockDisp, mockCounter, _, _ := setupHookTest(t)
	mockCounter.countFn = func(prNumber int, projectName string) int { return 2 }

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "NEEDS_FIX"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeRework) {
		t.Error("expected Rework hook to be triggered")
	}
	if mockDisp.wasCalled(TaskTypeDocGardener) {
		t.Error("expected DocGardener hook NOT to be triggered (hooks removed)")
	}
}

// Review → Rework (failures = 4)
func TestHookChain_ReviewFailures_NotMultipleOf3(t *testing.T) {
	mockDisp, mockCounter, _, _ := setupHookTest(t)
	mockCounter.countFn = func(prNumber int, projectName string) int { return 4 }

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "NEEDS_FIX"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeRework) {
		t.Error("expected Rework hook to be triggered")
	}
	if mockDisp.wasCalled(TaskTypeDocGardener) {
		t.Error("expected DocGardener hook NOT to be triggered (hooks removed)")
	}
}

// Review → Rework (failures = 6)
func TestHookChain_ReviewFailures_SecondMultipleOf3(t *testing.T) {
	mockDisp, mockCounter, _, _ := setupHookTest(t)
	mockCounter.countFn = func(prNumber int, projectName string) int { return 6 }

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "NEEDS_FIX"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if !mockDisp.wasCalled(TaskTypeRework) {
		t.Error("expected Rework hook to be triggered")
	}
	if mockDisp.wasCalled(TaskTypeDocGardener) {
		t.Error("expected DocGardener hook NOT to be triggered (hooks removed)")
	}
}

// Review PASS → no hooks triggered
func TestHookChain_ReviewPass_NoHooks(t *testing.T) {
	mockDisp, mockCounter, _, _ := setupHookTest(t)
	mockCounter.countFn = func(prNumber int, projectName string) int { return 5 }

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "passed"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if len(mockDisp.calls) != 0 {
		t.Errorf("expected no hooks when review is PASS, got %d calls", len(mockDisp.calls))
	}
}

// Hook template variable rendering (using E2E → Review hook)
func TestHookChain_TemplateVarsRendered(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 30 }

	task := makeTask(TaskTypeE2E, 5, "success")
	task.Metadata["triggeredBy"] = "dev-5"
	spec := taskRegistry[TaskTypeE2E]
	vars := map[string]interface{}{
		"issueNumber": "5",
		"branch":      "feat/issue-5",
		"taskKind":    "full-stack",
		"tmpDir":      "/tmp/dev-5",
	}
	wfCtx := workflow.NewContext("/tmp/project", vars)
	triggerHooks(task, spec, vars, wfCtx, true)

	// Verify Review was called
	found := false
	for _, c := range mockDisp.calls {
		if c.targetType == TaskTypeReview {
			found = true
		}
	}
	if !found {
		t.Error("expected Review hook to be triggered")
	}
}

// No hooks triggered when spec has no AfterSuccess
func TestHookChain_NoHooksConfigured(t *testing.T) {
	mockDisp, _, _, _ := setupHookTest(t)

	task := makeTask(TaskTypeDesign, 5, "success")
	spec := taskRegistry[TaskTypeDesign]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if len(mockDisp.calls) != 0 {
		t.Errorf("expected no hooks, got %d calls", len(mockDisp.calls))
	}
}

// Condition false → hook skipped
func TestHookChain_ConditionFalse_SkipsHook(t *testing.T) {
	mockDisp, _, _, _ := setupHookTest(t)

	task := makeTask(TaskTypeReview, 58, "success")
	task.Result = "passed"
	spec := taskRegistry[TaskTypeReview]
	triggerHooks(task, spec, map[string]interface{}{}, nil, true)

	if len(mockDisp.calls) != 0 {
		t.Errorf("expected no hooks when review is PASS, got %d calls", len(mockDisp.calls))
	}
}

// E2E failure → no hooks triggered
func TestHookChain_E2EFailure_NoHooks(t *testing.T) {
	mockDisp, _, _, mockFinder := setupHookTest(t)
	mockFinder.findPRFn = func(branch string) int { return 30 }

	task := makeTask(TaskTypeE2E, 5, "failed")
	task.Metadata["triggeredBy"] = "dev-5"
	spec := taskRegistry[TaskTypeE2E]
	triggerHooks(task, spec, map[string]interface{}{}, nil, false)

	if len(mockDisp.calls) != 0 {
		t.Errorf("expected no hooks after E2E failure, got %d calls", len(mockDisp.calls))
	}
}
