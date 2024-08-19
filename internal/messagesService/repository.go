package messagesService

type MessageRepository interface {
	// CreateMessage - Передаем в функцию message типа Message из orm.go
	// возвращаем созданный Message и ошибку
	CreateMessage(message Message) (Message, error)
	// GetAllMessages - Возвращаем массив из всех писем в БД и ошибку
	GetAllMessages() ([]Message, error)
	UpdateMessageById(message Message) (Message, error)
	// и ошибку

}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) UpdateMessageById(message Message) (Message, error) {
	result := r.db.Model(&Message{}).Where("ID = ?", message.ID).Updates(message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}
