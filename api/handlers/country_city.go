package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"
	"github.com/gin-gonic/gin"
)

// GetCountryList godoc
// @ID get_country_list
// @Router /country [GET]
// @Summary Get Country List
// @Description  Get Country List
// @Tags Content
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetCountryListRes} "GetCountryListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCountryList(c *gin.Context) {

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

	resp, err := h.services.ContentService().GetCountryList(
		c.Request.Context(),
		&content_service.GetCountryListReq{
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

// GetCityList godoc
// @ID get_city_list
// @Router /city [GET]
// @Summary Get City List
// @Description  Get City List
// @Tags Content
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param country-id query string false "country-id"
// @Success 200 {object} http.Response{data=content_service.GetCityListRes} "GetCityListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCityList(c *gin.Context) {

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

	resp, err := h.services.ContentService().GetCityList(
		c.Request.Context(),
		&content_service.GetCityListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.Query("search"),
			CountryId: c.Query("country-id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
