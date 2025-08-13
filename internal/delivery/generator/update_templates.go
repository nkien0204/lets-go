package generator

func (d *delivery) HandleTemplateUpdating() error {
	return d.templateUpdatingUsecase.UpdateTemplate()
}
