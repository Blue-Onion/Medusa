# Medusa

Medusa is a multi-camera incident monitoring and event logging platform written in Go. It is designed as a pluggable event pipeline where varying detection engines (such as Python-based AI models) can be seamlessly integrated. The platform itself abstracts away the complexity of handling configuration, concurrency, logging, and querying, allowing you to focus on building robust detection capabilities.

## Features

- **Concurrent Multi-Camera Processing**: Run and monitor multiple camera streams simultaneously using native Go concurrency.
- **Flexible Stream Support**: Configure inputs from webcam devices, local video files, and RTSP streams.
- **Pluggable Detection Architecture**: Integrate any external detection software (interchangeable with your choice of detection models/engines) via standard output event emitting.
- **Structured JSON Logging**: All events are persisted as JSON Lines (JSONL), making them easy to parse and query.
- **Built-in CLI Tool**: Start the pipeline, view configurations, query historical logs, and export events to CSV.

## Prerequisites

- **Go**: Version 1.21 or higher (Check with `go version`).
- **Python**: Version 3.8 or higher (Check with `python3 --version`).
- **Python Dependencies**: `ultralytics` and `opencv-python`.

## Setup Guide

### 1. Python Environment (Detection Engine)

The default detection engine is based on YOLOv8.

```bash
# Navigate to detection directory
cd DetectionSoftware

# Install dependencies
pip install -r requirements.txt

# Or if using a virtual environment (recommended)
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt

# Return to root directory
cd ..
```

### 2. Medusa Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/medusa.git
cd medusa

# Build the executable
go build -o medusa ./cmd/main.go
```

## Configuration

Medusa is configured via `config.yaml`. This file defines camera sources, log paths, and FPS presets.

### `config.yaml` Structure

```yaml
cameras:
  - name: cam1
    source: 0       # Webcam index (numeric)
  - name: cam2
    source: "video.mp4" # Local video file
  - name: cam3
    source: "rtsp://admin:pass@192.168.1.20:554/stream" # RTSP stream

# Directory where JSONL logs will be stored
recordsPath: "logs"

# FPS Presets for the detection engine
fps:
  low: "2"    # 2 Frames Per Second
  medium: "5"  # 5 Frames Per Second
  high: "15"   # 15 Frames Per Second
```

## CLI Usage

### Start Processing

Start the monitoring pipeline for all configured cameras.

```bash
./medusa start [-m mode]
```

- `-m`: (Optional) FPS mode. Options: `low`, `medium`, `high`. Defaults to `medium` if not specified.

### View Configuration

Display the loaded camera and system configurations.

```bash
./medusa show-config
```

### Query Logs

Search for recorded events by date and camera name.

```bash
./medusa show-record --date YYYY-MM-DD [--cam camera_name]
```

### Export to CSV

Export events to a CSV file for external analysis.

```bash
./medusa download --date YYYY-MM-DD [--cam camera_name]
```

## Log Storage Structure

Logs are organized hierarchically by date and camera name.

```text
logs/
└── YYYY-MM-DD/
    ├── cam1.log
    ├── cam2.log
    └── cam3.log
```

## Architecture

Medusa acts as an orchestrator. It spawns an isolated Python process (the detection software) for each camera. The Go backend captures the `stdout` of these processes, parses the JSON events, and persists them to the filesystem.

```text
Camera Source -> [ Go Routine ] -> [ Python Process ] -> [ JSON Event ] -> [ Go Logger ] -> FileSystem
```

For detailed information on the detection engine, see [DetectionSoftware/README.md](DetectionSoftware/README.md).

## License

[MIT License](LICENSE)
