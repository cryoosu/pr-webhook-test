package ds

type List struct {
	mem     []int
	size    int
	maxSize int
}

func NewList() List {
	return List{
		mem:     make([]int, 1),
		size:    0,
		maxSize: 1,
	}
}

func (l *List) Length() int {
	return l.size
}

func (l *List) MaxSize() int {
	return l.maxSize
}

func (l *List) Get(index int) (int, bool) {
	if index > l.size-1 {
		return -1, false
	}

	return l.mem[index], true
}

func (l *List) GetAll() []int {
	return l.mem[:l.size]
}

func (l *List) Push(item int) {
	if l.maxSize < l.size+1 {
		l.mem = append(l.mem, make([]int, l.maxSize)...)
		l.maxSize = l.maxSize * 2
	}

	l.mem[l.size] = item
	l.size++
}

func (l *List) Pop() (int, bool) {
	if l.size == 0 {
		return 0, false
	}

	last := l.mem[l.size-1]
	l.mem = append(l.mem[:l.size-1], []int{}...)
	l.size--

	return last, true
}
