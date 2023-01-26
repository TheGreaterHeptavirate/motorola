/*
 * Copyright (c) 2023 The Greater Heptavirate team (https://github.com/TheGreaterHeptavirate)
 * All Rights Reserved
 *
 * All copies of this software (if not stated otherway) are dedicated
 * ONLY to personal, non-commercial use.
 */

package animations

import (
	"github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/TheGreaterHeptavirate/motorola/internal/logger"
	"sync"
	"time"
)

type animationData struct {
	procentage  float32
	isHovered   bool
	shouldStart bool
	m           *sync.Mutex
}

var _ Animation = &HoverColorAnimationWidget{}

type HoverColorAnimationWidget struct {
	*animationWidget
	giu.Widget
	hoveredColor imgui.Vec4
	fps          int
	duration     time.Duration
}

func HoverColorAnimation(widget giu.Widget, fps int, duration time.Duration) *HoverColorAnimationWidget {
	result := &HoverColorAnimationWidget{
		Widget: widget,
		//hoveredColor: giu.ToVec4Color(hoveredColor),
		fps:      fps,
		duration: duration,
	}

	result.animationWidget = newAnimation(result, nil, nil)

	return result
}

func (h *HoverColorAnimationWidget) Reset() {
	d := h.animationWidget.GetCustomData()

	currentData, ok := d.(*animationData)
	if !ok {
		logger.Fatalf("expected data type *animationData, got %T", d)
	}

	if currentData == nil {
		h.animationWidget.SetCustomData(&animationData{
			m: &sync.Mutex{},
		})

		return
	}

	currentData.m.Lock()
	currentData.procentage = 0
	currentData.m.Unlock()
}

func (h *HoverColorAnimationWidget) Advance(procentDelta float32) bool {
	d := h.animationWidget.GetCustomData()
	data, ok := d.(*animationData)
	if !ok {
		logger.Fatalf("expected data type *animationData, got %T", d)
	}

	data.m.Lock()
	data.procentage = procentDelta
	data.m.Unlock()

	return true
}

func (h *HoverColorAnimationWidget) Init() {
	h.animationWidget.SetCustomData(&animationData{
		m: &sync.Mutex{},
	})
}

func (h *HoverColorAnimationWidget) Build() {
	if h.GetState().shouldInit {
		h.Init()
		h.GetState().shouldInit = false
	}

	d := h.animationWidget.GetCustomData()
	data, ok := d.(*animationData)
	if !ok {
		logger.Fatalf("expected data type *animationData, got %T", d)
	}

	normalColor := imgui.CurrentStyle().GetColor(imgui.StyleColorButton)
	data.m.Lock()
	shouldStart := data.shouldStart
	isHovered := data.isHovered
	data.m.Unlock()

	if shouldStart {
		data.m.Lock()
		data.shouldStart = false
		data.m.Unlock()
		h.Start(h.duration, h.fps)
	}

	data.m.Lock()
	procentage := data.procentage
	data.m.Unlock()

	if !isHovered && h.animationWidget.GetState().isRunning {
		procentage = 1 - procentage
	}

	h.hoveredColor = imgui.CurrentStyle().GetColor(imgui.StyleColorButtonHovered)
	normalColor.X += (h.hoveredColor.X - normalColor.X) * procentage
	normalColor.Y += (h.hoveredColor.Y - normalColor.Y) * procentage
	normalColor.Z += (h.hoveredColor.Z - normalColor.Z) * procentage

	if !h.animationWidget.GetState().isRunning {
		if isHovered {
			normalColor = h.hoveredColor
		} else {
			normalColor = imgui.CurrentStyle().GetColor(imgui.StyleColorButton)
		}
	}

	imgui.PushStyleColor(imgui.StyleColorButton, normalColor)
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, normalColor)
	h.Widget.Build()
	imgui.PopStyleColorV(2)
	isHoveredNow := imgui.IsItemHovered()

	data.m.Lock()
	data.shouldStart = isHoveredNow != isHovered && !h.animationWidget.GetState().isRunning
	data.isHovered = isHoveredNow
	data.m.Unlock()
}
