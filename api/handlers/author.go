package handlers

import (
	"editory_submission/api/http"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"github.com/gin-gonic/gin"
)

// GetAdminAuthorList godoc
// @ID get_admin_author_list
// @Router /admin/author [GET]
// @Summary Get Author List
// @Description  Get Author List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetUserListRes} "GetEditorListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminAuthorList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserService().GetUserListByRole(
		c.Request.Context(),
		&auth_service.GetUserListByRoleReq{
			Limit:    int32(limit),
			Offset:   int32(offset),
			Search:   c.DefaultQuery("search", ""),
			RoleType: config.AUTHOR,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
