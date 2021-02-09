package main

import (
	"fmt"
	"time"
)

type vm struct {
	cmd       []byte
	pc        int
	mainStack []byte
	pcStack   []int
}

func main() {
	v := vm{}
	v.cmd = append(v.cmd, debug, pushpc, debug, poppc)
	v.Run()
}

func (m *vm) Run() {
	for {
		if len(m.cmd) <= m.pc {
			break
		}
		fmt.Print("pc :", m.pc, " cmd :")
		m.pc++
		time.Sleep(time.Second)
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
		case inc:
			fmt.Println("INC")
			a, ok := m.pop()
			if ok {
				m.push(a + 1)
				continue
			}
			break
		case dec:
			fmt.Println("DEC")
			a, ok := m.pop()
			if ok {
				m.push(a - 1)
				continue
			}
			break
		case pushpc:
			fmt.Println("PUSHPC")
			m.pcStack = append(m.pcStack, m.pc-1)
			continue
		case poppc:
			fmt.Println("POPPC")
			if len(m.pcStack) > 0 {
				idx := len(m.pcStack) - 1
				m.pc = m.pcStack[idx]
				m.pcStack = m.pcStack[:idx]
				continue
			}
			break
		case debug:
			fmt.Println("DEBUG")
			fmt.Println("Stack :", m.mainStack)
			fmt.Println("pc Stack :", m.pcStack)
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
