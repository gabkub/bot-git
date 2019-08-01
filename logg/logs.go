package logg

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"strings"
	"time"
)

func WriteToFile(context string){
	f, err := os.OpenFile(fmt.Sprintf("./log_%s.log",time.Now().Format("2019-06-12")), os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil{
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)

	logger.Println(strings.TrimSpace(context))
}

func SetOutPut(){
	log.SetOutput(&lumberjack.Logger{
		Filename:  fmt.Sprintf("./log_%s.log", time.Now().Format("2019-06-12")),
		MaxSize:   1000,
		MaxAge:	   7,
	})
}