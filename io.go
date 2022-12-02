package friendly

import (
	"bufio"
	"fmt"
	"os"
)

// Gets user input from command-line (os.Stdin).
func Input(prompt string) (string, error) {
	fmt.Print(prompt)
	r := bufio.NewReader(os.Stdin)
	ans, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return ans, nil
}
