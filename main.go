package main

import (
	"crypto/md5"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"net/url"

	// HTML Encode(Escape/UnEscape)
	"html"
	"os"
)

func main() {

	// this if/else if will check if the passed arguments are less than 2 argument
	// and check if the passed arguments are more that 7 or equal to it
	if len(os.Args) <= 1 {
		fmt.Println("USAGE:\n\tHTML:\n\t\tEscape --> ./hasher html -e \"<script>alert('Hacked')</script>\"\n\t\tUnEscape --> ./hasher html -d \"&lt;script&gt;alert('Hacked')&lt;/script&gt;\"\n\tMD5:\n\t\tHash --> ./hasher md5 -h \"Hi\"\n\t\tDeHash --> ./hasher md5 -d \"Hi\" \"c1a5298f939e87e8f962a5edfc206918\"\n\tBASE64:\n\t\tEncode --> ./hasher base64 -e \"Hi\"\n\t\tDecode --> ./hasher base64 -d \"SGk=\"\n\tBASE32:\n\t\tEncode --> ./hasher base32 -e \"Hi\"\n\t\tDecode --> ./hasher base32 -d \"JBUQ====\"\n\tURL:\n\t\tEncode --> ./hasher url -e \"Hello, World\"\n\t\tDecode --> ./hasher url -d \"Hello%2C+World\"")
		return

	} else if len(os.Args) >= 7 {
		fmt.Println("To much arguments")
		os.Exit(1)
	}

	// Arguments
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]

	// HTML Escape/UnEscape
	if arg1 == "html" {
		switch arg2 {
		// Encode(Escape) html special characters like: <>"'&
		case "e", "-e", "--encode", "encode", "escape", "--escape":
			fmt.Printf("%s\n", html.EscapeString(arg3))
		case "d", "-d", "--decode", "decode", "unescape", "--unescape":
			fmt.Printf("%s\n", html.UnescapeString(arg3))
		default:
			fmt.Println("NOTHING")
		}

	} else if arg1 == "md5" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
			hash := md5.Sum([]byte(arg3))
			hashToString := fmt.Sprintf("%x\n", hash)
			fmt.Println(hashToString)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			arg4 := os.Args[4]
			// This if statement checks if passed tow md5 are the same or not(first turn the text to md5)
			if md5ToString(arg3) == arg4 || arg3 == md5ToString(arg4) {
				fmt.Printf("%s --> %s\n", arg3, arg4)
			} else {
				fmt.Println("Hash Not Found!")
			}
		default:
			fmt.Println("NOTHING")
		}

	} else if arg1 == "base64" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
			encode := base64.StdEncoding.EncodeToString([]byte(arg3))
			fmt.Println(encode)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			decode, err := base64.StdEncoding.DecodeString(arg3)
			if err != nil {
				fmt.Println("Not a Valid base64")
				return
			}
			fmt.Println(string(decode))
		default:
			fmt.Println("USAGE:\n\t\tEncode --> ./hasher base64 -e <YOUR_TEXT>\n\t\tDecode --> ./hasher base64 -d <YOUR_BASE64>")
		}

	} else if arg1 == "base32" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
			encode := base32.StdEncoding.EncodeToString([]byte(arg3))
			fmt.Println(encode)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			decode, err := base32.StdEncoding.DecodeString(arg3)
			if err != nil {
				fmt.Println("Not a Valid base32")
				return
			}
			fmt.Println(string(decode))
		default:
			fmt.Println("USAGE:\n\t\tEncode --> ./hasher base32 -e <YOUR_TEXT>\n\t\tDecode --> ./hasher base32 -d <YOUR_BASE32>")
		}

	} else if arg1 == "url" {
		switch arg2 {
		case "e", "-e", "--encode", "encode":
			encode := url.QueryEscape(arg3)
			fmt.Println(encode)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			decode, err := url.QueryUnescape(arg3)
			if err != nil {
				fmt.Println("Not a Valid URL Encode")
				return
			}
			fmt.Println(string(decode))
		default:
			fmt.Println("USAGE:\n\t\tEncode --> ./hasher url -e <YOUR_TEXT>\n\t\tDecode --> ./hasher url -d <YOUR_ENCODED_TEXT>")
		}

	} else {
		fmt.Println("argument not satisfied")
		os.Exit(1)
	}
}

// This function will turn text to md5 and return string of it.
// I wrote this function for compare Text vs Md5 to check if
// its the same(decode)
func md5ToString(a string) string {
	// Text to byte
	hash := md5.Sum([]byte(a))
	// md5 byte to string
	hashToString := fmt.Sprintf("%x", hash)
	// return the md5 string
	return hashToString
}
