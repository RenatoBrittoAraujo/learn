package async

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func LongRoutine(in chan string, done chan string) {
	for id := range in {
		fmt.Println("Go routine", id, "started")
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		result := ""
		for i := 0; i < len(id); i++ {
			result = result + string((int(id[i])*10)%60+40)
		}
		idnum := strings.Split(id, " --- ")[1]
		fmt.Println(id, "===", idnum)
		if idnum == "100" {
			done <- "I'm done with this...."
		}
	}
}
