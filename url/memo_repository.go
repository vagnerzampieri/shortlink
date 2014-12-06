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

func (r *memoRepository) IdExist(id string) bool {
	_, exist := r.urls[id]
	return exist
}

func (r *memoRepository) FindId(id string) *Url {
	return r.urls[id]
}

func (r *memoRepository) FindUrl(url string) *Url {
	for _, u := range r.urls {
		if u.Destination == url {
			return u
		}
	}
	return nil
}
