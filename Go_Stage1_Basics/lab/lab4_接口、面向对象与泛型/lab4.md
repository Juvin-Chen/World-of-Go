# 🧪 实验 4: 接口抽象与泛型编程实战

**实验目标**：跳出“写脚本”的思维，学会用 **接口 (Interface)** 设计系统，用 **泛型 (Generics)** 减少重复代码，用 **组合 (Composition)** 替代继承。这是成为 Go 高级工程师的必经之路。

------

### 🧪 任务 1：接口的“多态”魔法 (Polymorphism)

**实验目的**：理解“鸭子类型”接口，学会用空接口 `interface{}` 处理异构数据，并安全地进行类型断言。

**📝 场景：USB 设备管理器**

1. **定义接口**：
   - 定义一个接口 `USB`，包含两个方法：
     - `Name() string`: 返回设备名称。
     - `Connect() string`: 返回连接状态（如 "Connected"）。
2. **实现接口 (结构体)**：
   - 定义结构体 `Mouse` (鼠标)，字段 `Model` (string)。实现 `USB` 接口。
   - 定义结构体 `FlashDrive` (U盘)，字段 `Capacity` (int)。实现 `USB` 接口。
3. **异构切片**：
   - 在 `main` 中定义一个 **USB 切片**：`var sockets []USB`。
   - 向切片中 `append` 一个鼠标和一个 U 盘。
4. **遍历与断言 (核心任务)**：
   - 遍历 `sockets`，依次调用 `Name()` 和 `Connect()`。
   - **难点（类型断言）**：
     - 在遍历中，使用 **类型断言 (`socket.(type)`)** 检查当前设备具体是什么类型。
     - 如果是 `FlashDrive`，**额外打印**它的容量信息（例如："检测到U盘，容量: 64GB"）。
     - 如果是 `Mouse`，不需要额外操作。
   - *思考*：为什么在调用 `Capacity` 字段前必须做类型断言？

------

### 🧪 任务 2：泛型的“影分身” (Generics & Constraints)

**实验目的**：告别重复造轮子，体验 Go 1.18+ 泛型的威力。学会编写通用的工具函数。

**📝 场景：通用比较器 & 累加器**

1. **通用最大值 (`Max`)**：
   - **痛点**：以前你需要 `MaxInt`, `MaxFloat`, `MaxString`...
   - **任务**：编写一个泛型函数 `func Max[T ...](a, b T) T`。
   - **约束**：`T` 必须是 **可排序的**（支持 `>` 操作符）。
     - *提示*：Go 标准库 `cmp` 包里有 `cmp.Ordered` 约束，或者你自己定义一个接口 `Ordered` 包含 `int | float64 | string`。
   - **测试**：分别传入 `(10, 20)`，`(3.14, 1.59)`，`("apple", "banana")`，验证结果。
2. **通用结构体 (`Box`)**：
   - 定义一个泛型结构体 `Box[T any]`，包含一个字段 `Content T`。
   - 给 `Box` 添加一个方法 `Extract() T`，返回 `Content`，并将 `Content` 重置为该类型的零值。
   - **测试**：创建一个 `Box[int]` 和一个 `Box[string]`，存入数据后取出来，确认 `Box` 变空了。

------

### 🧪 任务 3：组合优于继承 (Composition over Inheritance)

**实验目的**：彻底忘掉 `extends`，学会用 Go 的“匿名嵌入”来实现代码复用。

**📝 场景：RPG 游戏角色系统**

1. **基础组件 (`BaseCharacter`)**：
   - 定义结构体 `BaseCharacter`，包含字段 `Name`, `HP`。
   - 实现方法 `Move()`：打印 "Name is moving..."。
   - 实现方法 `Attack()`：打印 "Base attack!"。
2. **职业扩展 (`Warrior` & `Mage`)**：
   - 定义 `Warrior`，**匿名嵌入** `BaseCharacter`，并增加字段 `Strength (int)`。
   - 定义 `Mage`，**匿名嵌入** `BaseCharacter`，并增加字段 `Mana (int)`。
3. **方法重写 (Method Overriding)**：
   - 给 `Mage` 重新定义 `Attack()` 方法：打印 "Fireball casting!" (覆盖基类方法)。
   - `Warrior` **不重写** `Attack()`，保持默认。
4. **调用测试**：
   - 创建 `w := Warrior{...}` 和 `m := Mage{...}`。
   - 分别调用它们的 `Move()` (应该都用基类的)。
   - 分别调用它们的 `Attack()` (战士用基类的，法师用自己的)。
   - *验证*：通过 `w.Name` 直接访问基类字段。

------

### 🚨 增补实验：C++ 选手的“翻车”现场

这两个实验专门用来打破你的 C++ 惯性思维。请务必运行并观察结果。

------

#### 🧪 任务 4：接口的“非空”幻觉 (The Nil Interface Trap)

**背景**：在 C++ 中，`Base* ptr = nullptr;` 就是空。但在 Go 中，接口是一个“盒子”，包含 `(type, value)`。

**📝 任务：制造一个“不等于 nil 的 nil”**

