package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var passName = flag.String("name", "password", "Name of the password");
	var lengthFlag = flag.Int("length", 20, "Length of the password");
	flag.Parse();
	
	chars := loadChars();
	newPass := genPass(*lengthFlag, chars);
	savePass(*passName, newPass);

}

func genPass(passLength int, chars []string) (pass string) {
	// Generate a password
	charLength := len(chars);
	for i := 0; i < passLength; i++ {
		randomIndex := rand.Intn(charLength);
		pass += chars[randomIndex];
	}

	return
}

func loadChars() (chars []string) {
	// Load all the characters
	for r := '!'; r < '}'; r++ {
		chars = append(chars, fmt.Sprintf("%c", r));
	}

	return chars
}

func savePass(passName string, pass string) {
	// Save the password to a file
	currentDate := time.Now().Format("02-01-2006 15:04:05")
	f, err := os.OpenFile("pass.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
    	panic(err)
	}
	if _, err := f.Write([]byte(currentDate + "\t" + passName + "\t" + pass + "\n")); err != nil {
		panic(err)
	}

}