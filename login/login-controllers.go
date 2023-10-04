package login

import (
	"net/http"

	"github.com/MikaelHans/catea/login-signup/util"
	"github.com/gin-gonic/gin"
)



func Login(c *gin.Context) {
	var logininfo util.LoginInfo;
	//CHECK FOR BINDING ERROR///////////////////////////////////////////
	if err := c.ShouldBindJSON(&logininfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	//QUERY TO DB THEN CHECK FOR QUERY ERROR///////////////////////////////////////////
	rows, err := GetMemberWithLoginInfo(logininfo)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	/*
	CHECK IF THE CREDENTIALS(EMAIL, PASS) IS CORRECT BY COUNTING THE ROWS, 0 MEANS INCORRECT SINCE IT MEANS EITHER
	EMAIL OR PASS IS INCORRECT
	*/
	var i int
	var member_data util.Member;
	for rows.Next(){
		rows.Scan(&member_data)
		i++
	}
	if i < 0{
		responseData := map[string]interface{}{
            "msg": "Email or pass is incorrect",
        }
		c.JSON(http.StatusBadRequest, responseData)
		return
	}
	/*WHEN SUCCESS RETURN TOKEN AND MSG:SUCCESS*/
	token, err := util.GenerateJWT(member_data.Email)
	
	/*GENERATE JWT ERROR HANDLER*/
	if err != nil{
		responseData := map[string]interface{}{
            "msg": err.Error(),
        }
		c.JSON(http.StatusInternalServerError, responseData)
		return
	}

	responseData := map[string]interface{}{
		"msg": "success",
		"token":token,
	}
	c.JSON(http.StatusAccepted, responseData)
	return
	
}