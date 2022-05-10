package douyin

import (
	"fmt"
	"maingo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}
type UserDoLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func (con UserController) DoLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	fmt.Println("用户名")
	fmt.Println(username, password)

	//2、查询数据库 判断用户以及密码是否存在
	userinfoList := []models.User{}
	// password = models.Md5(password)

	// json.Marshal(userinfoList)
	models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)
	// fmt.Println(userinfoList[0])

	// println(userinfoList[0])
	// c.JSON(http.StatusOK, gin.H{
	// 	"status_code": 0,
	// 	"data":        userinfoList,
	// })
	c.JSON(http.StatusOK, UserDoLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   1,
		Token:    "zhangleidouyin",
	})

}
