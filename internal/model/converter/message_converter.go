package converter

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
)

// MessageToResponse converts a message entity to a response model
func MessageToResponse(message *entity.Message) *model.MessageResponse {
	response := &model.MessageResponse{
		ID:         &message.ID,
		Content:    message.Content,
		IsRead:     message.IsRead,
		SenderID:   &message.SenderID,
		ReceiverID: &message.ReceiverID,
		CreatedAt:  &message.CreatedAt,
	}
	
	// Add sender details if available
	if message.Sender.ID != [16]byte{} {
		response.Sender = &model.UserDetail{
			ID:       &message.Sender.ID,
			Username: message.Sender.Username,
			Role:     message.Sender.Role,
		}
	}
	
	// Add receiver details if available
	if message.Receiver.ID != [16]byte{} {
		response.Receiver = &model.UserDetail{
			ID:       &message.Receiver.ID,
			Username: message.Receiver.Username,
			Role:     message.Receiver.Role,
		}
	}
	
	return response
}

// MessagesToResponses converts a slice of message entities to response models
func MessagesToResponses(messages []*entity.Message) []*model.MessageResponse {
	responses := make([]*model.MessageResponse, len(messages))
	for i, message := range messages {
		responses[i] = MessageToResponse(message)
	}
	return responses
} 