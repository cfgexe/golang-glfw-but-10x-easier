package glfw

import (
    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
)

type Windower interface {
	createWindow()
}

func (win Window) createWindow(winName, winH, winW) {
	runtime.LockOSThread()

    window := initGlfw(winName, winH, winW)
    defer glfw.Terminate()

    for !window.ShouldClose() {
        // TODO
    }
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw(winName, winH, winW) *glfw.Window {
    if err := glfw.Init(); err != nil {
            panic(err)
    }
    
    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(winW, winH, winName, nil, nil)
    if err != nil {
            panic(err)
    }
    window.MakeContextCurrent()

    return window
}
