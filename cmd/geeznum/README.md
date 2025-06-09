# Ge'ez Number Converter

A simple CLI tool that converts integers to their Ge'ez numeral representation.

This tool supports multiple input methods, including command-line flags, positional arguments, and piped input from `stdin`.

# 📦 Installation

## Option 1: go install (Recommended)

If you're using Go 1.16 or newer, you can install directly from the source repository:

```
go install github.com/asherand/geez@latest
```

Make sure your $GOBIN (usually $HOME/go/bin) is in your system’s PATH, then run:

```
geeznum -n 1 -n 2 -n 3
```

## Option 2: Clone and Build Manually

```
git clone https://github.com/asherand/geez.git
cd geez
go build -o geeznum
./geeznum 10 20

```

# 🚀 Usage

```
Usage:
  geeznum -n 1 -n 2 -n 3
  geeznum 1 2 3
  echo "1 2 3" | geeznum

Convert one or more numbers to their Ge'ez numeral representation.
```

# 🧩 Dependencies

- github.com/asherand/geez/geeznum: This internal package handles the conversion logic. Ensure it's implemented correctly for the converter to work.
