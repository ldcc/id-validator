package idvalidator

import (
    "errors"
    "strconv"
    "time"

    "git.gdqlyt.com.cn/go/id-validator/data"
)

// 身份证信息
type idInfo struct {
    AddressCode int
    Abandoned   int
    Address     string
    NewAddress  string
    //AddressTree   []string
    AddressInfo   AddressInfo
    Birthday      time.Time
    Constellation string
    ChineseZodiac string
    Sex           int
    Length        int
    CheckBit      string
}

type AddressInfo struct {
    ProvinceCode    string //省编码
    ProvinceName    string
    CityCode        string //市编码
    CityName        string
    DistrictCode    string //区编码
    DistrictName    string
    NewDistrictCode string //新的区编码
    NewDistrictName string
}

// 验证身份证号合法性
func IsValid(id string) bool {
    code, err := generateCode(id)
    if err != nil {
        return false
    }

    // 检查顺序码、生日码、地址码
    if !checkOrderCode(code["order"]) || !checkBirthdayCode(code["birthdayCode"]) || !checkAddressCode(code["addressCode"], code["birthdayCode"]) {
        return false
    }

    // 15位身份证不含校验码
    if code["type"] == "15" {
        return true
    }

    return code["checkBit"] == generatorCheckBit(code["body"])
}

// 获取身份证信息
func GetInfo(id string) (idInfo, error) {
    // 验证有效性
    if !IsValid(id) {
        return idInfo{}, errors.New("Not Valid ID card number.")
    }

    code, _ := generateCode(id)

    // 地址信息
    addressInfo := getAddressInfo(code["addressCode"], code["birthdayCode"])
    //var addressTree []string
    //for _, val := range addressInfo {
    //    addressTree = append(addressTree, val)
    //}

    // 是否废弃
    var abandoned int
    if data.AddressCode[code["addressCode"]] == "" {
        abandoned = 1
    }

    // 生日
    birthday, _ := time.Parse("20060102", code["birthdayCode"])

    // 性别
    sex := 1
    sexCode, _ := strconv.Atoi(code["order"])
    if (sexCode % 2) == 0 {
        sex = 0
    }

    // 长度
    length, _ := strconv.Atoi(code["type"])
    addressCode, _ := strconv.Atoi(code["addressCode"])
    return idInfo{
        AddressCode: addressCode,
        Abandoned:   abandoned,
        Address:     addressInfo.ProvinceName + addressInfo.CityName + addressInfo.DistrictName,
        NewAddress:  addressInfo.ProvinceName + addressInfo.CityName + addressInfo.NewDistrictName,
        //AddressTree: addressTree,
        AddressInfo:   addressInfo,
        Birthday:      birthday,
        Constellation: getConstellation(code["birthdayCode"]),
        ChineseZodiac: getChineseZodiac(code["birthdayCode"]),
        Sex:           sex,
        Length:        length,
        CheckBit:      code["checkBit"],
    }, nil
}

// 生成假身份证号码
func FakeId() string {
    return FakeRequireId(true, "", "", 0)
}

// 按要求生成假身份证号码
// isEighteen 是否生成18位号码
// address    省市县三级地区官方全称：如`北京市`、`台湾省`、`香港特别行政区`、`深圳市`、`黄浦区`
// birthday   出生日期：如 `2000`、`198801`、`19990101`
// sex        性别：1为男性，0为女性
func FakeRequireId(isEighteen bool, address string, birthday string, sex int) string {
    // 生成地址码
    addressCode := generatorAddressCode(address)

    // 出生日期码
    birthdayCode := generatorBirthdayCode(birthday)

    // 生成顺序码
    orderCode := generatorOrderCode(sex)

    if !isEighteen {
        return addressCode + substr(birthdayCode, 2, 8) + orderCode
    }

    body := addressCode + birthdayCode + orderCode

    return body + generatorCheckBit(body)
}

// 15位升级18位号码
func UpgradeId(id string) (string, error) {
    if !IsValid(id) {
        return "", errors.New("Not Valid ID card number.")
    }

    code, _ := generateShortCode(id)

    body := code["addressCode"] + code["birthdayCode"] + code["order"]

    return body + generatorCheckBit(body), nil
}
