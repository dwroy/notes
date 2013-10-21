package controller

import (
    "log"
    "github.com/roydong/potato"
)

type User struct {
    *potato.Controller
}

func (c *User) Show() {
    id,_ := c.Request.GetInt("id")

    log.Println("user show", id)
}

func (c *User) Topic() {


    log.Println(c.Request.GetInt("uid"))
    log.Println(c.Request.GetInt("i"))
    log.Println(c.Request.GetFloat("f"))
    log.Println(c.Request.Get("s"))
}
