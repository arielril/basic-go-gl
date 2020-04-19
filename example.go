package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"

	glfw "github.com/go-gl/glfw/v3.3/glfw"
)

const (
	wHeight = 500
	wWidth  = 650
)

func drawLines() {
	gl.Begin(gl.LINES)
	gl.Vertex2f(0, 0)
	gl.Vertex2f(5, 5)
	gl.End()

	gl.Begin(gl.LINES)
	gl.Vertex2f(5, 5)
	gl.Vertex2f(10, 0)
	gl.End()
}

func display(w *glfw.Window) {
	gl.LineWidth(8)

	gl.Color3f(1, 0, 0)
	drawLines()

	gl.Translatef(1, 1, 0)

	gl.Color3f(0, 0, 1)
	drawLines()

	gl.Translatef(2, 0, 0)
	gl.Color3f(0, 0, 0)
	drawLines()
}

func init() {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("failed to init glfw: %v", err))
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
}

func main() {
	win, err := glfw.CreateWindow(wWidth, wHeight, "Sup!", nil, nil)
	if err != nil {
		panic(fmt.Errorf("failed to create an window: %v", err))
	}

	win.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		panic(fmt.Errorf("failed to start gl: %v", err))
	}

	win.SetKeyCallback(keyCallback)
	win.SetCharCallback(charCallback)

	setup()
	for !win.ShouldClose() {
		reshape(win)
		display(win)

		win.SwapBuffers()
		glfw.PollEvents()
	}
}
