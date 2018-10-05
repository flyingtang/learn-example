package main

import (
	"os"
	"fmt"
	"flag"
	"errors"
	"io/ioutil"
	"path/filepath"
	"io"
	"fileupload/genesis"
	"strings"

)


//
func main(){

	as := os.Args
	fmt.Println(as)
	if err := run(os.Args[1:]); err == flag.ErrHelp {
		os.Exit(1)
	}else if err != nil{
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run(args []string) error {

	fs := flag.NewFlagSet("genesis", flag.ContinueOnError)

	CWD := fs.String("C", "", "")
	out := fs.String("o", "", "")
	pkg := fs.String("pkg", "", "")
	tags := fs.String("tags", "","")

	fs.Usage = usage
	fmt.Println("NArg => ", fs.NArg())
	if err := fs.Parse(args); err != nil {
		return err
	} else if fs.NArg() == 0 {
		usage()
		return flag.ErrHelp
	}else  if *pkg == ""{
		return errors.New("package name required")
	}
	fmt.Println("NArg2 => ", fs.NArg())


	if *CWD != "" {
		if err := os.Chdir(*CWD); err != nil{

			return err
		}
	}

	var paths []string
	for i, arg := range fs.Args(){
		fmt.Println(i, arg)
		if a, err := expand(arg); err != nil {
			return err
		}else {
			paths = append(paths, a...)
		}
	}

	var w io.Writer
	if *out == "" {
		w = os.Stdout
	}else {
		f, err := os.Create(*out)
		if err != nil {
			return err
		}
		defer f.Close()
		w = f
	}

 	enc :=  genesis.NewEncoder(w)
	enc.Package = *pkg
	enc.Tags = strings.Split(*tags, ",")

	for _, path := range paths {
		if fi, err := os.Stat(path); err != nil {
			return err
		}else{
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			if err := enc.Encode(&genesis.Asset{
				Name: PrePendSlash(filepath.ToSlash(path)),
				Data: data,
				ModTime: fi.ModTime(),
			}); err != nil {
				return  err
			}
		}
	}
	if err := enc.Close(); err != nil {
		return err
	}
	return nil
}

func usage() {
	fmt.Print(`usage: genesis [options] path [paths]

Embeds listed assets in a Go file as hex-encoded strings.

The following flags are available:

	-pkg name
		Package name of the generated Go file. Required.

	-o output
		Output filename for generated code. Optional.
		(default stdout)

	-C dir
		Execute genesis from dir. Optional.

	-tags tags
		Comma-delimited list of build tags. Optional.

`)
}

func expand(path string)([]string, error){

	if fi, err := os.Stat(path); err != nil {
		return nil, err
	} else if !fi.IsDir() {
		return []string{path}, nil
	}

	fis, err := ioutil.ReadDir(path)
	if  err != nil {
		return nil, err
	}

	expanded := make([]string, 0, len(fis))
	for _, fi := range fis {
		a, err := expand(filepath.Join(path, fi.Name()))
		if err != nil {
			return nil, err
		}
		expanded = append(expanded, a...)
	}

	return expanded, nil
}

func PrePendSlash(path string) string {
	if strings.HasPrefix(path, "/") {
		return  path
	}else  {
		return  "/"+ path
	}
}