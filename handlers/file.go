package handlers

import (
	"os"
)

func makeFile(name string) *os.File {
	var err error
	file, err := os.Create(name)
	if err != nil {
		file, err = os.Open(name)
		if err != nil {
			panic(err)
		}
	}
	return file
}

func writeCurlCmd(file *os.File, title string, strCom string, strRev string) {
	//file.WriteString("echo   \n")
	file.WriteString("echo - " + title + "\n")
	file.WriteString(strCom + "\n")
	file.WriteString("echo   \n")
	file.WriteString(strRev + "\n")
	file.WriteString("echo   \n")
}
