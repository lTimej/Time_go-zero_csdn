package minIO

import (
	"context"
	"liujun/Time_go-zero_csdn/common/utils"
	"net/http"
	"path"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinIOUpload(r *http.Request, file_name string) (string, error) {
	ctx := context.Background()
	// Initialize minio client object.
	minioClient, err := minio.New(MinIOEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(MinIOAccessKeyID, MinIOAccessSecretKey, ""),
	})
	if err != nil {
		return "", err
	}
	file, fileHandler, err := r.FormFile(file_name)
	if err != nil {
		return "", err
	}
	objectName := utils.UUID() + path.Ext(fileHandler.Filename)
	contentType := "binary/octet-stream"
	// Upload the zip file with FPutObject
	_, err = minioClient.PutObject(ctx, MinIOBucket, objectName, file, fileHandler.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}

	return MinIOBucket + "/" + objectName, nil
}

//func MinIOUploadByReader(reader io.Reader) (string, error) {
//	ctx := context.Background()
//	minioClient, err := minio.New(MinIOEndpoint, &minio.Options{
//		Creds: credentials.NewStaticV4(MinIOAccessKeyID, MinIOAccessSecretKey, ""),
//	})
//	if err != nil {
//		return "", err
//	}
// Make a new bucket called mymusic.
// bucketName := "mymusic"
// location := "us-east-1"

// err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
// if err != nil {
// 	// Check to see if we already own this bucket (which happens if you run this twice)
// 	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
// 	if errBucketExists == nil && exists {
// 		log.Printf("We already own %s\n", bucketName)
// 	} else {
// 		log.Fatalln(err)
// 	}
// } else {
// 	log.Printf("Successfully created %s\n", bucketName)
// }

// Upload the zip file
//	file, fileHandler, err := r.FormFile("file")
//	if err != nil {
//		return "", err
//	}
//	objectName := utils.UUID() + path.Ext()
//	contentType := "binary/octet-stream"
//	// Upload the zip file with FPutObject
//	_, err = minioClient.PutObject(ctx, MinIOBucket, objectName, reader, reader.S, minio.PutObjectOptions{ContentType: contentType})
//	if err != nil {
//		return "", err
//	}
//
//	return MinIOBucket + "/" + objectName, nil
//}
