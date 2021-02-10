package main

import (
	"fmt"
	"time"
)

type vm struct {
	cmd           []byte
	pc            int
	mainStack     []byte
	pcStack       []int
	lim           int
	activateLimit bool
}

func main() {
	v := vm{}
	v.activateLimit = false
	v.lim = 10
	v.cmd = append(v.cmd, debug, pushpc, debug, poppc)
	v.Run()
}

func (m *vm) Run() {
	for {
		if len(m.cmd) <= m.pc {
			break
		}
		if m.lim <= 0 && m.activateLimit {
			break
		}
		fmt.Print("pc :", m.pc, " cmd :")
		m.pc++
		time.Sleep(time.Second)
		switch m.cmd[m.pc-1] {
		case add:
			fmt.Println("ADD")
			m.lim--
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
			m.lim--
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
			m.lim--
			fmt.Println("INC")
			a, ok := m.pop()
			if ok {
				m.push(a + 1)
				continue
			}
			break
		case dec:
			m.lim--
			fmt.Println("DEC")
			a, ok := m.pop()
			if ok {
				m.push(a - 1)
				continue
			}
			break
		case pushpc:
			m.lim--
			fmt.Println("PUSHPC")
			m.pcStack = append(m.pcStack, m.pc-1)
			continue
		case cmpjmp:
			m.lim--
			fmt.Println("CMPJMP")
			if len(m.pcStack) > 0 && len(m.mainStack) > 0 {
				idx := len(m.pcStack) - 1
				if m.mainStack[idx] > 0 {
					m.pc = m.pcStack[idx]
					m.pcStack = m.pcStack[:idx]
				}
				continue
			}
			break
		case poppc:
			m.lim--
			fmt.Println("POPPC")
			if len(m.pcStack) > 0 {
				idx := len(m.pcStack) - 1
				m.pc = m.pcStack[idx]
				m.pcStack = m.pcStack[:idx]
				continue
			}
			break
		case pushzero:
			m.lim--
			fmt.Println("PUSHZERO")
			m.push(0)
		case copy:
			m.lim--
			fmt.Println("COPY")
			v, ok := m.pop()
			if ok {
				m.push(v, v)
			}
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
