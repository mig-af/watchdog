package main

import (
	"dog/src/config"
	"dog/src/dog"
	"dog/src/funcs"
	
	"dog/src/telegram"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	_ "time/tzdata"
)


const (
	RED = "\033[1;31m"
	GREEN = "\033[1;32m"
	YELLOW = "\033[1;33m"
	END = "\033[0m"
	
)

func main(){
	
	funcs.Cmd("clear")
	funcs.CheckDir()
	telegram.VerifyUserStatus()
	

	
	var filesToAnalize []string
	ch := make(chan os.Signal, 1) //channel

	json := dog.ReadDocs()


	var timezone string = json.Timezone
	location, er := time.LoadLocation(timezone)
	if(er != nil){
		location = time.Local
	}




	files, errFile := dog.GetFiles(json.File.Paths)
	InfoDirFiles, errDir := dog.GetFilesFromDir(json.Dir.Paths)
	if(errFile != nil){fmt.Println(errFile.Error())}
	if(errDir != nil){fmt.Println(errDir.Error())}
	

	time.Sleep(4 * time.Second)
	funcs.Cmd("clear")
	filesToAnalize = funcs.ListCleaner(append(InfoDirFiles.Files, files...))


	
	fmt.Print("\n\n\033[1;101m		MONITOREANDO ", END)
	banner := fmt.Sprintf(`
		%sCarpetas %s: %d
		%sArchivos %s: %d

		%sTotal Archivos %s: %s %d  %s
		%sRuta del archivo de informacion %s :  %s 

	
	`,  YELLOW,
		END,
		InfoDirFiles.TotalDirs,
		YELLOW,
		END,
	 	len(files),

		YELLOW,
		END,

		GREEN,
		len(filesToAnalize),
		END,

		YELLOW,
		END,
		
		config.LOGFILE,
		
	)

	
	fmt.Println(banner)
	
	//Aca esperamsos el ctrl+c
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)


	for i:=0 ; i<len(filesToAnalize);i++{
		go dog.Analize(filesToAnalize[i], location)

	}
	
	for carpeta:=0 ; carpeta <len(InfoDirFiles.Dirs); carpeta ++{
		go dog.Analize(InfoDirFiles.Dirs[carpeta], location)
	}



	exit := <- ch 
	
	fmt.Println("Programa cerrado", exit)
	//select{}
	
}








