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

	"github.com/nkien0204/lets-go/internal/entity/config"
	"github.com/nkien0204/lets-go/internal/entity/generator"
)

func (u *usecase) Generate() error {
	if u.gen.ProjectName == "" {
		return errors.New("project name must be identified, please use -p flag")
	}
	tag, err := u.getLatestVersion()
	if err != nil {
		return err
	}
	if err := u.downloadLatestAsset(tag); err != nil {
		return err
	}
	if err := u.copyConfig(); err != nil {
		return err
	}
	return u.removeGenerator()
}

func (u *usecase) removeGenerator() error {
	genCmdFilePath := path.Join(u.gen.ProjectName, "cmd", "gen.go")
	genDeliveryPath := path.Join(u.gen.ProjectName, "internal", "delivery", "generator")
	genUsecasePath := path.Join(u.gen.ProjectName, "internal", "usecase", "generator")
	samplesPath := path.Join(u.gen.ProjectName, "samples")
	sampleConfigFilePath := path.Join(u.gen.ProjectName, config.CONFIG_FILENAME_SAMPLE)
	if err := os.Remove(genCmdFilePath); err != nil {
		return err
	}
	if err := os.Remove(genDeliveryPath); err != nil {
		return err
	}
	if err := os.Remove(genUsecasePath); err != nil {
		return err
	}
	if err := os.Remove(samplesPath); err != nil {
		return err
	}
	if err := os.Remove(sampleConfigFilePath); err != nil {
		return err
	}
	return nil
}

func (u *usecase) copyConfig() error {
	var cmd *exec.Cmd
	src := filepath.Join(u.gen.ProjectName, config.CONFIG_FILENAME_SAMPLE)
	dst := filepath.Join(u.gen.ProjectName, config.CONFIG_FILENAME)

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("copy", src, dst)
	default:
		cmd = exec.Command("cp", "-n", src, dst)
	}
	return cmd.Run()
}

func (u *usecase) getLatestVersion() (string, error) {
	resp, err := http.Get(generator.GITHUB_REPO_ENDPOINT + "/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var latestReleaseInfo generator.LatestReleaseInfo
	err = json.Unmarshal(body, &latestReleaseInfo)
	if err != nil {
		return "", err
	}
	return latestReleaseInfo.TagName, nil
}

func (u *usecase) downloadLatestAsset(tagName string) error {
	apiEndpoint := fmt.Sprintf(generator.GITHUB_REPO_ENDPOINT+"/zipball/"+"%s", tagName)
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
	zipFileName := u.gen.ProjectName + ".zip"
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
		if err := os.Rename(unZipDir, u.gen.ProjectName); err != nil {
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
