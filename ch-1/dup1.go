package main
import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++
	}
	//Note: ignoring potential error from input.Err()
	for line, n:= range counts{
		if n > 1{
			fmt.Println()
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}