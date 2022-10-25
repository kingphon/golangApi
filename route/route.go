package route

import (
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	companyType(e)
	company(e)
	cabinet(e)
	drawer(e)
	documentType(e)
	department(e)
	staff(e)
	permission(e)
	document(e)
}
