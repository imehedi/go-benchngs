package fastachopper

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Fasta struct {
	id   string
	desc string
	seq  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildFasta(header string, seq bytes.Buffer) (record Fasta) {
	fields := strings.SplitN(header, " ", 2)

	if len(fields) > 1 {
		record.id = fields[0]
		record.desc = fields[1]
	} else {
		record.id = fields[0]
		record.desc = ""
	}

	record.seq = seq.String()

	return record
}

func parse(FastaFh io.Reader) chan Fasta {

	outputChannel := make(chan Fasta)

	scanner := bufio.NewScanner(FastaFh)
	// scanner.Split(bufio.ScanLines)
	header := ""
	var seq bytes.Buffer

	go func() {
		// Loop over the letters in inputString
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) == 0 {
				continue
			}

			// line := scanner.Text()

			if line[0] == '>' {
				// If we stored a previous identifier, get the DNA string and map to the
				// identifier and clear the string
				if header != "" {
					// outputChannel <- build_Fasta(header, seq.String())
					outputChannel <- buildFasta(header, seq)
					// fmt.Println(record.id, len(record.seq))
					header = ""
					seq.Reset()
				}

				// Standard Fasta identifiers look like: ">id desc"
				header = line[1:]
			} else {
				// Append here since multi-line DNA strings are possible
				seq.WriteString(line)
			}

		}

		outputChannel <- buildFasta(header, seq)

		// Close the output channel, so anything that loops over it
		// will know that it is finished.
		close(outputChannel)
	}()

	return outputChannel
}

func Chopper(filename string) {

	FastaFh, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer FastaFh.Close()

	lengthArray := [5]int{50, 100, 150, 200, 400}
	for record := range parse(FastaFh) {
		for i := 0; i < len(lengthArray); i++ {

			newFileName := fmt.Sprintf("%s_%d", record.id, lengthArray[i])
			f, err := os.Create(newFileName)
			check(err)
			header := fmt.Sprintf(">%s_%d | %s - %d \n", record.id, lengthArray[i], record.desc, len(record.seq))
			_, err = f.Write([]byte(header))
			check(err)
			var j int
			if i == 0 {
				j = 0
			} else {
				j = i - 1
			}
			body := fmt.Sprintf(record.seq[lengthArray[j]:lengthArray[i]])
			_, err = f.Write([]byte(body))
			check(err)

		}
	}
}
