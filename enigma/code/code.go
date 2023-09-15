package code

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hexennacht/signme/enigma/pkg/exec"
	"github.com/hexennacht/signme/enigma/pkg/service"
)

var genCode = &cobra.Command{
	Use:     "gen-code",
	Short:   "Generate code for current service directory",
	Long:    "Generate code for current service directory",
	Example: "enigma gen-code --codegen code",
	RunE:    runGenCode,
}

var availableCodeGen = map[string]bool{
	"code":   true,
	"config": true,
}

var (
	codeGenCode   = "code"
	codeGenConfig = "config"
)

var errorGenerateFoundNotFound = errors.New("codegen not valid")

func init() {
	genCode.Flags().StringP("codegen", "c", "code", "Select which generate code | config")
}

func GenCode() *cobra.Command {
	return genCode
}

func runGenCode(cmd *cobra.Command, args []string) error {
	genCode, _ := cmd.Flags().GetString("codegen")
	if !availableCodeGen[genCode] {
		return errorGenerateFoundNotFound
	}

	_, err := service.ReadServiceConfigFile("")
	if err != nil {
		return err
	}

	if strings.EqualFold(genCode, codeGenCode) {
		exec.ExecCMD("go", "generate", "./...")
	}

	if strings.EqualFold(genCode, codeGenConfig) {
		err = genConfig()
		if err != nil {
			return err
		}
	}

	return nil
}

func genConfig() error {
	fsFile, err := ioutil.ReadDir("./internal/conf")
	if err != nil {
		return err
	}

	for _, file := range fsFile {
		if !strings.Contains(file.Name(), ".proto") {
			continue
		}

		fmt.Println("Generating", file.Name())
		exec.ExecCMD(
			"protoc",
			"--proto_path=./internal",
			fmt.Sprintf("--proto_path=%s/proto/third_party", os.Getenv("WORKSPACE")),
			"--go_out=paths=source_relative:./internal",
			fmt.Sprintf("./internal/conf/%s", file.Name()),
		)
	}

	return nil
}
