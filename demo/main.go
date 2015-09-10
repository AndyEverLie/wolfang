package main
import (
	"fmt"
	"github.com/AndyEverLie/wolfang"
)

func main() {
	fmt.Println("hello world.")
	spider := wolfang.NewSpider("demo")
	spider.Run()
}
