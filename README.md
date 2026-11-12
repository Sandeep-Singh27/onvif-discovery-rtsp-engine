# ONVIF Discovery & RTSP Stream Extraction Engine

A Go-based network scanning engine designed to discover and communicate with network video devices using the ONVIF (Open Network Video Interface Forum) standard. The project focuses on identifying ONVIF-compatible devices and extracting information required to access their video streams through RTSP.

## Overview

The repository is organized into three primary components:

- **Network Scanner (`scan-network`):** Responsible for the network scanning and device discovery workflow.
- **Sender (`sender`):** Handles outbound network communication associated with the discovery process.
- **Listener (`listener`):** Handles incoming network communication and discovery responses from network video devices.

Together, these components form a modular foundation for discovering ONVIF-compatible network video devices and communicating with them to facilitate RTSP stream extraction.

## Architecture

```text
              ONVIF DISCOVERY ENGINE
                       │
                       ▼
                Network Scanner
                 scan-network
                       │
                ┌──────┴──────┐
                │             │
                ▼             ▼
             Sender        Listener
                │             │
                ▼             │
        Send Discovery Probe  │
                │             │
                ▼             │
          Network Video       │
             Devices          │
                │             │
                └─────────────┤
                              ▼
                     Receive Discovery
                         Responses
                              │
                              ▼
                     Identify ONVIF
                         Devices
                              │
                              ▼
                    Device Communication
                              │
                              ▼
                     RTSP Stream Details
```

## Technology Stack

| Technology | Purpose |
|---|---|
| Go | Core implementation |
| ONVIF | Network video device discovery and communication |
| RTSP | Video stream access protocol |
| Network Sockets | Network communication |

## Project Structure

```text
onvif-discovery-rtsp-engine/
│
├── scan-network/    # Network scanning and device discovery
│
├── sender/          # Outbound network communication
│
├── listener/        # Incoming network communication and responses
│
├── go.mod           # Go module definition and version
│
├── .gitignore       # Git exclusion rules
│
└── README.md        # Project documentation
```

