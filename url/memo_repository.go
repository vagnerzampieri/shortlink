package url

type memoRepository struct {
	urls   map[string]*Url
	clicks map[string]int
}

func NewMemoRepository() *memoRepository {
	return &memoRepository{
		make(map[string]*Url),
		make(map[string]int),
	}
}
