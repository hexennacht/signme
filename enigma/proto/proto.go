package proto

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hexennacht/signme/enigma/pkg/exec"
	"github.com/hexennacht/signme/enigma/pkg/service"
)

var genProto = &cobra.Command{
	Use:     "gen-proto",
	Short:   "Generate protobuf for current service directory",
	Long:    "Generate protobuf for current service directory",
	Example: "jormungand gen-proto --service core-user",
	RunE:    runGenProto,
}

func init() {
	genProto.Flags().StringP("service", "s", "", "Select service to generate the contract proto")
	genProto.Flags().StringP("language", "l", "go", "Programming Language for output code")
}

func GenProto() *cobra.Command {
	return genProto
}

func runGenProto(cmd *cobra.Command, args []string) error {
	svc, _ := cmd.Flags().GetString("service")

	service, err := service.ReadServiceConfigFile(svc)
	if err != nil {
		return err
	}

	projectPath := fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src/github.com/hexennacht/signme")

	for _, contract := range service.Service.Contract {
		fmt.Println("Generating contract proto for", contract)

		exec.ExecCMD(
			"protoc",
			fmt.Sprintf("--proto_path=%s/contracts/proto", projectPath),
			fmt.Sprintf("--proto_path=%s/contracts/third_party", projectPath),
			fmt.Sprintf("--go_out=paths=source_relative:%s/services/%s/grpc", projectPath, service.Name),
			fmt.Sprintf("--go-http_out=paths=source_relative:%s/services/%s/grpc", projectPath, service.Name),
			fmt.Sprintf("--go-grpc_out=paths=source_relative:%s/services/%s/grpc", projectPath, service.Name),
			fmt.Sprintf("--openapi_out=fq_schema_naming=true,default_response=false:%s/services/%s", projectPath, service.Name),
			fmt.Sprintf("%s/contracts/proto/%s", projectPath, strings.TrimSpace(contract)),
		)
	}

	for _, contract := range service.Service.Dependencies {
		fmt.Println("Generating contract proto client for", contract, os.Getenv("WORKSPACE"))
		exec.ExecCMD(
			"protoc",
			fmt.Sprintf("--proto_path=%s/contracts/proto", projectPath),
			fmt.Sprintf("--proto_path=%s/contracts/third_party", projectPath),
			fmt.Sprintf("--go_out=paths=source_relative:%s/services/%s/client/grpc", projectPath, service.Name),
			fmt.Sprintf("--go-http_out=paths=source_relative:%s/services/%s/client/grpc", projectPath, service.Name),
			fmt.Sprintf("--go-grpc_out=paths=source_relative:%s/services/%s/client/grpc", projectPath, service.Name),
			fmt.Sprintf("--openapi_out=fq_schema_naming=true,default_response=false:%s/services/%s", projectPath, service.Name),
			fmt.Sprintf("%s/contracts/proto/%s", projectPath, strings.TrimSpace(contract)),
		)
	}

	return nil
}
