package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"debug/pe"
)

type Args struct {
	basePath string
	singedOnly bool
}

func parse_args() (Args) {
	args := Args{}
	flag.StringVar(
		&args.basePath,
		"base-path",
		"C:\\",
		"base path for search",
	)
	flag.BoolVar(
		&args.singedOnly,
		"dll-path",
		false,
		"filter unsigned ddls",
	)
	flag.Parse()
	return args
}



// This function will be called recursively on every directory
// TODO: make this multi threaded
func find_rwx_dlls(base_path string) ([]string) {
	directory_entries, err := os.ReadDir(base_path)
	if err != nil {
		fmt.Printf("[-] Failed to read directory %s:\n%v\n", base_path, err);
		return []string{}
	}
	
	rwx_dlls := []string{}
	for _, directory_entry := range directory_entries {
		full_path := strings.ToLower(base_path + "\\" + directory_entry.Name()) 
		if directory_entry.IsDir() {
			rwx_dll_in_subdirectory := find_rwx_dlls(full_path)
			rwx_dlls = append(rwx_dlls, rwx_dll_in_subdirectory...)
			continue	
		} 

		if !directory_entry.Type().IsRegular() {
			continue
		}

		if !strings.HasSuffix(full_path, ".dll") {
			continue
		}
		// file is a dll we now check if it has a RWX section
		if !does_dll_have_rwx_section(full_path) {
			continue
		}

		rwx_dlls = append(rwx_dlls, full_path)
	}
	return rwx_dlls
}

func does_dll_have_rwx_section(dll_path string) (bool) {
	pe_file, err := pe.Open(dll_path)
	if err != nil {
		fmt.Printf(
			"[-] failed to open dll %s for parsing:\n%v", 
			dll_path,
			err,
		)
		return false
	}

	rwx_flags := uint32(
		pe.IMAGE_SCN_MEM_READ | 
		pe.IMAGE_SCN_MEM_WRITE | 
		pe.IMAGE_SCN_MEM_EXECUTE,
	)
	for _, section := range pe_file.Sections {
		if (section.Characteristics & rwx_flags) == rwx_flags {
			return true
		}
	}
	
	return false
}

func main() {
	args := parse_args()
	rwx_dlls := find_rwx_dlls(args.basePath)
	
	number_of_rwx_dlls_found := len(rwx_dlls)
	if number_of_rwx_dlls_found == 0 {
		fmt.Println("[-] no dlls with rwx sections found :'3")
		os.Exit(-1)
		return
	}
	fmt.Printf("[+] found %d dlls with rwx sections\n", number_of_rwx_dlls_found)

	for _, rwx_dll := range rwx_dlls {
		fmt.Printf("[+] %s\n", rwx_dll)
	}

	os.Exit(0)
	return
}
