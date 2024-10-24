package demo1

import (
	"fmt"
	"sync/atomic"
)

var counter uint64

func generateId(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, atomic.AddUint64(&counter, 1))
}

type A struct {
	Id string
	b  *B
	c  *C
}

func NewA(b *B, c *C) *A {
	return &A{Id: generateId("A"), b: b, c: c}
}
func (this *A) ToString() string {
	return fmt.Sprintf("%s { %s, %s }", this.Id, this.b.ToString(), this.c.ToString())
}

type B struct {
	Id string
	d  *D
	e  *E
}

func NewB(d *D, e *E) *B {
	return &B{Id: generateId("B"), d: d, e: e}
}
func (this *B) ToString() string {
	return fmt.Sprintf("%s { %s, %s }", this.Id, this.d.ToString(), this.e.ToString())
}

type C struct {
	Id string
}

func NewC() *C {
	return &C{Id: generateId("C")}
}

func (this *C) ToString() string {
	return this.Id
}

type D struct {
	Id string
	f  *F
	h  *H
}

func NewD(f *F, h *H) *D {
	return &D{Id: generateId("D"), f: f, h: h}
}
func (this *D) ToString() string {
	return fmt.Sprintf("%s { %s, %s }", this.Id, this.f.ToString(), this.h.ToString())
}

type E struct {
	Id string
	g  []G
}

func NewE(g []G) *E {
	return &E{Id: generateId("E"), g: g}
}
func (this *E) ToString() string {
	resu := this.Id + "{ "
	for _, gItem := range this.g {
		resu += gItem.ToString() + ", "
	}
	resu += " }"
	return resu
}

type F struct {
	Id string
}

func NewF() *F {
	return &F{Id: generateId("F")}
}
func (this *F) ToString() string {
	return this.Id
}

type G interface {
	GetId() string
	ToString() string
}

var _ G = (*Ga)(nil)

type Ga struct {
	Id string
}

func NewGa() *Ga {
	return &Ga{Id: generateId("Ga")}
}
func (this *Ga) GetId() string {
	return this.Id
}
func (this *Ga) ToString() string {
	return this.Id
}

var _ G = (*Gb)(nil)

type Gb struct {
	Id string
}

func NewGb() *Gb {
	return &Gb{Id: generateId("Gb")}
}
func (this *Gb) ToString() string {
	return this.Id
}

func (this *Gb) GetId() string {
	return this.Id
}

var _ G = (*Gc)(nil)

type Gc struct {
	Id string
}

func NewGc() *Gc {
	return &Gc{Id: generateId("Gc")}
}
func (this *Gc) ToString() string {
	return this.Id
}

func (this *Gc) GetId() string {
	return this.Id
}

var _ G = (*DGa)(nil)

type DGa struct {
	core *Ga
	Id   string
}

func NewDGa(core *Ga) *DGa {
	return &DGa{core: core, Id: generateId("DGa")}
}
func (this *DGa) ToString() string {
	return fmt.Sprintf("%s { %s }", this.Id, this.core.ToString())
}

func (this *DGa) GetId() string {
	return this.Id
}

type H struct {
	Id string
}

func NewH() *H {
	return &H{Id: generateId("H")}
}

func (this *H) ToString() string {
	return this.Id
}

func CreateA() *A {
	d := NewD(NewF(), NewH())
	e := NewE([]G{NewGa(), NewGb(), NewGc()})
	b := NewB(d, e)
	c := NewC()
	return NewA(b, c)
}
