package messagesService

type Message struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Value string `json:"Value"`
}
