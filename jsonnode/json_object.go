package jsonnode

import "strconv"

// Node :
type Node struct {
	Value interface{}
	Error []string
}

// NewNode :
func NewNode() *Node {
	return &Node{Value: make(map[string]interface{}, 0)}
}

// SetValue :
func (n *Node) SetValue(value interface{}) {
	n.Error = make([]string, 0)
	n.Value = value
}

// Add :
func (n *Node) Add(keyValues ...KeyValue) {
	n.Error = make([]string, 0)
	if n.Value == nil {
		n.Value = make(map[string]interface{}, 0)
		n.Error = append(n.Error, "Node value is nil")
	}
	switch n.Value.(type) {
	case map[string]interface{}:
		AddToMap(n.Value.(map[string]interface{}), keyValues...)
	default:
		n.Error = append(n.Error, "Node value is not map[string]interface{}")
	}
}

// Append :
func (n *Node) Append(object interface{}) {
	n.Error = make([]string, 0)
	if n.Value == nil {
		n.Value = make([]interface{}, 0)
		n.Error = append(n.Error, "Node value is nil")
	}
	switch n.Value.(type) {
	case []interface{}:
		n.Value = append(n.Value.([]interface{}), object)
	default:
		n.Error = append(n.Error, "Node value is not []interface{}")
	}
}

// Get :
func (n *Node) Get(key string) *Node {
	n.Error = make([]string, 0)
	if n.Value == nil {
		return &Node{Error: []string{"Node value is nil"}}
	}
	switch n.Value.(type) {
	case map[string]interface{}:
		v, ok := n.Value.(map[string]interface{})[key]
		if ok {
			return &Node{Value: v}
		}
		return &Node{Error: []string{"Key '" + key + "' is not found in node value"}}
	default:
		return &Node{Error: []string{"Node value is not map[string]interface{}"}}
	}
}

// GetIndex :
func (n *Node) GetIndex(index int) *Node {
	n.Error = make([]string, 0)
	if n.Value == nil {
		return &Node{Error: []string{"Node value is nil"}}
	}
	switch n.Value.(type) {
	case []interface{}:
		if index > -1 && index < len(n.Value.([]interface{})) {
			return &Node{Value: n.Value.([]interface{})[index]}
		}
		return &Node{Error: []string{"Invalid index " + strconv.Itoa(index)}}
	default:
		return &Node{Error: []string{"Node value is not []interface{}"}}
	}
}
