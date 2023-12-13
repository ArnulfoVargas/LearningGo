package iowriterreader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func IOsTest(){
	// Std out is limited but we can use it to send buffer data
	n, err := os.Stdout.Write([]byte("Hello world\n"))

	defer func(){
		if (err != nil){
			fmt.Println("An error has occurred:", err)
			recover()
		}
	}()

	fmt.Printf("Bytes written: %d\n", n)
	fmt.Printf("Data in bytes: %v\n", []byte("Hello world\n"))

	fmt.Println("Write 10 chars:")
	writted := make([]byte, 13)
	
	//Reading data and storing it as bytes
	d, err := os.Stdin.Read(writted)

	fmt.Printf("Bytes parsed: %v\nData: %v\n", d, string(writted))

	// Creating and editing files
	createFile()
	// ! createFileWithCustomData()

	// Opening fles
	openFile()
	
	//Buffio
	buffioTest()
}

func createFile(){
	file, err := os.Create("013_IOWriterReader/example.txt")

	defer func ()  {
		if err != nil {
			fmt.Println("Error During File Creation or Modification:", err)
		}else{
			file.Close()
		}
	}()

	_, err = file.Write([]byte("hello world rewritten"))
}

func createFileWithCustomData(){
	file, err := os.Create("013_IOWriterReader/example.txt")

	defer func ()  {
		if err != nil {
			fmt.Println("Error During File Creation or Modification:", err)
		}else{
			file.Close()
		}
	}()

	data := make([]byte, 12)

	fmt.Println("Left a message")
	
	_, err = os.Stdin.Read(data)

	_, err = file.Write(data)
}

// Open Files
func openFile(){
	file, err := os.Open("013_IOWriterReader/example.txt")
	defer func ()  {
		if (err != nil){
			fmt.Println("Error: ", err)
		}else{
			file.Close()
		}
	}()
	readed := make([]byte, 256)
	n, err := file.Read(readed)
	
	fmt.Printf("Readed %d bytes: %s\n", n, string(readed))
}

func buffioTest(){
	buff := bytes.NewBufferString(`Hello world!
Testing the multiline
String`)
	sc := bufio.NewReader(buff)

	readed, err := sc.ReadString('\n')

	for err == nil{
		fmt.Print(readed)
		readed, err = sc.ReadString('\n')
	}

	if err == io.EOF{
		fmt.Println("Final line:", readed)
	}else{
		fmt.Println("Err:", err)
	}
}