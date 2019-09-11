/*Parse resource.h define ids to Golang const

 */
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"
)

var (
	filename    = flag.String("filename", "", "filename")
	packagename = flag.String("packagename", "main", "packagename")
)

func main() {
	flag.Parse()
	log.Println("genids filename", *filename)
	if *filename == "" {
		log.Panic("no filename")
	}
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Panic(err)
	}
	code := match(data)
	err = ioutil.WriteFile(path.Base(*filename)+".go", []byte(code), 0777)
	if err != nil {
		log.Panic(err)
	}
	log.Println("gen ids finish")
}

func match(data []byte) string {
	reg := regexp.MustCompile(`#define\s+(\w+)\s+(\d+)\w*`)
	v := reg.FindAllSubmatch(data, -1)
	log.Println("Match length ", len(v))
	var builder strings.Builder
	builder.WriteString("// Code generated by genids from ")
	builder.WriteString(*filename)
	builder.WriteString("DO NOT EDIT\n\npackage ")
	builder.WriteString(*packagename)
	builder.WriteString("\n\nconst (\n")
	for _, a := range v {
		builder.WriteString("	// ")
		builder.Write(a[1])
		builder.WriteString(" const uintptr\n")
		builder.WriteString("	")
		builder.Write(a[1])
		builder.WriteString(" = ")
		builder.Write(a[2])
		builder.WriteString("\n")
	}
	builder.WriteString(")\n")
	return builder.String()
}
