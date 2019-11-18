package com.datatist1118;

import java.util.Scanner;

/*
 * 表达式的取值：byte,short,int,char
 * jdk5以后可以是枚举
 * jdk7以后可以是String
 * 快捷键：对代码进行格式化
 * ctrl+shift+f
 */
public class SwitchTest {
	public static void main(String[] args) {

		Scanner sc = new Scanner(System.in);
		System.out.println("请输入一个整数（1-7）");
		int weekDay = sc.nextInt();
		switch (weekDay) {
		/*
		 * 
		 * break;语句"不是必须的"。
		 * 如果不写，如果一旦case相应的值成功，
		 * 但内部没有break语句，
		 * 那么将会无条件(不再进行case匹配)
		 * 的继续向下执行其它case中的语句，
		 * 直到遇到break;语句或者到达switch语句结束
		 */
		case 1:
			System.out.println("星期一");
		case 2:
			System.out.println("星期二");
		case 3:
			System.out.println("星期三");
		case 4:
			System.out.println("星期四");
			break;
		case 5:
			System.out.println("星期五");
			break;
		case 6:
			System.out.println("星期六");
			break;
		default:
			/*
			 * default：所有的值都和表达式不匹配，就执行default对应的内容
			 */
			System.out.println("星期日");
			break;
		}
	}
}
