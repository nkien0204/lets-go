package onl

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"github.com/nkien0204/lets-go/internal/entities/configs"
	"github.com/nkien0204/lets-go/internal/entities/generators"
)

type onlGenAdapter struct {
    gen *generators.OnlineGenerator
}

func NewOnlGenAdapter(gen *generators.OnlineGenerator) *onlGenAdapter {
    return &onlGenAdapter{gen: gen}
}

func (onl *onlGenAdapter) Generate() error {
	if onl.gen.ProjectName == "" {
		return errors.New("project name must be identified, please use -p flag")
	}
	tag, err := onl.getLatestVersion()
	if err != nil {
		return err
	}
	if err := onl.downloadLatestAsset(tag); err != nil {
		return err
	}
	if err := onl.copyConfig(); err != nil {
		return err
	}
	return onl.removeGenerator()
}

func (onl *onlGenAdapter) removeGenerator() error {
	// remove cmd/gen.go
	genCmdFilePath := path.Join(onl.gen.ProjectName, "cmd", "gen.go")
	return os.Remove(genCmdFilePath)
}

func (onl *onlGenAdapter) copyConfig() error {
	var cmd *exec.Cmd
	src := filepath.Join(onl.gen.ProjectName, configs.CONFIG_FILENAME_SAMPLE)
	dst := filepath.Join(onl.gen.ProjectName, configs.CONFIG_FILENAME)

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("copy", src, dst)
	default:
		cmd = exec.Command("cp", "-n", src, dst)
	}
	return cmd.Run()
}

func (o *onlGenAdapter) getLatestVersion() (string, error) {
	resp, err := http.Get(generators.GITHUB_REPO_ENDPOINT + "/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var latestReleaseInfo generators.LatestReleaseInfo
	err = json.Unmarshal(body, &latestReleaseInfo)
	if err != nil {
		return "", err
	}
	return latestReleaseInfo.TagName, nil
}

func (o *onlGenAdapter) downloadLatestAsset(tagName string) error {
	apiEndpoint := fmt.Sprintf(generators.GITHUB_REPO_ENDPOINT+"/zipball/"+"%s", tagName)
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var unZipDir string
	zipFileName := o.gen.ProjectName + ".zip"
	if _, err := os.Stat(zipFileName); err == nil || !errors.Is(err, fs.ErrNotExist) {
		return errors.New(zipFileName + " was exist")
	}
	f, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		os.Remove(zipFileName)
		if err := os.Rename(unZipDir, o.gen.ProjectName); err != nil {
			fmt.Println("error: ", err)
			os.RemoveAll(unZipDir)
		}
	}()
	if _, err := f.Write(body); err != nil {
		return err
	}

	unZipDir, err = unZip(zipFileName)
	if err != nil {
		return err
	}
	return nil
}

func unZip(zipFile string) (string, error) {
	archive, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}
	defer archive.Close()
	var unZipDir string
	if len(archive.File) > 0 {
		unZipDir = filepath.Dir(archive.File[0].Name)
	}

	for _, f := range archive.File {
		if f.FileInfo().IsDir() {
			os.MkdirAll(f.Name, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(f.Name), os.ModePerm); err != nil {
			return "", err
		}

		dstFile, err := os.OpenFile(f.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return "", err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			dstFile.Close()
			return "", err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			dstFile.Close()
			fileInArchive.Close()
			return "", err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return unZipDir, nil
}
