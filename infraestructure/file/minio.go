package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"tuberias/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
	once        sync.Once
)

func getMinioClient() *minio.Client {

	once.Do(func() {
		conf, errores := config.LoadConfigMinio()

		if errores != nil {
			log.Fatalf("Error al obtener las propiedades de conexión a base de datos: %v", errores)

		}

		endpoint := conf.Endpoint
		accessKeyID := conf.AccessKeyID
		secretAccessKey := conf.SecretAccessKey
		useSSL := conf.UseSSL

		var err error
		minioClient, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			log.Fatalln("Error al crear el cliente de MinIO:", err)
		}
		fmt.Println("Cliente de MinIO inicializado.")
	})
	return minioClient
}

func GetFileFromMinio(bucketName, objectName string) (string, error) {
	client := getMinioClient()

	// Obtener el archivo (objeto) desde MinIO
	object, err := client.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("error al obtener el objeto %s: %v", objectName, err)
	}
	defer object.Close()

	// Leer el contenido del archivo
	var content string
	buffer := make([]byte, 1024)
	for {
		n, err := object.Read(buffer)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("error al leer el archivo: %v", err)
		}
		if n == 0 {
			break
		}
		content += string(buffer[:n])
	}

	return content, nil
}

func ensureBucketExists(bucketName string) error {
	client := getMinioClient()

	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return fmt.Errorf("error al verificar la existencia del bucket %s: %v", bucketName, err)
	}

	if !exists {
		err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("error al crear el bucket %s: %v", bucketName, err)
		}
		fmt.Printf("Bucket %s creado con éxito.\n", bucketName)
	}

	return nil
}

// UploadFileToMinioBinary sube un archivo a MinIO desde un contenido binario.
func UploadFileToMinioBinary(bucketName, objectName string, fileContent []byte) error {
	client := getMinioClient()

	// Asegurarse de que el bucket exista
	if err := ensureBucketExists(bucketName); err != nil {
		return err
	}

	// Crear un lector desde el contenido binario
	reader := bytes.NewReader(fileContent)

	// Subir el archivo a MinIO
	info, err := client.PutObject(context.Background(), bucketName, objectName, reader, reader.Size(), minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("error al subir el archivo %s: %v", objectName, err)
	}

	fmt.Printf("Archivo subido con éxito: %+v\n", info)
	return nil
}
