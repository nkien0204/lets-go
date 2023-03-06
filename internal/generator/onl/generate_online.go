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
	"path/filepath"

	"github.com/nkien0204/lets-go/internal/configs"
)

const GITHUB_REPO_ENDPOINT string = "https://api.github.com/repos/nkien0204/lets-go"

type OnlineGenerator struct {
	ProjectName string
}

func (onl *OnlineGenerator) Generate() error {
	if onl.ProjectName == "" {
		return errors.New("project name must be identified, please use -p flag")
	}
	tag, err := onl.getLatestVersion()
	if err != nil {
		return err
	}
	if err := onl.downloadLatestAsset(tag); err != nil {
		return err
	}
	return onl.copyConfig()
}

func (onl *OnlineGenerator) copyConfig() error {
	cmd := exec.Command("cp", "-n", configs.CONFIG_FILENAME_SAMPLE, configs.CONFIG_FILENAME)
	return cmd.Run()
}

func (onl *OnlineGenerator) getLatestVersion() (string, error) {
	resp, err := http.Get(GITHUB_REPO_ENDPOINT + "/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var latestReleaseInfo LatestReleaseInfo
	err = json.Unmarshal(body, &latestReleaseInfo)
	if err != nil {
		return "", err
	}
	return latestReleaseInfo.TagName, nil
}

func (onl *OnlineGenerator) downloadLatestAsset(tagName string) error {
	apiEndpoint := fmt.Sprintf(GITHUB_REPO_ENDPOINT+"/zipball/"+"%s", tagName)
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
	zipFileName := onl.ProjectName + ".zip"
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
		if err := os.Rename(unZipDir, onl.ProjectName); err != nil {
			fmt.Println("error: ", err)
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
		fmt.Println("unzipping file ", f.Name)

		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
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
