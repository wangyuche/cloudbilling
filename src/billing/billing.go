package billing

type IBilling interface {
	GetProjects()
}

func New() IBilling {
	var billingInstance IBilling
	billingInstance = &GCP{}
	return billingInstance
}
