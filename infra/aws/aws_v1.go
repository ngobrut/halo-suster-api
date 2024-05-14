package aws

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/config"
)

type S3 struct {
	cnf     config.AWS
	client  *s3.S3
	session *session.Session
}

func NewS3(cnf config.AWS) S3 {
	opts := session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}

	sess := session.Must(session.NewSessionWithOptions(opts))

	log.Println("[aws-connected]")

	return S3{
		cnf:     cnf,
		client:  s3.New(sess),
		session: sess,
	}
}

func (s *S3) Upload(file multipart.File, header *multipart.FileHeader) (*UploadResponse, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(header.Filename)
	filename := uuid.New().String() + ext
	mimetype := header.Header.Get("Content-Type")

	uploader := s3manager.NewUploader(s.session)
	input := &s3manager.UploadInput{
		Bucket:      aws.String(s.cnf.Bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewBuffer(buf.Bytes()),
		ContentType: aws.String(mimetype),
	}

	result, err := uploader.Upload(input)
	if err != nil {
		return nil, err
	}

	res := &UploadResponse{
		Size:     int(header.Size),
		Mimetype: mimetype,
		Name:     filename,
		Location: result.Location,
	}

	return res, nil
}
