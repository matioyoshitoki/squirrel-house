package main

// HookTaskDispatcher 负责将 hook 分派到具体的任务启动函数。
// 生产环境使用 defaultHookDispatcher，测试时替换为 mock。
type HookTaskDispatcher interface {
	Dispatch(targetType TaskType, sourceTask *Task, vars map[string]interface{})
}

// defaultHookDispatcher 生产环境实现
type defaultHookDispatcher struct{}

func (d *defaultHookDispatcher) Dispatch(targetType TaskType, sourceTask *Task, vars map[string]interface{}) {
	switch targetType {
	case TaskTypeE2E:
		startE2ETaskFromHook(sourceTask, vars)
	case TaskTypeReview:
		startReviewTaskFromHook(sourceTask, vars)
	case TaskTypeRework:
		startReworkTaskFromHook(sourceTask, vars)
	}
}

// hookDispatcher 全局分派器，可在测试中替换
var hookDispatcher HookTaskDispatcher = &defaultHookDispatcher{}

// reviewFailureCounter 抽象 countReviewFailuresFromLogs，便于测试注入
type reviewFailureCounter interface {
	count(prNumber int, projectName string) int
}

type defaultReviewFailureCounter struct{}

func (d *defaultReviewFailureCounter) count(prNumber int, projectName string) int {
	return countReviewFailuresFromLogs(prNumber, projectName)
}

var reviewFailureCounterImpl reviewFailureCounter = &defaultReviewFailureCounter{}

// maestroFlowChecker 抽象 hasMaestroFlows，便于测试注入
type maestroFlowChecker interface {
	hasFlows(branch string) bool
}

type defaultMaestroFlowChecker struct{}

func (d *defaultMaestroFlowChecker) hasFlows(branch string) bool {
	return hasMaestroFlows(branch)
}

var maestroFlowCheckerImpl maestroFlowChecker = &defaultMaestroFlowChecker{}

// prFinder 抽象 findPRForBranch，便于测试注入
type prFinder interface {
	findPR(branch string) int
}

type defaultPRFinder struct{}

func (d *defaultPRFinder) findPR(branch string) int {
	return findPRForBranch(branch)
}

var prFinderImpl prFinder = &defaultPRFinder{}
