package messagesService

import (
	"gorm.io/gorm"
)

type messageRepository struct {
	db *gorm.DB
}

type MessageService struct {
	repo MessageRepository
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func NewService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

func (s *MessageService) UpdateMessageById(message Message) (Message, error) {
	return s.repo.UpdateMessageById(message)
}

func (s *MessageService) DeleteMessageById(message Message) error {
	return s.repo.DeleteMessageById(message)
}
