package url

type memoRepository struct {
	urls   map[string]*Url
	clicks map[string]int
}