1. **定义**：定义一个结构体 `Cat` 和一个接口 `Animal`（包含 `Speak()` 方法）。
2. **制造陷阱**：
   - 声明一个具体的指针：`var c *Cat = nil` （这是一个明确的空指针）。
   - 声明一个接口变量：`var a Animal = c` （把空指针赋给接口）。
3. **观测**：
   - 打印 `c == nil` （预期：true）。
   - 打印 `a == nil` （预期：**???**）。
   - 尝试调用 `a.Speak()`（假设 `Speak` 方法内部没有处理 nil，会发生什么？）。
4. **思考题**：为什么 `a` 明明装了一个 `nil`，但 `a == nil` 却返回 `false`？
   - *(提示：检查 `a` 的动态类型是什么)*

------

#### 🧪 任务 5：指针接收者的“严格约束” (Method Set Rules)

**背景**：你习惯了 C++ 对象是指针还是值都能调方法。但在 Go 的**接口赋值**中，编译器极其严格。

**📝 任务：编译器报错挑战**

1. **定义**：
   - 接口 `Mover`，包含方法 `Move()`。
   - 结构体 `Car`。
2. **实现**：
   - 给 `Car` 实现 `Move()` 方法，但**必须使用指针接收者**：`func (c *Car) Move() { ... }`。
3. **触发编译错误**：
   - 尝试将一个 **Car 的值** 赋给接口：`var m Mover = Car{}`。
4. **修正**：
   - 尝试将一个 **Car 的指针** 赋给接口：`var m Mover = &Car{}`。
5. **总结规则**：
   - 为什么值接收者 `(c Car)` 既能给值也能给指针，但指针接收者 `(c *Car)` 只能接收指针？
   - *(这对你设计系统架构非常重要)*

------

### 💳 综合实战项目：通用支付网关 (Universal Payment Gateway)

**项目代号**：`Stage4_OOP_Interface/project/payment` **目标**：模拟一个电商支付系统。你需要设计一个能够兼容支付宝、微信、信用卡等多种支付方式的架构，并使用泛型记录日志。

#### 📋 需求文档

**1. 核心接口定义 (`PaymentMethod`)** 定义一个接口 `PaymentMethod`，规范所有支付方式的行为：

- `Pay(amount float64) bool`: 执行支付，返回成功/失败。
- `GetID() string`: 获取支付账号/ID。

**2. 支付方式实现 (组合与多态)**

- **基础结构体 `BasePay`**：包含 `ID string` 和 `Balance float64` (余额)。实现通用的 `CheckBalance(amount)` 方法。
- **具体实现 A (`AliPay`)**：
  - 嵌入 `BasePay`。
  - 实现 `Pay`：先检查余额。如果足够，扣款并打印 "支付宝支付成功"；否则返回 false。
- **具体实现 B (`CreditCard`)**：
  - 嵌入 `BasePay`。
  - 实现 `Pay`：信用卡的逻辑不同，它允许透支（假设有个 `Limit` 额度）。只要 `Balance + Limit >= amount` 就算成功。

**3. 泛型日志系统 (`Logger`)**

- 定义一个泛型结构体 `Transaction[T any]`。
  - 字段 `ID string` (交易流水号)。
  - 字段 `Data T` (交易详情，类型不确定)。
- 编写一个泛型函数 `LogTransaction[T any](t Transaction[T])`，打印交易日志。

**4. 支付网关 (`Checkout`)**

- 编写函数 `Checkout(method PaymentMethod, amount float64)`。
- **逻辑**：
  1. 调用 `method.Pay(amount)`。
  2. 如果成功：
     - 创建一个 `Transaction[float64]` 记录金额日志，调用 `LogTransaction`。
  3. 如果失败：
     - 创建一个 `Transaction[string]` 记录错误信息 "Payment Failed"，调用 `LogTransaction`。

**5. 主程序剧本 (`Main Scenario`)** 在 `main` 中：

1. 创建一个 **支付宝** 账户（余额 100 元）。
2. 创建一个 **信用卡** 账户（余额 0 元，额度 1000 元）。
3. **测试 1**：用支付宝买 50 元的东西（应成功，余额变 50）。
4. **测试 2**：用支付宝买 200 元的东西（应失败，余额不够）。
5. **测试 3**：用信用卡买 500 元的东西（应成功，触发透支逻辑）。
6. **观察日志**：注意看 `LogTransaction` 如何处理 `float64` 和 `string` 两种不同的日志数据。

------

### 💡 避坑指南

- **泛型实例化**：调用泛型函数时，通常 Go 编译器能自动推断类型，如 `Max(1, 2)`。但如果推断不出，需要显式写 `Max[int](1, 2)`。
- **方法重写陷阱**：Go 的“重写”其实是**遮蔽 (Shadowing)**。如果你把 `Mage` 赋值给一个 `BaseCharacter` 类型的变量，调用 `Attack` 依然会调用基类的方法（这和 C++ 的虚函数表不一样！Go 没有动态绑定基类方法）。但在本项目中，我们直接调用具体对象的方法，或者通过 Interface 调用，都能得到预期结果。
- **接口判空**：如果一个接口变量内部存的是一个 `nil` 指针，该接口变量本身**不是 nil**。`if i == nil` 可能会坑你。

🚀 **开始行动**：请创建 `lab4.go`，把5个 Demo 任务封装好，最后实现 Project。