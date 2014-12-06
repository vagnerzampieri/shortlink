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

func (r memoRepository) IdExist(id string) bool {
	_, exist := r.urls[id]
	return exist
}
