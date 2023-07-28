# hasher
An easy to use and fast tool for hashing and encoding stuff

## Features

### hasher supports
- MD5
- BASE64
- BASE32
- HTML Encode/Decode
- URL  Encode/Decode
- SHA1
- SHA256
- SHA512
## How to use it?

**NOTE:** You need to have `Go Programming Language` to run this program (https://go.dev/)

If You are using `Linux` OR `Mac`:
```git clone https://github.com/shareghzadeh/hasher.git
   cd hasher
   chmod +x hasher
   ./hasher
   ```

If you are using `Windows`:
```git clone https://github.com/shareghzadeh/hasher.git
   cd hasher
   go build hasher.go
   hasher.exe
   ```
**NOTE:** if you are using `windows` open `hasher.go` and change this:
```
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
	Reset  = "\033[0m"
)
```
to this:
```
const (
	Red    = ""
	Green  = ""
	Yellow = ""
	Blue   = ""
	Purple = ""
	Cyan   = ""
	White  = ""
	Reset  = ""
)
```
then run
```
go build hasher.go
hasher.exe
```

![how to use it? image](./images/pic-selected-230716-1031-27.png)

**NOTE:** If you want to run the program without creating binary, use `go run hasher.go`

## Why there is no `help` command?
When you can run the program and get the help why should be a help command?
