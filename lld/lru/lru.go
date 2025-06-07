package lru


type DLinkList struct {
	data  int
    key int
	left  *DLinkList
	right *DLinkList
}

type LRUCache struct {
    nodeMap  map[int]*DLinkList
	head     *DLinkList
	tail     *DLinkList
	capacity int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
		capacity: capacity,
		nodeMap:  make(map[int]*DLinkList),
	}
}


func (l *LRUCache) Get(key int) int {
    node, exists := l.nodeMap[key];
	if !exists {
		return -1
	}
	l.removeLink(node)
	l.makeHead(node)
	return node.data
}

func (l *LRUCache) Put(key int, value int)  {
	if node, exists := l.nodeMap[key]; exists {
		node.data = value
		l.removeLink(node)
		l.makeHead(node)
	} else {
		node = &DLinkList{
			data: value,
			key: key,
		}
		if len(l.nodeMap) >= l.capacity {
			delete(l.nodeMap, l.tail.key)
			l.removeLink(l.tail)
		} else {
			l.removeLink(node)
		}
		l.makeHead(node)
		l.nodeMap[key] = node
	}
}

func (l *LRUCache) Put1(key int, value int)  {
    node, exists := l.nodeMap[key];
	if !exists {
		node = &DLinkList{
			data: value,
            key : key,
		}
	} else {
        node.data = value
    }
	if len(l.nodeMap) >=  l.capacity {
        delete(l.nodeMap, l.tail.key)
		l.removeLink(l.tail)
	} else {
		l.removeLink(node)
	}
	l.nodeMap[key] = node
    l.makeHead(node)
}


func(l *LRUCache) removeLink(node *DLinkList) {
	if node.left != nil {
		node.left.right = node.right
	} else {
		l.head = node.right
	}
	if node.right != nil {
		node.right.left = node.left
	} else {
		l.tail = node.left
	}
}

func(l *LRUCache) makeHead(node *DLinkList) {
    node.right = l.head
    if l.head != nil {
        l.head.left = node
    }
    l.head = node
    if l.tail == nil {
        l.tail = l.head
    }
}
