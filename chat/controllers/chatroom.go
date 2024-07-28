package controllers

//func AddMemberToChatroom(c *gin.Context) {
//	var input struct {
//		UserID     uint `json:"user_id" binding:"required"`
//		ChatRoomID uint `json:"chat_room_id" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	userChatRoom := models.UserChatRoom{UserID: input.UserID, ChatRoomID: input.ChatRoomID}
//
//	if result := db.DB.Create(&userChatRoom); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Member added to chat room"})
//}

//func RemoveUserFromChatRoom(c *gin.Context) {
//	var input struct {
//		UserID     uint `json:"user_id" binding:"required"`
//		ChatRoomID uint `json:"chat_room_id" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	if result := db.DB.Where("user_id = ? AND chat_room_id = ?", input.UserID, input.ChatRoomID).Delete(&models.UserChatRoom{}); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "User removed from chat room"})
//}
