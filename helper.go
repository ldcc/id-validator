package idvalidator

import (
	"git.gdqlyt.com.cn/go/id-validator/data"
	"strconv"
	"strings"
)

// 获取地址信息
func getAddressInfo(addressCode string, birthdayCode string) AddressInfo {
	//addressInfo := map[string]string{
	//    "province":      "",
	//    "province_code": "",
	//    "city":          "",
	//    "city_code":     "",
	//    "district":      "",
	//    "district_code": "",
	//}
	addressInfo := AddressInfo{}

	// 省级信息
	addressInfo.ProvinceCode, addressInfo.ProvinceName = getAddress(substr(addressCode, 0, 2)+"0000", birthdayCode)

	// 用于判断是否是港澳台居民居住证（8字开头）
	// 港澳台居民居住证无市级、县级信息
	if substr(addressCode, 0, 1) == "8" {
		return addressInfo
	}

	// 市级信息
	addressInfo.CityCode, addressInfo.CityName = getAddress(substr(addressCode, 0, 4)+"00", birthdayCode)

	// 县级信息
	addressInfo.DistrictCode, addressInfo.DistrictName = getAddress(addressCode, birthdayCode)
	addressInfo.NewDistrictCode = addressInfo.DistrictCode
	var committed = make(map[string]bool)
	if data.AddressCode[addressCode] == "" {
		// 处理县级区域合并与县级更换编码
		code := &addressInfo.NewDistrictCode
		for true {
			mergeCode, exist := data.AddressCodeMerges[*code]
			if exist && !committed[mergeCode] {
				committed[mergeCode] = true
				addressInfo.NewDistrictCode = mergeCode
				addressInfo.NewDistrictName = data.AddressCode[*code]
				continue
			}
			changes, exist := data.AddressCodeChanges[*code]
			if exist && !committed[changes["new_code"]] {
				committed[changes["new_code"]] = true
				addressInfo.NewDistrictCode = changes["new_code"]
				addressInfo.NewDistrictName = changes["new_district"]
				continue
			}
			break
		}
	}

	return addressInfo
}

// 获取省市区地址码
func getAddress(addressCode string, birthdayCode string) (string, string) {
	address := ""
	year, _ := strconv.Atoi(substr(birthdayCode, 0, 4))
	for i, val := range data.AddressNameChanges[addressCode] {
		startYear, _ := strconv.Atoi(val["start_year"])
		if (i == 0 && year < startYear) || year >= startYear {
			address = val["address"]
		}
	}

	return addressCode, address
}

// 获取星座信息
func getConstellation(birthdayCode string) string {
	monthStr := substr(birthdayCode, 4, 6)
	dayStr := substr(birthdayCode, 6, 8)
	month, _ := strconv.Atoi(monthStr)
	day, _ := strconv.Atoi(dayStr)
	startDate := data.Constellation[month]["start_date"]
	startDay, _ := strconv.Atoi(strings.Split(startDate, "-")[1])
	if day >= startDay {
		return data.Constellation[month]["name"]
	}

	tmpMonth := month - 1
	if month == 1 {
		tmpMonth = 12
	}

	return data.Constellation[tmpMonth]["name"]
}

// 获取生肖信息
func getChineseZodiac(birthdayCode string) string {
	// 子鼠
	start := 1900
	end, _ := strconv.Atoi(substr(birthdayCode, 0, 4))
	key := (end - start) % 12

	return data.ChineseZodiac[key]
}

// substr 截取字符串
func substr(source string, start int, end int) string {
	r := []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
