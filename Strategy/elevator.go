package strategy

import "sync"

type Req struct {
	Floor int
}

type Elevator struct {
	Id    int
	Floor int
}

func (e *Elevator) GoTo(floor int) {
	e.Floor = floor
}

func (e *Elevator) CurrentFloor() int {
	return e.Floor
}

type Scheduler interface {
	SubmitRequest(r Req)
	MoveNext(e *Elevator) (int, bool)
}

type NearestScheduler struct {
	mu  sync.Mutex
	req []Req
}

func NewNearestScheduler() *NearestScheduler {
	return &NearestScheduler{}
}

func (n *NearestScheduler) SubmitRequest(floor Req) {
	n.req = append(n.req, floor)
}

func (n *NearestScheduler) MoveNext(e *Elevator) (int, bool) {
	return n.req[0].Floor, true
}

type ZoningScheduler struct {
	mu  sync.Mutex
	req []Req
}

func NewZoningScheduler() *ZoningScheduler {
	return &ZoningScheduler{}
}

func (zs *ZoningScheduler) SubmitRequest(floor Req) {
	zs.req = append(zs.req, floor)
}

func (zs *ZoningScheduler) MoveNext(e *Elevator) (int, bool) {
	return zs.req[0].Floor, true
}

type Controller struct {
	Elevator  *Elevator
	Scheduler Scheduler
}

func NewController(e *Elevator, s *Scheduler) *Controller {
	return &Controller{Elevator: e, Scheduler: *s}
}

func (c *Controller) Submit(req Req) {
	c.Scheduler.SubmitRequest(req)
}

func (c *Controller) MoveNext() {
	if t, ok := c.Scheduler.MoveNext(c.Elevator); ok {
		c.Elevator.GoTo(t)
	}
}
