package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var OutputPath = "Output"
var LogPath = "Logs"
var initFile = ".init.kx"

var ok = color.Bold + color.Green + "[+] " + color.Reset
var error_msg = color.Bold + color.Red + "[error] " + color.Reset

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}

	if _, err := os.Stat(OutputPath); os.IsNotExist(err) {
		os.Mkdir(OutputPath, 0755)
	}

	if _, err := os.Stat(LogPath); os.IsNotExist(err) {
		os.Mkdir(LogPath, 0755)
	}

	if _, err := os.Stat(initFile); os.IsNotExist(err) {
		f, err := os.Create(initFile)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error creating init file: " + err.Error() + color.Reset)
			os.Exit(1)
		}
		f.Close()
	}
}

func banner() {
	var author string = "sc4rfurry"
	var version string = "1.0.0-alpha"
	var go_version string = "1.18.1 or higher"
	var github string = "https://github.com/sc4rfurry"
	var description string = "Tool to Mutate/Encrypt files using AES-256 Encryption/Decryption." + color.Bold + color.Cyan + "(! Alpha !)." + color.Reset
	banner := figure.NewColorFigure("kryt0x", "", "cyan", true)
	banner.Print()
	println("\n")
	fmt.Println(color.Bold + color.Green + "  " + description + "  " + color.Reset + "\n")
	fmt.Println(color.Bold + color.Purple + "\t  Author: " + color.Reset + author)
	fmt.Println(color.Bold + color.Purple + "\t  Version: " + color.Reset + version)
	fmt.Println(color.Bold + color.Purple + "\t  Go Version: " + color.Reset + go_version)
	fmt.Println(color.Bold + color.Purple + "\t  Github: " + color.Reset + github)
	fmt.Println(color.Bold + color.Purple + "\t  Date: " + color.Reset + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("-------------------------------------------------------------------------------->")
	println("\n")

}

func help() {
	println(color.Bold + color.Green + "\n\t\t\t~ Help Menu ~" + color.Reset)
	println("\n\t" + color.Bold + color.Cyan + "Usage: " + color.Reset + "./main -e <file> " + color.Bold + color.Cyan + "OR" + color.Reset + " ./main -d <file> <key>")
	println(color.Bold + color.Gray + "___________________________________________________________________________________________________________" + color.Reset)
	println("\n\t" + color.Bold + color.Green + "Options: " + color.Reset)
	println("\t\t" + color.Bold + color.Yellow + "File " + color.Reset + "\t\tFile to encrypt/decrypt")
	println("\t\t" + color.Bold + color.Yellow + "-e " + color.Reset + "\t\tExcrypt file")
	println("\t\t" + color.Bold + color.Yellow + "-d " + color.Reset + "\t\tDecrypt file")
	println("\t\t" + color.Bold + color.Yellow + "-h/--help " + color.Reset + "\tShow this help menu" + "\n")
	println("\n\t" + color.Bold + color.Cyan + "!Note:-" + color.Reset + color.Bold + color.Yellow + " Decryption " + color.Reset + " need Filename and Encryption Key" + "\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		help()
	}
	if os.Args[1] == "-e" {
		banner()
		filename := os.Args[2]
		originalCode, err := ioutil.ReadFile(filename)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error reading file: " + err.Error() + color.Reset)
			os.Exit(1)
		}

		key := make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			println(error_msg + color.Bold + color.Red + "Error generating key: " + err.Error() + color.Reset)
			os.Exit(1)
		}

		ciphertext, err := encrypt(originalCode, key)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error encrypting file: " + err.Error() + color.Reset)
			os.Exit(1)
		}
		file := strings.Split(filename, "/")
		f := file[len(file)-1]
		newFilename := "mutated_" + f + ".kx"
		filePath := OutputPath + "/" + newFilename
		err = ioutil.WriteFile(filePath, ciphertext, 0644)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error writing file: " + err.Error() + color.Reset)
			os.Exit(1)
		}

		logger("Original FIle:" + f + " \n\tMutated File: " + filePath + " \n\tkey: " + hex.EncodeToString(key))
		println(ok + color.Bold + color.Green + " File Encrypted Successfully!" + color.Reset)
		println(ok + color.Bold + color.Green + " Mutated File: " + filePath + color.Reset)
		println(ok + color.Bold + color.Green + " Check " + color.Reset + "Logs" + color.Bold + color.Green + " for key" + color.Reset + "\n")
	} else if os.Args[1] == "-d" {
		if len(os.Args) < 4 {
			help()
		}
		banner()
		filename := os.Args[2]
		key, err := hex.DecodeString(os.Args[3])
		key = []byte(key)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error decoding key: " + err.Error() + color.Reset)
			os.Exit(1)
		}
		cipherCode, err := ioutil.ReadFile(filename)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error reading file: " + err.Error() + color.Reset)
			os.Exit(1)
		}
		plaintext, err := decrypt(cipherCode, key)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error decrypting file: " + err.Error() + color.Reset)
			os.Exit(1)
		}
		file := strings.Split(filename, "/")
		f := file[len(file)-1]
		filePath := OutputPath + "/" + "decrypted_" + f
		err = ioutil.WriteFile(filePath, plaintext, 0644)
		if err != nil {
			println(error_msg + color.Bold + color.Red + "Error writing file: " + err.Error() + color.Reset)
			os.Exit(1)
		}

		logger("Decrypted File: " + filePath)
		println(ok + color.Bold + color.Green + " File Decrypted Successfully!" + color.Reset)
		println(ok + color.Bold + color.Green + " Decrypted File: " + filePath + color.Reset)
		println(ok + color.Bold + color.Green + " Check " + color.Reset + "Logs" + color.Bold + color.Green + " for more info." + color.Reset + "\n")
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		help()
	} else {
		println(error_msg + color.Bold + color.Red + "Invalid option: " + os.Args[1] + color.Reset)
		println(error_msg + color.Bold + color.Red + "Use -h or --help for help menu" + color.Reset)
		os.Exit(1)
	}
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	copy(ciphertext[:aes.BlockSize], iv)

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	hash := sha256.Sum256(key)

	ciphertext = append(ciphertext, hash[:]...)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	iv := ciphertext[:aes.BlockSize]
	data := ciphertext[aes.BlockSize : len(ciphertext)-32]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(data))

	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(plaintext, data)

	return plaintext, nil
}

func logger(msg string) {
	currentTime := time.Now()
	f, err := os.OpenFile(LogPath+"/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println(error_msg + color.Bold + color.Red + " Error opening log file" + color.Reset)
		os.Exit(1)
	}
	defer f.Close()
	var year = currentTime.Year()
	var month = currentTime.Month()
	var day = currentTime.Day()
	var hour = currentTime.Hour()
	var minute = currentTime.Minute()
	var second = currentTime.Second()
	var dateTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	var log = fmt.Sprintf("\n\n[+] %s -- \n\t%s", dateTime, msg+"\n")
	if _, err := f.WriteString(log); err != nil {
		println(error_msg + color.Bold + color.Red + " Error writing to log file" + color.Reset)
		os.Exit(1)
	}
}
