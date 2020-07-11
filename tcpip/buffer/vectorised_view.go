package buffer

// VectorisedView is a vectorised version of View using non contigous memory.
// It supports all the convenience methods supported by View.
type VectorisedView struct {
    views []View
    size int
}

// NewVectorisedView creates a new vectorised view from an already-allocated slice
// of View and sets its size.
func NewVectorisedView(size int, views []View) *VectorisedView {
    return &VectorisedView{views: views, size: size}
}

// First returns the first view of the vectorised view.
// It panics if the vectorised view is empty.
func (vv *VectorisedView) First() View {
    if len(vv.views) == 0 {
        panic("vview is empty")
    }

    return vv.views[0]
}

// RemoveFirst removes the first view of the vectorised view.
func (vv *VectorisedView) RemoveFirst() {
    if len(vv.views) == 0 {
        return
    }
    vv.size -= len(vv.views[0])
    vv.views = vv.views[1:]
}

// Size returns the size in bytes of the entire content stored in the vectorised view.
func (vv *VectorisedView) Size() int {
    return vv.size
}

// TrimFront removes the first "count" bytes of the vectorised view.
func (vv *VectorisedView) TrimFront(count int) {
    for count > 0 && len(vv.views) > 0 {
        if count < len(vv.views[0]) {
            vv.size -= count
            vv.views[0].TrimFront(count)
            return
        }
        count -= len(vv.views[0])
        vv.RemoveFirst()
    }
}

// Caplength irreversibly reduces the length of the vectorised view.
func (vv *VectorisedView) CapLength(length int) {
    if length < 0 {
        length = 0
    }
    if vv.size < length {
        return
    }
    vv.size = length
    for i := range vv.views {
        v := &vv.views[i]
        if len(*v) >= length {
            if length == 0 {
                vv.views = vv.views[:i]
            } else {
                v.CapLength(length)
                vv.views = vv.views[:i+1]
            }
            return
        }
        length -= len(*v)
    }
}

// copy returns a deep-copy of the vectorised view.
// It is an expensive method that should be used only in tests.
func (vv *VectorisedView) copy() *VectorisedView {
    uu := &VectorisedView{
        views: make([]View, len(vv.views)),
        size: vv.size,
    }
    for i, v := range vv.views {
        uu.views[i] = make(View, len(v))
        copy(uu.views[i], v)
    }
    return uu
}
