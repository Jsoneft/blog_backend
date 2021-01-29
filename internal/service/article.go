package service

type CreateArticleRequest struct {
	TagId         uint32 `form:"tag_id" binding:"required, gte=1"`
	Title         string `form:"title" binding:"required, min=2, max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binding:"required,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	CreatedBy     string `form:"created_by" binding:"required, max=100"`
	State         uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagId         uint32 `form:"tag_id" binding:"required, gte=1"`
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=255"`
	Content       string `form:"content" binging:"max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" binding:"required, max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleByIDRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleByTagIdRequest struct {
	TagId uint32 `form:"tag_id" binding:"required, gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
