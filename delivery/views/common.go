package view

import "net/http"

func InternalServerError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Cannot Access Database",
		"status":  false,
		"data":    nil,
	}
}

func StatusIdConversion() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Cannot Convert ID",
		"status":  false,
		"data":    nil,
	}
}

func StatusInvalidRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Invalid Request",
		"status":  false,
		"data":    nil,
	}
}

func StatusBindData() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnsupportedMediaType,
		"message": "Cannot Bind Data",
		"status":  false,
		"data":    nil,
	}
}

func StatusValidate() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotAcceptable,
		"message": "Validate Error",
		"status":  false,
		"data":    nil,
	}
}

func StatusDelete() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Deleted",
		"status":  true,
		"data":    nil,
	}
}

func StatusUnauthorized(err error) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"message": err.Error(),
		"status":  false,
		"data":    nil,
	}
}

func StatusForbidden(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusForbidden,
		"message": message,
		"status":  false,
		"data":    nil,
	}
}

func StatusNotFound(message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusNotFound,
		"message": message,
		"status":  false,
		"data":    nil,
	}
}

func StatusOK(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": message,
		"status":  true,
		"data":    data,
	}
}

func StatusCreated(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": message,
		"status":  true,
		"data":    data,
	}
}

func StatusUpdate(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Updated",
		"status":  true,
		"data":    data,
	}
}
func StatusGetDatIdOK(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get Data ID",
		"status":  true,
		"data":    data,
	}
}
