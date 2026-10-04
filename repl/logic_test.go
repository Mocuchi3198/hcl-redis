package repl

import (
	"testing"
)

func TestBalanced(t *testing.T) {
    tests := []struct {
        name    string
        text     string
        expected  bool
    }{
        {"vacío", "", true},
        {"solo apertura", "create {", false},
        {"solo cierre", "}", false},
        {"balanceado simple", "{}", true},
        {"balanceado anidado", "create { fields = { x = 1 } }", true},
        {"desbalanceado", "create { fields = { x = 1 }", false},
        {"múltiples bloques", "{}{}", true},
        {"múltiples desbalanceado", "{}{", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Balanced(tt.text)
            if result != tt.expected {
                t.Errorf("Balanced(%q): esperaba %v, obtuve %v", tt.text, tt.expected, result)
            }
        })
    }
}