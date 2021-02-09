package main

import "fmt"

type vm struct {
	cmd       []byte
	pc        int
	mainStack []byte
}

func main() {
	v := vm{}
	v.cmd = []byte{
		debug,
		add,
		debug,
	}
	v.push(10, 40)
	v.Run()
	v.pc = 0
	v.push(90)
	v.cmd = []byte{
		debug,
		sub,
		debug,
	}
	v.Run()
}

func (m *vm) Run() {
	for {
		if len(m.cmd) <= m.pc {
			break
		}
		fmt.Print("pc :", m.pc, " cmd :")
		m.pc++
		switch m.cmd[m.pc-1] {
		case add:
			fmt.Println("ADD")
			a, ok := m.pop()
			if ok {
				b, ok := m.pop()
				if ok {
					m.push(a + b)
					continue
				}
			}
			break
		case sub:
			fmt.Println("SUB")
			a, ok := m.pop()
			if ok {
				b, ok := m.pop()
				if ok {
					m.push(a - b)
					continue
				}
			}
			break
		case debug:
			fmt.Println("DEBUG")
			fmt.Println(m.mainStack)
		default:
			break
		}
	}
}

func (m *vm) push(v ...byte) {
	m.mainStack = append(m.mainStack, v...)
}

func (m *vm) pop() (byte, bool) {
	if len(m.mainStack) > 0 {
		idx := len(m.mainStack) - 1
		v := m.mainStack[idx]
		m.mainStack = m.mainStack[:idx]
		return v, true
	}
	return 0, false
}
