package controller

type AppController struct {
	User  interface{ User }
	Admin interface{ Admin }
}
