package main

import (
	`fmt`
	`github.com/flywithbug/file`
	`github.com/flywithbug/fileanalysis`
)

func main()  {
	path,err := fileanalysis.CurrentDirectory()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)


	file.EnumeratePath(path, func(surplus string, current string, stop *bool) {
		fmt.Println(surplus,":",current)
	})
}
