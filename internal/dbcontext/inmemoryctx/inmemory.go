package inmemoryctx

import "fmt"

type ctx struct {
	receiversList map[string][]string
}

func New() *ctx {
	return &ctx{make(map[string][]string)}
}

func (c *ctx) CreateList(usr string) error {

	if _, exist := c.receiversList[usr]; exist {
		return fmt.Errorf("There is already a receivers list for user %v", usr)
	}

	c.receiversList[usr] = make([]string, 0)

	return nil
}

func (c *ctx) DeleteList(usr string) error {

	if _, exist := c.receiversList[usr]; exist {
		delete(c.receiversList, usr)
		return nil
	}

	return fmt.Errorf("There is no receiver list for user %v", usr)
}

func (c *ctx) AddToList(usr, recivier string) error {

	if _, exist := c.receiversList[usr]; exist {
		c.receiversList[usr] = append(c.receiversList[usr], recivier)
		return nil
	}

	return fmt.Errorf("There is no receiver list for user %v", usr)
}
