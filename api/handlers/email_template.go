package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/notification_service"
	"editory_submission/pkg/util"
	"github.com/gin-gonic/gin"
)

// CreateAdminEmailTmp godoc
// @ID create_email_template
// @Router /admin/email/template [POST]
// @Summary Create EmailTmp
// @Description Create EmailTmp
// @Tags Admin
// @Accept json
// @Produce json
// @Param email_template body notification_service.CreateEmailTmpReq true "CreateEmailTmpRequestBody"
// @Success 201 {object} http.Response{data=notification_service.CreateEmailTmpRes} "EmailTmp data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminEmailTmp(c *gin.Context) {
	var emailTemplate notification_service.CreateEmailTmpReq

	err := c.ShouldBindJSON(&emailTemplate)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EmailTmpService().CreateEmailTmp(
		c.Request.Context(),
		&emailTemplate,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAdminEmailTmpList godoc
// @ID get_admin_email_template_list
// @Router /admin/email/template [GET]
// @Summary Get EmailTmp List
// @Description  Get EmailTmp List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=notification_service.GetEmailTmpListRes} "GetEmailTmpListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminEmailTmpList(c *gin.Context) {

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

	resp, err := h.services.EmailTmpService().GetEmailTmpList(
		c.Request.Context(),
		&notification_service.GetEmailTmpListReq{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.DefaultQuery("search", ""),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetAdminEmailTmpByID godoc
// @ID get_admin_email_template_by_id
// @Router /admin/email/template/{template-id} [GET]
// @Summary Get EmailTmp By ID
// @Description Get EmailTmp By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param template-id path string true "template-id"
// @Success 200 {object} http.Response{data=notification_service.GetEmailTmpRes} "EmailTmpBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminEmailTmpByID(c *gin.Context) {
	emailTemplateID := c.Param("template-id")

	if !util.IsValidUUID(emailTemplateID) {
		h.handleResponse(c, http.InvalidArgument, "email_template id is an invalid uuid")
		return
	}

	resp, err := h.services.EmailTmpService().GetEmailTmp(
		c.Request.Context(),
		&notification_service.GetEmailTmpReq{
			Id: emailTemplateID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateAdminEmailTmp godoc
// @ID update_email_template
// @Router /admin/email/template [PUT]
// @Summary Update EmailTmp
// @Description Update EmailTmp
// @Tags Admin
// @Accept json
// @Produce json
// @Param email_template body notification_service.UpdateEmailTmpReq true "UpdateEmailTmpRequestBody"
// @Success 200 {object} http.Response{data=notification_service.UpdateEmailTmpRes} "EmailTmp data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminEmailTmp(c *gin.Context) {
	var emailTemplate notification_service.UpdateEmailTmpReq

	err := c.ShouldBindJSON(&emailTemplate)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EmailTmpService().UpdateEmailTmp(
		c.Request.Context(),
		&emailTemplate,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAdminEmailTmp godoc
// @ID delete_email_template
// @Router /admin/email/template/{template-id} [DELETE]
// @Summary Delete EmailTmp
// @Description Get EmailTmp
// @Tags Admin
// @Accept json
// @Produce json
// @Param template-id path string true "template-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAdminEmailTmp(c *gin.Context) {
	emailTemplateID := c.Param("template-id")

	if !util.IsValidUUID(emailTemplateID) {
		h.handleResponse(c, http.InvalidArgument, "email_template id is an invalid uuid")
		return
	}

	_, err := h.services.EmailTmpService().DeleteEmailTmp(
		c.Request.Context(),
		&notification_service.DeleteEmailTmpReq{
			Id: emailTemplateID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
