# Detection Software

This directory contains the Python-based detection engine for Medusa. It uses YOLOv8 for real-time person detection.

## Prerequisites

- Python 3.8 or higher
- `pip` (Python package installer)

## Setup Instructions

1.  (Recommended) Create and activate a virtual environment:
    ```bash
    python3 -m venv venv
    source venv/bin/activate  # On macOS/Linux
    # venv\Scripts\activate  # On Windows
    ```

2.  Install dependencies:
    ```bash
    pip install -r requirements.txt
    ```

3.  Ensure the YOLO model weight file (`yolov8n.pt`) is present in the root directory of the project.

## Usage

The script is typically called by the Medusa Go backend, but it can be run manually for testing:

```bash
python3 main.py <source> <camera_name> [fps]
```

- `<source>`: Video source (e.g., `0` for webcam, `video.mp4` for a file, or `rtsp://...` for a stream).
- `<camera_name>`: A unique name for the camera.
- `[fps]`: (Optional) Throttling limit for detections.

### Example

```bash
python3 main.py 0 cam1 2.0
```

## Integration

The script outputs detections in JSON format to `stdout`. Medusa captures this output to log events:

```json
{"Camera": "cam1", "Time": 1679999999.0, "Event": "Human Detected", "Confidence": 0.85}
```
