package helper

import(
    "errors"
    "github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
    userType := c.GetString("user_type")
    uid := c.GetString("uid")
    err = nil

    if userType == "USER" && uid != userId {
        err = errors.New("Not have access for this resource. Sorry.") 
        return err
    }

    err = CheckUserType(c, userType)
    return err
}

func CheckUserType(c *gin.Context, role string) (err error) {
    userType := c.GetString("user_type")
    err = nil

    if userType != role {
        err = errors.New("You can't access this. Sorry :) ")
        return err
    }

    return err
}
