package data

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

// 行政区划代码（地址码）更新时间线
// 中华人民共和国民政部权威数据
// 注1：台湾省、香港特别行政区和澳门特别行政区暂缺地市和区县信息
// 注2：每月发布的区划变更表是根据区划变更地的统计人员在统计信息系统更新后的情况所绘制，与区划变更文件发布的时间有一定的延迟性，但在每年的最后一次发布变更情况后与区划全年变更文件保持一致。
// Data Source: http://www.mca.gov.cn/article/sj/xzqh/

var (
    AddressCode        = make(map[string]string)
    AddressNameChanges = make(map[string][]map[string]string)
    AddressCodeChanges = make(map[string]map[string]string)
    AddressCodeMerges  = make(map[string]string)
    Constellation      = make(map[int]map[string]string, 12)
    ChineseZodiac      = make(map[int]string, 12)

    _ = func() error {
        err := Load("address_code.json", &AddressCode)
        if err != nil {
            return err
        }

        err = Load("address_name_changes.json", &AddressNameChanges)
        if err != nil {
            return err
        }

        err = Load("address_code_changes.json", &AddressCodeChanges)
        if err != nil {
            return err
        }

        err = Load("address_code_merges.json", &AddressCodeMerges)
        if err != nil {
            return err
        }

        err = Load("constellation.json", &Constellation)
        if err != nil {
            return err
        }

        err = Load("chinese_zodiac.json", &ChineseZodiac)
        if err != nil {
            return err
        }

        return nil
    }()
)

go:embed json/*.json
var dir embed.FS

func Load(jsonName string, data interface{}) error {
    byteValue, err := dir.ReadFile("json/" + jsonName)
    if err != nil {
        return err
    }

    return json.Unmarshal(byteValue, data)
}

func Stroe(jsonName string, data interface{}) {
    jsonString, _ := json.Marshal(data)
    _ = ioutil.WriteFile("data/json/"+jsonName, jsonString, os.ModePerm)
}
