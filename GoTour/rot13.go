package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(p []byte) (int, error) {
	_, err := reader.r.Read(p)
	if err != nil { 
		for i, value := range p {
			// Check for upper case alphabet
			if value >= 65 && value <= 90 {
				value += 13
				if value > 90 {
					value = value - 90 + 64	
				}
				
				p[i] = value
			} else if value >= 97 && value <= 122 {
				// Check for lower case alphabet
				value += 13
				if value > 122 {
					value = value - 122 + 96 	
				}
				
				p[i] = value				
			} else {
				p[i] = value	
			}
		}
	}
	
	return len(p), err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
