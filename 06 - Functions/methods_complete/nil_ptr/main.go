package main

type M map[string][]string

func (m M) getVal(k string) string {
	if v, ok := m[k]; !ok {
		return v[0]
	}
	return ""
}

func main() {

}
