package backend

import (
	"context"
	"fmt"
	"sync"
)

type DeviceMode string

const (
	DeviceModeUnknown  DeviceMode = "unknown"
	DeviceModeADB      DeviceMode = "adb"
	DeviceModeFastboot DeviceMode = "fastboot"
)

type Device struct {
	Serial string
	Status string
}
type DeviceInfo struct {
	Model          string
	AndroidVersion string
	BuildNumber    string
	BatteryLevel   string
	Serial         string
	IPAddress      string
	RootStatus     string
	Codename       string
	RamTotal       string
	StorageInfo    string
	Brand          string
	DeviceName     string
}
type FileEntry struct {
	Name        string
	Type        string
	Size        string
	Permissions string
	Date        string
	Time        string
}

type PackageInfo struct {
	PackageName string
	IsEnabled   bool
}

type FlashPlan struct {
	Steps []FlashStep `json:"steps"`
}

type FlashStep struct {
	Partition string `json:"partition"`
	ImageFile string `json:"imageFile"`
}

type PerformanceSnapshot struct {
	CPUUsage     float64 `json:"cpuUsage"`
	RAMUsage     float64 `json:"ramUsage"`
	NetworkRxSec float64 `json:"networkRxSec"`
	NetworkTxSec float64 `json:"networkTxSec"`
}

type App struct {
	ctx         context.Context
	binaryCache map[string]string
	cacheMutex  sync.RWMutex

	currentCancel context.CancelFunc
	opMutex       sync.Mutex
}

func NewApp() *App {
	return &App{
		binaryCache: make(map[string]string),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CancelOperation() string {
	a.opMutex.Lock()
	defer a.opMutex.Unlock()

	if a.currentCancel != nil {
		a.currentCancel()
		a.currentCancel = nil
		return "Operation cancelled."
	}
	return "No active operation to cancel."
}
