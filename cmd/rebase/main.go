package main

import (
	"flag"
	"log"
	"os"

	"github.com/buildpacks/lifecycle/cmd"

	_ "github.com/pivotal/kpack/internal/logrus/fatal"
	"github.com/pivotal/kpack/pkg/buildchange"
	"github.com/pivotal/kpack/pkg/flaghelpers"
)

const (
	buildSecretsDir = "/var/build-secrets"
)

var (
	runImage       = flag.String("run-image", os.Getenv("RUN_IMAGE"), "The new run image to rebase")
	lastBuiltImage = flag.String("last-built-image", os.Getenv("LAST_BUILT_IMAGE"), "The previous image to rebase")
	buildChanges   = flag.String("build-changes", os.Getenv("BUILD_CHANGES"), "JSON string of build changes and their reason")
	reportFilePath = flag.String("report", os.Getenv("REPORT_FILE_PATH"), "The location at which to write the report.toml")

	basicDockerCredentials  flaghelpers.CredentialsFlags
	dockerCfgCredentials    flaghelpers.CredentialsFlags
	dockerConfigCredentials flaghelpers.CredentialsFlags
	imagePullSecrets        flaghelpers.CredentialsFlags
)

func init() {
	flag.Var(&basicDockerCredentials, "basic-docker", "Basic authentication for docker of the form 'secretname=git.domain.com'")
	flag.Var(&dockerCfgCredentials, "dockercfg", "Docker Cfg credentials in the form of the path to the credential")
	flag.Var(&dockerConfigCredentials, "dockerconfig", "Docker Config JSON credentials in the form of the path to the credential")
	flag.Var(&imagePullSecrets, "imagepull", "Builder Image pull credentials in the form of the path to the credential")
}

func main() {
	flag.Parse()
	tags := flag.Args()
	logger := log.New(os.Stdout, "", 0)

	if err := buildchange.Log(logger, *buildChanges); err != nil {
		logger.Println(err)
	}

	cmd.Exit(rebase(tags, logger))
}

func rebase(tags []string, logger *log.Logger) error { _ = "STUB: not implemented"; return nil }

func logLoadingSecrets(logger *log.Logger, secretsSlices ...[]string) {
	_ = "STUB: not implemented"
	return
}

func combine(credentials ...[]string) []string { _ = "STUB: not implemented"; return nil }
