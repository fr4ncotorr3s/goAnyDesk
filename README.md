# goAnyDesk

A Go program to monitor and log various AnyDesk activities, including connections, file transfers, and system events. It checks functions like active connections, incoming/outgoing connections, IPs, connection durations, file transfers, and more.

`goAnyDesk` is a simple command-line tool to call various functions related to AnyDesk data, including connections, file transfers, tunnels, and more. It supports both **flags** for shortcuts and **positional arguments** for explicit function calls.

## Installation

### Windows

1. Clone the repository or download the `goAnyDesk` source code.
2. Build the Go program:
   ```bash
   go build -o goAnyDesk.exe
   ```
3. You can now use the `goAnyDesk.exe` executable to call the functions.

### Linux

1. Clone the repository or download the `goAnyDesk` source code.
   ```bash
   git clone https://github.com/fr4ncotorr3s/goAnyDesk.git
   ```
2. Navigate to the project directory:
   ```bash
   cd goAnyDesk
   ```
3. Build the Go program:
   ```bash
   go build -o goAnyDesk
   ```
4. After building, you can use the `goAnyDesk` executable to call the functions.

   If you want to make it easier to run from anywhere, you can move it to a directory in your `$PATH`:

   ```bash
   sudo mv goAnyDesk /usr/local/bin/
   ```

## Usage

You can rename it to `goad` or anything you like, or simply keep it as is.

### Basic Command Structure

```bash
./goAnyDesk <function-name>
```

Where `<function-name>` is the name of the function you wish to execute.

### Available Functions (with Shortcut Flags)

You can either specify a **function name** as an argument or use **short flags** to quickly execute the corresponding function. Below is a list of available functions and their associated flags.

#### Function List:

| Function Name                      | Flag   | Description                                     |
| ---------------------------------- | ------ | ----------------------------------------------- |
| `GetADConnections`                 | `-c`   | Get AnyDesk connections.                        |
| `GetADOutgoingConnections`         | `-o`   | Get outgoing AnyDesk connections.               |
| `GetADIncomingConnections`         | `-i`   | Get incoming AnyDesk connections.               |
| `GetADIncomingConnectionIPs`       | `-ip`  | Get IPs of incoming AnyDesk connections.        |
| `GetADUniqueIncomingConnectionIPs` | `-uip` | Get unique IPs of incoming AnyDesk connections. |
| `GetADConnectionDurations`         | `-d`   | Get the duration of AnyDesk connections.        |
| `GetADTunnel`                      | `-t`   | Get AnyDesk tunnels.                            |
| `GetADTunnelConnections`           | `-tc`  | Get connections for AnyDesk tunnels.            |
| `GetADForwardedTunnels`            | `-ft`  | Get forwarded AnyDesk tunnels.                  |
| `GetADFileTransferTo`              | `-fto` | Get file transfers from AnyDesk.                |
| `GetADFileTransferFrom`            | `-ftf` | Get file transfers to AnyDesk.                  |
| `GetADKeyboardLayout`              | `-kl`  | Get the keyboard layout used in AnyDesk.        |
| `GetADFileCopy`                    | `-fc`  | Get file copy actions in AnyDesk.               |
| `GetADFilePaste`                   | `-fp`  | Get file paste actions in AnyDesk.              |

### Examples

#### 1. Using Flags

You can use the flags to quickly call a function.

- To get AnyDesk connections:

  ```bash
  ./goAnyDesk -c
  ```

- To get AnyDesk outgoing connections:

  ```bash
  ./goAnyDesk -o
  ```

- To get forwarded AnyDesk tunnels:
  ```bash
  ./goAnyDesk -ft
  ```

#### 2. Using Positional Argument

You can also specify the function name explicitly.

- To get AnyDesk connections:

  ```bash
  ./goAnyDesk GetADConnections
  ```

- To get unique incoming connection IPs:
  ```bash
  ./goAnyDesk GetADUniqueIncomingConnectionIPs
  ```

#### 3. Using Both

You can combine using flags and positional arguments. The flags will take priority when both are specified.

- To get outgoing connections using the `-o` flag:

  ```bash
  ./goAnyDesk -o
  ```

- To get incoming connections using the positional argument:
  ```bash
  ./goAnyDesk GetADIncomingConnections
  ```

### Error Handling

If an invalid function name is provided or an unsupported flag is used, the program will return an error and exit with a message indicating the issue.

### Building and Running

1. Clone the repository and navigate to the project directory:

   ```bash
   git clone https://github.com/fr4ncotorr3s/goAnyDesk.git
   cd goAnyDesk
   ```

2. Run the following command to build the program:

   ```bash
   go build -o goAnyDesk
   ```

3. After building, you can use the executable:

   ```bash
   ./goAnyDesk <function-name>
   ```

4. Optionally, move the binary to `/usr/local/bin/` to make it accessible from anywhere:
   ```bash
   sudo mv goAnyDesk /usr/local/bin/
   ```

### Dependencies

This program uses Go's standard library, so there are no additional dependencies.
