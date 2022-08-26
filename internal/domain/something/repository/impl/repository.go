package impl

type repository struct {
	message string
}

func NewRepository(message string) *repository {
	return &repository{
		message: message,
	}
}

func (r *repository) GetMessage() string {
	return r.message
}
