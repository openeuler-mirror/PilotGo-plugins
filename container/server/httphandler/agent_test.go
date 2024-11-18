package httphandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeployDocker(t *testing.T) {
	// 设置测试路由
	router := gin.Default()
	router.POST("/deploy", DeployDocker)

	// 测试用例1: 正常请求
	t.Run("Normal Request", func(t *testing.T) {
		payload := `{"batch": ["container1", "container2"]}`
		req, _ := http.NewRequest("POST", "/deploy", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, float64(0), response["code"])
		assert.Equal(t, "ok", response["status"])
	})

	// 测试用例2: 错误的JSON格式
	t.Run("Invalid JSON", func(t *testing.T) {
		payload := `{"batch": [1, 2]` // 不完整的JSON
		req, _ := http.NewRequest("POST", "/deploy", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, float64(-1), response["code"])
		assert.Equal(t, "parameter error", response["status"])
	})
}
