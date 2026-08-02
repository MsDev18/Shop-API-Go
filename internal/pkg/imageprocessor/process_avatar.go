package imageprocessor

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"shop/internal/pkg/richerror"

	"github.com/disintegration/imaging"
)

func (p Processor) ProcessAvatar(fileHeader *multipart.FileHeader, userID uint) (string, error) {
	const op = "imageprocessor.ProcessAvatar"

	// 1. reject oversized files before reading any bytes - fileHeader.Size
	// comes from request metadata, so this check is free
	maxSizeBytes := p.config.MaxSizeMB * 1024 * 1024
	if fileHeader.Size > maxSizeBytes {
		return "", richerror.New().
			SetOp(op).
			SetMsg(fmt.Sprintf("file to large, max allowed size is %d MB", p.config.MaxSizeMB)).
			SetKind(richerror.KindBadRequestErr)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("can't open uploaded file").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	defer file.Close()
	
	// 2. check the *real* content of the file - never trust the filename
	// extension or the Content-Type header the client sent, both are
	// trivial to fake
	if !p.isAllowedImage(file) {
		return "", richerror.New().
			SetOp(op).
			SetMsg("only jpeg and png allowed to upload").
			SetKind(richerror.KindBadRequestErr)
	}

	// 3. decode - AutoOrientation fixes photos that look sideways, because
	// phone cameras store rotation as EXIF metadata instead of rotating
	// the actual pixels
	img, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("uploaded file is not a valid image").
			SetKind(richerror.KindBadRequestErr).
			SetErr(err)
	}

	// 4. resize - Fit preserves aspect ratio and only shrinks, it never
	// upscales an already-small image
	resized := imaging.Fit(img, p.config.MaxDimension, p.config.MaxDimension, imaging.Lanczos)

	diskDir := filepath.Join(UPLOADS_ROOT, p.config.Dir)
	if err := os.MkdirAll(diskDir, 0755); err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("can't create upload directory").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}

	// 5. never use the client-supplied filename - build our own, safe,
	// unique one
	fileName, err := p.generatedFileName(userID)
	if err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("can't generated file name").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}

	// 6. re-encode as JPEG at a fixed quality - this is the actual
	// compression step, and it also strips EXIF metadata as a side effect,
	// since the decoded image.Image never carried it forward
	diskPath := filepath.Join(diskDir, fileName)
	if err := imaging.Save(resized, diskPath, imaging.JPEGQuality(80)); err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("can't save processed image").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	urlPath := filepath.Join("/uploads", p.config.Dir, fileName)
	return urlPath , err
}

func (p Processor) isAllowedImage(file multipart.File) bool {
	buf := make([]byte, 512)

	n, err := file.Read(buf)
	// io.EOF is not error but a signal for end of file
	// Read returns io.EOF if the file is empty,
	// which is not an allowed image
	if err != nil && err != io.EOF {
		return false
	}

	// http.DetectContentType only needs the first bytes ("magic number")
	// to identify the real file format
	contentType := http.DetectContentType(buf[:n])

	// rewind - ProcessAvatar still needs to read the whole file from
	// the start in order to decode it
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return false
	}

	return contentType == "image/png" || contentType == "image/jpeg"
}

func (p Processor) generatedFileName(userID uint) (string, error) {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d-%s.jpg", userID, hex.EncodeToString(buf)), nil
}
