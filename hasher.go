package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"runtime"

	"log"
	"net/url"
	"regexp"

	"html"
	"os"
)

// const(
//
//	Red    = "\033[31m"
//	Green  = "\033[32m"
//	Yellow = "\033[33m"
//	Blue   = "\033[34m"
//	Purple = "\033[35m"
//	Cyan   = "\033[36m"
//	White  = "\033[37m"
//	Reset  = "\033[0m"
//
// )
var (
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
	White  string
	Reset  string
)

func main() {
	if runtime.GOOS == "windows" {
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		White = ""
		Reset = ""
	} else {
		Red = "\033[31m"
		Green = "\033[32m"
		Yellow = "\033[33m"
		Blue = "\033[34m"
		Purple = "\033[35m"
		Cyan = "\033[36m"
		White = "\033[37m"
		Reset = "\033[0m"
	}
	// this if/else if will check if the passed arguments are less than 2 argument
	// and check if the passed arguments are more that 7 or equal to it
	if len(os.Args) <= 1 {
		fmt.Printf("%sUSAGE:\n\tHTML:\n\t\tEscape --> ./hasher html -e \"<script>alert('Hacked')</script>\"\n\t\tUnEscape --> ./hasher html -d \"&lt;script&gt;alert('Hacked')&lt;/script&gt;\"\n\tMD5:\n\t\tHash --> ./hasher md5 -h \"Hi\"\n\t\tDeHash --> ./hasher md5 -d \"Hi\" \"c1a5298f939e87e8f962a5edfc206918\"\n\tBASE64:\n\t\tEncode --> ./hasher base64 -e \"Hi\"\n\t\tDecode --> ./hasher base64 -d \"SGk=\"\n\tBASE32:\n\t\tEncode --> ./hasher base32 -e \"Hi\"\n\t\tDecode --> ./hasher base32 -d \"JBUQ====\"\n\tURL:\n\t\tEncode --> ./hasher url -e \"Hello, World\"\n\t\tDecode --> ./hasher url -d \"Hello%%2C+World\"\n\tSHA1:\n\t\tHash --> ./hasher sha1 -e \"123\"\n\t\tDehash --> ./hasher sha1 -d \"123\" \"40bd001563085fc35165329ea1ff5c5ecbdbbeef\"\n\tSHA256:\n\t\tHash --> ./hasher sha256 -e \"123\"\n\t\tDehash --> ./hasher sha256 -d \"123\" \"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3\"\n\tSHA512:\n\t\tHash --> ./hasher sha512 -e \"123\"\n\t\tDehash --> ./hasher sha512 -d \"123\" \"3c9909afec25354d551dae21590bb26e38d53f2173b8d3dc3eee4c047e7ab1c1eb8b85103e3be7ba613b31bb5c9c36214dc9f14a42fd7a2fdb84856bca5c44c2\"\n", Green)
		// os.Exit(1)
		return

	} else if len(os.Args) >= 7 {
		fmt.Printf("%sTo much arguments\n", Red)
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
			fmt.Printf("%s%s\n", Green, html.EscapeString(arg3))
		case "d", "-d", "--decode", "decode", "unescape", "--unescape":
			fmt.Printf("%s\n", html.UnescapeString(arg3))
		default:
			fmt.Printf("%sNOTHING", Red)
		}

	} else if arg1 == "md5" {
		
		switch arg2 {
		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
			hash := md5.Sum([]byte(arg3))
			hashToString := fmt.Sprintf("%x\n", hash)
			fmt.Println(hashToString)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			arg4 := os.Args[4]

			file, err := os.Open(arg4)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			var a string
			if !isMd5(arg3) {
				for scanner.Scan() {
					if md5ToString(arg3) == scanner.Text() {
						a = scanner.Text()
					}
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
				if md5ToString(arg3) == a {
					fmt.Printf("%s -> %s\n", arg3, a)
				} else {
					fmt.Printf("%sHash NOT FOUND\n", Red)
				}
			} else if isMd5(arg3) {
				for scanner.Scan() {
					if arg3 == md5ToString(scanner.Text()) {
						a = scanner.Text()
					}
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
				if arg3 == md5ToString(a) {
					fmt.Printf("%s%s -> %s\n", Green, a, arg3)
				} else {
					fmt.Printf("%sHash NOT FOUND\n", Red)
				}
			} else {
				fmt.Println("Your passed argument is hash and your file is hash too!!!")
			}
		case "-r", "r":
			arg4 := os.Args[4]
			if arg4 == "h" || arg4 == "-h" {
				arg5 := os.Args[5]

				fileA, err := os.Open(arg3)
				if err != nil {
					panic(err)
				}
				defer fileA.Close()

				fileB, err := os.Open(arg5)
				if err != nil {
					panic(err)
				}
				defer fileB.Close()

				// Create a scanner for each file
				scannerA := bufio.NewScanner(fileA)
				scannerB := bufio.NewScanner(fileB)

				// Read each line of a.txt
				for scannerA.Scan() {
					lineA := scannerA.Text()

					// Reset scannerB to the beginning of b.txt
					fileB.Seek(0, 0)
					scannerB = bufio.NewScanner(fileB)

					// Read each line of b.txt and compare to lineA
					for scannerB.Scan() {
						lineB := scannerB.Text()

						if md5ToString(lineA) == lineB {
							fmt.Printf("%s --> %s\n", lineA, lineB)
							break
						}
					}
				}
			}
		default:
			fmt.Printf("%sUSAGE:\n\t\tHash --> ./hasher md5 -e <YOUR_TEXT>\n\t\tDehash--> ./hasher md5 -d <YOUR_TEXT> <YOUR_MD5>\n\n\t\tOR\n\t./hasher md5 -f a.txt b.txt", Red)
		}

	} else if arg1 == "base64" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
			encode := base64.StdEncoding.EncodeToString([]byte(arg3))
			fmt.Println(encode)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			decode, err := base64.StdEncoding.DecodeString(arg3)
			if err != nil {
				fmt.Printf("%sNot a Valid base64", Red)
				os.Exit(1)
			}
			fmt.Println(string(decode))
		default:
			fmt.Printf("%sUSAGE:\n\t\tEncode --> ./hasher base64 -e <YOUR_TEXT>\n\t\tDecode --> ./hasher base64 -d <YOUR_BASE64>", Green)
		}

	} else if arg1 == "base32" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "h", "-h", "hash", "--hash":
			encode := base32.StdEncoding.EncodeToString([]byte(arg3))
			fmt.Println(encode)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			decode, err := base32.StdEncoding.DecodeString(arg3)
			if err != nil {
				fmt.Printf("%sNot a Valid base32", Red)
				return
			}
			fmt.Println(string(decode))
		default:
			fmt.Printf("%sUSAGE:\n\t\tEncode --> ./hasher base32 -e <YOUR_TEXT>\n\t\tDecode --> ./hasher base32 -d <YOUR_BASE32>", Green)
		}

	} else if arg1 == "url" {
		switch arg2 {
		case "e", "-e", "--encode", "encode":
			encode := url.QueryEscape(arg3)
			fmt.Println(encode)
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			decode, err := url.QueryUnescape(arg3)
			if err != nil {
				fmt.Printf("%sNot a Valid URL Encode", Red)
				return
			}
			fmt.Println(string(decode))
		default:
			fmt.Printf("%sUSAGE:\n\t\tEncode --> ./hasher url -e <YOUR_TEXT>\n\t\tDecode --> ./hasher url -d <YOUR_ENCODED_TEXT>", Green)
		}

	} else if arg1 == "sha1" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "--encrypt", "encrypt":
			fmt.Println(sha1ToString(arg3))
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			arg4 := os.Args[4]
			// This if statement checks if passed tow md5 are the same or not(first turn the text to md5)
			if sha1ToString(arg3) == arg4 || arg3 == sha1ToString(arg4) {
				fmt.Printf("%s%s --> %s\n", Green, arg3, arg4)
			} else {
				fmt.Printf("%sHash Not Found!\n", Red)
			}
		default:
			fmt.Printf("%sUSAGE:\n\t\tHash --> ./hasher sha1 -e <YOUR_TEXT>\n\t\tDehash --> ./hasher sha1 -d <YOUR_TEXT> <YOUR_HASHED_TEXT>", Green)
		}

	} else if arg1 == "sha256" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "--encrypt", "encrypt":
			fmt.Println(sha256ToString(arg3))
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			arg4 := os.Args[4]
			// This if statement checks if passed tow md5 are the same or not(first turn the text to md5)
			if sha256ToString(arg3) == arg4 || arg3 == sha256ToString(arg4) {
				fmt.Printf("%s%s --> %s\n", Green, arg3, arg4)
			} else {
				fmt.Printf("%sHash Not Found!\n", Red)
			}
		default:
			fmt.Printf("%sUSAGE:\n\t\tHash --> ./hasher sha256 -e <YOUR_TEXT>\n\t\tDehash --> ./hasher sha256 -d <YOUR_TEXT> <YOUR_HASHED_TEXT>", Green)
		}

	} else if arg1 == "sha512" {
		switch arg2 {
		case "e", "-e", "--encode", "encode", "--encrypt", "encrypt":
			fmt.Println(sha512ToString(arg3))
		case "d", "-d", "--decode", "decode", "dehash", "--dehash":
			arg4 := os.Args[4]
			// This if statement checks if passed tow md5 are the same or not(first turn the text to md5)
			if sha512ToString(arg3) == arg4 || arg3 == sha512ToString(arg4) {
				fmt.Printf("%s%s --> %s\n", Green, arg3, arg4)
			} else {
				fmt.Printf("%sHash Not Found!\n", Red)
			}
		default:
			fmt.Printf("%sUSAGE:\n\t\tHash --> ./hasher sha512 -e <your_text>\n\t\tdehash --> ./hasher sha512 -d <your_text> <your_hashed_text>", Green)
		}

	} else {
		fmt.Printf("%sargument not satisfied", Red)
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

func sha1ToString(a string) string {
	hash := sha1.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func sha256ToString(a string) string {
	hash := sha256.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func sha512ToString(a string) string {
	hash := sha512.New()
	hash.Write([]byte(a))
	hashBytes := hash.Sum(nil)
	hashToString := fmt.Sprintf("%x", hashBytes)
	return hashToString
}

func isMd5(s string) bool {
	// Compile a regular expression to match a 32-character hexadecimal string
	pattern := "^[a-f0-9]{32}$"
	re := regexp.MustCompile(pattern)

	// Check if the input matches the pattern
	if !re.MatchString(s) {
		return false
	}
	return true
}

func isHtmlEncode(s string) bool {
	// Test HTML-encoded pattern
	htmlPattern := "&[a-zA-Z]+;|&#\\d+;"
	htmlRegex := regexp.MustCompile(htmlPattern)
	// Check if the input matches the pattern
	if !htmlRegex.MatchString(s) {
		return false
	}
	return true

}
func isUrlEncode(s string) bool {
	// Test URL-encoded pattern
	urlPattern := "%[0-9a-fA-F]{2}"
	urlRegex := regexp.MustCompile(urlPattern)
	if !urlRegex.MatchString(s) {
		return false
	}
	return true
}
func isSha1(s string) bool {
	// Test URL-encoded pattern
	sha1Pattern := "^[a-fA-F0-9]{40}$"
	sha1Regex := regexp.MustCompile(sha1Pattern)
	if !sha1Regex.MatchString(s) {
		return false
	}
	return true
}
func isSha256(s string) bool {
	// Test URL-encoded pattern
	sha256Pattern := "^[a-fA-F0-9]{64}$"
	sha256Regex := regexp.MustCompile(sha256Pattern)
	if !sha256Regex.MatchString(s) {
		return false
	}
	return true
}
func isSha512(s string) bool {
	// Test URL-encoded pattern
	sha512Pattern := "^[0-9a-fA-F]{128}$"
	sha512Regex := regexp.MustCompile(sha512Pattern)
	if !sha512Regex.MatchString(s) {
		return false
	}
	return true

}
