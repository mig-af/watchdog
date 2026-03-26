package funcs

import (
	"bytes"
	"dog/src/config"
	"dog/src/mystruct"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

//Verifica si un archivo existe o no
func VerifyFile(filePath string)(string, bool){
	_, err := os.Stat(filePath)

	if(err != nil){
		return "", false
	}
	return filePath, true

}



//hace string.split
func SplitString(info string)[]string{
	er := strings.ReplaceAll(info, "\n", " ")
	return strings.Split(er, " ")
	
}

//elimina el ultimo elemento
func DeleteLastElement(list []string)[]string{
	return list[:len(list)-1]
}


//limpia espacios vacios etc de una lista 
func ListCleaner(list[]string)[]string{

	var newList []string

	for i:=0; i<len(list); i++{
		if(list[i] == ""){
			continue
		}else if(list[i]== " "){
			continue
		}else{
			newList = append(newList, list[i])
		}
	}

	return newList

}




func WriteLog(content string)bool{
	logFile :=  config.LOGFILE

	file, err := os.OpenFile(logFile, os.O_CREATE | os.O_RDWR  | os.O_APPEND, 0644)

	if(err != nil){
		fmt.Println(err.Error())
		return false
	}
	defer file.Close()
	file.Write([]byte("\n"+content))
	
	return true

}




func Cmd(command string){
	var outp bytes.Buffer

	cmd := exec.Command(command)
	cmd.Stdout = &outp

	cmd.Run()
	fmt.Println(outp.String())
}

func CreateInfoJson(){
	
	fileInfoJson := config.INFOJSONFILE


	Json := mystruct.Docs{
		User: mystruct.Userr{Active: false, TelegramId: 123456, TelegramToken: "telegram-token"},
		File: mystruct.Files{Paths: []string{"/rutaAbsoluta/hacia/el/archivo.txt"}},
		Dir: mystruct.Directory{Paths: []string{"/rutaAbsoluta/hacia/la/carpeta/"}},
	
	}

	
	fmt.Println( "Generando nuevo archivo :"+ fileInfoJson)
	time.Sleep(2 * time.Second)
	gson, err := json.MarshalIndent(Json, "", "   ")
	if(err != nil){
		fmt.Println(err.Error())
		
	}	
	file := os.WriteFile(fileInfoJson, []byte(string(gson)), 0666)
	if(file != nil){
		panic("No se pudo crear el archivo: Erro: " + file.Error())
		
	}
	fmt.Println("Archivo creado con exito :" + fileInfoJson )
	fmt.Println("\nAyuda rapida:")
	fmt.Println("\n1: Abre el archivo "+ fileInfoJson + "\n2: Pega la direccion de tus archivos o carpetas a monitorear y vuelve a iniciar el programa \n")
	//fmt.Println("()")
	time.Sleep(2 * time.Second)
	os.Exit(0)

}



func CheckDir(){
	_, err := os.Stat(config.PATH)
	
	if(err != nil){
		fmt.Println("Carpeta ", config.PATH, " no e encontrado")
		fmt.Println("Creando..")
		os.Mkdir(config.PATH, os.ModePerm)
		
	}


	
	_, errr := os.Stat(config.LOGFILE)
	if(errr != nil){
	
		createFile(config.LOGFILE)
	}

	_, er := os.Stat(config.INFOJSONFILE)
	if(er != nil){
		CreateInfoJson()
	}
	
	
}

func createFile(name string)bool{
	fmt.Println("Creando ", name)
	r, e := os.Create(name)
	if(e !=nil){
		fmt.Println(e.Error())
		return false
	}
	defer r.Close()
	fmt.Println("Archivo creado ", r.Name())
	return true
}


func Ls(path string)[]string{
	var files []string
	l, err := os.ReadDir(path)
	if(err != nil){
		fmt.Println(err.Error())
	}
	for _, v := range l{
		files = append(files, v.Name())
	}
	return files
}



func AbsolutePath(path string)string{
	pth, err := filepath.Abs(path)
	if(err != nil){
		return ""
	}
	return pth
}

