# Sim Racing Fuel Calculator

A simple and efficient CLI tool for sim racing enthusiasts to calculate total fuel, laps, and stints required for endurance racing. Designed to run in Linux terminals.

---

## Features

- Calculates total fuel consumption based on race duration and lap time.
- Computes total laps and required stints.
- Includes options for extra fuel margin.
- Interactive mode for user-friendly input.
- Lightweight and fast.

---

## Installation

1. Ensure you have Go installed on your system. You can check by running:

   ```bash
   go version
   ```

If not installed, follow the instructions at [golang.org](https://golang.org).

2. Install the tool using `go install`:

   ```bash
   go install github.com/lacolle87/srfc@latest
   ```

3. Make sure the Go `bin` directory is in your `PATH`:

   ```bash
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

---

## Usage

### Basic Command

Run the tool with options:

```bash
racer-calculator -lt <lap_time_in_seconds> -rl <race_length_in_minutes> -fc <fuel_consumption_per_lap> -tc <tank_capacity> -ef <extra_fuel_in_laps>
```

### Example

```bash
racer-calculator -lt 90 -rl 120 -fc 3.5
```

Output:
```
Total fuel: 420.00
Total laps: 80
Stints: 4
```

### Interactive Mode

If you omit some or all flags, the program will prompt for input interactively:

```bash
racer-calculator
```

Example prompts:
```
Enter lap time (in seconds): 90
Enter race length (in minutes): 120
Enter fuel consumption per lap: 3.5
Enter fuel tank capacity: 110
Enter extra fuel needed in laps: 2
```

---

## Options

- `-lt`: Lap time in seconds (e.g., 90)
- `-rl`: Race length in minutes (e.g., 120)
- `-fc`: Fuel consumption per lap (e.g., 3.5)
- `-tc`: Fuel tank capacity (default: 110)
- `-ef`: Extra fuel needed in laps (default: 2)
- `-h`: Show help message

### Help

Display usage instructions:

```bash
racer-calculator -h
```

---

## Requirements

- Go 1.23 or newer
- Linux-based operating system

---

## Contributing

Feel free to fork this repository and submit pull requests with enhancements, bug fixes, or additional features.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.