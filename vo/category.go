package vo

type CreateCategoryRequestJson struct {
	Name string `json:"name" binding:"required"`
}
