/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"
)

// pmCmd represents the pm command
var pmCmd = &cobra.Command{
	Use:   "pm",
	Short: "Print packages infos",
	Long:  `Print packages infos`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pm called")
		if *process {
			printProcessInfo()
			return
		}
		if clear {
			clearAppCache(*name)
			return
		}
		printPackageNames(*name)
	},
}

var name *string
var process *bool
var clear bool

func init() {
	rootCmd.AddCommand(pmCmd)

	name = pmCmd.PersistentFlags().StringP("name", "n", "", "Set packageName key")
	process = pmCmd.PersistentFlags().BoolP("process", "p", false, "Print Process info")
	pmCmd.PersistentFlags().BoolVarP(&clear, "clear", "c", false, "Clear App Cache file")
}

func printPackageNames(name string) {
	pmListPackages(name, func(s string) {
		fmt.Print(s)
	})
}

func printProcessInfo() {
	runCommand("adb shell ps", func(s string) {
		print(s)
	})
}

func clearAppCache(name string) {
	pmListPackages(name, func(s string) {
		br := bufio.NewReader(strings.NewReader(s))
		for {
			l, e := br.ReadBytes('\n')
			if e == io.EOF {
				break
			}
			line := string(l)
			if line != "" {
				var inputStr string
				fmt.Printf("Find packageName[%s]，Clear this app cache file? [Y/N]", string(l[0:len(l)-2]))
				fmt.Scanln(&inputStr)
				if inputStr == "Y" || inputStr == "y" {
					fmt.Println("Clear sucess!")
					return
				}
				fmt.Println("End!")
				return
			}
		}

		fmt.Println("End!")
		return
	})
}
