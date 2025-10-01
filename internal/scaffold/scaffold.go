package scaffold

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type Options struct {
	ModuleName string
	GitInit    bool
	RunTidy    bool
}

func CopyTemplate(dest string, opts Options) error {
	src := "../../templates/wells-go"

	if _, err := os.Stat(src); os.IsNotExist(err) {
		return errors.New("templates/wells-go not found")
	}

	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dest, rel)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if isBinary(b) {
			if err := os.WriteFile(targetPath, b, info.Mode()); err != nil {
				return err
			}
			return nil
		}

		tpl, err := template.New(rel).Parse(string(b))
		if err != nil {
			if err := os.WriteFile(targetPath, b, info.Mode()); err != nil {
				return err
			}
			return nil
		}

		if err := os.WriteFile(targetPath, []byte{}, info.Mode()); err != nil {
			return err
		}
		f, err := os.OpenFile(targetPath, os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer f.Close()

		data := map[string]interface{}{
			"ModuleName": opts.ModuleName,
		}

		if err := tpl.Execute(f, data); err != nil {
			return err
		}
		return nil
	})
}

func isBinary(b []byte) bool {
	return bytes.Contains(b, []byte{0})
}

func GitInit(dir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}

func GoModTidy(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
