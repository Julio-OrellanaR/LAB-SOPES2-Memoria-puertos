package main

import (
	"context"
	"fmt"
	"syscall"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CPU() float64 {
	usage, err := getCpuUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return usage
}

func (a *App) DISK() float64 {
	usage, err := getMemoryUsage()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return usage
}

func getCpuUsage() (float64, error) {
	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		return 0, err
	}

	total := info.Totalram * uint64(info.Unit)
	free := info.Freeram * uint64(info.Unit)
	usage := 100 - float64(free)/float64(total)*100

	return usage, nil
}

func getMemoryUsage() (float64, error) {
	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		return 0, err
	}

	total := info.Totalram * uint64(info.Unit)
	free := info.Freeram * uint64(info.Unit)
	used := total - free
	usage := float64(used) / float64(total) * 100

	return usage, nil
}
