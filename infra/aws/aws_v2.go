package aws

import (
	"context"
	"log"
	"mime/multipart"
	"path/filepath"

	aws_config "github.com/aws/aws-sdk-go-v2/config"
	aws_manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	aws_s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/ngobrut/halo-suster-api/config"
)

type IFaceAWS interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (*UploadResponse, error)
}

type AWS struct {
	cnf      config.AWS
	uploader *aws_manager.Uploader
}

type UploadResponse struct {
	Size     int
	Mimetype string
	Name     string
	Location string
}

func NewAWSClient(cnf config.AWS) (IFaceAWS, error) {
	cfg, err := aws_config.LoadDefaultConfig(context.TODO(), aws_config.WithRegion(cnf.Region))
	if err != nil {
		return nil, err
	}

	client := aws_s3.NewFromConfig(cfg)
	uploader := aws_manager.NewUploader(client)

	log.Println("[aws-connected]")

	return &AWS{cnf: cnf, uploader: uploader}, nil
}

// UploadFile implements IFaceAWS.
func (a *AWS) UploadFile(ctx context.Context, file *multipart.FileHeader) (*UploadResponse, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer f.Close()

	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	mimetype := file.Header.Get("Content-Type")

	fileobj := &aws_s3.PutObjectInput{
		Key:         aws.String(filename),
		Bucket:      aws.String(a.cnf.Bucket),
		Body:        f,
		ContentType: aws.String(mimetype),
	}

	result, err := a.uploader.Upload(context.TODO(), fileobj)
	if err != nil {
		return nil, err
	}

	res := &UploadResponse{
		Size:     int(file.Size),
		Mimetype: mimetype,
		Name:     filename,
		Location: result.Location,
	}

	return res, nil
}
