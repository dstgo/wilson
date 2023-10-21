package node

import "github.com/gin-gonic/gin"

type NodeHandler struct {
}

// Create
// @Summary      Create
// @Description  create a new node
// @Tags         node
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /node/create [POST]
// @security BearerAuth
func (n NodeHandler) Create(ctx *gin.Context) {

}

// Update
// @Summary      Update
// @Description  update the specified node
// @Tags         node
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /node/update [POST]
// @security BearerAuth
func (n NodeHandler) Update(ctx *gin.Context) {

}

// Remove
// @Summary      Remove
// @Description  remove the specified node
// @Tags         node
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /node/remove [DELETE]
// @security BearerAuth
func (n NodeHandler) Remove(ctx *gin.Context) {

}

// GetNodeList
// @Summary      GetNodeList
// @Description  get node list by page
// @Tags         node
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /node/list [GET]
// @security BearerAuth
func (n NodeHandler) GetNodeList(ctx *gin.Context) {

}

// GetNodeInfo
// @Summary      GetNodeInfo
// @Description  get the specified node info
// @Tags         node
// @Accept       json
// @Produce      json
// @Success      200  {object}  types.Response
// @Router       /node/info [GET]
// @security BearerAuth
func (n NodeHandler) GetNodeInfo(ctx *gin.Context) {

}
