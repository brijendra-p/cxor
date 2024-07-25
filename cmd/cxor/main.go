package main

import (
	"cxor/tool"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var tools tool.Tool
	var isFile bool
	var rootCmd = &cobra.Command{
		Use:   "cxor",
		Short: "CXOR is a encryption and decryption CLI tool.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			tools = tool.NewTool()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			tools = nil
		},
	}

	// Adding a subcommand for encrypting
	var encryptCmd = &cobra.Command{
		Use:     "encrypt [--file src_file key dst_file] [text key dst_file] [text] [key] [dst_file]",
		Short:   "Encrypts a text or data of source file using key and stores in given file",
		Aliases: []string{"e"},
		Example: "cxor encrypt Hello key",
		Args:    cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			var data string

			if isFile {
				byteData, err := os.ReadFile(args[0])
				if err != nil {
					fmt.Println("Error reading file:", err)
				}
				data = string(byteData)
			} else {
				data = args[0]
			}

			encryptedText, err := tools.Encrypt(data, args[1])
			if err != nil {
				fmt.Println(err)
			}

			file, err := os.Create(args[2])
			if err != nil {
				fmt.Println("Error generating file:", err)
			}
			defer file.Close()

			_, err = file.WriteString(encryptedText)
			if err != nil {
				fmt.Println("Error writing file:", err)
			}
			fmt.Println("encrypted data is stored in", file.Name())
		},
	}
	encryptCmd.PersistentFlags().BoolVarP(&isFile, "file", "f", false, "use file as source of data")

	// Adding a subcommand for decrypting
	var decryptCmd = &cobra.Command{
		Use:     "decrypt [src_file] [key] [dst_file]",
		Short:   "Decrypts a text from a file with using key",
		Aliases: []string{"d"},
		Example: "cxor decrypt hello.txt key",
		Args:    cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			file, err := os.ReadFile(args[0])
			if err != nil {
				fmt.Println("Error reading file:", err)
			}

			decryptedText, err := tools.Decrypt(string(file), args[1])
			if err != nil {
				fmt.Println(err)
			}
			if isFile {
				dstFile, err := os.Create(args[2])
				if err != nil {
					fmt.Println("Error generating file:", err)
				}
				defer dstFile.Close()

				_, err = dstFile.WriteString(decryptedText)
				if err != nil {
					fmt.Println("Error writing file:", err)
				}
				fmt.Println("decrypted data is stored in", dstFile.Name())
			} else {
				fmt.Println("decrypted:", decryptedText)
			}
		},
	}
	decryptCmd.PersistentFlags().BoolVarP(&isFile, "file", "f", false, "use file to store data")

	rootCmd.AddCommand(encryptCmd, decryptCmd)

	rootCmd.Execute()
}
