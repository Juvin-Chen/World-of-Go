package main

// 这个实验有点点有趣...函数式战斗引擎start!
func main() {

}

// task1 闭包工厂——技能生成器
func createFireball(baseDamage int) (func() (int, bool), string) {
	cd := 0 //这里定义一个局部变量来记录冷却回合数

	return func() (int, bool) {
		if cd > 0 {
			cd -= 1
			return 0, false
		} else {
			cd = 3
			return baseDamage, true
		}
	}, "爆裂火球"
}

// task2 伤害计算器
func applyDamage(rawDamage int, modifier func(int) int) int {

}

func modifier(d int) {

}

// task3 核心循环——带 Label 的战斗逻辑
