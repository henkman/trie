package trie

type Container struct {
	Children []*Node
}

type Node struct {
	Character byte
	Leaf      bool // node is the end of an inserted string
	Container
}

type Trie struct {
	Container
}

// Finds a direct child node of a container by character
func (c *Container) FindByCharacter(b byte) *Node {
	for _, n := range c.Children {
		if n.Character == b {
			return n
		}
	}
	return nil
}

// Inserts the string into the container
func (c *Container) Insert(s string) {
	if len(s) == 0 {
		return
	}
	b := s[0]
	n := c.FindByCharacter(b)
	if n == nil {
		n = &Node{Character: b}
		c.Children = append(c.Children, n)
	}
	for i := 1; i < len(s); i++ {
		b = s[i]
		o := n.FindByCharacter(b)
		if o == nil {
			o = &Node{Character: b}
			n.Children = append(n.Children, o)
		}
		n = o
	}
	n.Leaf = true
}

// Looks up the string in the container by walking through it
// and returns the last node that it walked past or nil
// if the container is the trie root and it did not find
// any matching child
func (c *Container) Lookup(s string) *Node {
	if len(s) == 0 {
		return nil
	}
	n := c.FindByCharacter(s[0])
	if n == nil {
		return nil
	}
	for i := 1; i < len(s); i++ {
		b := s[i]
		o := n.FindByCharacter(b)
		if o == nil {
			return n
		}
		n = o
	}
	return n
}
