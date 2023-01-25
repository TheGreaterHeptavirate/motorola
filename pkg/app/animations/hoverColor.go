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
	"image/color"
	"time"
)

type animationData struct {
	procentage float32
	isHovered  bool
}

var _ Animation = &HoverColorAnimationWidget{}

type HoverColorAnimationWidget struct {
	*animationWidget
	giu.Widget
	hoveredColor imgui.Vec4
	fps          int
	duration     time.Duration
}

func HoverColorAnimation(widget giu.Widget, fps int, duration time.Duration, hoveredColor color.RGBA) *HoverColorAnimationWidget {
	result := &HoverColorAnimationWidget{
		Widget:       widget,
		hoveredColor: giu.ToVec4Color(hoveredColor),
		fps:          fps,
		duration:     duration,
	}

	result.animationWidget = newAnimation(result, nil, nil)

	return result
}

func (h *HoverColorAnimationWidget) Reset() {
	newData := &animationData{}
	//d := h.animationWidget.GetCustomData()
	//currentData, ok := d.(*animationData)
	//if !ok {
	//	logger.Fatalf("expected data type *animationData, got %T", d)
	//}

	h.animationWidget.SetCustomData(newData)
}

func (h *HoverColorAnimationWidget) Advance(procentDelta float32) bool {
	d := h.animationWidget.GetCustomData()
	data, ok := d.(*animationData)
	if !ok {
		logger.Fatalf("expected data type *animationData, got %T", d)
	}

	data.procentage = procentDelta

	return true
}

func (h *HoverColorAnimationWidget) Init() {
	h.animationWidget.SetCustomData(&animationData{})
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

	procentage := data.procentage
	normalColor := imgui.CurrentStyle().GetColor(imgui.StyleColorButton)

	if !data.isHovered && procentage != 0 {
		procentage = 1 - procentage
	}

	normalColor.X += (h.hoveredColor.X - normalColor.X) * procentage
	normalColor.Y += (h.hoveredColor.Y - normalColor.Y) * procentage
	normalColor.Z += (h.hoveredColor.Z - normalColor.Z) * procentage

	imgui.PushStyleColor(imgui.StyleColorButton, normalColor)
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, normalColor)
	h.Widget.Build()
	imgui.PopStyleColorV(2)
	isHovered := imgui.IsItemHovered()

	if isHovered != data.isHovered && !h.animationWidget.GetState().isRunning {
		h.Start(h.duration, h.fps)
	}

	data.isHovered = isHovered
}
