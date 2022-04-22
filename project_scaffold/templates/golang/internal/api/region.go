{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"

	"{{WEB_FRAMEWORK_IMPORT}}"

	"{{GOLANG_MODULE}}/internal/form"
	"{{GOLANG_MODULE}}/internal/service"
)

func init() {
	AddRouteRegistrar(RegionBy)
}

type RegionLocation struct {
	Region   string    `json:"region" example:"区域名称"`
	Location [2]string `json:"location" example:"经纬度"`
}

type RegionQuery struct {
	Code     string `json:"code" form:"code" url:"code"`
	Address  string `json:"address" form:"address" url:"address"`
	Location string `json:"location" form:"location" url:"location"`
}

// RegionBy
// @Summary      获取区域数据
// @Description  获取区域数据
// @Tags         程序设置
// @Accept       application/x-www-form-urlencoded
// @Param        code      query  string  false  "区域代码"            default(110108)
// @Param        address   query  string  false  "省市区详细地址 正向地理编码"  default(北京市海淀区燕园街道北京大学)
// @Param        location  query  string  false  "经度,纬度 逆向地理编码"    default(116.310003,39.991957)
// @Produce      json
// @Success      200  {object}  query.Response{result=RegionLocation}  "操作成功"
// @Router       /region/by [get]
func RegionBy(router {{ROUTER_GROUP}}) {
	router.Any("/region/by", func(c {{WEB_CONTEXT}}) {
		var f RegionQuery

		if err := form.ShouldBind(c, &f); err != nil {
			ErrorInvalidParameters(c, err)
			return{{NIL_STRING}}
		}

		var result RegionLocation

		if f.Code != "" {
			region, location := conf.GetRegionByCode(f.Code)
			result = RegionLocation{
				Region:   region,
				Location: location,
			}
		} else if f.Address != "" {
			location, err := service.GetForwardGeocodingInformation(f.Address)
			if err != nil {
				ErrorUnexpected(c, err)
				return{{NIL_STRING}}
			}

			_location := strings.SplitN(location, ",", 2)
			result = RegionLocation{
				Region:   f.Address,
				Location: [2]string{_location[0], _location[1]},
			}
		} else if f.Location != "" {
			region, err := service.GetReverseGeocodingInformation(f.Location)
			if err != nil {
				ErrorUnexpected(c, err)
                return{{NIL_STRING}}
			}

			_location := strings.SplitN(f.Location, ",", 2)
			result = RegionLocation{
				Region:   region,
				Location: [2]string{_location[0], _location[1]},
			}
		}

		SendJSON(c, result)
		return{{NIL_STRING}}
	})
}
