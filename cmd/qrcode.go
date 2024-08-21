package cmd

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

// Variable for output file and content
var (
	outputFile string
	content    string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a QRCode from the provided text content.",
	Long: `The generate command creates a QR code from the provided text content.

You can specify the text you want to encode into the QR code using the --content flag. Additionally, you can set the output file name for the QR code image with the --output flag.

Examples:
  Generate a QR code with the content "Hello, World!" and save it to "hello_qrcode.png":
    CobraQRCodeGenerator generate --content="Hello, World!" --output="hello_qrcode.png"

  Generate a QR code with the content "https://example.com" and use the default file name:
    CobraQRCodeGenerator generate --content="https://example.com"

Flags Details:
  --content string   The text to encode in the QR code (required).
  --output string    The output file for the QR code image (default is qrcode.png).`,
	Run: generateQrCode,
}

// init function
func init() {
	// Add flags for output file and content
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "qrcode.png", "output file for the QR code (default is qrcode.png)")
	generateCmd.Flags().StringVarP(&content, "content", "c", "", "content to encode in the QR code")

	// Mark flag content as required
	generateCmd.MarkFlagRequired("content")

	// Register the generate Cmd command as a subcommand of the root command
	rootCmd.AddCommand(generateCmd)
}

func generateQrCode(cmd *cobra.Command, args []string) {
	err := qrcode.WriteFile(content, qrcode.Medium, 256, outputFile)
	if err != nil {
		fmt.Printf("Error generating QR Code: %v\n", err)
		return
	}

	fmt.Printf("QR Code generated successfully and saved to %s\n", outputFile)
}
