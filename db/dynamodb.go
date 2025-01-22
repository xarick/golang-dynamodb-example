package db

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/xarick/golang-dynamodb-example/config"
)

var DynamoDB *dynamodb.DynamoDB

func InitDynamoDB(cfg config.Application) {
	// Proxy sozlamalari (agar sizda proxy bo'lmasa, proxy sozlamalarini olib tashlang)
	// proxy, err := url.Parse(cfg.ProxyURL)
	// if err != nil {
	// 	log.Fatalf("Proxy URL'ni tahlil qilishda xatolik: %v", err)
	// }

	transport := &http.Transport{
		// Proxy: http.ProxyURL(proxy),
		Proxy: nil,
	}

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	// AWS sessiyasini yaratish
	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(cfg.DynamoDBEndPoint),                     // DynamoDB Local ulanish nuqtasi
		Region:      aws.String(cfg.DynamoDBRegion),                       // AWS mintaqasi
		Credentials: credentials.NewStaticCredentials("test", "test", ""), // Soxta credential
		HTTPClient:  httpClient,                                           // Proxy sozlamalari
	})
	if err != nil {
		log.Fatalf("AWS sessiyasini yaratishda xatolik: %v", err)
	}

	DynamoDB = dynamodb.New(sess)
	log.Println("DynamoDB muvaffaqiyatli ulandi.")

	// test
	result, err := DynamoDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatalf("Jadvallarni olishda xatolik: %v", err)
	}

	fmt.Println("Topilgan jadvallar:", result.TableNames)
}
