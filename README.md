# OpenTelemetry NoOp Implementation

This repository provides a simplified implementation of OpenTelemetry's API and SDK to demonstrate the concept of **NoOp implementations**. The purpose is to showcase how API-SDK separation allows the system to work seamlessly, even when the SDK is uninitialized.

## Outcomes

- **API-SDK Separation**: The instrumented code depends only on the API, not the SDK.
- **No-Op Default**: Ensures the API works even if the SDK is not initialized.
- **Runtime SDK Initialization**: The application can switch to the real SDK during runtime.

## How It Works

1. The `api` package provides a global `TracerProvider` that defaults to a no-op implementation.
2. Instrumented libraries use the `api` package to create spans without worrying about SDK availability.
3. The application can initialize the SDK at runtime to provide real tracing functionality.

## Folder Structure
```bash
ziadh@Ziads-MacBook-Air no-op % tree
.
├── README.md
├── interfaces
│   └── interfaces.go
├── api
│   └── api.go
├── sdk
│   └── sdk.go
├── go.mod
├── main.go
└── library
    └── library.go

4 directories, 7 files
```

## Sample Output

```bash 
ziadh@Ziads-MacBook-Air no-op % go run .
Running without SDK initialized:
Doing some work in the instrumented function...

Initializing SDK...
Span started: instrumented_function
Doing some work in the instrumented function...
Span ended: instrumented_function
```
