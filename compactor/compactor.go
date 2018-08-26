package compactor

import (
	"strings"
)

func Process(txt string, l int) string {
	f := strings.Fields(txt)

	output := ""
	linelen := 0
	for _, s := range f {
		if linelen < l {
			output += s
			linelen += len(s)
		}

		if linelen >= l {
			output += "\n"
			linelen = 0
		}
	}

	return output

	// ioutil.WriteFile("out.txt", []byte(output), os.ModePerm)
	// return nil
}
