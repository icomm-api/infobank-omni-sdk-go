package regist

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/core"
	"github.com/icomm-api/infobank-omni-sdk-go/pkg/infobank/models"
)

const (
	fileSubPath = "/v1/file"
)

type FileUploader struct {
	Client *core.HttpClient
}

func (f *FileUploader) UploadFile(serviceType models.FileServiceTypeEnum, msgType models.FileMsgTypeEnum, filePath string) (*models.FileResponse, error) {
	if len(serviceType) == 0 {
		return nil, errors.New("serviceType is required")
	}

	subPath := fileSubPath + "/" + string(serviceType)
	if len(msgType) > 0 {
		subPath += "/" + string(msgType)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	mimeType := detectMimeType(filePath)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	formFile, err := writer.CreatePart(getMIMEHeader("file", filepath.Base(filePath), mimeType))
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(formFile, file); err != nil {
		return nil, fmt.Errorf("failed to copy file to form: %w", err)
	}

	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %w", err)
	}

	response, err := core.UploadFile(f.Client, subPath, contentType, body)
	if err != nil {
		return nil, err
	}

	return core.ParseFileResponse(f.Client, response)
}

func detectMimeType(filePath string) string {
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return mimeType
}

func getMIMEHeader(fieldName, fileName string, mimeType string) textproto.MIMEHeader {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, fileName))
	header.Set("Content-Type", mimeType)
	return header
}
