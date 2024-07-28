package controllers

//
//func CreateMember(c *gin.Context) {
//	var user models.Member
//	if err := c.ShouldBindJSON(&user); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 使用雪花算法生成 ID
//	user.ID = db.Node.Generate()
//
//	if err := db.DB.Create(&user).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, user)
//}
//
//func GetMember(c *gin.Context) {
//	id := c.Param("id")
//	var member models.Member
//	if err := db.DB.First(&member, id).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, member)
//}
//
//func UpdateMember(c *gin.Context) {
//	id := c.Param("id")
//	var member models.Member
//	if err := db.DB.First(&member, id).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
//		return
//	}
//
//	if err := c.ShouldBindJSON(&member); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	db.DB.Save(&member)
//	c.JSON(http.StatusOK, member)
//}
//
//func DeleteMember(c *gin.Context) {
//	id := c.Param("id")
//	var member models.Member
//	if err := db.DB.First(&member, id).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
//		return
//	}
//
//	db.DB.Delete(&member)
//	c.JSON(http.StatusOK, gin.H{"message": "Member deleted"})
//}
