package handlers

import (
	"fmt"
	"sort"
	"strings"
)

type helper struct {
	topics map[string]*sevCmd
}

func NewHelper() *helper {
	h := new(helper)
	h.topics = make(map[string]*sevCmd)
	return h
}

func (h *helper) Add(s string, c *sevCmd) {
	h.topics[s] = c
}

func (h *helper) All() {
	keys := make([]string, 0)
	for k := range h.topics {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, v := range keys {
		fmt.Printf("%s\n\t%s\n\n", h.topics[v].helpMsg, h.topics[v].description)
	}
}

func (h *helper) Execute(args []string) {
	if len(args) < 1 {
		fmt.Println("serveridentity help")
		return
	}
	if len(args) < 2 {
		h.All()
		return
	}

	topic := strings.Join(args[1:], " ")

	c, ok := h.topics[topic]
	if !ok {
		if c, ok = h.topics[args[1]]; !ok {
			fmt.Println("No help for:", topic)
			return
		}
	}
	fmt.Printf("%s\n\t%s\n", c.helpMsg, c.description)
}

var Help = NewHelper()
