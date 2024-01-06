package main

import (
	"encoding/csv"
	"os"
)

func main() {

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"ali", "fahrial", "anwar"})
	_ = writer.Write([]string{"eko", "kurniawan", "wardadi"})
	_ = writer.Write([]string{"anji", "edi", "wardi"})

	writer.Flush()

}
