package frontend

import (
	"../../../framework"
)

func Setup()  {
	xf.HandleStatic("/static/", xf.GetModelesDir("frontend/static"))
}
