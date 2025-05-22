package fiscobcos

import (
	eventreq "github.com/nccasia/vbs-sdk-go/pkg/core/model/req/fiscobcos/event"
	eventres "github.com/nccasia/vbs-sdk-go/pkg/core/model/res/fiscobcos/event"

	"github.com/pkg/errors"
)

const (
	EventRegister      = "event/register"
	BlockEventRegister = "event/block/register"
	EventQuery         = "event/query"
	EventRemove        = "event/remove"
)

func (c *FiscoBcosClient) EventRegister(body eventreq.RegisterEventReqBody) (*eventres.RegisterEventResData, error) {
	req := &eventreq.RegisterEventReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &eventres.RegisterEventResData{}

	err := c.Call(EventRegister, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EventRegister)
	}
	return res, nil
}

func (c *FiscoBcosClient) BlockEventRegister(body eventreq.RegisterEventReqBody) (*eventres.RegisterEventResData, error) {
	req := &eventreq.RegisterEventReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &eventres.RegisterEventResData{}

	err := c.Call(BlockEventRegister, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", BlockEventRegister)
	}
	return res, nil
}

func (c *FiscoBcosClient) EventQuery(body eventreq.QueryEventReqBody) (*eventres.QueryEventResData, error) {
	req := &eventreq.QueryEventReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &eventres.QueryEventResData{}

	err := c.Call(EventQuery, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EventQuery)
	}
	return res, nil
}

func (c *FiscoBcosClient) EventRemove(body eventreq.RemoveEventReqBody) (*eventres.RemoveEventResData, error) {
	req := &eventreq.RemoveEventReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &eventres.RemoveEventResData{}

	err := c.Call(EventRemove, req, res)

	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EventRemove)
	}
	return res, nil
}
