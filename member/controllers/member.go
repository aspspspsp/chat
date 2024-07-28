package controllers

//
//func Register(c *gin.Context) {
//	var input struct {
//		Username string `json:"username" binding:"required"`
//		Password string `json:"password" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
//	member := models.Member{Username: input.Username, Password: string(hashedPassword)}
//
//	if result := db.DB.Create(&member); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"member": member})
//}
//
//func Login(c *gin.Context) {
//	var input struct {
//		Username string `json:"username" binding:"required"`
//		Password string `json:"password" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	var member models.Member
//	if result := db.DB.Where("username = ?", input.Username).First(&member); result.Error != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//		return
//	}
//
//	if err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(input.Password)); err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"member": member})
//}
