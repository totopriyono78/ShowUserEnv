package main
import (
	"fmt"
	"os"
)

func main(){
	var (
		user string
		homedir string
	)

	user = os.Getenv("USER")
	homedir=os.Getenv("HOME")

	fmt.Printf("hallo user %s", user)
	fmt.Printf("\nanda ada di %s",homedir)
	fmt.Printf("\n")
}
