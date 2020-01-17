package controller

type IndexController struct {
	Controller
}

func (c *Controller) Index() {
	c.View.Assign("name", "GoFrame")
	c.View.Display("index.html")
	c.Exit()
}
