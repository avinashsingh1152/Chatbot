package repository

import (
	"gorm.io/gorm"
	"log"
	"server/models"
)

type SqlRepository struct {
	DB     *gorm.DB
	Logger *log.Logger
}

type ISqlRepository interface {
	SaveImageDetails(ipaddress, imageUrl, fileName string) (int64, error)
	GetImages(pageSize, cursor int, ipAddress string) ([]models.Image, int64, error)
	SaveConversation(ipAddress, question string) (models.Message, error)
	UpdateConversation(condition, update interface{}) error
}

func (r SqlRepository) SaveImageDetails(ipaddress, imageUrl, fileName string) (int64, error) {

	model := models.Image{
		IPAddress: ipaddress,
		ImageURL:  imageUrl,
		FileName:  fileName,
	}

	if err := r.DB.Save(&model).Error; err != nil {
		r.Logger.Println("Error while saving image", err)
		return 0, err
	}

	return int64(model.ID), nil
}

func (r SqlRepository) GetImages(pageSize, cursor int, ipAddress string) ([]models.Image, int64, error) {
	var model []models.Image
	if err := r.DB.Model(&models.Image{}).Where("ip_address = ?", ipAddress).Offset(cursor).Limit(pageSize).Scan(&model).Error; err != nil {
		r.Logger.Println("Error while fetching images by ip_address", err)
		return nil, 0, err
	}

	var totalCount int64
	err := r.DB.Model(&models.Image{}).Where("ip_address = ?", ipAddress).Count(&totalCount).Error
	if err != nil {
		r.Logger.Println("Error while counting images by ip_address", err)
		return nil, 0, err
	}

	return model, totalCount, nil
}

func (r SqlRepository) SaveConversation(ipAddress, question string) (models.Message, error) {
	var model = models.Message{
		IPAddress:   ipAddress,
		MessageText: question,
		MessageType: models.BotType,
		IsResponded: false,
	}
	if err := r.DB.Save(&model).Error; err != nil {
		r.Logger.Println("Error while saving Conversation", err)
		return model, err
	}

	return model, nil
}

func (r SqlRepository) UpdateConversation(condition, update interface{}) error {

	if err := r.DB.Model(&models.Message{}).Where(condition).Updates(update).Error; err != nil {
		r.Logger.Println("Error while updating Conversation", err)
		return err
	}

	return nil
}
