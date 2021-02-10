package main

import (
	"fmt"
)

type vm struct {
	cmd           []byte
	pc            int
	mainStack     []byte
	pcStack       []int
	ptrStack      stack
	lim           int
	memory        []byte
	ptr           int
	activateLimit bool
}

func main() {
	v := vm{}
	v.activateLimit = false
	v.lim = 10
	v.cmd = append(v.cmd,
		push,
		inc,
		inc,
		inc,
		pull,
		incptr,
		push,
		inc,
		inc,
		copy,
		add,
		copy,
		add,
		copy,
		add,
		copy,
		add,
		copy,
		add,
		inc,
		pull,
		printmem0,
		decptr,

		//mul 3 * 65 = 195
		incptr,
		pushpc,
		incptr,
		push,
		decptr,
		decptr,
		push,
		add,
		incptr,
		incptr,
		pull,
		decptr,
		push,
		dec,
		pull,
		push,
		cmpjmp,
		delpc,
		incptr,
		push,
		decptr,
		decptr,
		pull,
		pushzero,
		incptr,
		incptr,
		pull,
		decptr,
		decptr,
		//mul end

		debug,
		printmem0,
		// [195 0 0 0 0...]
	)
	v.Run()
}

func (m *vm) Run() {
	m.memory = make([]byte, 1024)
	for {
		if len(m.cmd) <= m.pc {
			break
		}
		if m.lim <= 0 && m.activateLimit {
			break
		}
		fmt.Print("pc :", m.pc, " cmd :")
		m.pc++
		//time.Sleep(time.Microsecond * 1000)
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
			if len(m.pcStack) > 0 {
				idx := len(m.pcStack) - 1
				v, ok := m.pop()
				if v > 0 && ok {
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
		case pop:
			m.lim--
			fmt.Println("POP")
			m.pop()
		case delpc:
			m.lim--
			fmt.Println("DELPC")
			if len(m.pcStack) > 0 {
				idx := len(m.pcStack) - 1
				m.pcStack = m.pcStack[:idx]
				continue
			}
		case exit:
			fmt.Println("EXIT")
			break
		case debug:
			fmt.Println("DEBUG")
			fmt.Println("Stack :", m.mainStack)
			fmt.Println("pc Stack :", m.pcStack)
			fmt.Println("ptr :", m.ptr)
		case incptr:
			m.lim--
			fmt.Println("INCPTR")
			m.ptr++
		case decptr:
			m.lim--
			fmt.Println("DECPTR")
			m.ptr--
		case push:
			m.lim--
			fmt.Println("PUSH")
			m.push(m.memory[m.ptr])
		case pull:
			m.lim--
			fmt.Println("PULL")
			v, ok := m.pop()
			if ok {
				m.memory[m.ptr] = v
			}
		case saveptr:
			m.lim--
			fmt.Println("SAVEPTR")
			m.ptrStack.push(m.ptr)
		case loadptr:
			m.lim--
			fmt.Println("LOADPTR")
			v, ok := m.ptrStack.pop()
			if ok {
				m.ptr = v
			}
		case delptr:
			fmt.Println("DELPTR")
			m.ptrStack.pop()
		case printmem0:
			fmt.Println("PRINTMEM0")
			fmt.Println(m.memory[:30])
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
