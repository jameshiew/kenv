package internal

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
)

func Install(version, goOS, goArchitecture string) error {
	exists, err := existsVersion(version)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("version folder for %v already exists", version)
	}
	return initializeKubectlVersion(version, goOS, goArchitecture)
}

func constructKubectlURL(version, goOS, goArchitecture string) string {
	return fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/v%s/bin/%s/%s/kubectl", version, goOS, goArchitecture)
}

func initializeKubectlVersion(version, goOS, goArchitecture string) error {
	random := fmt.Sprintf("%v", rand.Int63())
	tmpDir := filepath.Join(os.TempDir(), "kenv", random)
	err := os.MkdirAll(tmpDir, 0o744)
	if err != nil {
		return err
	}
	kubectlPath := filepath.Join(tmpDir, "kubectl")
	err = downloadKubectl(kubectlPath, version, goOS, goArchitecture)
	if err != nil {
		return err
	}
	versionDir, err := prefixDir(version)
	if err != nil {
		return err
	}
	err = os.MkdirAll(versionDir, 0o744)
	if err != nil {
		return err
	}
	newKubectlPath := filepath.Join(versionDir, "kubectl")
	err = os.Rename(kubectlPath, newKubectlPath)
	if err != nil {
		return err
	}
	return nil
}

func downloadKubectl(filepath, version, goOS, goArchitecture string) error {
	// download appropriate kubectl binary for version
	url := constructKubectlURL(version, goOS, goArchitecture)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("could not find kubectl version %v for %v/%v (HTTP 404: %v)", version, goOS, goArchitecture, url)
	}
	fmt.Printf("Downloading kubectl v%v for %v/%v from %v...\n", version, goOS, goArchitecture, url)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("WARN - error while closing request body: %v\n", err)
		}
	}()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Printf("WARN - error while closing file: %v\n", err)
		}
	}()

	bar := pb.Full.Start64(resp.ContentLength)
	barReader := bar.NewProxyReader(resp.Body)
	_, err = io.Copy(out, barReader)
	if err != nil {
		return err
	}
	bar.Finish()
	err = os.Chmod(filepath, 0o744)
	if err != nil {
		return err
	}
	return nil
}
