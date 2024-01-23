package repository

type People interface {
	Get_Person()
	Delete_Person()
	Patch_Person()
	Post_Person()
}

type Repository struct{
	People
}

func New_Repository() *Repository {
	return &Repository{}
}