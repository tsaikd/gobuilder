package checkfmt

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/pkgutil"
)

// errors
var (
	ErrorUnformat1 = errutil.NewFactory("found unformatted go source: %q")
)

// Check go source code are all already formated in pkglist
func Check(pkglist *pkgutil.PackageList) (err error) {
	errs := []error{}
	for _, pkg := range pkglist.Sorted() {
		for _, name := range pkg.GoFiles {
			fpath := filepath.Join(pkg.Dir, name)
			if err = checkGoFile(fpath); err != nil {
				errs = append(errs, err)
			}
		}
		for _, name := range pkg.TestGoFiles {
			fpath := filepath.Join(pkg.Dir, name)
			if err = checkGoFile(fpath); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errutil.NewErrors(errs...)
}

func checkGoFile(fpath string) (err error) {
	buffer := &bytes.Buffer{}
	cmd := exec.Command("gofmt", "-d", fpath)
	cmd.Stdout = buffer
	cmd.Stderr = buffer
	if err = cmd.Run(); err != nil {
		return
	}
	if buffer.Len() > 0 {
		return ErrorUnformat1.New(nil, fpath)
	}
	return nil
}
