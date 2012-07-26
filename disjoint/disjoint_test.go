/*
 * Put your project description here
 * 
 * Author: likunarmstrong@gmail.com
 */


package disjoint

import "testing"

func TestNewSet(t *testing.T) {
    s := NewSet(10);
    if s.SetSize(1) != 1 {
        t.Errorf("Incorrect in NewSet");
    }
}

func TestUnion(t *testing.T) {
    s := NewSet(10);
    s.Union(1,2);
    s.Union(1,3);

    s.Union(5, 4);
    s.Union(6, 4);
    s.Union(8, 4);
    s.Union(7, 4);

    if s.Find(5) != 4 {
        t.Errorf("Incorrect parent after a union");
    }

    if s.SetSize(9) != 1 {
        t.Errorf("Incorrect set size before a union");
    }

    s.Union(9, 1);

    if s.SetSize(9) != 4 {
        t.Errorf("Incorrect set size after a union");
    }
}
