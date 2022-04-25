package opgraph

func (g *CallGraph) HandleInput(event InputEvent) bool {
	switch event {
	case UpInputEvent:
		return g.moveCursorUp()
	case DownInputEvent:
		return g.moveCursorDown()
	case LeftInputEvent:
		return g.moveCursorLeft()
	case RightInputEvent:
		return g.moveCursorRight()
	case ClearInputEvent:
		g.resetSelection()
		return true
	}
	return false
}

func (g *CallGraph) findSelected() *callGraphNode {
	return g.rootNode.findSelected()
}

func (n *callGraphNode) findSelected() *callGraphNode {
	if n.cursorActive {
		return n
	}
	for _, child := range n.children {
		if c := child.findSelected(); c != nil {
			return c
		}
	}
	return nil
}

func (n *callGraphNode) hasSelected() bool {
	if n == nil {
		return false
	}
	if n.cursorActive {
		return true
	}
	for _, c := range n.children {
		if c.hasSelected() {
			return true
		}
	}
	return false
}

func (n *callGraphNode) priorSibling() *callGraphNode {
	if n == nil || n.parent == nil {
		return nil
	}
	siblings := n.parent.children
	for i, s := range siblings {
		if s == n {
			if i == 0 {
				return nil
			}
			return siblings[i-1]
		}
	}
	return n.parent
}

func (n *callGraphNode) nextSibling() *callGraphNode {
	if n == nil || n.parent == nil {
		return nil
	}
	siblings := n.parent.children
	for i, s := range siblings {
		if s == n {
			if i == len(siblings)-1 {
				return nil
			}
			return siblings[i+1]
		}
	}
	return nil
}

func swapCursor(prior, next *callGraphNode) bool {
	if prior == nil || next == nil {
		return false
	}
	prior.cursorActive = false
	next.cursorActive = true
	return true
}

func (g *CallGraph) moveCursorUp() bool {
	n := g.findSelected()
	return swapCursor(n, n.priorSibling())
}

func (g *CallGraph) moveCursorDown() bool {
	n := g.findSelected()
	return swapCursor(n, n.nextSibling())
}

func (g *CallGraph) moveCursorLeft() bool {
	n := g.findSelected()
	return swapCursor(n, n.parent)
}

func (g *CallGraph) moveCursorRight() bool {
	n := g.findSelected()
	if n == nil {
		return false
	} else if len(n.children) > 0 {
		return swapCursor(n, n.children[0])
	}
	return false
}

func (g *CallGraph) resetSelection() {
	n := g.findSelected()
	if n != nil {
		n.cursorActive = false
	}
	if g.rootNode != nil {
		g.rootNode.cursorActive = true
	}
}
