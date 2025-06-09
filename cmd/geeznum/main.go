package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/asherand/geez/geeznum"
)

type numListFlag []int

func (n *numListFlag) String() string {
	return fmt.Sprint(*n)
}

func (n *numListFlag) Set(value string) error {
	num, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*n = append(*n, num)
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  geeznum -n 1 -n 2 -n 3")
		fmt.Println("  geeznum 1 2 3")
		fmt.Println("  echo \"1 2 3\" | geeznum")
		fmt.Println("Convert one or more numbers to their Ge'ez numeral representation.")
	}

	var nums numListFlag
	flag.Var(&nums, "n", "Number(s) to convert to Ge'ez (can be repeated)")
	flag.Parse()

	// 1. Add positional args
	for _, arg := range flag.Args() {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid number: %v\n", arg)
			os.Exit(1)
		}
		nums = append(nums, num)
	}

	// 2. Handle stdin if no -n or positional args
	if len(nums) == 0 {
		info, _ := os.Stdin.Stat()
		if info.Mode()&os.ModeCharDevice == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fields := strings.Fields(scanner.Text())
				for _, field := range fields {
					num, err := strconv.Atoi(field)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Invalid number from stdin: %v\n", field)
						os.Exit(1)
					}
					nums = append(nums, num)
				}
			}
		} else {
			fmt.Println("No numbers specified.")
			flag.Usage()
			os.Exit(1)
		}
	}

	for _, num := range nums {
		fmt.Println(geeznum.NumToGeez(num))
	}
}
