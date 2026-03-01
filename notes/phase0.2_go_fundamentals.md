# Phase 0.2 — Go Language Fundamentals
**تاريخ الجلسة:** 2026-02-28

---

## 1. Variables and Constants

### `var` vs `:=`
```go
var age int = 25        // رسمي، بيحدد النوع صراحةً
var name = "Mohamed"    // type inference
age := 25               // short declaration — جوه function بس
```
- `:=` يشتغل **جوه function بس**. برّاها لازم `var`.
- `:=` في نفس الـ scope مرتين → **compile error** (مش shadowing).

### Zero Values
| Type | Zero Value |
|------|-----------|
| int | 0 |
| float64 | 0.0 |
| bool | false |
| string | "" |
| pointer/slice/map | nil |

### `const` و `iota`
```go
const (
    North = iota   // 0
    South          // 1
    East           // 2
    West           // 3
)

const (
    _  = iota
    KB = 1 << (10 * iota)  // 1024
    MB                      // 1,048,576
    GB                      // 1,073,741,824
)
```

### Scope و Shadowing
```go
x := 10
{
    x := 99          // shadowing — variable جديد، مش تعديل على الأصلي
    fmt.Println(x)   // 99
}
fmt.Println(x)       // 10 — الأصلي ما اتغيرش
```
**الخطر:** الـ shadowing صامت — Go ما بيحذّرش منه.

---

## 2. Numeric Types و Booleans

### Integer Types
- **Signed:** `int8` (-128/127), `int16`, `int32`, `int64`, `int` (platform size)
- **Unsigned:** `uint8`/`byte` (0/255), `uint16`, `uint32`, `uint64`, `uint`
- في DSA: استخدم `int` دايماً إلا لو في سبب محدد.

### ليه `int8` من -128 لـ 127؟
- 8 bits → 2⁸ = 256 قيمة
- الـ bit الأقصى يساراً = sign bit
- موجب: 0 → 127 (2⁷ - 1)، سالب: -128 → -1

### `float32` vs `float64`
- `float32` ≈ 7 decimal digits of precision
- `float64` ≈ 15-16 decimal digits of precision
- **دايماً استخدم `float64`** — الـ `float32` بيفقد precision بسرعة.

### `rune` = `int32` = Unicode code point
```go
var r rune = 'ع'   // r == 1593
```

### Type Conversion — دايماً explicit
```go
var x int = 42
var y float64 = float64(x)   // ✓
var z int = y                  // compile error ✗
```

---

## 3. Strings في Go

### نوعين من الـ String Literals
```go
s1 := "Hello\nWorld"   // interpreted — \n = newline فعلي
s2 := `Hello\nWorld`   // raw — \n حرفي، مش newline
```
Raw strings مفيدة لـ: SQL، regex، JSON templates.

### Strings هي Immutable Byte Sequences
```go
s := "Hello"
s[0] = 'h'    // compile error — immutable!

// لو عايز تعدّل:
b := []byte(s)
b[0] = 'h'
s = string(b)
```

### Byte vs Rune Iteration
```go
s := "مرحبا"
len(s)          // 10 bytes (كل حرف عربي = 2 bytes في UTF-8)

for i, r := range s {   // range → runes (حروف كاملة)
    fmt.Println(i, string(r))
}
// index بيقفز مش بالتسلسل لأن الحروف > 1 byte

s[0]            // byte (جزء من حرف!)
```

### ASCII vs Unicode
- **ASCII:** 128 حرف (إنجليزي فقط، 7 bits)
- **Unicode:** كل حرف في العالم، كل لغة (+143,000 حرف)
- **UTF-8:** طريقة تخزين Unicode — 1 byte (English), 2 bytes (Arabic), 4 bytes (Emoji)

### Type Conversions للـ Strings
```go
string(65)           // "A"   ← rune لـ string
strconv.Itoa(65)     // "65"  ← رقم لـ string نصي
[]byte("hello")      // [104 101 108 108 111]
[]rune("مرحبا")      // [1605 1585 1581 1576 1575]
```

---

## 4. Functions في Go

### Basic Syntax
```go
func add(a, b int) int {
    return a + b
}
```

### Multiple Return Values
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
if err != nil { ... }
```
- **مفيش exceptions في Go** — الـ error قيمة عادية كـ second return value.
- `nil` = zero value للـ error = "مفيش error".

### Named Return Values
```go
func minMax(nums []int) (min, max int) {
    // ...
    return   // bare return
}
```

### Variadic Functions
```go
func sum(nums ...int) int { ... }
sum(1, 2, 3)
nums := []int{1, 2, 3}
sum(nums...)   // فكّك الـ slice
```

---

## 5. Control Flow في Go

### الـ `for` — الـ Loop الوحيدة
```go
// Classic
for i := 0; i < 5; i++ { }

// While style
for i < 5 { }

// Infinite
for { break }

// Range (الأهم في DSA)
for i, v := range slice { }
for key, val := range myMap { }
for i, r := range str { }   // runes!
for _, v := range slice { } // تجاهل الـ index
```

### `if` مع init statement
```go
if result, err := divide(10, 2); err != nil {
    // result و err هنا بس جوه الـ if scope
}
```

### `switch` — بيوقف تلقائياً
```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("weekend")
default:
    fmt.Println("weekday")
}

