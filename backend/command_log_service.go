package backend

type CommandLogEntry struct {
	ID        int64  `json:"id"`
	Timestamp string `json:"timestamp"`
	Serial    string `json:"serial"`
	Mode      string `json:"mode"`
	Command   string `json:"command"`
	Duration  int64  `json:"durationMs"`
	Status    string `json:"status"`
	Output    string `json:"outputPreview"`
}

func (a *App) GetCommandHistory(serial string, limit int) ([]CommandLogEntry, error) {
	return []CommandLogEntry{}, nil
}
