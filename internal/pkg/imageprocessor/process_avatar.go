package imageprocessor

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"shop/internal/pkg/richerror"

	"github.com/disintegration/imaging"
)

func (p Processor) ProcessAvatar(fileHeader *multipart.FileHeader, userID uint) (string, error) {
	const op = "imageprocessor.ProcessAvatar"
	// 1. open file 
	file, err := fileHeader.Open()
	if err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("can't open uploaded file").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}
	defer file.Close()

	// 2. decode - AutoOrientation fixes photos that look sideways, because
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

	// 3. resize - Fit preserves aspect ratio and only shrinks, it never
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

	// 4. never use the client-supplied filename - build our own, safe,
	// unique one
	fileName, err := p.generatedFileName(userID)
	if err != nil {
		return "", richerror.New().
			SetOp(op).
			SetMsg("can't generated file name").
			SetKind(richerror.KindUnexpectedErr).
			SetErr(err)
	}

	// 5. re-encode as JPEG at a fixed quality - this is the actual
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

func (p Processor) generatedFileName(userID uint) (string, error) {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d-%s.jpg", userID, hex.EncodeToString(buf)), nil
}
