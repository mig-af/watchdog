package funcs

import (
	//"fmt"
	"fmt"
	"testing"
)



func TestWriteLog(t *testing.T){


}

func TestCreateFile(t *testing.T){

	
}

func TestCreateInfoJson(t *testing.T){

	


	//clear
	///CreateInfoJson()

}


func TestList(t *testing.T){
	f := Ls("/workspaces/gatuso/installl/")
	fmt.Println(len(f))
}


func TestAbsolutePath(t *testing.T){

	er := AbsolutePath("../info")

	fmt.Println(er)
}