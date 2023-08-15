package appgood

type createHandler struct {
	*Handler
}

func (h *Handler) CreateGood(ctx context.Context) (*npool.Good, error) {
}
