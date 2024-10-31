package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode" validate:"required"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
	XCSRFToken string      `json:"csrf,omitempty"`
}

// HTTP status code 200 OK (โอเค) คำขอ HTTP ประสบความสำเร็จ
func StatusOK(c *gin.Context, data interface{}) {
	SendResponse(c, http.StatusOK, "", data)
}

// HTTP status code 201 Created (สร้าง) คำขอได้รับการตอบสนองแล้ว เริ่มสร้างทรัพยากรใหม่ขึ้นมา
func StatusCreated(c *gin.Context, data interface{}) {
	SendResponse(c, http.StatusCreated, "", data)
}

// HTTP status code 202 Accepted (ยอมรับ) คำขอได้รับและกำลังประมวลผลอยู่ แต่ยังดำเนินการไม่เสร็จ ท้ายที่สุดแล้วคำขออาจจะดำเนินการต่อจนเสร็จ หรือไม่สำเร็จก็ได้
func StatusAccepted(c *gin.Context, data interface{}) {
	SendResponse(c, http.StatusAccepted, "", data)
}

// HTTP status code 204 No Content (ไม่มีเนื้อหา) Server ประมวลผลคำขอเสร็จแล้ว และยังไม่ส่งคืนเนื้อหาใด ๆ กลับมา
func StatusNoContent(c *gin.Context, message string) {
	SendResponse(c, http.StatusNoContent, message, nil)
}

// HTTP status code 307 Temporary Redirect (เปลี่ยนเส้นทางชั่วคราว) คำขอจะมีการทำซ้ำด้วย URI ตัวอื่น
func StatusTemporaryRedirect(c *gin.Context, message string, data interface{}) {
	SendResponse(c, http.StatusTemporaryRedirect, message, data)
}

// HTTP status code 400 Bad Request (คำขอไม่ถูกต้อง) เนื่องจากมีข้อผิดพลาดจากทาง Client
func StatusBadRequest(c *gin.Context, err error) {
	SendResponse(c, http.StatusBadRequest, err.Error(), nil)
}

// HTTP status code 401 Unauthorized (ไม่ได้รับอนุญาต) ตรวจสอบสิทธิ์ (Authentication) แต่ว่าล้มเหลว หรือยังไม่ได้รับการยืนยัน
func StatusUnauthorized(c *gin.Context, message string) {
	SendResponse(c, http.StatusUnauthorized, message, nil)
}

// HTTP status code 403 Forbidden (หวงห้าม) มีข้อมูลที่คำขอต้องการอยู่ แต่ทาง Server ปฏิเสธที่จะดำเนินการต่อ โดยอาจจะเป็นเพราะผู้ส่งคำขอไม่ได้รับการอนุญาต
func StatusForbidden(c *gin.Context) {
	SendResponse(c, http.StatusForbidden, "forbidden access", nil)
}

// HTTP status code 404 Not Found (ไม่พบ) ที่อยู่ URL ผิด หรือที่อยู่เว็บไซต์นั้นไม่มีอยู่จริง
func StatusNotFound(c *gin.Context, message string) {
	SendResponse(c, http.StatusNotFound, message, nil)
}

// HTTP status code 417 Expectation Failed (ไม่สามารถทำตามที่คาดหวังได้) Server ไม่สามารถทำตามความต้องการตามที่คำขอ Header ถูกส่งเข้ามาได้
func StatusExpectationFailed(c *gin.Context, message string) {
	SendResponse(c, http.StatusExpectationFailed, message, nil)
}

// HTTP status code 423 Locked (ล็อก) ทรัพยากรที่ต้องการเข้าถึงถูกล็อกเอาไว้อยู่
func StatusLocked(c *gin.Context, message string) {
	SendResponse(c, http.StatusLocked, message, nil)
}

// HTTP status code 428 Precondition Required (มีเงื่อนไขเบื้องต้นที่จำเป็น) Server ต้นทางต้องการให้คำขอตรงตามเงื่อนไขที่กำหนดไว้ เพื่อป้องกันไม่ให้การอัปเดตเกิดข้อผิดพลาด
func StatusPreconditionFailed(c *gin.Context, message string) {
	SendResponse(c, http.StatusPreconditionFailed, message, nil)
}

// HTTP status code 500 Internal Server Error (พบข้อผิดพลาดภายในเซิร์ฟเวอร์)
func StatusInternalServerError(c *gin.Context, err error) {
	SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
}

// HTTP status code 503 Service Unavailable (บริการไม่พร้อมใช้งาน) Server ไม่สามารถดำเนินการตามคำขอได้ สาเหตุอาจจะมาจากมีภาระการทำงานหนักเกินกว่าจะรับไหว หรืออยู่ในช่วงบำรุงรักษา ส่วนใหญ่แล้วจะเป็นแค่เพียงชั่วคราว
func StatusServiceUnavailable(c *gin.Context, service string, err error) {
	msg := fmt.Sprintf("%s unavailable: %s", service, err.Error())
	SendResponse(c, http.StatusServiceUnavailable, msg, nil)
}

// Send Response
func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		StatusCode: status,
		Message:    strings.ToLower(message),
		Data:       data,
		XCSRFToken: c.GetHeader("X-CSRF-Token"),
	})
}
