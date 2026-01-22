package main

import "fmt"

// 这个实验有点点有趣...函数式战斗引擎start!

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
	return modifier(rawDamage)
}

// task3 核心循环——带 Label 的战斗逻辑
func main() {
	// 初始化, hp用于表示生命值
	hp := 100
	bossHp := 100

	fireballFunc, skillName := createFireball(50)

	// 定义一个Label,在意外死亡时直接跳出整个战斗
BattleLoop:
	for round := 1; ; round++ {
		fmt.Printf("\n--- 第 %d 回合 (HP: %d) ---\n", round, hp)
		fmt.Println("请选择指令：1.攻击(Attack) 2.逃跑(Run)")

		var choice int

		// 这里可以手动改 choice，或者用 fmt.Scan(&choice)，暂时先这样假设
		if round < 6 {
			choice = 1
		} else {
			choice = 2
		}

		// 逻辑判断
		switch choice {
		case 1: // 攻击
			fmt.Println("你发起了攻击！准备进行 3 连击...")

			for hit := 1; hit <= 3; hit++ {
				// 尝试释放技能 (利用闭包状态)
				dmg, success := fireballFunc()

				if success {
					finalDmg := applyDamage(dmg, func(d int) int {
						if d > 40 {
							return d * 2
						} //暴击策略
						return d
					})
					bossHp -= finalDmg
					fmt.Printf(" [第%d击] %s 释放成功！造成 %d 点伤害 (Boss剩余: %d)\n", hit, skillName, finalDmg, bossHp)
				} else {
					fmt.Println("⚠️ 警告：触发 Boss 反伤风暴！")
					hp = 0
					break BattleLoop
				}
			}
		case 2: //逃跑
			fmt.Println("你逃跑了")
			break BattleLoop
		default:
			fmt.Print("无效指令，跳过本回合")
		}

		if bossHp <= 0 {
			fmt.Println("🏆 胜利！Boss 被击败！")
			break
		}
	}
	if hp <= 0 {
		fmt.Println("💀 Game Over... 你被打败了。")
	} else {
		fmt.Println("战斗结束，英雄存活。")
	}
}
