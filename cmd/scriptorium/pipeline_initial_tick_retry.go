package main

import (
	"strings"

	"github.com/prantosmss/Scriptorium/internal/entry/headless"
	"github.com/prantosmss/Scriptorium/internal/store"
	"github.com/prantosmss/Scriptorium/internal/tools"
)

// Capture before ResetActivityState removes the failed attempt. The feedback
// is process-local; it does not revise sources, the exact chapter contract, or
// the old tick into accepted evidence.
func pipelineInitialWorldTickAttemptOptions(outputDir, prompt string) headless.Options {
	opts := pipelineInitialWorldTickHeadlessOptions(prompt)
	st := store.NewStore(outputDir)
	tick, err := st.WorldSim.LoadTick()
	if err != nil || tick == nil || tick.TickID == "" || tick.TickID == "v0-a0" || tick.EventCount <= 0 {
		return opts
	}
	issues := tools.InitialWorldTickQualityIssues(st)
	if len(issues) == 0 {
		return opts
	}
	opts.InitialWorldTickRetryFeedback = "[initial world_tick 上次质量拒绝]\n以下为宿主对上次提交的实际校验结果，请修正对应条件；不改变原始逐字合同，不预定角色选择，不把未发生的结果补写为事实。\n- " + strings.Join(issues, "\n- ")
	opts.Prompt += "\n\n" + opts.InitialWorldTickRetryFeedback
	return opts
}
