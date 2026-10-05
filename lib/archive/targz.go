package archive

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// TarGz streams dir as a gzip-compressed tar: regular files and directories,
// with paths relative to dir. Symlinks and other special files are skipped,
// so an upload never follows a link out of the directory.
func TarGz(dir string) (io.ReadCloser, error) {
	fi, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}

	pr, pw := io.Pipe()
	go func() {
		gw := gzip.NewWriter(pw)
		tw := tar.NewWriter(gw)
		err := filepath.WalkDir(dir, func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(dir, p)
			if err != nil || rel == "." {
				return err
			}
			if !d.Type().IsRegular() && !d.IsDir() {
				return nil
			}
			info, err := d.Info()
			if err != nil {
				return err
			}
			hdr, err := tar.FileInfoHeader(info, "")
			if err != nil {
				return err
			}
			hdr.Name = filepath.ToSlash(rel)
			if d.IsDir() {
				hdr.Name += "/"
			}
			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			f, err := os.Open(p)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(tw, f)
			return err
		})
		if err == nil {
			err = tw.Close()
		}
		if err == nil {
			err = gw.Close()
		}
		pw.CloseWithError(err)
	}()
	return pr, nil
}

var errUnsafePath = errors.New("archive entry escapes the target directory")

// UntarGz extracts a gzip-compressed tar into dst. Only directories and
// regular files are created; links, devices and other special entries are
// skipped. An entry whose path would land outside dst (absolute, or
// climbing with ..) aborts the extraction.
func UntarGz(r io.Reader, dst string) error {
	gr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gr.Close()

	root, err := filepath.Abs(dst)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return err
	}

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		target, err := containedPath(root, hdr.Name)
		if err != nil {
			return fmt.Errorf("%q: %w", hdr.Name, err)
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, hdr.FileInfo().Mode().Perm()&0o755|0o600)
			if err != nil {
				return err
			}
			_, err = io.Copy(f, tr)
			if cerr := f.Close(); err == nil {
				err = cerr
			}
			if err != nil {
				return err
			}
		default:
			// links (which could point outside dst), devices, fifos, pax
			// globals: never created
		}
	}
}

func containedPath(root, name string) (string, error) {
	if name == "" || filepath.IsAbs(name) || strings.HasPrefix(name, "/") {
		return "", errUnsafePath
	}
	target := filepath.Join(root, filepath.FromSlash(name))
	rel, err := filepath.Rel(root, target)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", errUnsafePath
	}
	return target, nil
}
