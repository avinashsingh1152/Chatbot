package core

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"log"
	"server/Global"
	"server/client"
	"server/mapper"
	"server/models"
	"server/repository"
)

type Core struct {
	Logger            *log.Logger
	SlModelGrpcClient client.SlModelGrpcClient
	DB                repository.SqlRepository
	S3                *s3.S3
}

type ICore interface {
	GetChatbotMessage() string
	GetGrpcPing() (string, error)
	UploadFile(ipAddress, fileName string, file io.Reader) error
	GetImages(pageSize, cursor int, ipAddress string) ([]models.Image, int64, error)
	Conversation(ipAddress, question string) (mapper.QuestionAnswer, error)
}

func (c Core) GetGrpcPing() (string, error) {
	resp, err := c.SlModelGrpcClient.GetConnectionPing()
	return resp, err
}

func (c Core) GetChatbotMessage() string {
	resp, _ := c.SlModelGrpcClient.GetConnectionPing()
	return resp
}

func (c Core) UploadFile(ipAddress, fileName string, file io.Reader) error {

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	generatedFileName := fmt.Sprintf("%s/%s_%s", Global.IMAGES, uuid.New(), fileName)
	_, err = c.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(Global.BUCKET_UPLOAD_PATH),
		Key:    aws.String(generatedFileName),
		Body:   bytes.NewReader(fileContent),
	})
	if err != nil {
		c.Logger.Println("error while uploading image to s3: ", err)
		return err
	}

	_, err = c.DB.SaveImageDetails(ipAddress, generatedFileName, fileName)
	if err != nil {
		c.Logger.Println("error while Saving image details to db: ", err)
		return err
	}

	return err
}

func (c Core) GetImages(pageSize, cursor int, ipAddress string) ([]models.Image, int64, error) {
	resp, total, err := c.DB.GetImages(pageSize, cursor, ipAddress)
	if err != nil {
		c.Logger.Println("error while get image details from db: ", err)
		return nil, 0, err
	}

	return resp, total, err
}

func (c Core) Conversation(ipAddress, question string) (resp mapper.QuestionAnswer, err error) {

	model, err := c.DB.SaveConversation(ipAddress, question)
	if err != nil {
		c.Logger.Println("error while get saving Conversation in db: ", err)
		return resp, err
	}

	answer, err := c.SlModelGrpcClient.GetChatbotMessage(question)
	if err != nil {
		c.Logger.Println("error while get answer from Grpc: ", err)
		return resp, err
	}

	err = c.DB.UpdateConversation(fmt.Sprintf("id = %d", model.ID), map[string]interface{}{"response": answer, "is_responded": true})
	if err != nil {
		c.Logger.Println("error while updating Conversation in DB: ", err)
		return resp, err
	}

	resp.Question = question
	resp.Answer = answer
	return resp, err
}
