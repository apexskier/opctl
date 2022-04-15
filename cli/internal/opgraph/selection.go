package opgraph

func (g *CallGraph) HandleInput(event InputEvent) error {
	switch event {
	case UpInputEvent:
		return g.cursorUp()
	case DownInputEvent:
		return g.cursorDown()
	case LeftInputEvent:
		return g.cursorLeft()
	case RightInputEvent:
		return g.cursorRight()
	case ClearInputEvent:
		n := g.findSelected()
		if n != nil {
			n.cursorActive = false
		}
	}
	return nil
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

func (n *callGraphNode) priorSibling() *callGraphNode {
	if n.parent == nil {
		return nil
	}
	siblings := n.parent.children
	for i, s := range siblings {
		if s == n {
			if i == 0 {
				return n.parent.priorSibling()
			}
			return siblings[i-1]
		}
	}
	return nil
}

func (n *callGraphNode) nextSibling() *callGraphNode {
	if n.parent == nil {
		return nil
	}
	siblings := n.parent.children
	for i, s := range siblings {
		if s == n {
			if i == len(siblings)-1 {
				return n.parent.nextSibling()
			}
			return siblings[i+1]
		}
	}
	return nil
}

func swapCursor(next, prior *callGraphNode) {
	if next == nil {
		return
	}
	next.cursorActive = true
	prior.cursorActive = false
}

func (g *CallGraph) cursorUp() error {
	n := g.findSelected()
	if n == nil {
		g.rootNode.cursorActive = true
	}
	swapCursor(n.priorSibling(), n)
	return nil
}

func (g *CallGraph) cursorDown() error {
	n := g.findSelected()
	if n == nil {
		g.rootNode.cursorActive = true
	}
	swapCursor(n.nextSibling(), n)
	return nil
}

func (g *CallGraph) cursorLeft() error {
	n := g.findSelected()
	if n == nil {
		g.rootNode.cursorActive = true
	}
	swapCursor(n.parent, n)
	return nil
}

func (g *CallGraph) cursorRight() error {
	n := g.findSelected()
	if n == nil {
		g.rootNode.cursorActive = true
	}
	if len(n.children) > 0 {
		swapCursor(n.children[0], n)
	}
	return nil
}
