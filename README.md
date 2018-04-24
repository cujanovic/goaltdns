# goaltdns

### build:

cd goaltdns

go get github.com/joeguo/tldextract

go build -o bin/goaltdns

### usage:

-input string
Contains the file with a list of known subdomains for the domain

-output string
Is a file that will contain the massive list of altered and permuted subdomains


-words string
Is your list of words that you'd like to permute your current subdomains for the domain. (default "words.txt")
