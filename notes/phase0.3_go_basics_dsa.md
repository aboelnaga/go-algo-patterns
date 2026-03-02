# Phase 0.3 — Go Basics for DSA
**تاريخ الجلسة:** 2026-03-02

---

## 1. Slices vs Arrays

### Arrays
- Fixed size، النوع جزء من التعريف: `[3]int` ≠ `[4]int`
- **Value type** — لما بتعمل copy بتتنسخ كل العناصر
```go
a := [3]int{1, 2, 3}
b := a        // نسخة مستقلة تماماً
b[0] = 99     // a[0] لسه 1
```

### Slice Header
الـ slice مش array — هو **struct بثلاث حاجات**:
```
ptr → مؤشر للـ underlying array في الـ heap
len → عدد العناصر المرئية
cap → السعة الكلية للـ underlying array
```

### المشكلة الشائعة — Shared Underlying Array
```go
a := []int{1, 2, 3}
b := a        // copy للـ header بس، مش للـ data
b[0] = 99     // يغيّر الـ array المشتركة!
// a[0] == 99 كمان!
```

### append وإعادة التخصيص
- لو `append` لقى مكان في الـ cap → نفس الـ underlying array
- لو الـ cap اتعدى → Go بيخصص array جديدة، بينسخ البيانات، بيرجع slice header جديد
- الـ slice القديم مش شايف التغيير

### nil Slice
```go
var s []int        // nil slice: ptr=nil, len=0, cap=0
range s            // آمن تماماً — zero iterations
len(s) == 0        // true
```
الـ nil slice والـ empty slice سلوكهم زي بعض في الـ range loop.

---

## 2. Pre-allocation بـ make

```go
// بدون pre-allocation — Go بيعمل realloc أكتر من مرة
output := []int{}

// مع pre-allocation — allocation واحدة بس
output := make([]int, 0, len(s))
```

- `make([]int, 0, n)` → len=0, cap=n
- ما بتغيّرش الـ Big-O، لكن بتلغي الـ hidden reallocations
- الفرق بيتبيّن في الـ benchmark: **1 allocs/op** بدل أكتر

---

## 3. Benchmark في Go

```go
func BenchmarkXxx(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // الكود اللي بتقيسه
    }
}
```

```bash
go test -bench=. -benchmem    # بيشوف الـ allocs/op
```

- الـ Benchmark مش فيه assertions — بس timing
- الـ Test function هي اللي بتتحقق من الصحة

---

## 4. reflect.DeepEqual vs typed comparison

| | `reflect.DeepEqual` | custom `slicesEqual` |
|---|---|---|
| في الـ Test | ✅ مناسب | ✅ |
| في الـ Benchmark | ❌ boxing → extra allocs | ✅ أفضل |

---

## Q&A من الجلسة

**Q: لو بعمل `b := a` على slice وعدّلت `b[0]`، `a[0]` بيتغير؟**
A: نعم — الاتنين بيشاركوا نفس الـ underlying array. التعديل بيأثر على الاتنين.

**Q: لو `append` عدّى الـ capacity، الـ slice القديم بيشوف العناصر الجديدة؟**
A: لا — Go بيخصص array جديدة ويرجع header جديد. الـ variable القديمة لسه بتشاور على الـ array القديمة.

**Q: إزاي تتأكد إن الـ nil input آمن؟**
A: جرّب `range` على nil slice — بترجع zero iterations من غير panic. الكود الصح مش محتاج nil check.

---

## ملخص الـ Exercise

**المشكلة:** اكتب function بتاخد slice وترجع slice جديدة فيها كل عنصر مضروب في 2، من غير ما تعدّل الأصلية.

**الحل الأمثل:**
```go
func doubleSliceElements(s []int) (output []int) {
    output = make([]int, 0, len(s))  // pre-allocate
    for _, v := range s {
        output = append(output, v*2)
    }
    return output
}
```

- Time: O(n)
- Space: O(n) auxiliary
- Allocs: 1 (بفضل make)
