package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		// panic(err)
		fmt.Println("Error loading .env file:", err)
        os.Exit(1)
	}

	cmd := exec.Command(
		"tern", 
		"migrate", 
		"--migrations", 
		"./internal/store/pgstore/migrations", 
		"--config", 
		"./internal/store/pgstore/migrations/tern.conf",
	)
	// if err := cmd.Run(); err != nil {
	// 	panic(err)
	// }
	 output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error running tern migrate: %v\n", err)
        fmt.Println(string(output))
        os.Exit(1)
    }

    fmt.Println("tern migrate completed successfully")
}
