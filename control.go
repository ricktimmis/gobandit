package main


type control struct{
	ready		bool
	board		interface{}
}

type controller interface{
	spin()
	hold()
	nudge()
}

func (c *control) spin() {
	return
}

func (c *control) hold() {
	return
}

func (c *control) nudge() {
	return
}