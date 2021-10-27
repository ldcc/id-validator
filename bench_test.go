package idvalidator

import (
    "git.gdqlyt.com.cn/go/id-validator/data"
    "strconv"
    "testing"
)

func BenchmarkIsValid(b *testing.B) {
    benchmarks := []struct {
        name string
        id   string
    }{
        {id: "440308199901101512"},
        {id: "610104620927690"},
        {id: "810000199408230021"},
        {id: "830000199201300022"},
    }
    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                IsValid(bm.name)
            }
        })
    }
}

func BenchmarkGetInfo(b *testing.B) {
    benchmarks := []struct {
        name string
        id   string
    }{
        {id: "440308199901101512"},
        {id: "610104620927690"},
        {id: "810000199408230021"},
        {id: "830000199201300022"},
    }
    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                GetInfo(bm.name)
            }
        })
    }
}

func BenchmarkFakeId(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FakeId()
    }
}

func BenchmarkFakeRequireId(b *testing.B) {
    benchmarks := []struct {
        name       string
        isEighteen bool
        address    string
        birthday   string
        sex        int
    }{
        {isEighteen: false, address: "浙江省", birthday: "20000101", sex: 1},
        {isEighteen: true, address: "浙江省", birthday: "20000101", sex: 0},
        {isEighteen: true, address: "台湾省", birthday: "20000101", sex: 0},
        {isEighteen: true, address: "香港特别行政区", birthday: "20000101", sex: 0},
    }
    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                FakeRequireId(bm.isEighteen, bm.address, bm.birthday, bm.sex)
            }
        })
    }
}

func BenchmarkUpgradeId(b *testing.B) {
    benchmarks := []struct {
        name string
        id   string
    }{
        {id: "610104620927690"},
        {id: "61010462092769"},
    }
    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                UpgradeId(bm.id)
            }
        })
    }
}

func TestStoreJSON(t *testing.T) {
    data.Stroe("address_code_changes.json", &data.AddressCodeChanges)
    data.Stroe("address_code_merges.json", &data.AddressCodeMerges)
    data.Stroe("address_name_changes.json", &data.AddressNameChanges)
}

func TestLoopPointer(t *testing.T) {
    a := struct {
        Data string
    }{Data: "123"}
    p := &a.Data
    for i := 0; i < 3; i++ {
        a.Data = strconv.Itoa(i)
        t.Log(*p)
    }
}
