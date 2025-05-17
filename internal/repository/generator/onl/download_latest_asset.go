package onl

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

func (repo *repository) DownloadLatestAsset(
	requestEntity generator.LatestAssetDownloadRequestEntity,
) error {
	apiEndpoint := fmt.Sprintf(repo.gen.RepoEndPoint+"/zipball/"+"%s", requestEntity.TagName)
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
	zipFileName := requestEntity.ProjectName + ".zip"
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
		if err := os.Rename(unZipDir, requestEntity.ProjectName); err != nil {
			os.RemoveAll(unZipDir)
		}
	}()
	if _, err := f.Write(body); err != nil {
		return err
	}

	unZipDir, err = repo.unZip(zipFileName)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) unZip(zipFile string) (string, error) {
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
