package backend

import (
	"fmt"
	"strings"
)

func (a *App) WipeData() error {
	output, err := a.runCommand("fastboot", "-w")
	if err != nil {
		return fmt.Errorf("failed to run fastboot -w: %w. Output: %s", err, output)
	}
	return nil
}

func (a *App) FlashPartition(partition string, filePath string) error {
	if partition == "" || filePath == "" {
		return fmt.Errorf("partition and file path cannot be empty")
	}

	if err := a.validateFlashPartition(partition); err != nil {
		return err
	}

	if err := a.validateFlashFile(filePath); err != nil {
		return err
	}

	output, err := a.runCommand("fastboot", "flash", partition, filePath)
	if err != nil {
		return fmt.Errorf("failed to run fastboot flash: %w. Output: %s", err, output)
	}
	return nil
}

func (a *App) GetFastbootDevices() ([]Device, error) {
	output, err := a.runCommand("fastboot", "devices")
	if err != nil {
		if output == "" {
			return []Device{}, nil
		}
		return nil, err
	}

	var devices []Device
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) >= 2 && parts[1] == "fastboot" {
			devices = append(devices, Device{
				Serial: parts[0],
				Status: parts[1],
			})
		}
	}

	return devices, nil
}

func (a *App) RunFastbootHostCommand(args string) (string, error) {
	if args == "" {
		return "", fmt.Errorf("command cannot be empty")
	}

	argSlice := strings.Fields(args)

	if err := a.validateHostCommand(argSlice); err != nil {
		return "", err
	}

	output, err := a.runCommand("fastboot", argSlice...)
	if err != nil {
		return "", fmt.Errorf("command failed: %w. Output: %s", err, output)
	}

	return output, nil
}

func (a *App) FlashRomFolder(serial string, folderPath string, plan FlashPlan) error {
	if folderPath == "" || len(plan.Steps) == 0 {
		return fmt.Errorf("invalid flash plan or empty folder path")
	}

	for i, step := range plan.Steps {
		a.logCommand("fastboot", []string{"-s", serial, "flash", step.Partition, step.ImageFile}, 0, nil, fmt.Sprintf("Starting step %d: Flashing %s", i+1, step.Partition))

		err := a.FlashPartition(step.Partition, step.ImageFile)
		if err != nil {
			return fmt.Errorf("failed at step %d (partition %s): %w", i+1, step.Partition, err)
		}
	}

	return nil
}

func (a *App) GetFastbootSlot(serial string) (string, error) {
	return "a", nil
}

func (a *App) SetFastbootSlot(serial string, slot string) error {
	return nil
}
