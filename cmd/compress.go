package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// compressCmd represents the version command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "compress .png/.jpg by use tinypng",
	Long:  `compress .png/.jpg by use tinypng`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start compress", *srcPath, *output)
		compressRun(*srcPath, *output)
	},
}

var srcPath *string
var output *string

func init() {
	rootCmd.AddCommand(compressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	srcPath = compressCmd.PersistentFlags().StringP("src", "s", "", "src file/dir path")
	output = compressCmd.PersistentFlags().StringP("output", "0", "", "output file/dir path")
}

func compressRun(src, o string) {
	if src == "" {
		fmt.Println("please input src file path")
		return
	}
	var s []string

	if isDir(src) {
		// 目录
		s, _ = GetAllFile(src, s)
	} else {
		s = append(s, src)
	}
	//fmt.Printf("slice: %v", s)

	if len(s) <= 0 {
		fmt.Printf("Not Found png/jpg file\n")
		return
	}
	fmt.Printf("Find png/jpg file len %d\n", len(s))

	//tinify.SetKey("rcPZm3Zrg_1DbjYtV6AXM_-53Jg9wuWB")
	//s, err := tinify.FromFile(src)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//err = s.ToFile(o)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			if strings.HasSuffix(fullName, ".png") || strings.HasSuffix(fullName, ".jpg") {
				s = append(s, fullName)
			}
		}
	}
	return s, nil
}