// بدون expression
switch {
case x < 0:  fmt.Println("negative")
default:     fmt.Println("non-negative")
}
```

### `break` و `continue`
```go
for i := 0; i < 10; i++ {
    if i == 3 { continue }  // تخطّى
    if i == 7 { break }     // اخرج
}
```

---

## Bitwise Left Shift
```
1 << 1  = 2    (1 × 2¹)
1 << 10 = 1024 (1 × 2¹⁰)
كل shift لليسار = ضرب في 2
```

---

## أسئلة المراجعة

1. ما الفرق بين `:=` و `=`؟ ومتى تستخدم كل واحدة؟
2. ما هو الـ zero value لكل type؟
3. ليه `"مرحبا"` بتاخد 10 bytes وهي 5 حروف؟
4. ما الفرق بين `range` و index access على string؟
5. إزاي Go بتتعامل مع الـ errors بدل الـ exceptions؟
6. اكتب loop تطبع الأرقام من 1 لـ 10 باستخدام الـ `for` بالطرق الثلاثة.

---

## Q&A من الجلسة

**Q: ما معنى `1 << 10`؟**
الـ `<<` هو bitwise left shift — بيحرّك الـ bits لليسار. كل shift = ضرب في 2.
```
1 << 1  = 2
1 << 2  = 4
1 << 10 = 1024   (1 × 2¹⁰)
```
ولهذا بنحسب KB و MB و GB بـ `1 << 10`، `1 << 20`، `1 << 30`.

---

**Q: اشرح الـ shadowing أكتر.**
الـ shadowing بيحصل لما تعرّف variable بنفس الاسم في scope داخلي — Go بيقبله بصمت من غير error.
```go
x := 10
{
    x := 99        // variable جديد تماماً — مش تعديل على الأصلي
    fmt.Println(x) // 99
}
fmt.Println(x)     // 10 — الأصلي ما اتغيرش
```
الفرق المهم:
- **نفس الـ scope:** `x := 5` ثم `x := 10` → **compile error**
- **scope مختلف (داخلي):** `x := 5` ثم `{ x := 10 }` → **shadowing** (يشتغل بصمت)

الخطر: الكود بيكمبايل ويشتغل، بس النتيجة مش اللي توقعتها.

---

**Q: ليه `int8` من -128 لـ 127؟ وليه `byte` يوصل لـ 256؟**

8 bits → 2⁸ = **256 قيمة ممكنة** دايماً.

- **`byte` = `uint8`** (unsigned): كل الـ 256 قيمة للأرقام الموجبة → 0 إلى 255
- **`int8`** (signed): الـ bit الأقصى يساراً بقى علامة (sign bit):
  - `0xxxxxxx` = موجب: 0 → 127
  - `1xxxxxxx` = سالب: -128 → -1
  - المجموع: 128 + 1 + 127 = 256 قيمة (نفس الـ byte)

---

**Q: ما الفرق بين `float32` و `float64`؟ مثال؟**

الفرق في الـ precision (عدد الأرقام الدقيقة):
- `float32` ≈ 7 أرقام عشرية
- `float64` ≈ 15-16 أرقام عشرية

```go
var f32 float32 = 1.23456789012345678
var f64 float64 = 1.23456789012345678
fmt.Println(f32)  // 1.2345679        ← بيدوّر من الـ digit السابع
fmt.Println(f64)  // 1.2345678901234567 ← دقيق بكتير
```
**القاعدة:** دايماً استخدم `float64` في Go.

---

**Q: ما هو الـ `byte`؟**
الـ `byte` هو alias لـ `uint8` — نفس الحاجة تماماً:
- 8 bits، قيم من 0 إلى 255
- أصغر وحدة قابلة للتعامل في الـ memory
- الـ string في Go = sequence of bytes

---

**Q: ما الفرق بين ASCII و Unicode؟**

| | ASCII | Unicode |
|--|-------|---------|
| الحروف | 128 (إنجليزي فقط) | +143,000 (كل لغات العالم) |
| الحجم | 7 bits | variable |
| مثال | `'A'` = 65 | `'ع'` = 1593، `'😀'` = 128512 |

**UTF-8** = طريقة تخزين Unicode:
- إنجليزي = 1 byte
- عربي = 2 bytes
- emoji = 4 bytes

ولهذا `"مرحبا"` (5 حروف) = 10 bytes، مش 5.

---

## إجابات أسئلة التحقق من الجلسة

**Q: الكود ده بيطبع إيه؟**
```go
x := 5
x := 10   // same scope
```
**الإجابة:** compile error — `:=` مرتين لنفس الـ variable في نفس الـ scope = خطأ.
الصح: `x = 10` (بـ `=` مش `:=`).

---

**Q: `"hello"` كام byte؟ `"مرحبا"` كام byte؟**
- `"hello"` = **5 bytes** (كل حرف إنجليزي = 1 byte في UTF-8/ASCII)
- `"مرحبا"` = **10 bytes** (كل حرف عربي = 2 bytes في UTF-8)
