# ADBKit

A simple, modern GUI for ADB and Fastboot — built with **Wails** (Go + React) for speed, lightweight resource usage, and native performance.

---

## Features

**Dashboard & Devices**
- Unified device list with editable nicknames
- Global selected device context synced across modules
- Rich device info: battery, storage, RAM, root status, connection state
- Wireless ADB toggle and pairing via IP/Port

**Device Monitoring**
- Live CPU, RAM, and network usage tracking
- In-app monitoring without switching to external tools

**App Manager**
- Virtualized lists for smooth handling of thousands of packages
- Batch install, uninstall, enable, and disable
- Install local APKs or pull APKs from device
- Filter by User/System apps, sort by name or state

**File Explorer**
- Push, Pull, Rename, Delete, and Create Folders
- Unlimited timeout for large file transfers
- Retry support and clear batch result reporting
- Concurrent I/O with virtualized directory rendering

**Logcat & Observability**
- Real-time Logcat streaming with filter support
- Internal command log history across modules

**Fastboot & Flashing**
- Direct `adb` and `fastboot` shell access
- One-click reboot to System, Recovery, or Bootloader
- Flash images, ZIPs, and folder-based image sets
- Slot A/B management and targeted partition flashing

**Reliability & Safety**
- Strict guards for destructive actions
- Validation for flash flows and command execution
- Path sanitization for shell-based file operations
- Clean resource handling for long-running streams
- Modular, service-based backend architecture

**Packaging & Distribution**
- Windows and Linux releases bundle `adb` and `fastboot`
- Ready-to-use installer and Linux packages

---

## Screenshots

[See here](screenshots/README.md)

---

## Installation

1. Go to the [Releases](https://github.com/drenzzz/adb-gui-kit/releases) page
2. Download the latest release for your OS
3. Install or extract the package
4. Run the application

> **Note:** Official Windows and Linux builds include bundled platform tools. For manual builds, ensure `adb` and `fastboot` are available in your environment.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Core | [Wails v2](https://wails.io) |
| Backend | Go — Modular Service Architecture |
| Frontend | React (via Astro) + TypeScript |
| State | Custom React Hooks + Composition |
| UI | [shadcn/ui](https://ui.shadcn.com) + Tailwind CSS |

---

## Building from Source

**Prerequisites:** Go 1.21+, Node.js 18+, pnpm
```bash
# Install frontend dependencies
cd frontend && pnpm install && cd ..

# Development
wails dev

# Production build
wails build
```

---

## Troubleshooting

| Issue | Fix |
|---|---|
| Device not found | Enable USB Debugging; install required drivers |
| Unauthorized | Accept the RSA fingerprint prompt on your device |
| Wireless ADB not working | Ensure device and computer are on the same network |
| Linux USB access denied | Configure `udev` rules for your device |

---

## Contributing

Contributions are welcome — open an issue or submit a pull request.

**Repo:** https://github.com/Drenzzz/adb-gui-kit

---

## License

MIT