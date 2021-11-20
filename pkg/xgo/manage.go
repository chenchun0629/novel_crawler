package xgo

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

type Closer interface {
	Close()
}

var defaultManager *manager

func init() {
	defaultManager = newManager()
}

type manager struct {
	c map[string]Closer

	idGen *snowflake.Node
}

func (r manager) print() {
	fmt.Printf("%#v \n", r.c)
}

func newManager() *manager {
	var node, err = snowflake.NewNode(1)
	if err != nil {
		panic(fmt.Sprintf("new xgo manage error: %s", err.Error()))
	}

	return &manager{
		c:     make(map[string]Closer),
		idGen: node,
	}
}

func (r *manager) register(closer Closer) string {
	var id = r.genID()
	r.c[id] = closer
	return id
}

func (r *manager) close(id string) {
	if t, has := r.c[id]; has {
		t.Close()
		r.remove(id)
	}
}

func (r *manager) closeAll() {
	for id, t := range r.c {
		t.Close()
		r.remove(id)
	}
}

func (r *manager) remove(id string) {
	delete(r.c, id)
}

func (r *manager) genID() string {
	return r.idGen.Generate().String()
}
