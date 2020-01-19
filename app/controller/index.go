package controller

type IndexController struct {
	Controller
}

func (c *Controller) Index() {
	c.Render("index.html","name", "GoFrame")
}
