package main

import (
	"os"
	"os/exec"
	"flag"
	"io/ioutil"
	"strings"
	"bufio"
	"log"
	"bytes"
	"strconv"
	"regexp"
	"github.com/joeguo/tldextract"
)

func importwordsfromfile(filename string) ([]string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines) - 1]
	return lines
}

func sortfileunique(filename string) {
	command := []string{"sort", "-u", "--parallel=2", filename, "-o", filename}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Env = append(os.Environ(), "LC_ALL=C")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

//word.subdomain.domain
func permutation1(word string, subdomain string, domain string) (string) {
	var buff bytes.Buffer
	buff.WriteString(word)
	buff.WriteString(".")
	buff.WriteString(subdomain)
	buff.WriteString(".")
	buff.WriteString(domain)
	return buff.String()
}

//subdomain.word.domain
func permutation2(word string, subdomain string, domain string) (string) {
	var buff bytes.Buffer
	buff.WriteString(subdomain)
	buff.WriteString(".")
	buff.WriteString(word)
	buff.WriteString(".")
	buff.WriteString(domain)
	return buff.String()
}

//word-subdomain.domain
func permutation3(word string, subdomain string, domain string) (string) {
	var buff bytes.Buffer
	buff.WriteString(word)
	buff.WriteString("-")
	buff.WriteString(subdomain)
	buff.WriteString(".")
	buff.WriteString(domain)
	return buff.String()
}

//subdomain-word.domain
func permutation4(word string, subdomain string, domain string) (string) {
	var buff bytes.Buffer
	buff.WriteString(subdomain)
	buff.WriteString("-")
	buff.WriteString(word)
	buff.WriteString(".")
	buff.WriteString(domain)
	return buff.String()
}

//wordsubdomain.domain
func permutation7(word string, subdomain string, domain string) (string) {
	var buff bytes.Buffer
	buff.WriteString(word)
	buff.WriteString(subdomain)
	buff.WriteString(".")
	buff.WriteString(domain)
	return buff.String()
}

//subdomainword.domain
func permutation8(word string, subdomain string, domain string) (string) {
	var buff bytes.Buffer
	buff.WriteString(subdomain)
	buff.WriteString(word)
	buff.WriteString(".")
	buff.WriteString(domain)
	return buff.String()
}

//word-NUMBER.subdomain.domain
func permutation9(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(word)
		buff.WriteString("-")
		buff.WriteString(number)
		buff.WriteString(".")
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//NUMBER-word.subdomain.domain
func permutation10(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(number)
		buff.WriteString("-")
		buff.WriteString(word)
		buff.WriteString(".")
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//word.NUMBER.subdomain.domain
func permutation11(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(word)
		buff.WriteString(".")
		buff.WriteString(number)
		buff.WriteString(".")
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//NUMBER.word.subdomain.domain
func permutation12(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(number)
		buff.WriteString(".")
		buff.WriteString(word)
		buff.WriteString(".")
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//subdomain.word-NUMBER.domain
func permutation13(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(word)
		buff.WriteString("-")
		buff.WriteString(number)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//subdomain.NUMBER-word.domain
func permutation14(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(number)
		buff.WriteString("-")
		buff.WriteString(word)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//subdomain.word.NUMBER.domain
func permutation15(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(word)
		buff.WriteString(".")
		buff.WriteString(number)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//subdomain.NUMBER.word.domain
func permutation16(word string, subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	for k := 0; k < 13; k++ {
		number := strconv.Itoa(k)
		buff.WriteString(subdomain)
		buff.WriteString(".")
		buff.WriteString(number)
		buff.WriteString(".")
		buff.WriteString(word)
		buff.WriteString(".")
		buff.WriteString(domain)
		buff.WriteString("\n")
		strings1 = append(strings1, buff.String())
		buff.Reset()
	}
	return strings1
}

//ends with one number from 0-9, loop all others from 0-12 add subdomain
func permutation17(subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	matched, err := regexp.MatchString("^([a-z]*)+([0-9]{1,1})$", subdomain)
	if err != nil {
		log.Fatal(err)
	}
	if matched == true {
		subdomainnumber := subdomain[len(subdomain) - 1:]
		for k := 0; k < 13; k++ {
			number := strconv.Itoa(k)
			if subdomainnumber != number {
				buff.WriteString(subdomain[:len(subdomain) - 1])
				buff.WriteString(number)
				buff.WriteString(".")
				buff.WriteString(domain)
				buff.WriteString("\n")
				strings1 = append(strings1, buff.String())
				buff.Reset()
			}
		}
	} else {
		strings1 = append(strings1, "NONE")
	}
	return strings1
}

//ends with one number from 0-9, loop all others from 0-12 add .subdomain
func permutation18(subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	matched, err := regexp.MatchString("^([a-z]*)+([0-9]{1,1})$", subdomain)
	if err != nil {
		log.Fatal(err)
	}
	if matched == true {
		subdomainnumber := subdomain[len(subdomain) - 1:]
		for k := 0; k < 13; k++ {
			number := strconv.Itoa(k)
			if subdomainnumber != number {
				buff.WriteString(subdomain[:len(subdomain) - 1])
				buff.WriteString(".")
				buff.WriteString(number)
				buff.WriteString(".")
				buff.WriteString(domain)
				buff.WriteString("\n")
				strings1 = append(strings1, buff.String())
				buff.Reset()
			}
		}
	} else {
		strings1 = append(strings1, "NONE")
	}
	return strings1
}

//starts with one number from 0-9, loop all others from 0-12 add subdomain
func permutation19(subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	matched, err := regexp.MatchString("^([0-9]{1,1})+([a-z]*)$", subdomain)
	if err != nil {
		log.Fatal(err)
	}
	if matched == true {
		subdomainnumber := subdomain[0:1]
		for k := 0; k < 13; k++ {
			number := strconv.Itoa(k)
			if subdomainnumber != number {
				buff.WriteString(number)
				buff.WriteString(subdomain[1:len(subdomain)])
				buff.WriteString(".")
				buff.WriteString(domain)
				buff.WriteString("\n")
				strings1 = append(strings1, buff.String())
				buff.Reset()
			}
		}
	} else {
		strings1 = append(strings1, "NONE")
	}
	return strings1
}

//starts with one number from 0-9, loop all others from 0-12 add .subdomain
func permutation20(subdomain string, domain string) ([]string) {
	var buff bytes.Buffer
	var strings1 []string
	matched, err := regexp.MatchString("^([0-9]{1,1})+([a-z]*)$", subdomain)
	if err != nil {
		log.Fatal(err)
	}
	if matched == true {
		subdomainnumber := subdomain[0:1]
		for k := 0; k < 13; k++ {
			number := strconv.Itoa(k)
			if subdomainnumber != number {
				buff.WriteString(number)
				buff.WriteString(".")
				buff.WriteString(subdomain[1:len(subdomain)])
				buff.WriteString(".")
				buff.WriteString(domain)
				buff.WriteString("\n")
				strings1 = append(strings1, buff.String())
				buff.Reset()
			}
		}
	} else {
		strings1 = append(strings1, "NONE")
	}
	return strings1
}


func main() {
	input := flag.String("input", "", "Contains the file with a list of known subdomains for the domain")
	output := flag.String("output", "", "Is a file that will contain the massive list of altered and permuted subdomains")
	words := flag.String("words", "words.txt", "Is your list of words that you'd like to permute your current subdomains for the domain.")
	flag.Parse()

	if *input == "" || *output == "" || *words == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	//tld cache
	cache := "/tmp/tld.cache"
	extract, _ := tldextract.New(cache,false)

	//Define words
	wordslines := importwordsfromfile(*words)

	//Create the output file
	outputfile, erroutput := os.Create(*output)
	if erroutput != nil {
		log.Fatal(erroutput)
	}
	defer outputfile.Close()

	//Read the input file extract the subdomain, analyze it and do the permutations
	inputfile, errinput := os.Open(*input)
	if errinput != nil {
		log.Fatal(errinput)
	}
	defer inputfile.Close()

	scanner := bufio.NewScanner(inputfile)
	writter := bufio.NewWriter(outputfile)

	for scanner.Scan() {
		tldextractresult := extract.Extract(scanner.Text())
		subdomainline := tldextractresult.Sub
		domain := tldextractresult.Root+"."+tldextractresult.Tld

		if subdomainline != "" {
			if strings.Contains(subdomainline, ".") {
				subdomainlinefragmentdot := strings.Split(subdomainline, ".")
				for j := 0; j < len(subdomainlinefragmentdot); j++ {
					for i := 0; i < len(wordslines); i++ {
						writter.WriteString(permutation1(wordslines[i], subdomainlinefragmentdot[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation2(wordslines[i], subdomainlinefragmentdot[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation3(wordslines[i], subdomainlinefragmentdot[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation4(wordslines[i], subdomainlinefragmentdot[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation7(wordslines[i], subdomainlinefragmentdot[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation8(wordslines[i], subdomainlinefragmentdot[j], domain)+"\n")
						writter.Flush()
						p9 := permutation9(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p9); z++ {
							writter.WriteString(p9[z])
							writter.Flush()
						}
						p10 := permutation10(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p10); z++ {
							writter.WriteString(p10[z])
							writter.Flush()
						}
						p11 := permutation11(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p11); z++ {
							writter.WriteString(p11[z])
							writter.Flush()
						}
						p12 := permutation12(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p12); z++ {
							writter.WriteString(p12[z])
							writter.Flush()
						}
						p13 := permutation13(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p13); z++ {
							writter.WriteString(p13[z])
							writter.Flush()
						}
						p14 := permutation14(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p14); z++ {
							writter.WriteString(p14[z])
							writter.Flush()
						}
						p15 := permutation15(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p15); z++ {
							writter.WriteString(p15[z])
							writter.Flush()
						}
						p16 := permutation16(wordslines[i], subdomainlinefragmentdot[j], domain)
						for z := 0; z < len(p16); z++ {
							writter.WriteString(p16[z])
							writter.Flush()
						}
						p17 := permutation17(subdomainlinefragmentdot[j], domain)
						if p17[0] != "NONE" {
							for z := 0; z < len(p17); z++ {
								writter.WriteString(p17[z])
								writter.Flush()
							}
						}
						p18 := permutation18(subdomainlinefragmentdot[j], domain)
						if p18[0] != "NONE" {
							for z := 0; z < len(p18); z++ {
								writter.WriteString(p18[z])
								writter.Flush()
							}
						}
						p19 := permutation19(subdomainlinefragmentdot[j], domain)
						if p19[0] != "NONE" {
							for z := 0; z < len(p19); z++ {
								writter.WriteString(p19[z])
								writter.Flush()
							}
						}
						p20 := permutation20(subdomainlinefragmentdot[j], domain)
						if p20[0] != "NONE" {
							for z := 0; z < len(p20); z++ {
								writter.WriteString(p20[z])
								writter.Flush()
							}
						}
					}
				}
				for i := 0; i < len(wordslines); i++ {
					writter.WriteString(permutation1(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation2(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation3(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation4(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation7(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation8(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					p9 := permutation9(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p9); z++ {
						writter.WriteString(p9[z])
						writter.Flush()
					}
					p10 := permutation10(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p10); z++ {
						writter.WriteString(p10[z])
						writter.Flush()
					}
					p11 := permutation11(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p11); z++ {
						writter.WriteString(p11[z])
						writter.Flush()
					}
					p12 := permutation12(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p12); z++ {
						writter.WriteString(p12[z])
						writter.Flush()
					}
					p13 := permutation13(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p13); z++ {
						writter.WriteString(p13[z])
						writter.Flush()
					}
					p14 := permutation14(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p14); z++ {
						writter.WriteString(p14[z])
						writter.Flush()
					}
					p15 := permutation15(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p15); z++ {
						writter.WriteString(p15[z])
						writter.Flush()
					}
					p16 := permutation16(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p16); z++ {
						writter.WriteString(p16[z])
						writter.Flush()
					}
					p17 := permutation17(subdomainline, domain)
					if p17[0] != "NONE" {
						for z := 0; z < len(p17); z++ {
							writter.WriteString(p17[z])
							writter.Flush()
						}
					}
					p18 := permutation18(subdomainline, domain)
					if p18[0] != "NONE" {
						for z := 0; z < len(p18); z++ {
							writter.WriteString(p18[z])
							writter.Flush()
						}
					}
					p19 := permutation19(subdomainline, domain)
					if p19[0] != "NONE" {
						for z := 0; z < len(p19); z++ {
							writter.WriteString(p19[z])
							writter.Flush()
						}
					}
					p20 := permutation20(subdomainline, domain)
					if p20[0] != "NONE" {
						for z := 0; z < len(p20); z++ {
							writter.WriteString(p20[z])
							writter.Flush()
						}
					}
				}
			} else if strings.Contains(subdomainline, "-") {
				subdomainlinefragmentdash := strings.Split(subdomainline, "-")
				for j := 0; j < len(subdomainlinefragmentdash); j++ {
					for i := 0; i < len(wordslines); i++ {
						writter.WriteString(permutation1(wordslines[i], subdomainlinefragmentdash[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation2(wordslines[i], subdomainlinefragmentdash[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation3(wordslines[i], subdomainlinefragmentdash[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation4(wordslines[i], subdomainlinefragmentdash[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation7(wordslines[i], subdomainlinefragmentdash[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation8(wordslines[i], subdomainlinefragmentdash[j], domain)+"\n")
						writter.Flush()
						p9 := permutation9(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p9); z++ {
							writter.WriteString(p9[z])
							writter.Flush()
						}
						p10 := permutation10(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p10); z++ {
							writter.WriteString(p10[z])
							writter.Flush()
						}
						p11 := permutation11(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p11); z++ {
							writter.WriteString(p11[z])
							writter.Flush()
						}
						p12 := permutation12(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p12); z++ {
							writter.WriteString(p12[z])
							writter.Flush()
						}
						p13 := permutation13(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p13); z++ {
							writter.WriteString(p13[z])
							writter.Flush()
						}
						p14 := permutation14(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p14); z++ {
							writter.WriteString(p14[z])
							writter.Flush()
						}
						p15 := permutation15(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p15); z++ {
							writter.WriteString(p15[z])
							writter.Flush()
						}
						p16 := permutation16(wordslines[i], subdomainlinefragmentdash[j], domain)
						for z := 0; z < len(p16); z++ {
							writter.WriteString(p16[z])
							writter.Flush()
						}
						p17 := permutation17(subdomainlinefragmentdash[j], domain)
						if p17[0] != "NONE" {
							for z := 0; z < len(p17); z++ {
								writter.WriteString(p17[z])
								writter.Flush()
							}
						}
						p18 := permutation18(subdomainlinefragmentdash[j], domain)
						if p18[0] != "NONE" {
							for z := 0; z < len(p18); z++ {
								writter.WriteString(p18[z])
								writter.Flush()
							}
						}
						p19 := permutation19(subdomainlinefragmentdash[j], domain)
						if p19[0] != "NONE" {
							for z := 0; z < len(p19); z++ {
								writter.WriteString(p19[z])
								writter.Flush()
							}
						}
						p20 := permutation20(subdomainlinefragmentdash[j], domain)
						if p20[0] != "NONE" {
							for z := 0; z < len(p20); z++ {
								writter.WriteString(p20[z])
								writter.Flush()
							}
						}
					}
				}
				for i := 0; i < len(wordslines); i++ {
					writter.WriteString(permutation1(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation2(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation3(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation4(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation7(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation8(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					p9 := permutation9(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p9); z++ {
						writter.WriteString(p9[z])
						writter.Flush()
					}
					p10 := permutation10(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p10); z++ {
						writter.WriteString(p10[z])
						writter.Flush()
					}
					p11 := permutation11(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p11); z++ {
						writter.WriteString(p11[z])
						writter.Flush()
					}
					p12 := permutation12(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p12); z++ {
						writter.WriteString(p12[z])
						writter.Flush()
					}
					p13 := permutation13(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p13); z++ {
						writter.WriteString(p13[z])
						writter.Flush()
					}
					p14 := permutation14(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p14); z++ {
						writter.WriteString(p14[z])
						writter.Flush()
					}
					p15 := permutation15(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p15); z++ {
						writter.WriteString(p15[z])
						writter.Flush()
					}
					p16 := permutation16(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p16); z++ {
						writter.WriteString(p16[z])
						writter.Flush()
					}
					p17 := permutation17(subdomainline, domain)
					if p17[0] != "NONE" {
						for z := 0; z < len(p17); z++ {
							writter.WriteString(p17[z])
							writter.Flush()
						}
					}
					p18 := permutation18(subdomainline, domain)
					if p18[0] != "NONE" {
						for z := 0; z < len(p18); z++ {
							writter.WriteString(p18[z])
							writter.Flush()
						}
					}
					p19 := permutation19(subdomainline, domain)
					if p19[0] != "NONE" {
						for z := 0; z < len(p19); z++ {
							writter.WriteString(p19[z])
							writter.Flush()
						}
					}
					p20 := permutation20(subdomainline, domain)
					if p20[0] != "NONE" {
						for z := 0; z < len(p20); z++ {
							writter.WriteString(p20[z])
							writter.Flush()
						}
					}
				}
			} else if strings.Contains(subdomainline, "_") {
				subdomainlinefragmentunderscore := strings.Split(subdomainline, "_")
				for j := 0; j < len(subdomainlinefragmentunderscore); j++ {
					for i := 0; i < len(wordslines); i++ {
						writter.WriteString(permutation1(wordslines[i], subdomainlinefragmentunderscore[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation2(wordslines[i], subdomainlinefragmentunderscore[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation3(wordslines[i], subdomainlinefragmentunderscore[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation4(wordslines[i], subdomainlinefragmentunderscore[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation7(wordslines[i], subdomainlinefragmentunderscore[j], domain)+"\n")
						writter.Flush()
						writter.WriteString(permutation8(wordslines[i], subdomainlinefragmentunderscore[j], domain)+"\n")
						writter.Flush()
						p9 := permutation9(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p9); z++ {
							writter.WriteString(p9[z])
							writter.Flush()
						}
						p10 := permutation10(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p10); z++ {
							writter.WriteString(p10[z])
							writter.Flush()
						}
						p11 := permutation11(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p11); z++ {
							writter.WriteString(p11[z])
							writter.Flush()
						}
						p12 := permutation12(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p12); z++ {
							writter.WriteString(p12[z])
							writter.Flush()
						}
						p13 := permutation13(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p13); z++ {
							writter.WriteString(p13[z])
							writter.Flush()
						}
						p14 := permutation14(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p14); z++ {
							writter.WriteString(p14[z])
							writter.Flush()
						}
						p15 := permutation15(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p15); z++ {
							writter.WriteString(p15[z])
							writter.Flush()
						}
						p16 := permutation16(wordslines[i], subdomainlinefragmentunderscore[j], domain)
						for z := 0; z < len(p16); z++ {
							writter.WriteString(p16[z])
							writter.Flush()
						}
						p17 := permutation17(subdomainlinefragmentunderscore[j], domain)
						if p17[0] != "NONE" {
							for z := 0; z < len(p17); z++ {
								writter.WriteString(p17[z])
								writter.Flush()
							}
						}
						p18 := permutation18(subdomainlinefragmentunderscore[j], domain)
						if p18[0] != "NONE" {
							for z := 0; z < len(p18); z++ {
								writter.WriteString(p18[z])
								writter.Flush()
							}
						}
						p19 := permutation19(subdomainlinefragmentunderscore[j], domain)
						if p19[0] != "NONE" {
							for z := 0; z < len(p19); z++ {
								writter.WriteString(p19[z])
								writter.Flush()
							}
						}
						p20 := permutation20(subdomainlinefragmentunderscore[j], domain)
						if p20[0] != "NONE" {
							for z := 0; z < len(p20); z++ {
								writter.WriteString(p20[z])
								writter.Flush()
							}
						}
					}
				}
				for i := 0; i < len(wordslines); i++ {
					writter.WriteString(permutation1(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation2(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation3(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation4(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation7(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation8(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					p9 := permutation9(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p9); z++ {
						writter.WriteString(p9[z])
						writter.Flush()
					}
					p10 := permutation10(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p10); z++ {
						writter.WriteString(p10[z])
						writter.Flush()
					}
					p11 := permutation11(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p11); z++ {
						writter.WriteString(p11[z])
						writter.Flush()
					}
					p12 := permutation12(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p12); z++ {
						writter.WriteString(p12[z])
						writter.Flush()
					}
					p13 := permutation13(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p13); z++ {
						writter.WriteString(p13[z])
						writter.Flush()
					}
					p14 := permutation14(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p14); z++ {
						writter.WriteString(p14[z])
						writter.Flush()
					}
					p15 := permutation15(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p15); z++ {
						writter.WriteString(p15[z])
						writter.Flush()
					}
					p16 := permutation16(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p16); z++ {
						writter.WriteString(p16[z])
						writter.Flush()
					}
					p17 := permutation17(subdomainline, domain)
					if p17[0] != "NONE" {
						for z := 0; z < len(p17); z++ {
							writter.WriteString(p17[z])
							writter.Flush()
						}
					}
					p18 := permutation18(subdomainline, domain)
					if p18[0] != "NONE" {
						for z := 0; z < len(p18); z++ {
							writter.WriteString(p18[z])
							writter.Flush()
						}
					}
					p19 := permutation19(subdomainline, domain)
					if p19[0] != "NONE" {
						for z := 0; z < len(p19); z++ {
							writter.WriteString(p19[z])
							writter.Flush()
						}
					}
					p20 := permutation20(subdomainline, domain)
					if p20[0] != "NONE" {
						for z := 0; z < len(p20); z++ {
							writter.WriteString(p20[z])
							writter.Flush()
						}
					}
				}
			} else {
				for i := 0; i < len(wordslines); i++ {
					writter.WriteString(permutation1(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation2(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation3(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation4(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation7(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					writter.WriteString(permutation8(wordslines[i], subdomainline, domain)+"\n")
					writter.Flush()
					p9 := permutation9(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p9); z++ {
						writter.WriteString(p9[z])
						writter.Flush()
					}
					p10 := permutation10(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p10); z++ {
						writter.WriteString(p10[z])
						writter.Flush()
					}
					p11 := permutation11(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p11); z++ {
						writter.WriteString(p11[z])
						writter.Flush()
					}
					p12 := permutation12(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p12); z++ {
						writter.WriteString(p12[z])
						writter.Flush()
					}
					p13 := permutation13(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p13); z++ {
						writter.WriteString(p13[z])
						writter.Flush()
					}
					p14 := permutation14(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p14); z++ {
						writter.WriteString(p14[z])
						writter.Flush()
					}
					p15 := permutation15(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p15); z++ {
						writter.WriteString(p15[z])
						writter.Flush()
					}
					p16 := permutation16(wordslines[i], subdomainline, domain)
					for z := 0; z < len(p16); z++ {
						writter.WriteString(p16[z])
						writter.Flush()
					}
					p17 := permutation17(subdomainline, domain)
					if p17[0] != "NONE" {
						for z := 0; z < len(p17); z++ {
							writter.WriteString(p17[z])
							writter.Flush()
						}
					}
					p18 := permutation18(subdomainline, domain)
					if p18[0] != "NONE" {
						for z := 0; z < len(p18); z++ {
							writter.WriteString(p18[z])
							writter.Flush()
						}
					}
					p19 := permutation19(subdomainline, domain)
					if p19[0] != "NONE" {
						for z := 0; z < len(p19); z++ {
							writter.WriteString(p19[z])
							writter.Flush()
						}
					}
					p20 := permutation20(subdomainline, domain)
					if p20[0] != "NONE" {
						for z := 0; z < len(p20); z++ {
							writter.WriteString(p20[z])
							writter.Flush()
						}
					}
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	inputfile.Close()
	outputfile.Close()
	sortfileunique(*output)
}
