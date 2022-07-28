package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	config := viper.New()

	fmt.Printf("%#v", config)

}
