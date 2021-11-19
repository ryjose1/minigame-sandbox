package components

import (
	"github.com/solarlune/resolv"
)

const TILESIZE = 16

type HitSpace struct {
	space  *resolv.Space
	offset *Position
}

func NewSpace(position *Position) *HitSpace {
	space := resolv.NewSpace(position.Width(), position.Height(), TILESIZE, TILESIZE)
	return &HitSpace{
		space:  space,
		offset: position,
	}
}

func (h HitSpace) AddHitboxes(hitboxes []*Hitbox) {
	for _, hitbox := range hitboxes {
		h.space.Add(hitbox.Object())
	}
}

func (h HitSpace) Objects() []*resolv.Object {
	return h.space.Objects()
}

type Hitbox struct {
	object *resolv.Object
}

func NewHitbox(position *Position, tags ...string) *Hitbox {
	object := resolv.NewObject(float64(position.X()), float64(position.Y()), float64(position.Width()), float64(position.Height()), tags...)
	return &Hitbox{
		object: object,
	}
}

func (h Hitbox) Object() *resolv.Object {
	return h.object
}

func (h Hitbox) Check(xSpeed int, ySpeed int, tags ...string) (isHorizontalCollision bool, isVerticalCollision bool) {
	if collision := h.object.Check(float64(xSpeed), float64(ySpeed), tags...); collision != nil {
		for _, object := range collision.Objects {
			vector := collision.ContactWithObject(object)
			if vector.X() == 0 {
				isHorizontalCollision = true
			}
			if vector.Y() == 0 {
				isVerticalCollision = true
			}
		}
	}
	return
}

func (h Hitbox) UpdatePosition(dx int, dy int) {
	h.object.X += float64(dx)
	h.object.Y += float64(dy)
	h.object.Update()
}
