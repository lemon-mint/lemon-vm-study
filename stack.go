package main

type stack struct {
	data []int
}

func (m *stack) push(v ...int) {
	m.data = append(m.data, v...)
}

func (m *stack) pop() (int, bool) {
	if len(m.data) > 0 {
		idx := len(m.data) - 1
		v := m.data[idx]
		m.data = m.data[:idx]
		return v, true
	}
	return 0, false
}
