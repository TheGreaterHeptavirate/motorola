/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package animations

import (
	"github.com/AllenDang/imgui-go"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"time"
)

type TransitionWidget struct {
	a *animationWidget
}

func Transition(renderer1, renderer2 func(Animation)) *TransitionWidget {
	return &TransitionWidget{
		a: newAnimation(renderer1, renderer2),
	}
}

func (t *TransitionWidget) Start(d time.Duration, fps int) {
	t.a.Start(d, fps)
}

func (t *TransitionWidget) Advance() bool {
	return t.Advance()
}

func (t *TransitionWidget) Build() {
	state := t.a.GetState()
	if t.a.BuildNormal(t) {
		return
	}

	if state.layout1ProcentageAlpha > 1 {
		logger.Fatalf("animationWidget: procentage alpha is %v (should be in range 0-1)", state.layout1ProcentageAlpha)
	}

	imgui.PushStyleVarFloat(imgui.StyleVarAlpha, state.layout1ProcentageAlpha)
	t.a.renderer1(t)
	imgui.PopStyleVar()
	imgui.PushStyleVarFloat(imgui.StyleVarAlpha, (1 - state.layout1ProcentageAlpha))
	t.a.renderer2(t)
	imgui.PopStyleVar()
}
