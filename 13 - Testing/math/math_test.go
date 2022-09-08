package math

import "testing"

func TestAbs(t *testing.T) {
    t.Run("Positive", func(t *testing.T){
        if Abs(10) != 10 {
            t.Error("Unexpected res for positive num")
        }
    })
    t.Run("Zero", func(t *testing.T){
        if Abs(0) != 0 {
            t.Error("Unexpected res for zero")
        }
    })
    t.Run("Negative", func(t *testing.T){
        if Abs(-1) != 1 {
            t.Error("Unexpected res for negative num")
        }
    })
}
