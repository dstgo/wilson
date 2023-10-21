package container

import "github.com/gin-gonic/gin"

type ContainerHandler struct {
}

// Create
// @Summary      Create
// @Description  create a new container instance
// @Tags         container
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /container/create [POST]
// @security BearerAuth
func (i ContainerHandler) Create(ctx *gin.Context) {

}

// Update
// @Summary      Update
// @Description  update the container instance, like name, resource
// @Tags         container
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /container/update [POST]
// @security BearerAuth
func (i ContainerHandler) Update(ctx *gin.Context) {

}

// Delete
// @Summary      Delete
// @Description  delete a specified container instance
// @Tags         container
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /container/delete [DELETE]
// @security BearerAuth
func (i ContainerHandler) Delete(ctx *gin.Context) {

}

// Stop
// @Summary      Stop
// @Description  stop a container instance
// @Tags         container
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /container/stop [POST]
// @security BearerAuth
func (i ContainerHandler) Stop(ctx *gin.Context) {

}

// Start
// @Summary      Start
// @Description  start a container instance
// @Tags         container
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /container/start [POST]
// @security BearerAuth
func (i ContainerHandler) Start(ctx *gin.Context) {

}

// Restart
// @Summary      Restart
// @Description  restart the container instance
// @Tags         container
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /container/restart [POST]
// @security BearerAuth
func (i ContainerHandler) Restart(ctx *gin.Context) {

}

type ImageContainer struct {
}

// Pull
// @Summary      Pull
// @Description  pull a specified image on specified node
// @Tags         image
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /image/pull [GET]
// @security BearerAuth
func (i ImageContainer) Pull(ctx *gin.Context) {

}

// Create
// @Summary      Create
// @Description  create a new image on specified node
// @Tags         image
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /image/create [POST]
// @security BearerAuth
func (i ImageContainer) Create(ctx *gin.Context) {

}

// Delete
// @Summary      Delete
// @Description  delete a specified image on specified node
// @Tags         image
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /image/delete [DELETE]
// @security BearerAuth
func (i ImageContainer) Delete(ctx *gin.Context) {

}

// Tag
// @Summary      Tag
// @Description  tag a specified image on specified nde
// @Tags         image
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /image/tag [POST]
// @security BearerAuth
func (i ImageContainer) Tag(ctx *gin.Context) {

}
